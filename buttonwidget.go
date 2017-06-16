package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SWS_ButtonWidget struct {
	SWS_CoreWidget
	label        string
	image        *sdl.Surface
	buttonState  bool
	cursorInside bool
	centered     bool
	textcolor    sdl.Color
	clicked      func()
}

func (self *SWS_ButtonWidget) SetTextColor(color sdl.Color) {
	self.textcolor=color
}

func (self *SWS_ButtonWidget) SetCentered(centered bool) {
	self.centered=centered
}

func (self *SWS_ButtonWidget) SetImage(image *sdl.Surface) {
	self.image=image
}

func (self *SWS_ButtonWidget) SetClicked(callback func()) {
	self.clicked = callback
}

func (self *SWS_ButtonWidget) MousePressDown(x, y int32, button uint8) {
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

func (self *SWS_ButtonWidget) Repaint() {
	self.SWS_CoreWidget.Repaint()
	self.SetDrawColor(0, 0, 0, 255)
	self.DrawLine(0, 1, 0, self.Height()-2)
	self.DrawLine(self.Width()-1, 1, self.Width()-1, self.Height()-2)
	self.DrawLine(1, 0, self.Width()-2, 0)
	self.DrawLine(1, self.Height()-1, self.Width()-2, self.Height()-1)
	
	// text
	var text *sdl.Surface
	var err error
	if (self.label!="") {
		if text, err = self.Font().RenderUTF8_Blended(self.label, self.textcolor); err != nil {
		//	fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
		}
        	defer text.Free()
        }

	if self.cursorInside == true {
		if (text!=nil && self.image == nil) {
			wGap := self.Width() - text.W
			hGap := self.Height() - text.H
			if self.centered==false {
				wGap=0
			}
			rectSrc := sdl.Rect{0, 0, text.W, text.H}
			rectDst := sdl.Rect{2+(wGap/2), 2+(hGap/2), self.Width()-2-(wGap/2), self.Height()-2-(hGap/2)}
			if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		} else if (text==nil && self.image !=nil) {
			wGap := self.Width() - self.image.W
			hGap := self.Height() - self.image.H
			if self.centered==false {
				wGap=0
			}
			rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
			rectDst := sdl.Rect{2+(wGap/2), 2+(hGap/2), self.Width()-2-(wGap/2), self.Height()-2-(hGap/2)}
			if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		} else if (text!=nil && self.image !=nil) {
			wTGap := self.Width() - text.W
			wIGap := self.Width() - self.image.W
			hGap := self.Height() - self.image.H - text.H
			if self.centered==false {
				wTGap=0
				wIGap=0
			}
			rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
			rectDst := sdl.Rect{2+(wIGap/2), 2+(hGap/2), self.Width()-2-(wIGap/2), self.Height()-2-(hGap/2)}
			if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
			rectSrc = sdl.Rect{0, 0, text.W, text.H}
			rectDst = sdl.Rect{2+(wTGap/2), 2+(hGap/2)+self.image.H, self.Width()-2-(wTGap/2), self.Height()-2-(hGap/2)-self.image.H}
			if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		}
		// dark
		self.SetDrawColor(50, 50, 50, 255)
		self.DrawLine(1, 1, 1, self.Height()-2)
		self.DrawLine(1, 1, self.Width()-2, 1)
		self.SetDrawColor(150, 150, 150, 255)
		self.DrawLine(2, 2, 2, self.Height()-3)
		self.DrawLine(2, 2, self.Width()-3, 2)
		//brigth
		self.SetDrawColor(255, 255, 255, 255)
		self.DrawLine(self.Width()-2, 1, self.Width()-2, self.Height()-2)
		self.DrawLine(1, self.Height()-2, self.Width()-2, self.Height()-2)
		self.SetDrawColor(240, 240, 240, 255)
		self.DrawLine(self.Width()-3, 2, self.Width()-3, self.Height()-3)
		self.DrawLine(2, self.Height()-3, self.Width()-3, self.Height()-3)
	} else {
		if (text!=nil && self.image == nil) {
			wGap := self.Width() - text.W
			hGap := self.Height() - text.H
			if self.centered==false {
				wGap=0
			}
			rectSrc := sdl.Rect{0, 0, text.W, text.H}
			rectDst := sdl.Rect{(wGap/2), (hGap/2), self.Width()-(wGap/2), self.Height()-(hGap/2)}
			if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		} else if (text==nil && self.image !=nil) {
			wGap := self.Width() - self.image.W
			hGap := self.Height() - self.image.H
			if self.centered==false {
				wGap=0
			}
			rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
			rectDst := sdl.Rect{(wGap/2), (hGap/2), self.Width()-(wGap/2), self.Height()-(hGap/2)}
			if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		} else if (text!=nil && self.image !=nil) {
			wTGap := self.Width() - text.W
			wIGap := self.Width() - self.image.W
			hGap := self.Height() - self.image.H - text.H
			if self.centered==false {
				wTGap=0
				wIGap=0
			}
			rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
			rectDst := sdl.Rect{(wIGap/2), (hGap/2), self.Width()-(wIGap/2), self.Height()-(hGap/2)}
			if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
			rectSrc = sdl.Rect{0, 0, text.W, text.H}
			rectDst = sdl.Rect{(wTGap/2), (hGap/2)+self.image.H, self.Width()-(wTGap/2), self.Height()-(hGap/2)-self.image.H}
			if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		}
		// bright
		self.SetDrawColor(255, 255, 255, 255)
		self.DrawLine(1, 1, 1, self.Height()-2)
		self.DrawLine(1, 1, self.Width()-2, 1)
		self.SetDrawColor(240, 240, 240, 255)
		self.DrawLine(2, 2, 2, self.Height()-3)
		self.DrawLine(2, 2, self.Width()-3, 2)
		//dark
		self.SetDrawColor(50, 50, 50, 255)
		self.DrawLine(self.Width()-2, 1, self.Width()-2, self.Height()-2)
		self.DrawLine(1, self.Height()-2, self.Width()-2, self.Height()-2)
		self.SetDrawColor(150, 150, 150, 255)
		self.DrawLine(self.Width()-3, 2, self.Width()-3, self.Height()-3)
		self.DrawLine(2, self.Height()-3, self.Width()-3, self.Height()-3)
	}
}

func CreateButtonWidget(w, h int32, s string) *SWS_ButtonWidget {
	corewidget := CreateCoreWidget(w, h)
	widget := &SWS_ButtonWidget{SWS_CoreWidget: *corewidget,
		label:        s,
		image:        nil,
		buttonState:  false,
		cursorInside: false,
		centered:     true,
		textcolor:    sdl.Color{0, 0, 0, 255}}
	return widget
}
