package sws

import (
)

//
// a list selection of SWS_FlatButtonWidget
//
type SWS_ListWidget struct {
	SWS_VBoxWidget
	currentItem int32
	callbacks   map[*SWS_FlatButtonWidget]func()
}

func (self *SWS_ListWidget) AddChild(child SWS_Widget) {
	panic("SWS_ListWidget do not honor this method")
}

func (self *SWS_ListWidget) RemoveChild(child SWS_Widget) {
	panic("SWS_ListWidget do not honor this method")
}

func (self *SWS_ListWidget) GetCurrentItem() (int32, *SWS_FlatButtonWidget) {
	return self.currentItem,self.children[self.currentItem].(*SWS_FlatButtonWidget)
}

func (self *SWS_ListWidget) AddItem(height int32,label string,image string,callback func()) {
	button:=CreateFlatButtonWidget(self.Width(),height,label)
	button.SetCentered(false)
	if image!="" { button.SetImage(image) }
	self.callbacks[button]=callback

	button.SetColor(0xffdddddd)
	button.SetClicked(func() {
		self.SetCurrentItem(button)
		if (self.callbacks[button]!=nil) {
			self.callbacks[button]()
		}
	})
	self.SWS_VBoxWidget.AddChild(button)
}

func (self *SWS_ListWidget) SetCurrentItem(child *SWS_FlatButtonWidget) {
	for i,c := range self.children {
		if c==child {
			self.currentItem=int32(i)
			c.(*SWS_FlatButtonWidget).SetColor(0xff8888ff)
		} else {
			c.(*SWS_FlatButtonWidget).SetColor(0xffdddddd)
		}
	}
	PostUpdate()
}

// not so usefull for the moment
func (self *SWS_ListWidget) RemoveItem(child *SWS_FlatButtonWidget) {
	if child==nil { return }
	self.SWS_VBoxWidget.RemoveChild(child)
	if self.currentItem==int32(len(self.children)) {
		self.currentItem--
		if (self.currentItem>0) {
			self.SetCurrentItem(self.children[self.currentItem].(*SWS_FlatButtonWidget))
		}
	}
	delete(self.callbacks,child)
}

func CreateListWidget(w, h int32) *SWS_ListWidget {
	vbox := CreateVBoxWidget(w, h)
	widget := &SWS_ListWidget{SWS_VBoxWidget: *vbox,
		currentItem: -1,
		callbacks: make(map[*SWS_FlatButtonWidget]func()),
	}
	return widget
}
