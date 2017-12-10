package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

type DropdownWidget struct {
	CoreWidget
	Choices      []string
	ActiveChoice int32
	buttonState  bool
	cursorInside bool
	clicked      func()
	menu         *MenuWidget
	hasfocus     bool
}

func (self *DropdownWidget) IsInputWidget() bool {
	return true
}

func (self *DropdownWidget) KeyDown(key sdl.Keycode, mod uint16) {
	if key == sdl.K_TAB {
		if self.focusOnNextInputWidgetCallback != nil {
			self.menu = nil
			hideMenu(nil)
			self.focusOnNextInputWidgetCallback(!(mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT))
		}
	}
	if key == sdl.K_UP || key == sdl.K_DOWN {
		if self.menu == nil {
			self.menu = NewMenuWidget()
			for i, choice := range self.Choices {
				index := i
				self.menu.AddItem(NewMenuItemLabel(choice, func() {
					self.ActiveChoice = int32(index)
					self.PostUpdate()
					if self.clicked != nil {
						self.clicked()
					}
				}))
			}
			var xx int32
			yy := self.height
			var widget Widget
			widget = self
			for widget != nil {
				xx += widget.X()
				yy += widget.Y()
				widget = widget.Parent()
			}
			self.menu.Move(xx, yy-2)
			ShowMenu(self.menu)
			self.PostUpdate()
		} else {
		}
	}
}

func (self *DropdownWidget) SetChoices(choices []string) {
	self.Choices = choices
	self.ActiveChoice = 0
	self.PostUpdate()
}

func (self *DropdownWidget) SetActiveChoice(choice int32) {
	if choice < int32(len(self.Choices)) {
		self.ActiveChoice = choice
		if self.clicked != nil {
			self.clicked()
		}
		self.PostUpdate()
	}
}

func (self *DropdownWidget) HasFocus(hasfocus bool) {
	self.hasfocus = hasfocus
	if hasfocus == false {
		menuInitiator = nil
	} else {
		hideMenu(nil)
		menuInitiator = self
	}

	self.PostUpdate()
}

func (self *DropdownWidget) SetClicked(callback func()) {
	self.clicked = callback
}

func (self *DropdownWidget) MousePressDown(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.buttonState = true
		self.cursorInside = true
		if self.menu != nil && self.menu.Parent() != nil {
			hideMenu(nil)
			self.menu = nil
		} else {
			self.menu = NewMenuWidget()
			for i, choice := range self.Choices {
				index := i
				self.menu.AddItem(NewMenuItemLabel(choice, func() {
					self.ActiveChoice = int32(index)
					self.PostUpdate()
					if self.clicked != nil {
						self.clicked()
					}
				}))
			}
			var xx int32
			yy := self.height
			var widget Widget
			widget = self
			for widget != nil {
				xx += widget.X()
				yy += widget.Y()
				widget = widget.Parent()
			}
			self.menu.Move(xx, yy-2)
			ShowMenu(self.menu)
			self.PostUpdate()
		}
	}
}

func (self *DropdownWidget) MousePressUp(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.buttonState = false
		if self.cursorInside == true {
			/*if self.clicked != nil {
			    self.clicked()
			}*/
			if self.menu != nil && self.menu.Parent() != nil {
				var xx int32
				yy := self.height
				var widget Widget
				widget = self
				for widget != nil {
					xx += widget.X()
					yy += widget.Y()
					widget = widget.Parent()
				}
				self.menu.Move(xx, yy-2)
				ShowMenu(self.menu)
			}
		}
		self.cursorInside = false
		self.PostUpdate()
	}
}

func (self *DropdownWidget) MouseMove(x, y, xrel, yrel int32) {
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

func (self *DropdownWidget) Repaint() {
	label := ""
	if self.ActiveChoice >= int32(len(self.Choices)) {
		self.ActiveChoice = int32(len(self.Choices) - 1)
	}
	if len(self.Choices) > 0 {
		label = self.Choices[self.ActiveChoice]
	}

	self.CoreWidget.Repaint()
	self.FillRect(3, 3, self.width-6, self.height-6, 0xffdddddd)
	self.SetDrawColor(0, 0, 0, 255)
	self.DrawLine(1, 2, 1, self.Height()-3)
	self.DrawLine(self.Width()-2, 2, self.Width()-2, self.Height()-3)
	self.DrawLine(2, 1, self.Width()-3, 1)
	self.DrawLine(2, self.Height()-2, self.Width()-3, self.Height()-2)
	self.WriteText(5, 1, label, sdl.Color{0, 0, 0, 255})
	self.FillRect(self.Width()-26, 3, 18, self.Height()-6, 0xffdddddd)
	self.SetDrawColor(0, 0, 0, 255)
	//self.DrawLine(self.Width()-25,4,self.Width()-25,self.Height()-6)
	for i := 0; i < 5; i++ {
		self.DrawLine(self.Width()-19+int32(i), 6+int32(i)*2, self.Width()-10-int32(i), 6+int32(i)*2)
		self.DrawLine(self.Width()-19+int32(i), 7+int32(i)*2, self.Width()-10-int32(i), 7+int32(i)*2)
	}
	// bright
	self.SetDrawColor(255, 255, 255, 255)
	self.DrawLine(2, 2, 2, self.Height()-3)
	self.DrawLine(2, 2, self.Width()-3, 2)
	self.SetDrawColor(240, 240, 240, 255)
	self.DrawLine(3, 3, 3, self.Height()-4)
	self.DrawLine(3, 3, self.Width()-4, 3)
	//dark
	self.SetDrawColor(50, 50, 50, 255)
	self.DrawLine(self.Width()-3, 2, self.Width()-3, self.Height()-3)
	self.DrawLine(2, self.Height()-3, self.Width()-3, self.Height()-3)
	self.SetDrawColor(150, 150, 150, 255)
	self.DrawLine(self.Width()-4, 3, self.Width()-4, self.Height()-4)
	self.DrawLine(3, self.Height()-4, self.Width()-4, self.Height()-4)

	if self.hasfocus {
		self.SetDrawColor(0x46, 0xc8, 0xe8, 255)
		self.DrawLine(0, 2, 0, self.Height()-3)
		self.DrawPoint(1, self.Height()-2)
		self.DrawLine(2, self.Height()-1, self.Width()-3, self.Height()-1)
		self.DrawPoint(self.Width()-2, self.Height()-2)
		self.DrawLine(self.Width()-1, self.Height()-3, self.Width()-1, 2)
		self.DrawPoint(self.Width()-2, 1)
		self.DrawLine(self.Width()-3, 0, 2, 0)
		self.DrawPoint(1, 1)
	}
}

func NewDropdownWidget(w, h int32, choices []string) *DropdownWidget {
	corewidget := NewCoreWidget(w, h)
	widget := &DropdownWidget{CoreWidget: *corewidget,
		buttonState:  false,
		cursorInside: false,
		Choices:      choices,
		ActiveChoice: 0,
	}
	return widget
}
