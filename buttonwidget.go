package sws

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

type SWS_ButtonWidget struct {
	SWS_CoreWidget
	label        string
	buttonState  bool
	cursorInside bool
	clicked      func()
}

func (self *SWS_ButtonWidget) SetClicked(callback func()) {
	self.clicked = callback
}

func (self *SWS_ButtonWidget) MousePressDown(x, y int32, button uint8) {
	fmt.Println("Button.PressDown")
	if button == sdl.BUTTON_LEFT {
		self.buttonState = true
		self.cursorInside = true
		PostUpdate()
	}
}

func (self *SWS_ButtonWidget) MousePressUp(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.buttonState = false
		if self.cursorInside == true {
			if self.clicked != nil {
				self.clicked()
			}
		}
		self.cursorInside = false
		PostUpdate()
	}
}

func (self *SWS_ButtonWidget) MouseMove(x, y, xrel, yrel int32) {
	oldCursorInside := self.cursorInside
	if self.buttonState == true {
		if (x >= 0 && x < self.Width() && y >= 0 && y < self.Height()) {
			self.cursorInside = true
		} else {
			self.cursorInside = false
		}
		if (oldCursorInside != self.cursorInside) {
			PostUpdate()
		}
	}
}

func (self *SWS_ButtonWidget) Repaint() {
	self.SWS_CoreWidget.Repaint()
	self.SetDrawColor(0, 0, 0, 255)
	self.DrawLine(0, 1, 0, self.Height() - 2)
	self.DrawLine(self.Width() - 1, 1, self.Width() - 1, self.Height() - 2)
	self.DrawLine(1, 0, self.Width() - 2, 0)
	self.DrawLine(1, self.Height() - 1, self.Width() - 2, self.Height() - 1)
	if self.cursorInside == true {
		self.WriteTextCenter(2, 2, self.label, sdl.Color{0, 0, 0, 255})
		// dark
		self.SetDrawColor(50, 50, 50, 255)
		self.DrawLine(1, 1, 1, self.Height() - 2)
		self.DrawLine(1, 1, self.Width() - 2, 1)
		self.SetDrawColor(150, 150, 150, 255)
		self.DrawLine(2, 2, 2, self.Height() - 3)
		self.DrawLine(2, 2, self.Width() - 3, 2)
		//brigth
		self.SetDrawColor(255, 255, 255, 255)
		self.DrawLine(self.Width() - 2, 1, self.Width() - 2, self.Height() - 2)
		self.DrawLine(1, self.Height() - 2, self.Width() - 2, self.Height() - 2)
		self.SetDrawColor(240, 240, 240, 255)
		self.DrawLine(self.Width() - 3, 2, self.Width() - 3, self.Height() - 3)
		self.DrawLine(2, self.Height() - 3, self.Width() - 3, self.Height() - 3)
	} else {
		self.WriteTextCenter(0, 0, self.label, sdl.Color{0, 0, 0, 255})
		// bright
		self.SetDrawColor(255, 255, 255, 255)
		self.DrawLine(1, 1, 1, self.Height() - 2)
		self.DrawLine(1, 1, self.Width() - 2, 1)
		self.SetDrawColor(240, 240, 240, 255)
		self.DrawLine(2, 2, 2, self.Height() - 3)
		self.DrawLine(2, 2, self.Width() - 3, 2)
		//dark
		self.SetDrawColor(50, 50, 50, 255)
		self.DrawLine(self.Width() - 2, 1, self.Width() - 2, self.Height() - 2)
		self.DrawLine(1, self.Height() - 2, self.Width() - 2, self.Height() - 2)
		self.SetDrawColor(150, 150, 150, 255)
		self.DrawLine(self.Width() - 3, 2, self.Width() - 3, self.Height() - 3)
		self.DrawLine(2, self.Height() - 3, self.Width() - 3, self.Height() - 3)
	}
}

func CreateButtonWidget(w, h int32, s string) *SWS_ButtonWidget {
	corewidget := CreateCoreWidget(w, h)
	widget := &SWS_ButtonWidget{SWS_CoreWidget: *corewidget,
		label: s,
		buttonState: false,
		cursorInside: false}
	return widget
}



