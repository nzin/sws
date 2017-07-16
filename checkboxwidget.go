package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SWS_CheckboxWidget struct {
	SWS_CoreWidget
	buttonState  bool
	cursorInside bool
	Selected     bool
	clicked      func()
}

func (self *SWS_CheckboxWidget) SetClicked(callback func()) {
	self.clicked = callback
}

func (self *SWS_CheckboxWidget) MousePressDown(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.buttonState = true
		self.cursorInside = true
		PostUpdate()
	}
}

func (self *SWS_CheckboxWidget) MousePressUp(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.buttonState = false
		if self.cursorInside == true {
			self.Selected=!self.Selected
			if self.clicked != nil {
				self.clicked()
			}
		}
		self.cursorInside = false
		PostUpdate()
	}
}

func (self *SWS_CheckboxWidget) MouseMove(x, y, xrel, yrel int32) {
	oldCursorInside := self.cursorInside
	if self.buttonState == true {
		if x >= 0 && x < self.Width() && y >= 0 && y < self.Height() {
			self.cursorInside = true
		} else {
			self.cursorInside = false
		}
		if oldCursorInside != self.cursorInside {
			PostUpdate()
		}
	}
}

func (self *SWS_CheckboxWidget) Repaint() {
	self.SWS_CoreWidget.Repaint()
	
	selected:=self.Selected
	if self.cursorInside {
		selected=!selected
	}
	
	self.SetDrawColor(0, 0, 0, 255)
	self.DrawLine(5,5,5,19)
	self.DrawLine(5,19,19,19)
	self.DrawLine(19,19,19,5)
	self.DrawLine(19,5,5,5)
	
	if selected {
		self.DrawLine(8,11,10,13)
		self.DrawLine(8,12,10,14)
		self.DrawLine(8,13,10,15)
		self.DrawLine(10,13,20,3)
		self.DrawLine(10,14,20,4)
		self.DrawLine(10,15,20,5)
	}
	
}

func CreateCheckboxWidget() *SWS_CheckboxWidget {
	corewidget := CreateCoreWidget(25, 25)
	widget := &SWS_CheckboxWidget{SWS_CoreWidget: *corewidget,
		buttonState:  false,
		cursorInside: false,
		Selected:     false,
	}
	return widget
}
