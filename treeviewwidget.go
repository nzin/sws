package sws

import (
	//"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type TreeViewItemWalkfunc func(level int32, item *TreeViewItem)

type TreeViewWidget struct {
	CoreWidget
	items        []*TreeViewItem
	currentfocus *TreeViewItem
}

func NewTreeViewWidget() *TreeViewWidget {
	tree := &TreeViewWidget{
		CoreWidget:   *NewCoreWidget(100, 100),
		items:        make([]*TreeViewItem, 0),
		currentfocus: nil,
	}
	return tree
}

// select a specific TreeViewItem (but do not open automatically til this item)
func (self *TreeViewWidget) SetFocusOn(item *TreeViewItem) {
	if self.currentfocus != nil {
		self.currentfocus.SetFocus(false)
	}
	self.currentfocus = item
	if self.currentfocus != nil {
		self.currentfocus.SetFocus(true)
	}
	PostUpdate()
}

func (self *TreeViewWidget) computeSize() {
	var width, height int32
	for _, item := range self.items {
		item.walkTreeView(0, true, func(level int32, i *TreeViewItem) {
			witem := i.getWidth(level * 25)
			if witem > width {
				width = witem
			}
			height += 25
		})
	}
	self.CoreWidget.Resize(width, height)
	PostUpdate()
}

func (self *TreeViewWidget) MousePressDown(x, y int32, button uint8) {
	var height int32
	for _, item := range self.items {
		item.walkTreeView(0, false, func(level int32, i *TreeViewItem) {
			if y > height && y < height+25 {
				i.MousePressDown(level*25, x, y-height, button)
			}
			height += 25
		})
	}

}

func (self *TreeViewWidget) Repaint() {
	self.CoreWidget.Repaint()
	var y int32
	for _, item := range self.items {
		item.walkTreeView(0, false, func(level int32, i *TreeViewItem) {
			i.Repaint(25 * level)

			rectSrc := sdl.Rect{0, 0, i.Surface().W, i.Surface().H}
			rectDst := sdl.Rect{0, y, i.Surface().W, i.Surface().H}
			if err := i.Surface().Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
			y += 25
		})
	}
}

func (self *TreeViewWidget) Resize(width, height int32) {
}

func (self *TreeViewWidget) AddItem(item *TreeViewItem) {
	self.items = append(self.items, item)
	item.walkTreeView(0, true, func(level int32, i *TreeViewItem) {
		i.treeview = self
	})
	self.computeSize()
}

func (self *TreeViewWidget) RemoveItem(item *TreeViewItem) {
	for i, vi := range self.items {
		if vi == item {
			self.items = append(self.items[:i], self.items[i+1:]...)
		}
	}
	self.computeSize()
}

type TreeViewItem struct {
	CoreWidget
	treeview      *TreeViewWidget
	subitems      []*TreeViewItem
	icon          *sdl.Surface
	label         string
	textcolor     sdl.Color
	opened        bool
	callback      func()
	focus         bool
	computedwidth int32
}

func NewTreeViewItem(label string, icon string, callback func()) *TreeViewItem {
	item := &TreeViewItem{
		CoreWidget:    *NewCoreWidget(100, 25),
		subitems:      make([]*TreeViewItem, 0),
		label:         label,
		opened:        false,
		callback:      callback,
		computedwidth: -1,
		textcolor:     sdl.Color{0, 0, 0, 255},
	}
	if img, err := img.Load(icon); err == nil {
		item.icon = img
	}
	return item
}

func (self *TreeViewItem) getWidth(shift int32) int32 {
	if self.computedwidth > 0 {
		return self.computedwidth
	}
	shift += 25
	if self.icon != nil {
		shift += 25
	}
	width, _, _ := self.Font().SizeUTF8(self.label)
	shift += int32(width)
	self.computedwidth = shift

	self.Resize(self.computedwidth, 25)
	return self.computedwidth
}

//
// helper function to set up the treeview pointer
//
func (self *TreeViewItem) walkTreeView(level int32, onclosed bool, walkfunc TreeViewItemWalkfunc) {
	walkfunc(level, self)
	for _, i := range self.subitems {
		if onclosed {
			i.walkTreeView(level+1, onclosed, walkfunc)
		} else {
			if self.opened == true {
				i.walkTreeView(level+1, onclosed, walkfunc)
			}
		}
	}
}

// we add & refresh
func (self *TreeViewItem) AddSubItem(item *TreeViewItem) {
	self.subitems = append(self.subitems, item)
	item.treeview = self.treeview
	if item.treeview != nil {
		self.treeview.computeSize()
	}
}

func (self *TreeViewItem) RemoveSubItem(item *TreeViewItem) {
	for i, vi := range self.subitems {
		if vi == item {
			self.subitems = append(self.subitems[:i], self.subitems[i+1:]...)
		}
	}
	self.treeview.computeSize()
}

// paint on shift*16, arrow, icon, text
func (self *TreeViewItem) Repaint(shift int32) {
	self.CoreWidget.Repaint()
	if len(self.subitems) > 0 {
		self.SetDrawColor(0, 0, 0, 255)
		if self.opened == false {
			self.DrawLine(shift+10, 5, shift+10, 17)
			self.DrawLine(shift+11, 6, shift+11, 16)
			self.DrawLine(shift+12, 7, shift+12, 15)
			self.DrawLine(shift+13, 8, shift+13, 14)
			self.DrawLine(shift+14, 9, shift+14, 13)
			self.DrawLine(shift+15, 10, shift+15, 12)
		} else {
			self.DrawLine(shift+5, 10, shift+17, 10)
			self.DrawLine(shift+6, 11, shift+16, 11)
			self.DrawLine(shift+7, 12, shift+15, 12)
			self.DrawLine(shift+8, 13, shift+14, 13)
			self.DrawLine(shift+9, 14, shift+13, 14)
			self.DrawLine(shift+10, 15, shift+12, 15)
		}
	}
	shift += 25
	if self.icon != nil {
		wimage := self.icon.W
		himage := self.icon.H
		if wimage > 25 {
			wimage = 25
		}
		if himage > 25 {
			himage = 25
		}
		rectSrc := sdl.Rect{0, 0, wimage, himage}
		rectDst := sdl.Rect{shift + (25-wimage)/2, (25 - himage) / 2, wimage, himage}
		if err := self.icon.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
		}
		shift += 25
	}
	var text *sdl.Surface
	var err error
	if text, err = self.Font().RenderUTF8_Blended(self.label, self.textcolor); err != nil {
	}
	defer text.Free()
	rectSrc := sdl.Rect{0, 0, text.W, text.H}
	rectDst := sdl.Rect{shift, 0, text.W, text.H}
	if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
	}
}

// if we click on arrow -> we open/close (and refresh)
// if we click on the rest -> we emit a callback (and grab the focus)
func (self *TreeViewItem) MousePressDown(shift, x, y int32, button uint8) {
	if len(self.subitems) > 0 && x > shift && x < shift+25 {
		self.opened = !self.opened
		self.treeview.computeSize()
	} else {
		self.treeview.SetFocusOn(self)
	}

}

// to know if we highlight (or not) this row
func (self *TreeViewItem) SetFocus(focus bool) {
	self.focus = focus
	if focus {
		self.bgColor = 0xffcccccc
		if self.callback != nil {
			self.callback()
		}
	} else {
		self.bgColor = 0xffdddddd
	}
}
