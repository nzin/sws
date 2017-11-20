package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

type TabWidget struct {
	CoreWidget
	labels []string
	tabs   []Widget
	active int32
}

func (self *TabWidget) AddTab(label string, tab Widget) {
	tab.Move(5, 35)
	if len(self.tabs) == 0 {
		self.AddChild(tab)
		tab.Resize(self.Width()-10, self.Height()-40)
	}
	self.labels = append(self.labels, label)
	self.tabs = append(self.tabs, tab)
	self.PostUpdate()
}

func (self *TabWidget) Resize(width, height int32) {
	self.CoreWidget.Resize(width, height)
	if len(self.tabs) != 0 {
		self.tabs[self.active].Resize(self.Width()-10, self.Height()-40)
	}
}

func (self *TabWidget) SelectTab(index int32) {
	if index < 0 || index >= int32(len(self.labels)) {
		return
	}
	if index != self.active {
		self.RemoveChild(self.tabs[self.active])
		self.active = index
		self.AddChild(self.tabs[self.active])
		self.tabs[self.active].Resize(self.Width()-10, self.Height()-40)
		self.PostUpdate()
	}
}

func (self *TabWidget) MousePressDown(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		active := self.active
		var offset int32
		offset = 5
		for i, label := range self.labels {
			w, _, _ := self.Font().SizeUTF8(label)
			if x >= offset && x < offset+int32(w)+10 && y < 30 && y > 5 {
				active = int32(i)
			}
			offset += int32(w) + 20
		}
		if active != self.active {
			self.RemoveChild(self.tabs[self.active])
			self.active = active
			self.AddChild(self.tabs[self.active])
			self.tabs[self.active].Resize(self.Width()-10, self.Height()-40)
			self.PostUpdate()
		}
	}
}

func (self *TabWidget) Repaint() {
	self.CoreWidget.Repaint()
	self.SetDrawColor(0x88, 0x88, 0x88, 255)
	self.DrawLine(0, self.Height()-1, self.Width()-1, self.Height()-1)
	self.DrawLine(self.Width()-1, self.Height()-1, self.Width()-1, 30)
	self.SetDrawColor(255, 255, 255, 255)
	self.DrawLine(0, self.Height()-1, 0, 30)
	self.DrawLine(0, 30, self.Width(), 30)
	if len(self.tabs) == 0 {
		self.DrawLine(5, 7, 5, 29)
		self.DrawLine(6, 6, 24, 6)
		self.SetDrawColor(0x88, 0x88, 0x88, 255)
		self.DrawLine(25, 7, 25, 29)

		self.SetDrawColor(0xdd, 0xdd, 0xdd, 255)
		self.DrawLine(6, 30, 24, 30)
	} else {
		var offset int32
		offset = 5
		for i, label := range self.labels {
			text, err := self.Font().RenderUTF8_Blended(label, sdl.Color{0, 0, 0, 255})
			if err == nil {
				rectSrc := sdl.Rect{0, 0, text.W, text.H}
				rectDst := sdl.Rect{offset + 10, 7, text.W, text.H}
				if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
				}
				self.SetDrawColor(255, 255, 255, 255)
				self.DrawLine(offset, 7, offset, 29)
				self.DrawLine(offset+1, 6, offset+text.W+18, 6)
				self.SetDrawColor(0x88, 0x88, 0x88, 255)
				self.DrawLine(offset+text.W+19, 7, offset+text.W+19, 29)
				if int32(i) == self.active {
					self.SetDrawColor(0xdd, 0xdd, 0xdd, 255)
					self.DrawLine(offset+1, 30, offset+text.W+18, 30)
				}
				offset += 20 + text.W
			} else {
				self.DrawLine(offset, 7, offset, 29)
				self.DrawLine(offset+1, 6, offset+18, 6)
				self.SetDrawColor(0x88, 0x88, 0x88, 255)
				self.DrawLine(offset+19, 7, offset+19, 29)
				if int32(i) == self.active {
					self.SetDrawColor(0x88, 0x88, 0x88, 255)
					self.DrawLine(offset+1, 30, offset+18, 30)
				}
				offset += 20
			}
		}
	}
}

func NewTabWidget(w, h int32) *TabWidget {
	corewidget := NewCoreWidget(w, h)
	widget := &TabWidget{CoreWidget: *corewidget,
		labels: make([]string, 0),
		tabs:   make([]Widget, 0),
		active: 0,
	}
	return widget
}
