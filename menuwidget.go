package sws

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
)

type MenuItem interface {
	Repaint(selected bool) *sdl.Surface
	WidthHeight() (int32, int32)
	SubMenu() *MenuWidget
	Clicked()
	Destroy()
}

type MenuItemLabel struct {
	font          *ttf.Font
	surface       *sdl.Surface
	Label         string
	ClickCallback func()
	subMenu       *MenuWidget
}

func (self *MenuItemLabel) SubMenu() *MenuWidget {
	return self.subMenu
}

func (self *MenuItemLabel) SetSubMenu(sub *MenuWidget) {
	self.subMenu = sub
}

func (self *MenuItemLabel) Destroy() {
	self.surface.Free()
}

func (self *MenuItemLabel) Repaint(selected bool) *sdl.Surface {

	if self.Label != "" {
		var err error
		var solid *sdl.Surface
		color := sdl.Color{0, 0, 0, 255}
		if selected {
			color = sdl.Color{255, 255, 255, 255}
		}

		if solid, err = self.font.RenderUTF8_Blended(self.Label, color); err != nil {
			fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
			return self.surface
		}
		self.surface.Free()
		self.surface = solid
	}

	return self.surface
}

func (self *MenuItemLabel) WidthHeight() (int32, int32) {
	w, h, _ := self.font.SizeUTF8(self.Label)
	return int32(w), int32(h)
}

func (self *MenuItemLabel) Clicked() {
	if self.ClickCallback != nil {
		self.ClickCallback()
	}
	hideMenu(nil)
}

func NewMenuItemLabel(label string, callback func()) *MenuItemLabel {
	w, h, _ := defaultFont.SizeUTF8(label)
	surface, err := sdl.CreateRGBSurface(0, int32(w), int32(h), 32, 0x00ff0000, 0x0000ff00, 0x000000ff, 0xff000000)
	if err != nil {
		panic(err)
	}
	menuitem := &MenuItemLabel{
		Label:         label,
		ClickCallback: callback,
		subMenu:       nil,
		surface:       surface,
		//renderer:      renderer,
		font: defaultFont}
	return menuitem
}

//
// MenuWidget: representation of a (floating) menu.
// It is composed of several MenuItem (that could have sub-menus)
//
// in the main loop
// if menuInitiator widget exists, then:
// - if MouseMove -> find the corresponding menuInitiator/menu and
//    - on the widget: destroy childs and treat the event (eventually create child)
// - if MouseDown and we are not on menuInitiator/menu -> destroy everything
// - if MouseUp
//   - if we are not on the menuInitiator/menu -> destroy everything
//   - else we send a MouseUp on the menuInitiator/menu
//
type MenuWidget struct {
	CoreWidget
	items         []MenuItem
	activeItem    int
	lastSubActive int
}

func (self *MenuWidget) Destroy() {
	// cannot do self.CoreWidget.RemoveChild(self) <- it cast it into CoreWidget (which is a WSW_Widget interface, but the data/struct are not the same)
	self.Parent().RemoveChild(self)
	self.surface.Free()
	for _, i := range self.items {
		i.Destroy()
	}
}

func (self *MenuWidget) AddItem(item MenuItem) {
	self.items = append(self.items, item)
	w, h := item.WidthHeight()
	w = w + 30 + 4 // add space for margins and UI border
	if w > self.width {
		self.width = w
	}
	self.height += h

	// recreate the surface
	var err error
	self.surface.Free()
	self.surface, err = sdl.CreateRGBSurface(0, self.width, self.height, 32, 0x00ff0000, 0x0000ff00, 0x000000ff, 0xff000000)
	if err != nil {
		panic(err)
	}
}

func (self *MenuWidget) MousePressDown(x, y int32, button uint8) {
}

func (self *MenuWidget) MousePressUp(x, y int32, button uint8) {
	if self.activeItem != -1 {
		submenu := self.items[self.activeItem].SubMenu()
		if submenu == nil {
			self.items[self.activeItem].Clicked()
		}
	}
}

func (self *MenuWidget) MouseMove(x, y, xrel, yrel int32) {
	previousActiveItem := self.activeItem
	x -= 2 // UI border
	y -= 2 // UI border
	self.activeItem = -1
	var yy int32
	yy = 0
	if x >= 0 && x < self.width {
		for i, item := range self.items {
			_, h := item.WidthHeight()
			if yy <= y && yy+h > y {
				self.activeItem = i
				break
			}
			yy += h
		}
	}
	if previousActiveItem != self.activeItem {
		//hideMenu(self)
		if previousActiveItem != -1 && self.activeItem != -1 {
			submenu := self.items[previousActiveItem].SubMenu()
			if submenu != nil {
				hideMenu(submenu)
			}
		}
		if self.lastSubActive != -1 && self.activeItem != -1 {
			submenu := self.items[self.lastSubActive].SubMenu()
			if submenu != nil {
				hideMenu(submenu)
			}
		}
		if self.activeItem != -1 {
			submenu := self.items[self.activeItem].SubMenu()
			if submenu != nil {
				self.lastSubActive = self.activeItem
				submenu.Move(self.X()+self.Width()-2, self.Y()+yy)
				ShowMenu(submenu)
			} else {
				self.lastSubActive = -1
			}
		}
		self.PostUpdate()
	}
}

func (self *MenuWidget) Repaint() {
	var y int32
	rect := sdl.Rect{0, 0, self.width, self.height}
	self.surface.FillRect(&rect, 0xffdddddd)

	renderer, err := sdl.CreateSoftwareRenderer(self.surface)
	if err != nil {
		panic(err)
	}
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.DrawRect(&rect)
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawLine(1, 1, int(self.width-2), 1)
	renderer.DrawLine(1, 1, 1, int(self.height)-2)
	renderer.SetDrawColor(0x88, 0x88, 0x88, 255)
	renderer.DrawLine(int(self.width)-2, 2, int(self.width)-2, int(self.height)-2)
	renderer.DrawLine(2, int(self.height)-2, int(self.width)-2, int(self.height)-2)

	for i, item := range self.items {
		w, h := item.WidthHeight()
		if i == self.activeItem || i == self.lastSubActive {
			rect := sdl.Rect{2, y + 2, self.width - 4, h}
			self.surface.FillRect(&rect, 0xff8888ff)
		}
		surface := item.Repaint(i == self.activeItem || i == self.lastSubActive)
		rectSrc := sdl.Rect{0, 0, w, h}
		rectDst := sdl.Rect{5 + 2, y + 2, w, h}
		surface.Blit(&rectSrc, self.surface, &rectDst)

		if item.SubMenu() != nil {
			var i int32
			for i = 0; i < 5; i++ {
				rect := sdl.Rect{self.width - 15 + 2, y + (h / 2) - i + 2, (5 - i) * 2, 2*i + 1}
				self.surface.FillRect(&rect, 0xff000000)
			}
		}

		y += h
	}
}

func NewMenuWidget() *MenuWidget {
	corewidget := NewCoreWidget(4, 4)
	widget := &MenuWidget{CoreWidget: *corewidget,
		items:         make([]MenuItem, 0, 0),
		activeItem:    -1,
		lastSubActive: -1}
	return widget
}

// menuInitiator when set to not nil, is the widget creating the menu,
// usually a bar menu, or a dropdown button
//
var menuInitiator Widget

//
// menuStack is the list of all shown menu/submenu currently
// the lowest ( [0] ) is the root
//
var menuStack []*MenuWidget

func ShowMenu(menu *MenuWidget) {
	if menuStack == nil {
		menuStack = append(make([]*MenuWidget, 0, 0), menu)
	} else {
		menuStack = append(menuStack, menu)
	}
	root.AddChild(menu)
	root.PostUpdate()
}

func findMenu(x int32, y int32) *MenuWidget {
	if menuStack == nil {
		return nil
	}
	for i := len(menuStack) - 1; i >= 0; i-- {
		menu := menuStack[i]
		if x >= menu.X() && x < menu.X()+menu.Width() && y >= menu.Y() && y < menu.Y()+menu.Height() {
			return menu
		}
	}
	return nil
}

//
// how do we show or hide a menu? Well simple
// we add them to the RootWidget as a child
//
func hideMenu(menu *MenuWidget) {
	if menuStack == nil {
		return
	}
	// destroy all menus
	if menu == nil {
		for _, m := range menuStack {
			//        m.Destroy()
			root.RemoveChild(m)
			m.SetParent(nil)
		}
		menuStack = nil
		root.PostUpdate()
		return
	}
	// destroy submenu
	for i, m := range menuStack {
		if m == menu {
			for _, s := range menuStack[i:] {
				//    s.Destroy()
				root.RemoveChild(s)
			}
			menuStack = menuStack[:i]
			root.PostUpdate()
			return
		}
	}
}

//
// special type of menu: the MenuBar.
//
// It is a regular widget, but that can spawn menus
//
type MenuBarWidget struct {
	MenuWidget
	hasfocus    bool
	clickonmenu bool
}

func (self *MenuBarWidget) HasFocus(hasfocus bool) {
	self.hasfocus = hasfocus
	if hasfocus == false {
		menuInitiator = nil
		self.activeItem = -1
		self.lastSubActive = -1
	} else {
		hideMenu(nil)
		menuInitiator = self
	}

	//menuStack=append(make([]*MenuWidget,0,0),&(self.MenuWidget))

	self.PostUpdate()
}

func (self *MenuBarWidget) Repaint() {
	var x int32
	rect := sdl.Rect{0, 0, self.width, self.height}
	self.surface.FillRect(&rect, 0xffdddddd)

	renderer, err := sdl.CreateSoftwareRenderer(self.surface)
	if err != nil {
		panic(err)
	}

	for i, item := range self.items {
		w, h := item.WidthHeight()
		w += 10
		if i == self.activeItem || i == self.lastSubActive {
			rect := sdl.Rect{x, 0, w, self.height}
			self.surface.FillRect(&rect, 0xff8888ff)
		}
		surface := item.Repaint(i == self.activeItem || i == self.lastSubActive)
		rectSrc := sdl.Rect{0, 0, w, h}
		rectDst := sdl.Rect{x + 5, 0, w, h}
		surface.Blit(&rectSrc, self.surface, &rectDst)

		x += w
	}
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawLine(0, int(self.height-1), int(self.width-1), int(self.height-1))
}

func (self *MenuBarWidget) MousePressDown(x, y int32, button uint8) {
	self.hasfocus = true
	self.activeItem = -1 // because we close the menu in the main sws.go loop
	self.clickonmenu = false
	var xx int32
	xx = 0
	if y >= 0 && y < self.height {
		for _, item := range self.items {
			w, _ := item.WidthHeight()
			w += 10
			if xx <= x && xx+w > x {
				self.clickonmenu = true
				break
			}
			xx += w
		}
	}

	// to open the corresponding menu
	self.MouseMove(x, y, 0, 0)
}

func (self *MenuBarWidget) MousePressUp(x, y int32, button uint8) {
	if self.activeItem != -1 {
		submenu := self.items[self.activeItem].SubMenu()
		if submenu == nil {
			self.items[self.activeItem].Clicked()
		}
	}
	self.PostUpdate()
}

func (self *MenuBarWidget) MouseMove(x, y, xrel, yrel int32) {
	if self.hasfocus == false || self.clickonmenu == false {
		return
	}
	previousActiveItem := self.activeItem
	self.activeItem = -1
	var xx int32
	xx = 0
	if y >= 0 && y < self.height {
		for i, item := range self.items {
			w, _ := item.WidthHeight()
			w += 10
			if xx <= x && xx+w > x {
				self.activeItem = i
				break
			}
			xx += w
		}
	}
	if previousActiveItem != self.activeItem {
		//hideMenu(self)
		if previousActiveItem != -1 && self.activeItem != -1 {
			submenu := self.items[previousActiveItem].SubMenu()
			if submenu != nil {
				hideMenu(submenu)
			}
		}
		if self.lastSubActive != -1 && self.activeItem != -1 {
			submenu := self.items[self.lastSubActive].SubMenu()
			if submenu != nil {
				hideMenu(submenu)
			}
		}
		if self.activeItem != -1 {
			submenu := self.items[self.activeItem].SubMenu()
			if submenu != nil {
				self.lastSubActive = self.activeItem

				yy := self.height
				var widget Widget
				widget = self
				for widget != nil {
					xx += widget.X()
					yy += widget.Y()
					widget = widget.Parent()
				}
				submenu.Move(xx, yy-2)
				ShowMenu(submenu)
			} else {
				self.lastSubActive = -1
			}
		}
		self.PostUpdate()
	}
}

func (self *MenuBarWidget) AddItem(item MenuItem) {
	self.items = append(self.items, item)
}

func NewMenuBarWidget() *MenuBarWidget {
	menuwidget := NewMenuWidget()
	widget := &MenuBarWidget{MenuWidget: *menuwidget,
		clickonmenu: false,
		hasfocus:    false}
	widget.height = 25
	return widget
}
