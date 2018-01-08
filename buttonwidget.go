package sws

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type ButtonWidget struct {
	CoreWidget
	label        string
	image        *sdl.Surface
	imageleft    bool
	buttonState  bool
	cursorInside bool
	centered     bool
	textcolor    sdl.Color
	buttonColor  uint32
	clicked      func()
	hasfocus     bool
}

func (self *ButtonWidget) IsInputWidget() bool {
	return true
}

func (self *ButtonWidget) HasFocus(focus bool) {
	self.hasfocus = focus
	self.PostUpdate()
}

func (self *ButtonWidget) KeyDown(key sdl.Keycode, mod uint16) {
	if key == sdl.K_TAB {
		self.buttonState = false
		self.cursorInside = false
		self.PostUpdate()
		if self.focusOnNextInputWidgetCallback != nil {
			self.focusOnNextInputWidgetCallback(!(mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT))
		}
	}
	if key == sdl.K_SPACE {
		self.buttonState = true
		self.cursorInside = true
		self.PostUpdate()
	}
}

func (self *ButtonWidget) KeyUp(key sdl.Keycode, mod uint16) {
	if key == sdl.K_SPACE {
		self.buttonState = false
		if self.cursorInside == true {
			if self.clicked != nil {
				self.clicked()
			}
		}
		self.cursorInside = false
		self.PostUpdate()
	}
}

func (self *ButtonWidget) AlignImageLeft(alignleft bool) {
	self.imageleft = alignleft
}

func (self *ButtonWidget) SetButtonColor(color uint32) {
	self.buttonColor = color
	self.PostUpdate()
}

func (self *ButtonWidget) SetTextColor(color sdl.Color) {
	self.textcolor = color
	self.PostUpdate()
}

func (self *ButtonWidget) SetText(text string) {
	self.label = text
	self.PostUpdate()
}

func (self *ButtonWidget) SetCentered(centered bool) {
	self.centered = centered
	self.PostUpdate()
}

func (self *ButtonWidget) SetImage(image string) {
	if img, err := img.Load(image); err == nil {
		self.image = img
	}
	self.PostUpdate()
}

func (self *ButtonWidget) SetImageSurface(img *sdl.Surface) {
	self.image = img
	self.PostUpdate()
}

func (self *ButtonWidget) SetClicked(callback func()) {
	self.clicked = callback
}

func (self *ButtonWidget) MousePressDown(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.buttonState = true
		self.cursorInside = true
		self.PostUpdate()
	}
}

func (self *ButtonWidget) MousePressUp(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.buttonState = false
		if self.cursorInside == true {
			if self.clicked != nil {
				self.clicked()
			}
		}
		self.cursorInside = false
		self.PostUpdate()
	}
}

func (self *ButtonWidget) MouseMove(x, y, xrel, yrel int32) {
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

func (self *ButtonWidget) Repaint() {
	self.CoreWidget.Repaint()
	if self.buttonColor != 0 {
		self.FillRect(3, 3, self.width-6, self.height-6, self.buttonColor)
	}
	self.SetDrawColor(0, 0, 0, 255)
	self.DrawLine(1, 2, 1, self.Height()-3)
	self.DrawLine(self.Width()-2, 2, self.Width()-2, self.Height()-3)
	self.DrawLine(2, 1, self.Width()-3, 1)
	self.DrawLine(2, self.Height()-2, self.Width()-3, self.Height()-2)

	// text
	var text *sdl.Surface
	var err error
	if self.label != "" {
		if text, err = self.Font().RenderUTF8Blended(self.label, self.textcolor); err != nil {
			//	fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
		}
		defer text.Free()
	}

	if self.cursorInside == true {
		if text != nil && self.image == nil {
			wGap := self.Width() - text.W
			hGap := self.Height() - text.H
			if self.centered == false {
				wGap = 0
			}
			rectSrc := sdl.Rect{0, 0, text.W, text.H}
			rectDst := sdl.Rect{3 + (wGap / 2), 3 + (hGap / 2), self.Width() - 6 - (wGap / 2), self.Height() - 6 - (hGap / 2)}
			if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		} else if text == nil && self.image != nil {
			wGap := self.Width() - self.image.W
			hGap := self.Height() - self.image.H
			if self.centered == false {
				wGap = 0
			}
			rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
			rectDst := sdl.Rect{3 + (wGap / 2), 3 + (hGap / 2), self.Width() - 6 - (wGap / 2), self.Height() - 6 - (hGap / 2)}
			if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		} else if text != nil && self.image != nil {
			if self.imageleft == false {
				wTGap := self.Width() - text.W
				wIGap := self.Width() - self.image.W
				hGap := self.Height() - self.image.H - text.H
				if self.centered == false {
					wTGap = 0
					wIGap = 0
				}
				rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
				rectDst := sdl.Rect{3 + (wIGap / 2), 3 + (hGap / 2), self.Width() - 6 - (wIGap / 2), self.Height() - 6 - (hGap / 2)}
				if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
				}
				rectSrc = sdl.Rect{0, 0, text.W, text.H}
				rectDst = sdl.Rect{3 + (wTGap / 2), 3 + (hGap / 2) + self.image.H, self.Width() - 6 - (wTGap / 2), self.Height() - 6 - (hGap / 2) - self.image.H}
				if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
				}

			} else {
				hTGap := self.Height() - text.H
				hIGap := self.Height() - self.image.H
				wGap := self.Width() - self.image.W - text.W
				if self.centered == false {
					wGap = 0
				}
				rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
				rectDst := sdl.Rect{3 + (wGap / 2), 3 + (hIGap / 2), self.Width() - 6 - (wGap / 2), self.Height() - 6 - (hIGap / 2)}
				if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
				}
				rectSrc = sdl.Rect{0, 0, text.W, text.H}
				rectDst = sdl.Rect{3 + (wGap / 2) + self.image.W, 3 + (hTGap / 2), self.Width() - 6 - (wGap / 2) - self.image.W, self.Height() - 6 - (hTGap / 2)}
				if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
				}
			}
		}
		// dark
		self.SetDrawColor(50, 50, 50, 255)
		self.DrawLine(2, 2, 2, self.Height()-3)
		self.DrawLine(2, 2, self.Width()-3, 2)
		self.SetDrawColor(150, 150, 150, 255)
		self.DrawLine(3, 3, 3, self.Height()-4)
		self.DrawLine(3, 3, self.Width()-4, 3)
		//brigth
		self.SetDrawColor(255, 255, 255, 255)
		self.DrawLine(self.Width()-3, 2, self.Width()-3, self.Height()-3)
		self.DrawLine(2, self.Height()-3, self.Width()-3, self.Height()-3)
		self.SetDrawColor(240, 240, 240, 255)
		self.DrawLine(self.Width()-4, 3, self.Width()-4, self.Height()-4)
		self.DrawLine(3, self.Height()-4, self.Width()-4, self.Height()-4)
	} else {
		if text != nil && self.image == nil {
			wGap := self.Width() - text.W
			hGap := self.Height() - text.H
			if self.centered == false {
				wGap = 0
			}
			rectSrc := sdl.Rect{0, 0, text.W, text.H}
			rectDst := sdl.Rect{1 + (wGap / 2), 1 + (hGap / 2), self.Width() - 2 - (wGap / 2), self.Height() - 2 - (hGap / 2)}
			if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		} else if text == nil && self.image != nil {
			wGap := self.Width() - self.image.W
			hGap := self.Height() - self.image.H
			if self.centered == false {
				wGap = 0
			}
			rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
			rectDst := sdl.Rect{1 + (wGap / 2), 1 + (hGap / 2), self.Width() - 2 - (wGap / 2), self.Height() - 2 - (hGap / 2)}
			if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		} else if text != nil && self.image != nil {
			wTGap := self.Width() - text.W
			wIGap := self.Width() - self.image.W
			hGap := self.Height() - self.image.H - text.H
			if self.centered == false {
				wTGap = 0
				wIGap = 0
			}
			rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
			rectDst := sdl.Rect{1 + (wIGap / 2), 1 + (hGap / 2), self.Width() - 2 - (wIGap / 2), self.Height() - 2 - (hGap / 2)}
			if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
			rectSrc = sdl.Rect{0, 0, text.W, text.H}
			rectDst = sdl.Rect{1 + (wTGap / 2), 1 + (hGap / 2) + self.image.H, self.Width() - 2 - (wTGap / 2), self.Height() - 2 - (hGap / 2) - self.image.H}
			if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
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
	}
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

func NewButtonWidget(w, h int32, s string) *ButtonWidget {
	corewidget := NewCoreWidget(w, h)
	widget := &ButtonWidget{CoreWidget: *corewidget,
		label:        s,
		image:        nil,
		imageleft:    false,
		buttonState:  false,
		cursorInside: false,
		centered:     true,
		textcolor:    sdl.Color{0, 0, 0, 255},
	}
	widget.buttonColor = 0xffdddddd
	return widget
}
