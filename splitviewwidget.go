package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SWS_SplitviewWidget struct {
	SWS_CoreWidget
	leftwidget  SWS_Widget
	rightwidget SWS_Widget
	splitwidget *SWS_SplitWidget
	horizontal  bool
}

func (self *SWS_SplitviewWidget) SplitBarMovable(movable bool) {
	self.splitwidget.SetMovable(movable)
}

func (self *SWS_SplitviewWidget) PlaceSplitBar(position int32) {
	self.splitwidget.Place(position)
}

func (self *SWS_SplitviewWidget) Resize(width, height int32) {
        if width < 45 {
                width = 45
        }
        if height < 45 {
                height = 45
        }
        self.SWS_CoreWidget.Resize(width, height)

	if self.horizontal {
		self.splitwidget.Resize(4,height)
		if self.leftwidget!=nil {
			self.leftwidget.Resize(self.splitwidget.X(),height)
		}
		if self.rightwidget!=nil {
			self.rightwidget.Resize(width-(self.splitwidget.X()+self.splitwidget.Width()),height)
		}
	} else {
		self.splitwidget.Resize(width,4)
		if self.leftwidget!=nil {
			self.leftwidget.Resize(width,self.splitwidget.Y())
		}
		if self.rightwidget!=nil {
			self.rightwidget.Resize(height-(self.splitwidget.Y()+self.splitwidget.Height()),width)
		}
	}
	PostUpdate()
}

func (self *SWS_SplitviewWidget) SetLeftWidget(widget SWS_Widget) bool {

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
		widget.Resize(self.splitwidget.X(),self.Height())
	} else {
		widget.Resize(self.Width(),self.splitwidget.Y())
	}
	PostUpdate()

	return true
}

func (self *SWS_SplitviewWidget) SetRightWidget(widget SWS_Widget) bool {

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
		widget.Resize(self.Width()-(self.splitwidget.X()+self.splitwidget.Width()),self.Height())
		widget.Move(self.splitwidget.X()+self.splitwidget.Width(),0)
	} else {
		widget.Resize(self.Width(),self.Height()-(self.splitwidget.Y()+self.splitwidget.Height()))
		widget.Move(0,self.splitwidget.Y()+self.splitwidget.Height())
	}
	PostUpdate()

	return true
}

func CreateSplitviewWidget(w, h int32, horizontal bool) *SWS_SplitviewWidget {
	corewidget := CreateCoreWidget(w, h)
	widget := &SWS_SplitviewWidget{SWS_CoreWidget: *corewidget,
		leftwidget:  nil,
		rightwidget: nil,
		horizontal:  horizontal,
		splitwidget: CreateSplitWidget(horizontal),
	}
	widget.AddChild(widget.splitwidget)
	if horizontal {
		widget.splitwidget.Move(w/2,0)
		widget.splitwidget.SetCallback(func() {
			rightwidth:=widget.Width()-(widget.splitwidget.X()+widget.splitwidget.Width())
			if rightwidth<25 { rightwidth=25 }
			if widget.leftwidget!=nil {
				widget.leftwidget.Resize(widget.splitwidget.X(),widget.Height())
			}
			if widget.rightwidget!=nil {
				widget.rightwidget.Move(widget.splitwidget.X()+widget.splitwidget.Width(),0)
				widget.rightwidget.Resize(rightwidth,widget.Height())
			}
			PostUpdate()
		})
	} else {
		widget.splitwidget.Move(0,h/2)
		widget.splitwidget.SetCallback(func() {
			rightheight:=widget.Height()-(widget.splitwidget.Y()+widget.splitwidget.Height())
			if rightheight<25 { rightheight=25 }
			if widget.leftwidget!=nil {
				widget.leftwidget.Resize(widget.Width(),widget.splitwidget.Y())
			}
			if widget.rightwidget!=nil {
				widget.rightwidget.Move(0,widget.splitwidget.Y()+widget.splitwidget.Height())
				widget.rightwidget.Resize(widget.Width(),rightheight)
			}
			PostUpdate()
		})
	}
	return widget
}

type SWS_SplitWidget struct {
	SWS_CoreWidget
	horizontal bool
	callback   func()
	buttondown bool
	initialPos int32
	movable    bool
}

func (self *SWS_SplitWidget) SetMovable(movable bool) {
	self.movable=movable
	PostUpdate()
}

func (self *SWS_SplitWidget) Repaint() {
	self.SWS_CoreWidget.Repaint()
	if (self.movable) {
		self.SetDrawColor(255, 255, 255, 255)
		self.DrawLine(0, 0, 0, self.Height()-1)
		self.DrawLine(0, 0, self.Width()-1,0)
		self.SetDrawColor(100, 100, 100, 100)
		self.DrawLine(1, self.Height()-1,self.Width()-1,self.Height()-1)
		self.DrawLine(self.Width()-1,1,self.Width()-1,self.Height()-1)
	}
}

func (self *SWS_SplitWidget) SetCallback(callback func()) {
	self.callback=callback
}

func (self *SWS_SplitWidget) MousePressDown(x, y int32, button uint8) {
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

func (self *SWS_SplitWidget) MouseMove(x, y, xrel, yrel int32) {
	if (self.buttondown==true) {
		if self.horizontal {
			move:=x-self.initialPos
			self.Move(self.X()+move,0)
			if (self.X()<25) {
				self.Move(25,0)
			}
		} else {
			move:=y-self.initialPos
			self.Move(0,self.Y()+move)
			if (self.Y()<25) {
				self.Move(0,25)
			}
		}
		if self.callback!=nil {
			self.callback()
		}
		PostUpdate()
	}
}

func (self *SWS_SplitWidget) MousePressUp(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.buttondown = false
	}
}

func (self *SWS_SplitWidget) Place(position int32) {
	if self.horizontal {
		self.Move(position,0)
	} else {
		self.Move(0,position)
	}
	if self.callback!=nil {
		self.callback()
	}
}

func CreateSplitWidget(horizontal bool) *SWS_SplitWidget {
	var corewidget *SWS_CoreWidget
	if horizontal {
		corewidget = CreateCoreWidget(4, 100)
	} else {
		corewidget = CreateCoreWidget(100, 4)
	}
	widget := &SWS_SplitWidget{SWS_CoreWidget: *corewidget,
		horizontal: horizontal,
		callback:   nil,
		buttondown: false,
		movable:    true}
	return widget
}
