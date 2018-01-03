package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

//
// ListWidget represent a vertical list of labels
type ListWidget struct {
	CoreWidget
	currentItem       *ListWidgetItem
	items             []*ListWidgetItem
	yoffset           int32
	scrollbar         *ScrollbarWidget
	mousedownonscroll bool
	hasfocus          bool
}

func (self *ListWidget) HasFocus(focus bool) {
	self.hasfocus = focus
	self.PostUpdate()
}

func (self *ListWidget) GetCurrentItem() *ListWidgetItem {
	return self.currentItem
}

func (self *ListWidget) AddItem(label string) {
	item := NewListWidgetItem(25, self.Width()-4, label, self)
	self.items = append(self.items, item)
	self.Resize(self.Width(), self.Height())
}

func (self *ListWidget) GetItems() []*ListWidgetItem {
	return self.items
}

func (self *ListWidget) RemoveItem(item *ListWidgetItem) {
	for i, l := range self.items {
		if l == item {
			self.items = append(self.items[:i], self.items[i+1:]...)
			if self.yoffset > int32(25*i) {
				self.yoffset -= 25
				if self.yoffset < 0 {
					self.yoffset = 0
				}
			}
			break
		}
	}
	if item == self.currentItem {
		self.currentItem = nil
	}
	self.PostUpdate()
}

func (self *ListWidget) SelectItem(item *ListWidgetItem) {
	if self.currentItem != nil {
		self.currentItem.Selected(false)
	}
	self.currentItem = item
	item.Selected(true)
	if self.valueChangedCallback != nil {
		self.valueChangedCallback()
	}
	self.PostUpdate()
}

func (self *ListWidget) Resize(width, height int32) {
	self.CoreWidget.Resize(width, height)

	self.scrollbar.Resize(15, height-4)
	if height < int32(25*len(self.items)) {
		self.CoreWidget.AddChild(self.scrollbar)
		self.scrollbar.SetMaximum(int32(25*len(self.items)) - self.Height())
		for _, i := range self.items {
			i.Resize(width-4-15, 25)
		}
	} else {
		self.CoreWidget.RemoveChild(self.scrollbar)
		for _, i := range self.items {
			i.Resize(width-4, 25)
		}
	}
	self.PostUpdate()
}

func (self *ListWidget) MousePressDown(x, y int32, button uint8) {
	if self.Height() < int32(25*len(self.items)) && x >= self.Width()-17 && x < self.Width()-2 && y > 2 && y < self.Height()-2 {
		self.mousedownonscroll = true
		self.scrollbar.MousePressDown(x-self.Width()-17, y-2, button)
	} else {
		if button == sdl.BUTTON_LEFT {
			y = (y - 2 + self.yoffset) / 25
			if y < int32(len(self.items)) {
				self.SelectItem(self.items[y])
			}
		}
	}
}

func (self *ListWidget) MousePressUp(x, y int32, button uint8) {
	if self.mousedownonscroll == true {
		self.scrollbar.MousePressUp(x-self.Width()-17, y-2, button)
		self.mousedownonscroll = false
	}
}

func (self *ListWidget) MouseMove(x, y, xrel, yrel int32) {
	if self.mousedownonscroll == true {
		self.scrollbar.MouseMove(x-self.Width()-17, y-2, xrel, yrel)
	}
}

func (self *ListWidget) IsInputWidget() bool {
	return true
}

func (self *ListWidget) KeyDown(key sdl.Keycode, mod uint16) {
	if key == sdl.K_TAB {
		if self.focusOnNextInputWidgetCallback != nil {
			self.focusOnNextInputWidgetCallback(!(mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT))
		}
	}
	if key == sdl.K_DOWN {
		if self.currentItem == nil && len(self.items) > 0 {
			self.SelectItem(self.items[0])
			return
		}
		if self.currentItem != nil && self.currentItem != self.items[len(self.items)-1] {
			for i, item := range self.items {
				if item == self.currentItem {
					self.SelectItem(self.items[i+1])

					if int32((i+2)*25) > self.yoffset+self.Height() {
						self.yoffset = int32((i+2)*25) - self.Height()
						self.scrollbar.SetPosition(self.yoffset)
						self.PostUpdate()
					}

					if int32((i+1)*25) < self.yoffset {
						self.yoffset = int32((i + 1) * 25)
						self.scrollbar.SetPosition(self.yoffset)
						self.PostUpdate()
					}
					break
				}
			}
		}
	}
	if key == sdl.K_UP {
		if self.currentItem == nil && len(self.items) > 0 {
			self.SelectItem(self.items[len(self.items)-1])
			return
		}
		if self.currentItem != nil && self.currentItem != self.items[0] {
			for i, item := range self.items {
				if item == self.currentItem {
					self.SelectItem(self.items[i-1])

					if int32((i)*25) > self.yoffset+self.Height() {
						self.yoffset = int32((i)*25) - self.Height()
						self.scrollbar.SetPosition(self.yoffset)
						self.PostUpdate()
					}

					if int32((i-1)*25) < self.yoffset {
						self.yoffset = int32((i - 1) * 25)
						self.scrollbar.SetPosition(self.yoffset)
						self.PostUpdate()
					}
					break
				}
			}
		}
	}
}

func (self *ListWidget) Repaint() {
	self.FillRect(0, 0, self.width, self.height, 0xffffffff)

	// do we show the scrollbar
	if self.Height() < int32(25*len(self.items)) {
		if self.scrollbar.IsDirty() {
			self.scrollbar.Repaint()
		}
		rectSrc := sdl.Rect{0, 0, self.scrollbar.Width(), self.scrollbar.Height()}
		rectDst := sdl.Rect{self.Width() - 2 - self.scrollbar.Width(), 2, self.scrollbar.Width(), self.scrollbar.Height()}
		self.scrollbar.Surface().Blit(&rectSrc, self.Surface(), &rectDst)
	}

	for i, item := range self.items {
		if item.IsDirty() {
			item.Repaint()
		}
		rectSrc := sdl.Rect{0, 0, item.Width(), item.Height()}
		rectDst := sdl.Rect{2, 2 - self.yoffset + int32(25*i), item.Width(), item.Height()}
		item.Surface().Blit(&rectSrc, self.Surface(), &rectDst)
	}

	self.SetDrawColor(0x88, 0x88, 0x88, 0xff)
	self.DrawLine(0, 0, 0, self.Height()-1)
	self.DrawLine(0, 0, self.Width()-1, 0)
	//	self.SetDrawColor(0xff, 0xff, 0xff, 0xff)
	//	self.DrawLine(self.Width()-1, self.Height()-1, self.Width()-1, 0)
	//	self.DrawLine(self.Width()-1, 0, 0, 0)
	self.DrawLine(1, 1, 1, self.Height()-2)
	self.DrawLine(1, self.Height()-2, self.Width()-2, self.Height()-2)
	self.DrawLine(self.Width()-2, self.Height()-2, self.Width()-2, 1)
	self.DrawLine(self.Width()-2, 1, 1, 1)
	if self.hasfocus {
		width := self.Width() - 3
		if self.Height() < int32(25*len(self.items)) {
			width = self.Width() - 3 - self.scrollbar.Width()
		}
		self.SetDrawColor(0x46, 0xc8, 0xe8, 255)
		self.DrawLine(2, 2, 2, self.Height()-3)
		self.DrawLine(2, self.Height()-3, width, self.Height()-3)
		self.DrawLine(width, self.Height()-3, width, 2)
		self.DrawLine(width, 2, 2, 2)
	}
}

func NewListWidget(w, h int32) *ListWidget {
	corewidget := NewCoreWidget(w, h)
	widget := &ListWidget{
		CoreWidget:        *corewidget,
		currentItem:       nil,
		items:             make([]*ListWidgetItem, 0, 0),
		yoffset:           0,
		scrollbar:         NewScrollbarWidget(25, h-4, false),
		mousedownonscroll: false,
		hasfocus:          false,
	}
	widget.scrollbar.Move(w-27, 2)
	widget.scrollbar.SetCallback(func(currentposition int32) {
		widget.yoffset = currentposition
		widget.PostUpdate()
	})
	return widget
}

type ListWidgetItem struct {
	CoreWidget
	label      string
	listwidget *ListWidget
	selected   bool
}

func NewListWidgetItem(h, w int32, label string, listwidget *ListWidget) *ListWidgetItem {
	corewidget := NewCoreWidget(h, w)
	return &ListWidgetItem{
		CoreWidget: *corewidget,
		label:      label,
		listwidget: listwidget,
		selected:   false,
	}
}

func (self *ListWidgetItem) GetText() string {
	return self.label
}

func (self *ListWidgetItem) Selected(selected bool) {
	self.selected = selected
	if self.selected {
		self.SetColor(0xffaaaaaa)
	} else {
		self.SetColor(0xffdddddd)
	}
	self.PostUpdate()
}

func (self *ListWidgetItem) Repaint() {
	self.CoreWidget.Repaint()
	// text
	var text *sdl.Surface
	var err error
	if self.label != "" {
		if text, err = self.Font().RenderUTF8Blended(self.label, sdl.Color{0, 0, 0, 255}); err != nil {
		}
		defer text.Free()
	}

	wGap := self.Width() - text.W
	hGap := self.Height() - text.H
	rectSrc := sdl.Rect{0, 0, text.W, text.H}
	rectDst := sdl.Rect{(wGap / 2), (hGap / 2), self.Width() - (wGap / 2), self.Height() - (hGap / 2)}
	if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
	}
}
