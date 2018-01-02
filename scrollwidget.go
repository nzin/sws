package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

//	"github.com/veandco/go-sdl2/sdl"

const (
	SCROLLBAR_WIDTH = 17
)

type ScrollWidget struct {
	CoreWidget
	subwidget  Widget
	hScrollbar *ScrollbarWidget
	vScrollbar *ScrollbarWidget
	corner     *CoreWidget
	xOffset    int32
	yOffset    int32
	showH      bool
	showV      bool
	showBezel  bool
}

func (self *ScrollWidget) ShowBezel(bezel bool) {
	self.showBezel = bezel
	self.PostUpdate()
}

func (self *ScrollWidget) SetHorizontalPosition(position int32) {
	self.hScrollbar.SetPosition(position)
	self.PostUpdate()
}

func (self *ScrollWidget) SetVerticalPosition(position int32) {
	self.vScrollbar.SetPosition(position)
	self.PostUpdate()
}

func (self *ScrollWidget) ShowVerticalScrollbar(showV bool) {
	self.showV = showV
	self.PostUpdate()
}

func (self *ScrollWidget) ShowHorizontalScrollbar(showH bool) {
	self.showH = showH
	self.PostUpdate()
}

func (self *ScrollWidget) PostUpdate() {
	self.Resize(self.Width(), self.Height())
	self.CoreWidget.PostUpdate()
}

func (self *ScrollWidget) Resize(width, height int32) {
	if width < 45 {
		width = 45
	}
	if height < 45 {
		height = 45
	}
	self.CoreWidget.Resize(width, height)

	self.RemoveChild(self.hScrollbar)
	self.RemoveChild(self.vScrollbar)
	self.RemoveChild(self.corner)
	self.hScrollbar.SetMaximum(self.subwidget.Width() - self.Width())
	self.vScrollbar.SetMaximum(self.subwidget.Height() - self.Height())
	if self.subwidget != nil {
		if self.showBezel == false {

			if self.subwidget.Width() > self.Width()-SCROLLBAR_WIDTH && self.subwidget.Height() > self.Height()-SCROLLBAR_WIDTH && self.showH == true && self.showV == true {
				self.hScrollbar.SetMaximum(self.subwidget.Width() - (self.Width() - SCROLLBAR_WIDTH))
				self.AddChild(self.hScrollbar)
				self.hScrollbar.Resize(self.Width()-SCROLLBAR_WIDTH, SCROLLBAR_WIDTH)
				self.hScrollbar.Move(0, self.Height()-SCROLLBAR_WIDTH)
				self.vScrollbar.SetMaximum(self.subwidget.Height() - (self.Height() - SCROLLBAR_WIDTH))
				self.AddChild(self.vScrollbar)
				self.vScrollbar.Resize(SCROLLBAR_WIDTH, self.Height()-SCROLLBAR_WIDTH)
				self.vScrollbar.Move(self.Width()-SCROLLBAR_WIDTH, 0)
				self.AddChild(self.corner)
				self.corner.Move(self.Width()-SCROLLBAR_WIDTH, self.Height()-SCROLLBAR_WIDTH)
			} else if self.subwidget.Width() > self.Width() && self.showH == true {
				self.hScrollbar.SetMaximum(self.subwidget.Width() - self.Width())
				self.AddChild(self.hScrollbar)
				self.hScrollbar.Move(0, self.Height()-SCROLLBAR_WIDTH)
				self.hScrollbar.Resize(self.Width(), SCROLLBAR_WIDTH)
			} else if self.subwidget.Height() > self.Height() && self.showV == true {
				self.vScrollbar.SetMaximum(self.subwidget.Height() - self.Height())
				self.AddChild(self.vScrollbar)
				self.vScrollbar.Resize(SCROLLBAR_WIDTH, self.Height())
				self.vScrollbar.Move(self.Width()-SCROLLBAR_WIDTH, 0)
			}
		} else {
			if self.subwidget.Width() > self.Width()-SCROLLBAR_WIDTH-4 && self.subwidget.Height() > self.Height()-SCROLLBAR_WIDTH-4 && self.showH == true && self.showV == true {
				self.hScrollbar.SetMaximum(self.subwidget.Width() - (self.Width() - SCROLLBAR_WIDTH - 4))
				self.AddChild(self.hScrollbar)
				self.hScrollbar.Resize(self.Width()-SCROLLBAR_WIDTH-4, SCROLLBAR_WIDTH)
				self.hScrollbar.Move(2, self.Height()-SCROLLBAR_WIDTH-2)
				self.vScrollbar.SetMaximum(self.subwidget.Height() - (self.Height() - SCROLLBAR_WIDTH - 4))
				self.AddChild(self.vScrollbar)
				self.vScrollbar.Resize(SCROLLBAR_WIDTH, self.Height()-SCROLLBAR_WIDTH-4)
				self.vScrollbar.Move(self.Width()-SCROLLBAR_WIDTH-2, 2)
				self.AddChild(self.corner)
				self.corner.Move(self.Width()-SCROLLBAR_WIDTH-2, self.Height()-SCROLLBAR_WIDTH-2)
			} else if self.subwidget.Width() > self.Width()-4 && self.showH == true {
				self.hScrollbar.SetMaximum(self.subwidget.Width() - self.Width() - 4)
				self.AddChild(self.hScrollbar)
				self.hScrollbar.Move(2, self.Height()-SCROLLBAR_WIDTH-2)
				self.hScrollbar.Resize(self.Width()-4, SCROLLBAR_WIDTH)
			} else if self.subwidget.Height() > self.Height()-4 && self.showV == true {
				self.vScrollbar.SetMaximum(self.subwidget.Height() - self.Height() - 4)
				self.AddChild(self.vScrollbar)
				self.vScrollbar.Resize(SCROLLBAR_WIDTH, self.Height()-4)
				self.vScrollbar.Move(self.Width()-SCROLLBAR_WIDTH-2, 2)
			}
		}

	}
	self.CoreWidget.PostUpdate()
}

func (self *ScrollWidget) Repaint() {
	if self.showBezel {
		self.FillRect(0, 0, self.width, self.height, 0xffffffff)

		self.SetDrawColor(0x88, 0x88, 0x88, 0xff)
		self.DrawLine(0, 0, 0, self.Height()-1)
		self.DrawLine(0, 0, self.Width()-1, 0)
		//	self.SetDrawColor(0xff, 0xff, 0xff, 0xff)
		//	self.DrawLine(self.Width()-1, self.Height()-1, self.Width()-1, 0)
		//	self.DrawLine(self.Width()-1, 0, 0, 0)
		self.DrawLine(1, 1, 1, self.Height()-2)
		self.DrawLine(1, self.Height()-2, self.Width()-2, self.Height()-2)
		self.DrawLine(self.Width()-2, self.Height()-2, self.Width()-2, 1)
		self.DrawLine(self.Width()-2, 1, 1, 1)

		for _, child := range self.children {
			// adjust the clipping to the current child
			if child.IsDirty() {
				child.Repaint()
			}
			widthDst := child.Width()
			heightDst := child.Height()
			if child.X()+widthDst > self.Width()-4 {
				widthDst = self.Width() - 4 - child.X()
			}
			if widthDst < 0 {
				widthDst = 0
			}
			if child.Y()+heightDst > self.Height()-4 {
				heightDst = self.Height() - 4 - child.Y()
			}
			if heightDst < 0 {
				heightDst = 0
			}
			rectSrc := sdl.Rect{0, 0, widthDst, heightDst}
			rectDst := sdl.Rect{child.X() + 2, child.Y() + 2, widthDst, heightDst}
			child.Surface().Blit(&rectSrc, self.Surface(), &rectDst)
		}
		self.dirty = false
	} else {
		self.CoreWidget.Repaint()
	}
}

func (self *ScrollWidget) SetInnerWidget(widget Widget) bool {

	if widget == nil {
		return false
	}
	if self.subwidget != nil {
		self.RemoveChild(self.subwidget)
	}
	self.subwidget = widget
	self.AddChild(widget)
	//widget.SetParent(self)
	self.Resize(self.Width(), self.Height()) // to refresh the scroll bars

	return true
}

func NewScrollWidget(w, h int32) *ScrollWidget {
	corewidget := NewCoreWidget(w, h)
	subwidget := NewCoreWidget(w, h)
	widget := &ScrollWidget{CoreWidget: *corewidget,
		subwidget:  subwidget,
		hScrollbar: NewScrollbarWidget(100, SCROLLBAR_WIDTH, true),
		vScrollbar: NewScrollbarWidget(SCROLLBAR_WIDTH, 100, false),
		corner:     NewCoreWidget(SCROLLBAR_WIDTH, SCROLLBAR_WIDTH),
		showH:      true,
		showV:      true}
	widget.AddChild(subwidget)
	subwidget.SetParent(widget)
	widget.hScrollbar.SetParent(widget)
	widget.hScrollbar.SetCallback(func(currentposition int32) {
		w := widget
		if w.subwidget != nil {
			w.xOffset = currentposition
			w.subwidget.Move(-w.xOffset, -w.yOffset)
		}
	})
	widget.vScrollbar.SetParent(widget)
	widget.vScrollbar.SetCallback(func(currentposition int32) {
		w := widget
		if w.subwidget != nil {
			w.yOffset = currentposition
			w.subwidget.Move(-w.xOffset, -w.yOffset)
		}
	})
	widget.corner.SetParent(widget)
	return widget
}
