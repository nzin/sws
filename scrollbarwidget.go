package sws

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
	"fmt"
)

type Scrollbarcallback func(currentposition int32)

type SWS_ScrollbarWidget struct {
	SWS_CoreWidget
	horizontal      bool
	minimum         int32
	maximum         int32
	callback        Scrollbarcallback
	buttondown      bool
	onelevator      bool
	initialpos      int32
	Currentposition int32 // in the targeted widget: so from minimum to maximum
	timerevent      *TimerEvent
}

func (self *SWS_ScrollbarWidget) SetCallback(callback Scrollbarcallback) {
	self.callback=callback
}

func (self *SWS_ScrollbarWidget) SetMinimum(m int32) {
	self.minimum=m
}

func (self *SWS_ScrollbarWidget) SetMaximum(m int32) {
	self.maximum=m
}

func (self *SWS_ScrollbarWidget) MousePressDown(x, y int32, button uint8) {
	fmt.Println("SWS_ScrollbarWidget::MousePressDown()")
	if button == sdl.BUTTON_LEFT {
		self.buttondown = true
		if self.horizontal {
			w := self.Width() * self.Width() / (self.maximum - self.minimum)
			if w < 20 {
				w = 20
			}
			offset := (self.Width() - w) * (self.Currentposition - self.minimum) / (self.maximum - self.minimum)
			if x < offset {
				// click before
				self.Currentposition -= (self.maximum - self.minimum) / 10
				if self.Currentposition < self.minimum {
					self.Currentposition = self.minimum
				}
				if (self.callback!=nil) { self.callback(self.Currentposition) }
				PostUpdate()
				self.timerevent = TimerAddEvent(time.Now().Add(500*time.Millisecond), 250*time.Millisecond, func() {
					self.Currentposition -= (self.maximum - self.minimum) / 10
					if self.Currentposition < self.minimum {
						self.Currentposition = self.minimum
					}
					if (self.callback!=nil) { self.callback(self.Currentposition) }
					PostUpdate()
				})
			} else if x > offset+w {
				// click after
				self.Currentposition += (self.maximum - self.minimum) / 10
				if self.Currentposition > self.maximum {
					self.Currentposition = self.maximum
				}
				if (self.callback!=nil) { self.callback(self.Currentposition) }
				PostUpdate()
				self.timerevent = TimerAddEvent(time.Now().Add(500*time.Millisecond), 250*time.Millisecond, func() {
					self.Currentposition += (self.maximum - self.minimum) / 10
					if self.Currentposition > self.maximum {
						self.Currentposition = self.maximum
					}
					if (self.callback!=nil) { self.callback(self.Currentposition) }
					PostUpdate()
				})
			} else {
				self.onelevator = true
				self.initialpos = x - offset
			}
		} else {
			h := self.Height() * self.Height() / (self.maximum - self.minimum)
			if h < 20 {
				h = 20
			}
			offset := (self.Height() - h) * (self.Currentposition - self.minimum) / (self.maximum - self.minimum)
			if y < offset {
				// click before
				self.Currentposition -= (self.maximum - self.minimum) / 10
				if self.Currentposition < self.minimum {
					self.Currentposition = self.minimum
				}
				if (self.callback!=nil) { self.callback(self.Currentposition) }
				PostUpdate()
				self.timerevent = TimerAddEvent(time.Now().Add(500*time.Millisecond), 250*time.Millisecond, func() {
					self.Currentposition -= (self.maximum - self.minimum) / 10
					if self.Currentposition < self.minimum {
						self.Currentposition = self.minimum
					}
					if (self.callback!=nil) { self.callback(self.Currentposition) }
					PostUpdate()
				})
			} else if y > offset+h {
				// click after
				self.Currentposition += (self.maximum - self.minimum) / 10
				if self.Currentposition > self.maximum {
					self.Currentposition = self.maximum
				}
				if (self.callback!=nil) { self.callback(self.Currentposition) }
				PostUpdate()
				self.timerevent = TimerAddEvent(time.Now().Add(500*time.Millisecond), 250*time.Millisecond, func() {
					self.Currentposition += (self.maximum - self.minimum) / 10
					if self.Currentposition > self.maximum {
						self.Currentposition = self.maximum
					}
					if (self.callback!=nil) { self.callback(self.Currentposition) }
					PostUpdate()
				})
			} else {
				self.onelevator = true
				self.initialpos = y - offset
			}
		}
	}
}

func (self *SWS_ScrollbarWidget) MousePressUp(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.buttondown = false
		self.onelevator = false
		if self.timerevent != nil {
			self.timerevent.StopRepeat()
			self.timerevent = nil
		}
	}
}

func (self *SWS_ScrollbarWidget) MouseMove(x, y, xrel, yrel int32) {
	if self.buttondown == true && self.onelevator {
		if self.horizontal {
			w := self.Width() * self.Width() / (self.maximum - self.minimum)
			if w < 20 {
				w = 20
			}
			//offset:=(self.Width()-w)*(self.Currentposition-self.minimum)/((self.maximum-self.minimum))
			xpos := x - self.initialpos
			if xpos < 0 {
				xpos = 0
			}
			if xpos > self.Width()-w {
				xpos = self.Width() - w
			}
			self.Currentposition = self.minimum + (self.maximum-self.minimum)*xpos/(self.Width()-w)
			if (self.callback!=nil) { self.callback(self.Currentposition) }
			PostUpdate()
		} else {
			h := self.Height() * self.Height() / (self.maximum - self.minimum)
			if h < 20 {
				h = 20
			}
			//offset:=(self.Height()-h)*(self.Currentposition-self.minimum)/((self.maximum-self.minimum))
			ypos := y - self.initialpos
			if ypos < 0 {
				ypos = 0
			}
			if ypos > self.Height()-h {
				ypos = self.Height() - h
			}
			self.Currentposition = self.minimum + (self.maximum-self.minimum)*ypos/(self.Height()-h)
			self.callback(self.Currentposition)
			PostUpdate()
		}
	}
}

func (self *SWS_ScrollbarWidget) Repaint() {
	self.SWS_CoreWidget.Repaint()
	self.SetDrawColor(50, 50, 50, 255)
	self.DrawLine(0, 0, 0, self.Height()-1)
	self.DrawLine(0, 0, self.Width()-1, 0)
	self.DrawLine(0, self.Height()-1, self.Width()-1, self.Height()-1)
	self.DrawLine(self.Width()-1, 0, self.Width()-1, self.Height()-1)
	self.SetDrawColor(100, 100, 100, 255)
	self.DrawLine(1, 1, 1, self.Height()-2)
	self.DrawLine(1, 1, self.Width()-2, 1)
	self.SetDrawColor(255, 255, 255, 255)
	self.DrawLine(1, self.Height()-2, self.Width()-2, self.Height()-2)
	self.DrawLine(self.Width()-2, 1, self.Width()-2, self.Height()-2)
	if self.horizontal {
		w := self.Width() * self.Width() / (self.maximum - self.minimum)
		if w < 20 {
			w = 20
		}
		offset := (self.Width() - w) * (self.Currentposition - self.minimum) / (self.maximum - self.minimum)
		self.SetDrawColor(50, 50, 50, 255)
		self.DrawLine(offset, 0, offset+w-1, 0)
		self.DrawLine(offset+w-1, 0, offset+w-1, self.Height()-1)
		self.DrawLine(offset+w-1, self.Height()-1, offset, self.Height()-1)
		self.DrawLine(offset, self.Height()-1, offset, 0)
		self.SetDrawColor(255, 255, 255, 255)
		self.DrawLine(offset+1, 1, offset+w-2, 1)
		self.DrawLine(offset+1, 1, offset+1, self.Height()-2)
		self.SetDrawColor(100, 100, 100, 255)
		self.DrawLine(offset+w-2, 0, offset+w-2, self.Height()-2)
		self.DrawLine(offset+1, self.Height()-2, offset+w-2, self.Height()-2)
		self.FillRect(offset+2, 2, w-4, self.Height()-4, 0xffdddddd)
	} else {
		h := self.Height() * self.Height() / (self.maximum - self.minimum)
		if h < 20 {
			h = 20
		}
		offset := (self.Height() - h) * (self.Currentposition - self.minimum) / (self.maximum - self.minimum)
		self.SetDrawColor(50, 50, 50, 255)
		self.DrawLine(0, offset, 0, offset+h-1)
		self.DrawLine(0, offset+h-1, self.Width()-1, offset+h-1)
		self.DrawLine(self.Width()-1, offset+h-1, self.Width()-1, offset)
		self.DrawLine(self.Width()-1, offset, 0, offset)
		self.SetDrawColor(255, 255, 255, 255)
		self.DrawLine(1, offset+1, 1, offset+h-2)
		self.DrawLine(1, offset+1, self.Width()-2, offset+1)
		self.SetDrawColor(100, 100, 100, 255)
		self.DrawLine(1, offset+h-2, self.Width()-2, offset+h-2)
		self.DrawLine(self.Width()-2, offset+1, self.Width()-2, offset+h-2)
		self.FillRect(2, offset+2, self.Width()-4, h-4, 0xffdddddd)
	}
}

func CreateScrollbarWidget(w, h int32, horizontal bool) *SWS_ScrollbarWidget {
	corewidget := CreateCoreWidget(w, h)
	corewidget.SetColor(0xffcccccc)
	widget := &SWS_ScrollbarWidget{SWS_CoreWidget: *corewidget,
		horizontal:      horizontal,
		minimum:         0,
		maximum:         0,
		callback:        nil,
		Currentposition: 0,
		buttondown:      false,
		onelevator:      false,
		timerevent:      nil}
	return widget
}
