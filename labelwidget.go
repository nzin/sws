package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SWS_Label struct {
	SWS_CoreWidget
	label string
}

func (self *SWS_Label) Repaint() {
	self.SWS_CoreWidget.Repaint()
	self.WriteText(0, 0, self.label, sdl.Color{0, 0, 0, 255})
}

func CreateLabel(w, h int32, s string) *SWS_Label {
	corewidget := CreateCoreWidget(w, h)
	widget := &SWS_Label{SWS_CoreWidget: *corewidget,
		label: s}
	return widget
}



