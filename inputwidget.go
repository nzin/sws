package sws

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type SWS_InputWidget struct {
	SWS_CoreWidget
	text                  string
	initialCursorPosition int
	endCursorPosition     int
	hasfocus              bool
	leftButtonDown        bool
	writeOffset           int32
}

func (self *SWS_InputWidget) HasFocus(focus bool) {
	fmt.Println("HasFocus")
	self.hasfocus = focus
	PostUpdate()
}

func (self *SWS_InputWidget) MousePressDown(x, y int32, button uint8) {

	if button == sdl.BUTTON_LEFT {
		self.leftButtonDown = true
		self.initialCursorPosition = 0
		for i := 1; i <= len(self.text); i++ {
			w, _, err := self.Font().SizeUTF8(self.text[:i])
			if err != nil {
				panic(err)
			}
			if w > int(x-2+self.writeOffset) {
				break
			}
			self.initialCursorPosition++
		}
		self.endCursorPosition = self.initialCursorPosition
		PostUpdate()
	}
}

func (self *SWS_InputWidget) MousePressUp(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.leftButtonDown = false
	}
}

func (self *SWS_InputWidget) MouseMove(x, y, xrel, yrel int32) {
	if self.leftButtonDown == true {
		if (self.writeOffset>0 && x<0) {
			self.writeOffset+=x/2
			if self.writeOffset<0 {
				self.writeOffset=0
			}
		}
		w, _, err := self.Font().SizeUTF8(self.text)
		if err != nil {
			panic(err)
		}
		if (self.writeOffset<int32(w)-self.Width()+4 && x>self.Width()) {
			self.writeOffset+=(x-self.Width())/2
			if self.writeOffset>int32(w)-self.Width()+4 {
				self.writeOffset=int32(w)-self.Width()+4
			}
		}
		
		self.initialCursorPosition = 0
		for i := 1; i <= len(self.text); i++ {
			w, _, err := self.Font().SizeUTF8(self.text[:i])
			if err != nil {
				panic(err)
			}
			if w > int(x-2+self.writeOffset) {
				break
			}
			self.initialCursorPosition++
		}
		PostUpdate()
	}
}

func (self *SWS_InputWidget) KeyDown(key sdl.Keycode, mod uint16) {
	if key == sdl.K_UP {
		if mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT {
			if self.initialCursorPosition > 0 {
				self.initialCursorPosition=0
			}
		} else {
			if self.initialCursorPosition > 0 {
				self.initialCursorPosition=0
			}
			self.endCursorPosition = self.initialCursorPosition
		}
		PostUpdate()
	}
	if key == sdl.K_LEFT {
		if mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT {
			if self.initialCursorPosition > 0 {
				self.initialCursorPosition--
			}
		} else {
			if self.initialCursorPosition > 0 {
				self.initialCursorPosition--
			}
			self.endCursorPosition = self.initialCursorPosition
		}
		PostUpdate()
	}
	if key == sdl.K_RIGHT {
		if mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT {
			if self.initialCursorPosition < len(self.text) {
				self.initialCursorPosition++
			}
		} else {
			if self.initialCursorPosition < len(self.text) {
				self.initialCursorPosition++
			}
			self.endCursorPosition = self.initialCursorPosition
		}
		PostUpdate()
	}

	if key == sdl.K_DOWN {
		if mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT {
			if self.initialCursorPosition < len(self.text) {
				self.initialCursorPosition=len(self.text)
			}
		} else {
			if self.initialCursorPosition < len(self.text) {
				self.initialCursorPosition=len(self.text)
			}
			self.endCursorPosition = self.initialCursorPosition
		}
		PostUpdate()
	}

	if key == sdl.K_BACKSPACE {
		if self.initialCursorPosition == self.endCursorPosition {
			if self.initialCursorPosition > 0 {
				self.initialCursorPosition--
				self.text = self.text[:self.initialCursorPosition] + self.text[self.initialCursorPosition+1:]
			}
		} else {
			i, e := self.initialCursorPosition, self.endCursorPosition
			if i > e {
				i, e = e, i
			}
			self.text = self.text[:i] + self.text[e:]
			self.initialCursorPosition = i
		}
		self.endCursorPosition = self.initialCursorPosition
		PostUpdate()
	}

	if mod == sdl.KMOD_LCTRL || mod == sdl.KMOD_RCTRL {
		if key=='a' {
			self.endCursorPosition = 0
			self.initialCursorPosition = len(self.text)
			PostUpdate()
		}
	} else if (key >= 'a' && key <= 'z') || (key >= '0' && key <= '9') || key == ' ' {
		if key >= 'a' && key <= 'z' {
			if mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT {
				key += 'A' - 'a'
			}
		}
		if self.initialCursorPosition == self.endCursorPosition {
			self.text = self.text[:self.initialCursorPosition] + string(key) + self.text[self.initialCursorPosition:]
			self.initialCursorPosition++
		} else {
			i, e := self.initialCursorPosition, self.endCursorPosition
			if i > e {
				i, e = e, i
			}
			self.text = self.text[:i] + string(key) + self.text[e:]
			self.initialCursorPosition = i + 1
		}
		self.endCursorPosition = self.initialCursorPosition
		PostUpdate()
	}
	// replace cursor
	w, _, err := self.Font().SizeUTF8(self.text[:self.initialCursorPosition])
	if err != nil {
		panic(err)
	}
	if self.writeOffset>int32(w) {
		self.writeOffset=int32(w)
		PostUpdate()
	}
	if self.writeOffset+self.Width()-4<int32(w) {
		self.writeOffset=int32(w)-self.Width()+4
		PostUpdate()
	}
}

func (self *SWS_InputWidget) Repaint() {
	self.SWS_CoreWidget.Repaint()
	// write text and cursor
	i := self.initialCursorPosition
	e := self.endCursorPosition
	if i > e {
		i, e = e, i
	}

	self.SetDrawColor(0, 0, 0, 255)
	if e > len(self.text) {
		e = len(self.text)
	}
	if i > len(self.text) {
		i = len(self.text)
	}
	strbefore := self.text[:i]
	strMiddle := self.text[i:e]
	strafter := self.text[e:]
	wMiddle, _, _ := self.Font().SizeUTF8(strMiddle)

	wbefore, _ := self.WriteText(2-self.writeOffset, 2, strbefore, sdl.Color{0, 0, 0, 255})
	//    fmt.Println(wbefore,wMiddle)
	self.FillRect(wbefore+2-self.writeOffset, 3, int32(wMiddle), self.Height()-2, 0xff8888ff)
	self.SetDrawColor(0, 0, 0, 255)
	self.WriteText(wbefore+2-self.writeOffset, 2, strMiddle, sdl.Color{0, 0, 0, 255})
	self.WriteText(wbefore+int32(wMiddle)+2-self.writeOffset, 2, strafter, sdl.Color{0, 0, 0, 255})
	if self.hasfocus {
		if self.initialCursorPosition < self.endCursorPosition {
			self.DrawLine(wbefore+2-self.writeOffset, 3, wbefore+2-self.writeOffset, self.Height()-4)
		} else {
			self.DrawLine(wbefore+int32(wMiddle)+2-self.writeOffset, 3, wbefore+int32(wMiddle)+2-self.writeOffset, self.Height()-4)
		}
	}

	// write boundaries
	self.SetDrawColor(0, 0, 0, 255)
	self.DrawLine(0, 2, 0, self.Height()-3)
	self.DrawLine(self.Width()-1, 2, self.Width()-1, self.Height()-3)
	self.DrawLine(2, 0, self.Width()-3, 0)
	self.DrawLine(2, self.Height()-1, self.Width()-3, self.Height()-1)
	self.DrawPoint(1, self.Height()-2)
	self.DrawPoint(1, 1)
	self.DrawPoint(self.Width()-2, 1)
	self.DrawPoint(self.Width()-2, self.Height()-2)

	self.SetDrawColor(0xdd, 0xdd, 0xdd, 255)
	self.DrawLine(2, 1, self.Width()-3, 1)
	self.DrawLine(1, 2, 1, self.Height()-3)
	self.DrawPoint(0, 0)
	self.DrawPoint(0, 1)
	self.DrawPoint(1, 0)

	self.DrawPoint(self.Width()-1, 0)
	self.DrawPoint(self.Width()-1, 1)
	self.DrawPoint(self.Width()-2, 0)

	self.DrawPoint(0, self.Height()-1)
	self.DrawPoint(0, self.Height()-2)
	self.DrawPoint(1, self.Height()-1)

	self.DrawPoint(self.Width()-1, self.Height()-1)
	self.DrawPoint(self.Width()-1, self.Height()-2)
	self.DrawPoint(self.Width()-2, self.Height()-1)


}

func CreateInputWidget(w, h int32, s string) *SWS_InputWidget {
	corewidget := CreateCoreWidget(w, h)
	widget := &SWS_InputWidget{SWS_CoreWidget: *corewidget,
		text: s,
		initialCursorPosition: 0,
		hasfocus:              false,
		leftButtonDown:        false,
		writeOffset:           0}
	widget.SetColor(0xffffffff)
	return widget
}
