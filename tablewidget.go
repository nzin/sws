package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

const HEADER_HEIGHT = 25

// TableData is an interface use to provide the Model (MVC)
// aka the table datas
type TableData interface {
	GetNbColumns() int32
	GetNbRows() int32
	GetHeader(column int32) (string, int32)
	GetCell(colum, row int32) string
	// when the table grows/shrink
	SetRowUpdateCallback(callback func())
	// when the table need to be refreshed
	SetDataChangeCallback(callback func())
}

type TableWidget struct {
	CoreWidget
	tabledata       TableData
	vertical        *ScrollbarWidget
	headertextcolor sdl.Color
	textcolor       sdl.Color
	yoffset         int32
}

// NewTableWidget create a table widget, you need to provide a Model i.e. a TableData object
// When you create a table widget, the number of column is fixed (you cannot add/remove columns for the moment)
func NewTableWidget(w, h int32, tabledata TableData) *TableWidget {
	corewidget := NewCoreWidget(w, h)
	widget := &TableWidget{
		CoreWidget:      *corewidget,
		tabledata:       tabledata,
		vertical:        NewScrollbarWidget(15, 20, false),
		yoffset:         0,
		textcolor:       sdl.Color{0, 0, 0, 255},
		headertextcolor: sdl.Color{0, 0, 0, 255},
	}

	widget.vertical.Move(w-15, HEADER_HEIGHT)
	widget.vertical.SetMinimum(0)
	max := int32(0)
	if h < 25*tabledata.GetNbRows()+HEADER_HEIGHT {
		max = 25*tabledata.GetNbRows() + HEADER_HEIGHT - h
		widget.AddChild(widget.vertical)
	}
	widget.vertical.SetMaximum(max)
	widget.vertical.SetCallback(func(currentposition int32) {
		widget.yoffset = currentposition
		widget.PostUpdate()
	})

	tabledata.SetDataChangeCallback(func() {
		widget.PostUpdate()
	})

	tabledata.SetRowUpdateCallback(func() {
		widget.Resize(widget.Width(), widget.Height())
	})

	return widget
}

func (self *TableWidget) SetHeaderTextColor(color sdl.Color) {
	self.headertextcolor = color
	self.PostUpdate()
}

func (self *TableWidget) SetTextColor(color sdl.Color) {
	self.textcolor = color
	self.PostUpdate()
}

func (self *TableWidget) Resize(width, height int32) {
	self.CoreWidget.Resize(width, height)

	// sanity
	if height < HEADER_HEIGHT {
		height = HEADER_HEIGHT
	}

	self.vertical.Resize(15, height-HEADER_HEIGHT)
	if height < 25*self.tabledata.GetNbRows()+HEADER_HEIGHT {
		self.vertical.Move(width-15, HEADER_HEIGHT)
		self.AddChild(self.vertical)
		self.vertical.SetMaximum(25*self.tabledata.GetNbRows() + HEADER_HEIGHT - self.Height())
	} else {
		self.RemoveChild(self.vertical)
	}
	self.PostUpdate()
}

func (self *TableWidget) renderText(x, y, width, height int32, label string, textcolor sdl.Color) {
	text, err := self.Font().RenderUTF8Blended(label, textcolor)
	if err != nil {
		return
	}
	wGap := width - text.W
	hGap := 25 - text.H
	if wGap < 0 {
		wGap = 0
	}
	if hGap < 0 {
		hGap = 0
	}
	maxwidth := text.W
	maxheight := text.H
	if maxwidth > width {
		maxwidth = width
	}
	if maxheight > height {
		maxheight = height
	}
	rectSrc := sdl.Rect{0, 0, maxwidth, maxheight}
	rectDst := sdl.Rect{x + (wGap / 2), y + (hGap / 2), width - (wGap / 2), height - (hGap / 2)}
	if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
	}
}

func (self *TableWidget) Repaint() {
	self.FillRect(0, 0, self.width, self.height, 0xffffffff)

	columnSize := make([]int32, self.tabledata.GetNbColumns(), self.tabledata.GetNbColumns())

	// get the column size
	nbcolumns := self.tabledata.GetNbColumns()
	for i := int32(0); i < nbcolumns; i++ {
		_, size := self.tabledata.GetHeader(i)
		columnSize[i] = size
	}

	// print the cells
	for y := int32(0); y < self.tabledata.GetNbRows(); y++ {
		xoffset := int32(0)
		for i := int32(0); i < nbcolumns; i++ {
			label := self.tabledata.GetCell(i, y)
			self.renderText(xoffset, y*25+HEADER_HEIGHT-self.yoffset, columnSize[i], 25, label, self.textcolor)
			xoffset += columnSize[i]
		}
	}

	// headers
	self.FillRect(0, 0, self.width, HEADER_HEIGHT, self.bgColor)
	xoffset := int32(0)
	for i := int32(0); i < nbcolumns; i++ {
		label, size := self.tabledata.GetHeader(i)
		self.renderText(xoffset, 0, size, HEADER_HEIGHT, label, self.headertextcolor)

		//bezel
		self.SetDrawColorHex(0xffffffff)
		self.DrawLine(xoffset, 0, xoffset+size-1, 0)
		self.DrawLine(xoffset, 0, xoffset, HEADER_HEIGHT-1)
		self.SetDrawColor(50, 50, 50, 255)
		self.DrawLine(xoffset+size-1, 1, xoffset+size-1, HEADER_HEIGHT-1)
		self.DrawLine(xoffset+1, HEADER_HEIGHT-1, xoffset+size-1, HEADER_HEIGHT-1)

		xoffset += size
	}

	for _, child := range self.children {
		// adjust the clipping to the current child
		if child.IsDirty() {
			child.Repaint()
		}
		rectSrc := sdl.Rect{0, 0, child.Width(), child.Height()}
		rectDst := sdl.Rect{child.X(), child.Y(), child.Width(), child.Height()}
		child.Surface().Blit(&rectSrc, self.Surface(), &rectDst)
	}
	self.dirty = false
}
