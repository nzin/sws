package sws

import (
//	"github.com/veandco/go-sdl2/sdl"
)

//
// a simple vertical box container
//
type SWS_VBoxWidget struct {
	SWS_CoreWidget
}


func (self *SWS_VBoxWidget) Resize(width, height int32) {
	if width < 20 {
		width = 20
	}
	if height < 20 {
		height = 20
	}
	self.SWS_CoreWidget.Resize(width, height)

	for _,child := range self.children {
		child.Resize(width,child.Height())
	}
	PostUpdate()
}


func (self *SWS_VBoxWidget) AddChild(child SWS_Widget) {
	self.SWS_CoreWidget.AddChild(child)
	var width,height int32
	for _,child := range self.children {
		child.Move(0,height)
		if width<child.Width() {
			width=child.Width()
		}
		height+=child.Height()
	}
	self.height=height
	self.width=width
	self.SWS_CoreWidget.Resize(width, height)
	PostUpdate()
}

func (self *SWS_VBoxWidget) RemoveChild(child SWS_Widget) {
	self.SWS_CoreWidget.RemoveChild(child)
	var width,height int32
	for _,child := range self.children {
		child.Move(0,height)
		if width<child.Width() {
			width=child.Width()
		}
		height+=child.Height()
	}
	self.height=height
	self.width=width
	PostUpdate()
}

func CreateVBoxWidget(w, h int32) *SWS_VBoxWidget {
	corewidget := CreateCoreWidget(w, h)
	widget := &SWS_VBoxWidget{SWS_CoreWidget: *corewidget,
	}
	return widget
}
