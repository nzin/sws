package sws

import ()

//
// a list selection of FlatButtonWidget
//
type ListWidget struct {
	VBoxWidget
	currentItem int32
	callbacks   map[*FlatButtonWidget]func()
}

func (self *ListWidget) AddChild(child Widget) {
	panic("ListWidget do not honor this method")
}

func (self *ListWidget) RemoveChild(child Widget) {
	panic("ListWidget do not honor this method")
}

func (self *ListWidget) GetCurrentItem() (int32, *FlatButtonWidget) {
	return self.currentItem, self.children[self.currentItem].(*FlatButtonWidget)
}

func (self *ListWidget) AddItem(height int32, label string, image string, callback func()) {
	button := NewFlatButtonWidget(self.Width(), height, label)
	button.SetCentered(false)
	if image != "" {
		button.SetImage(image)
	}
	self.callbacks[button] = callback

	button.SetColor(0xffdddddd)
	button.SetClicked(func() {
		self.SetCurrentItem(button)
		if self.callbacks[button] != nil {
			self.callbacks[button]()
		}
	})
	self.VBoxWidget.AddChild(button)
}

func (self *ListWidget) SetCurrentItem(child *FlatButtonWidget) {
	for i, c := range self.children {
		if c == child {
			self.currentItem = int32(i)
			c.(*FlatButtonWidget).SetColor(0xff8888ff)
		} else {
			c.(*FlatButtonWidget).SetColor(0xffdddddd)
		}
	}
	self.PostUpdate()
}

// not so usefull for the moment
func (self *ListWidget) RemoveItem(child *FlatButtonWidget) {
	if child == nil {
		return
	}
	self.VBoxWidget.RemoveChild(child)
	if self.currentItem == int32(len(self.children)) {
		self.currentItem--
		if self.currentItem > 0 {
			self.SetCurrentItem(self.children[self.currentItem].(*FlatButtonWidget))
		}
	}
	delete(self.callbacks, child)
}

func NewListWidget(w, h int32) *ListWidget {
	vbox := NewVBoxWidget(w, h)
	widget := &ListWidget{VBoxWidget: *vbox,
		currentItem: -1,
		callbacks:   make(map[*FlatButtonWidget]func()),
	}
	return widget
}
