package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SplitviewWidget struct {
	CoreWidget
	leftwidget  Widget
	rightwidget Widget
	splitwidget *SplitWidget
	horizontal  bool
}

func (self *SplitviewWidget) SplitBarMovable(movable bool) {
	self.splitwidget.SetMovable(movable)
}

func (self *SplitviewWidget) PlaceSplitBar(position int32) {
	self.splitwidget.Place(position)
}

func (self *SplitviewWidget) Resize(width, height int32) {
	if width < 45 {
		width = 45
	}
	if height < 45 {
		height = 45
	}
	self.CoreWidget.Resize(width, height)

	if self.horizontal {
		self.splitwidget.Resize(4, height)
		if self.leftwidget != nil {
			self.leftwidget.Resize(self.splitwidget.X(), height)
		}
		if self.rightwidget != nil {
			self.rightwidget.Resize(width-(self.splitwidget.X()+self.splitwidget.Width()), height)
		}
	} else {
		self.splitwidget.Resize(width, 4)
		if self.leftwidget != nil {
			self.leftwidget.Resize(width, self.splitwidget.Y())
		}
		if self.rightwidget != nil {
			self.rightwidget.Resize(width, height-(self.splitwidget.Y()+self.splitwidget.Height()))
		}
	}
	PostUpdate()
}

func (self *SplitviewWidget) SetLeftWidget(widget Widget) bool {

	if widget == nil {
		return false
	}
	if self.leftwidget != nil {
		self.RemoveChild(self.leftwidget)
	}
	self.leftwidget = widget
	self.AddChild(widget)
	widget.SetParent(self)
	if self.horizontal {
		widget.Resize(self.splitwidget.X(), self.Height())
	} else {
		widget.Resize(self.Width(), self.splitwidget.Y())
	}
	PostUpdate()

	return true
}

func (self *SplitviewWidget) SetRightWidget(widget Widget) bool {

	if widget == nil {
		return false
	}
	if self.rightwidget != nil {
		self.RemoveChild(self.rightwidget)
	}
	self.rightwidget = widget
	self.AddChild(widget)
	widget.SetParent(self)
	if self.horizontal {
		widget.Resize(self.Width()-(self.splitwidget.X()+self.splitwidget.Width()), self.Height())
		widget.Move(self.splitwidget.X()+self.splitwidget.Width(), 0)
	} else {
		widget.Resize(self.Width(), self.Height()-(self.splitwidget.Y()+self.splitwidget.Height()))
		widget.Move(0, self.splitwidget.Y()+self.splitwidget.Height())
	}
	PostUpdate()

	return true
}

func NewSplitviewWidget(w, h int32, horizontal bool) *SplitviewWidget {
	corewidget := NewCoreWidget(w, h)
	widget := &SplitviewWidget{CoreWidget: *corewidget,
		leftwidget:  nil,
		rightwidget: nil,
		horizontal:  horizontal,
		splitwidget: NewSplitWidget(horizontal),
	}
	widget.AddChild(widget.splitwidget)
	if horizontal {
		widget.splitwidget.Move(w/2, 0)
		widget.splitwidget.SetCallback(func() {
			rightwidth := widget.Width() - (widget.splitwidget.X() + widget.splitwidget.Width())
			if rightwidth < 25 {
				rightwidth = 25
			}
			if widget.leftwidget != nil {
				widget.leftwidget.Resize(widget.splitwidget.X(), widget.Height())
			}
			if widget.rightwidget != nil {
				widget.rightwidget.Move(widget.splitwidget.X()+widget.splitwidget.Width(), 0)
				widget.rightwidget.Resize(rightwidth, widget.Height())
			}
			PostUpdate()
		})
	} else {
		widget.splitwidget.Move(0, h/2)
		widget.splitwidget.SetCallback(func() {
			rightheight := widget.Height() - (widget.splitwidget.Y() + widget.splitwidget.Height())
			if rightheight < 25 {
				rightheight = 25
			}
			if widget.leftwidget != nil {
				widget.leftwidget.Resize(widget.Width(), widget.splitwidget.Y())
			}
			if widget.rightwidget != nil {
				widget.rightwidget.Move(0, widget.splitwidget.Y()+widget.splitwidget.Height())
				widget.rightwidget.Resize(widget.Width(), rightheight)
			}
			PostUpdate()
		})
	}
	return widget
}

type SplitWidget struct {
	CoreWidget
	horizontal bool
	callback   func()
	buttondown bool
	initialPos int32
	movable    bool
}

func (self *SplitWidget) SetMovable(movable bool) {
	self.movable = movable
	PostUpdate()
}

func (self *SplitWidget) Repaint() {
	self.CoreWidget.Repaint()
	if self.movable {
		self.SetDrawColor(255, 255, 255, 255)
		self.DrawLine(0, 0, 0, self.Height()-1)
		self.DrawLine(0, 0, self.Width()-1, 0)
		self.SetDrawColor(100, 100, 100, 100)
		self.DrawLine(1, self.Height()-1, self.Width()-1, self.Height()-1)
		self.DrawLine(self.Width()-1, 1, self.Width()-1, self.Height()-1)
	}
}

func (self *SplitWidget) SetCallback(callback func()) {
	self.callback = callback
}

func (self *SplitWidget) MousePressDown(x, y int32, button uint8) {
	if self.movable {
		if button == sdl.BUTTON_LEFT {
			self.buttondown = true
			if self.horizontal {
				self.initialPos = x
			} else {
				self.initialPos = y
			}
		}
	}
}

func (self *SplitWidget) MouseMove(x, y, xrel, yrel int32) {
	if self.buttondown == true {
		if self.horizontal {
			move := x - self.initialPos
			self.Move(self.X()+move, 0)
			if self.X() < 25 {
				self.Move(25, 0)
			}
		} else {
			move := y - self.initialPos
			self.Move(0, self.Y()+move)
			if self.Y() < 25 {
				self.Move(0, 25)
			}
		}
		if self.callback != nil {
			self.callback()
		}
		PostUpdate()
	}
}

func (self *SplitWidget) MousePressUp(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.buttondown = false
	}
}

func (self *SplitWidget) Place(position int32) {
	if self.horizontal {
		self.Move(position, 0)
	} else {
		self.Move(0, position)
	}
	if self.callback != nil {
		self.callback()
	}
}

func NewSplitWidget(horizontal bool) *SplitWidget {
	var corewidget *CoreWidget
	if horizontal {
		corewidget = NewCoreWidget(4, 100)
	} else {
		corewidget = NewCoreWidget(100, 4)
	}
	widget := &SplitWidget{CoreWidget: *corewidget,
		horizontal: horizontal,
		callback:   nil,
		buttondown: false,
		movable:    true}
	return widget
}
