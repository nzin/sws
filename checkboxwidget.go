package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

type CheckboxWidget struct {
	CoreWidget
	buttonState  bool
	cursorInside bool
	Selected     bool
	clicked      func()
}

func (self *CheckboxWidget) SetSelected(selected bool) {
	self.Selected = selected
	self.PostUpdate()
}

func (self *CheckboxWidget) SetClicked(callback func()) {
	self.clicked = callback
}

func (self *CheckboxWidget) MousePressDown(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.buttonState = true
		self.cursorInside = true
		self.PostUpdate()
	}
}

func (self *CheckboxWidget) MousePressUp(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.buttonState = false
		if self.cursorInside == true {
			self.Selected = !self.Selected
			if self.clicked != nil {
				self.clicked()
			}
		}
		self.cursorInside = false
		self.PostUpdate()
	}
}

func (self *CheckboxWidget) MouseMove(x, y, xrel, yrel int32) {
	oldCursorInside := self.cursorInside
	if self.buttonState == true {
		if x >= 0 && x < self.Width() && y >= 0 && y < self.Height() {
			self.cursorInside = true
		} else {
			self.cursorInside = false
		}
		if oldCursorInside != self.cursorInside {
			self.PostUpdate()
		}
	}
}

func (self *CheckboxWidget) Repaint() {
	self.CoreWidget.Repaint()

	selected := self.Selected
	if self.cursorInside {
		selected = !selected
	}
	
	self.FillRect(6, 6, 12, 12, 0xffffffff)

	self.SetDrawColor(0, 0, 0, 255)
	self.DrawLine(5, 5, 5, 19)
	self.DrawLine(5, 19, 19, 19)
	self.DrawLine(19, 19, 19, 5)
	self.DrawLine(19, 5, 5, 5)

	if selected {
		self.DrawLine(8, 11, 10, 13)
		self.DrawLine(8, 12, 10, 14)
		self.DrawLine(8, 13, 10, 15)
		self.DrawLine(10, 13, 20, 3)
		self.DrawLine(10, 14, 20, 4)
		self.DrawLine(10, 15, 20, 5)
	}

}

func NewCheckboxWidget() *CheckboxWidget {
	corewidget := NewCoreWidget(25, 25)
	widget := &CheckboxWidget{CoreWidget: *corewidget,
		buttonState:  false,
		cursorInside: false,
		Selected:     false,
	}
	return widget
}
