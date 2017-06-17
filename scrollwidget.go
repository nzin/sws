package sws

import (
//	"github.com/veandco/go-sdl2/sdl"
)

const (
    SCROLLBAR_WIDTH = 17
)

type SWS_ScrollWidget struct {
	SWS_CoreWidget
	subwidget  SWS_Widget
	hScrollbar *SWS_ScrollbarWidget
	vScrollbar *SWS_ScrollbarWidget
	corner     *SWS_CoreWidget
	xOffset    int32
	yOffset    int32
	showH      bool
	showV      bool
}

func (self *SWS_ScrollWidget) SetHorizontalPosition(position int32) {
	self.hScrollbar.SetPosition(position)
}

func (self *SWS_ScrollWidget) SetVerticalPosition(position int32) {
	self.vScrollbar.SetPosition(position)
}

func (self *SWS_ScrollWidget) ShowVerticalScrollbar(showV bool) {
	self.showV=showV
}

func (self *SWS_ScrollWidget) ShowHorizontalScrollbar(showH bool) {
	self.showH=showH
}

func (self *SWS_ScrollWidget) Resize(width, height int32) {
	if width < 45 {
		width = 45
	}
	if height < 45 {
		height = 45
	}
	self.SWS_CoreWidget.Resize(width, height)

	self.RemoveChild(self.hScrollbar)
	self.RemoveChild(self.vScrollbar)
	self.RemoveChild(self.corner)
	self.hScrollbar.SetMaximum(self.subwidget.Width() - self.Width())
	self.vScrollbar.SetMaximum(self.subwidget.Height() - self.Height())
	if self.subwidget != nil {
		if self.subwidget.Width() > self.Width()-SCROLLBAR_WIDTH && self.subwidget.Height() > self.Height()-SCROLLBAR_WIDTH && self.showH==true && self.showV==true {
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
		} else if self.subwidget.Width() > self.Width()-SCROLLBAR_WIDTH && self.showH==true {
			self.hScrollbar.SetMaximum(self.subwidget.Width() - (self.Width() - SCROLLBAR_WIDTH))
			self.AddChild(self.hScrollbar)
			self.hScrollbar.Move(0, self.Height()-SCROLLBAR_WIDTH)
			self.hScrollbar.Resize(self.Width(), SCROLLBAR_WIDTH)
		} else if self.subwidget.Height() > self.Height()-SCROLLBAR_WIDTH && self.showV==true {
			self.vScrollbar.SetMaximum(self.subwidget.Height() - (self.Height() - SCROLLBAR_WIDTH))
			self.AddChild(self.vScrollbar)
			self.vScrollbar.Resize(SCROLLBAR_WIDTH, self.Height())
			self.vScrollbar.Move(self.Width()-SCROLLBAR_WIDTH, 0)
		}
	}
	PostUpdate()
}

func (self *SWS_ScrollWidget) SetInnerWidget(widget SWS_Widget) bool {

	if widget == nil {
		return false
	}
	if self.subwidget != nil {
		self.RemoveChild(self.subwidget)
	}
	self.subwidget = widget
	self.AddChild(widget)
	widget.SetParent(self)
	self.Resize(self.Width(), self.Height()) // to refresh the scroll bars

	return true
}

func CreateScrollWidget(w, h int32) *SWS_ScrollWidget {
	corewidget := CreateCoreWidget(w, h)
	subwidget := CreateCoreWidget(w, h)
	widget := &SWS_ScrollWidget{SWS_CoreWidget: *corewidget,
		subwidget:  subwidget,
		hScrollbar: CreateScrollbarWidget(100, SCROLLBAR_WIDTH, true),
		vScrollbar: CreateScrollbarWidget(SCROLLBAR_WIDTH, 100, false),
		corner:     CreateCoreWidget(SCROLLBAR_WIDTH, SCROLLBAR_WIDTH),
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
