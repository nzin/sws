package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

type CheckboxWidget struct {
	CoreWidget
	buttonState  bool
	cursorInside bool
	Selected     bool
	disabled     bool
	clicked      func()
	hasfocus     bool
}

func (self *CheckboxWidget) IsInputWidget() bool {
	if self.disabled == true {
		return false
	}
	return true
}

func (self *CheckboxWidget) HasFocus(focus bool) {
	self.hasfocus = focus
	self.PostUpdate()
}

func (self *CheckboxWidget) KeyDown(key sdl.Keycode, mod uint16) {
	if key == sdl.K_TAB {
		if self.focusOnNextInputWidgetCallback != nil {
			self.focusOnNextInputWidgetCallback(!(mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT))
		}
	}
	if key == sdl.K_SPACE && self.disabled == false {
		self.Selected = !self.Selected
		if self.clicked != nil {
			self.clicked()
		}
		self.PostUpdate()
	}
}

func (self *CheckboxWidget) SetDisabled(disabled bool) {
	self.disabled = disabled
	self.PostUpdate()
}

func (self *CheckboxWidget) SetSelected(selected bool) {
	self.Selected = selected
	self.PostUpdate()
	if self.valueChangedCallback != nil {
		self.valueChangedCallback()
	}
}

func (self *CheckboxWidget) SetClicked(callback func()) {
	self.clicked = callback
}

func (self *CheckboxWidget) MousePressDown(x, y int32, button uint8) {
	if self.disabled {
		return
	}
	if button == sdl.BUTTON_LEFT {
		self.buttonState = true
		self.cursorInside = true
		self.PostUpdate()
	}
}

func (self *CheckboxWidget) MousePressUp(x, y int32, button uint8) {
	if self.disabled {
		return
	}
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
		if self.valueChangedCallback != nil {
			self.valueChangedCallback()
		}
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

	bgcheckbox := uint32(0xffffffff)
	if self.disabled {
		bgcheckbox = self.bgColor
	}
	self.FillRect(6, 6, 12, 12, bgcheckbox)

	self.SetDrawColor(0, 0, 0, 255)
	self.DrawLine(5, 5, 5, 19)
	self.DrawLine(5, 19, 19, 19)
	self.DrawLine(19, 19, 19, 5)
	self.DrawLine(19, 5, 5, 5)

	if self.hasfocus && self.disabled == false {
		self.SetDrawColor(0x46, 0xc8, 0xe8, 255)
		self.DrawLine(4, 5, 4, 19)
		self.DrawLine(5, 20, 19, 20)
		self.DrawLine(20, 19, 20, 5)
		self.DrawLine(19, 4, 5, 4)
	}

	if selected {
		self.SetDrawColor(0, 0, 0, 255)
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
