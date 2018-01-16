package sws

import "github.com/veandco/go-sdl2/sdl"

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
	tabledata TableData
	vertical  *ScrollbarWidget
	yoffset   int32
}

// NewTableWidget create a table widget, you need to provide a Model i.e. a TableData object
// When you create a table widget, the number of column is fixed (you cannot add/remove columns for the moment)
func NewTableWidget(w, h int32, tabledata TableData) *TableWidget {
	corewidget := NewCoreWidget(w, h)
	widget := &TableWidget{
		CoreWidget: *corewidget,
		tabledata:  tabledata,
		vertical:   NewScrollbarWidget(15, 20, false),
		yoffset:    0,
	}

	widget.vertical.Move(w-15, 30)
	widget.vertical.SetMinimum(0)
	max := int32(0)
	if h < 25*tabledata.GetNbRows()+30 {
		max = 25*tabledata.GetNbRows() + 30 - h
		widget.AddChild(widget.vertical)
	}
	widget.vertical.SetMaximum(max)
	widget.vertical.SetCallbackValueChanged(func() {
		widget.yoffset = widget.vertical.Currentposition
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

func (self *TableWidget) Resize(width, height int32) {
	self.CoreWidget.Resize(width, height)

	// sanity
	if height < 30 {
		height = 30
	}

	self.vertical.Resize(15, height-30)
	if height < 25*self.tabledata.GetNbRows()+30 {
		self.CoreWidget.AddChild(self.vertical)
		self.vertical.SetMaximum(25*self.tabledata.GetNbRows() + 30 - self.Height())
	} else {
		self.CoreWidget.RemoveChild(self.vertical)
	}
	self.PostUpdate()
}

func (self *TableWidget) renderText(x, y, width, height int32, label string) {
	text, err := self.Font().RenderUTF8Blended(label, sdl.Color{0x0, 0x0, 0x0, 0xff})
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
			self.renderText(xoffset, y*25+30-self.yoffset, columnSize[i], 25, label)
			xoffset += columnSize[i]
		}
	}

	// headers
	self.FillRect(0, 0, self.width, 30, 0xffdddddd)
	xoffset := int32(0)
	for i := int32(0); i < nbcolumns; i++ {
		label, size := self.tabledata.GetHeader(i)
		self.renderText(xoffset, 0, size, 30, label)
		xoffset += size
	}

	// do we show the scrollbar
	if self.Height() < 25*self.tabledata.GetNbRows()+30 {
		if self.vertical.IsDirty() {
			self.vertical.Repaint()
		}
		rectSrc := sdl.Rect{0, 0, self.vertical.Width(), self.vertical.Height()}
		rectDst := sdl.Rect{self.Width() - 1 - self.vertical.Width(), 30, self.vertical.Width(), self.vertical.Height()}
		self.vertical.Surface().Blit(&rectSrc, self.Surface(), &rectDst)
	}

}
