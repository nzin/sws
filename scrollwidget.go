package sws

import (
//	"github.com/veandco/go-sdl2/sdl"
)

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
}

func (self *ScrollWidget) SetHorizontalPosition(position int32) {
	self.hScrollbar.SetPosition(position)
}

func (self *ScrollWidget) SetVerticalPosition(position int32) {
	self.vScrollbar.SetPosition(position)
}

func (self *ScrollWidget) ShowVerticalScrollbar(showV bool) {
	self.showV = showV
}

func (self *ScrollWidget) ShowHorizontalScrollbar(showH bool) {
	self.showH = showH
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
	}
	PostUpdate()
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
