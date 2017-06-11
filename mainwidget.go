package sws

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

//
// This specific widget is a "main" widget, that float on top of the
// RootWidget, and have a title bar (in grey/yellow)
// This is decorator, in the sense that, it holds a sub-widget (that you can
// customize with SetInnerWidget() )
//
// You can also set (or not) a menu bar (SetMenuBar() )
//
type SWS_MainWidget struct {
	SWS_CoreWidget
	label              string // title
	hasfocus           bool
	expandable         bool // can we full screen
	resizable          bool // can be resized
	inmove             bool // to know if we are currently "in move" state
	Close              func()
	buttonOnClose      bool // to know if we click down on the close button
	cursorInsideClose  bool // to know if we are over the close button
	buttonOnExpand     bool // to know if we click down on the fullscreen button
	cursorInsideExpand bool // to know if we are over the full screen button
	onResize           bool // to know if we are resizing
	subwidget          SWS_Widget
	menubar            *SWS_MenuBarWidget
}

func (self *SWS_MainWidget) SetInnerWidget(widget SWS_Widget) bool {
	if widget == nil {
		return false
	}
	self.RemoveChild(self.subwidget)
	self.subwidget = widget
	self.SWS_CoreWidget.AddChild(widget)
	if self.menubar==nil {
		widget.Move(6, 26)
		widget.Resize(self.Width()-12, self.Height()-32)
	} else {
		widget.Move(6, 26+self.menubar.Height())
		widget.Resize(self.Width()-12, self.Height()-32-self.menubar.Height())
	}
	PostUpdate()
	return true
}

func (self *SWS_MainWidget) SetMenuBar(menubar *SWS_MenuBarWidget) {
	if self.menubar != nil {
		self.RemoveChild(self.menubar)
	}
	self.menubar = menubar
	self.SWS_CoreWidget.AddChild(menubar)
	menubar.Move(6, 26)
	menubar.Resize(self.Width()-12, menubar.Height())
	self.subwidget.Resize(self.Width()-12, self.Height()-32-menubar.Height())
	self.subwidget.Move(6, 26+self.menubar.Height())
	PostUpdate()
}

func (self *SWS_MainWidget) HasFocus(focus bool) {
	if self.hasfocus != focus {
		self.hasfocus = focus
		PostUpdate()
	}
}

func (self *SWS_MainWidget) IsInside(x, y int32) bool {
	if y < 20 {
		wText, _, _ := self.font.SizeUTF8(self.label)
		return x >= 0 && y >= 0 && x < int32(wText)+40
	} else {
		return x >= 0 && y >= 0 && x < self.Width() && y < self.Height()
	}
}

func (self *SWS_MainWidget) Repaint() {
	// paint header
	var solid *sdl.Surface
	var err error
	var headbgcolor = self.bgColor

	var lightcolorR uint8 = 0xff
	var lightcolorG uint8 = 0xff
	var lightcolorB uint8 = 0xff
	var darkcolorR uint8 = 0x88
	var darkcolorG uint8 = 0x88
	var darkcolorB uint8 = 0x88

	if self.hasfocus {
		headbgcolor = 0xfffff10b
		lightcolorR = 0xff
		lightcolorG = 0xf9
		lightcolorB = 0x96
		darkcolorR = 0xbd
		darkcolorG = 0xb2
		darkcolorB = 0x00
	}

	color := sdl.Color{0, 0, 0, 255}
	if self.label == "" {
		self.label = "unnamed"
	}
	if solid, err = self.Font().RenderUTF8_Blended(self.label, color); err != nil {
		fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
	}
	defer solid.Free()

	maxW := solid.W + 40
	if maxW > self.Width() {
		maxW = self.Width()
	}
	self.FillRect(0, 0, maxW, 21, headbgcolor)

	rectSrc := sdl.Rect{0, 0, maxW - 40, solid.H}
	rectDst := sdl.Rect{20, 0, maxW - 40, 20}
	if err = solid.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
	}

	// high bezel
	self.SetDrawColor(darkcolorR, darkcolorG, darkcolorB, 0xff)
	self.DrawLine(0, 0, 0, 20)
	self.DrawLine(0, 0, maxW-1, 0)
	self.DrawLine(maxW-1, 0, maxW-1, 20)
	self.DrawLine(maxW-2, 1, maxW-2, 21)
	self.SetDrawColor(lightcolorR, lightcolorG, lightcolorB, 0xff)
	self.DrawLine(1, 1, 1, 20)
	self.DrawLine(1, 1, maxW-2, 1)

	// low bezel
	self.SetDrawColor(0x88, 0x88, 0x88, 0xff)
	self.DrawLine(0, 20, 0, self.Height()-2)
	self.DrawLine(0, self.Height()-1, self.Width()-1, self.Height()-1)
	self.DrawLine(self.Width()-1, self.Height()-1, self.Width()-1, 20)
	self.DrawLine(self.Width()-1, 20, maxW-1, 20)
	self.SetDrawColor(0xff, 0xff, 0xff, 0xff)
	self.DrawLine(1, 20, 1, self.Height()-2)
	self.DrawLine(self.Width()-2, 21, maxW-2, 21)
	self.SetDrawColor(0xdd, 0xdd, 0xdd, 0xff)
	self.DrawLine(2, 21, maxW-3, 21)
	self.SetDrawColor(0x88, 0x88, 0x88, 0xff)
	self.DrawLine(1, self.Height()-2, self.Width()-2, self.Height()-2)
	self.DrawLine(self.Width()-2, self.Height()-2, self.Width()-2, 22)

	// low bezel interior
	self.SetDrawColor(0xdd, 0xdd, 0xdd, 0xff)
	self.DrawLine(2, 22, self.Width()-3, 22)
	self.DrawLine(self.Width()-3, 22, self.Width()-3, self.Height()-3)
	self.DrawLine(self.Width()-3, self.Height()-3, 2, self.Height()-3)
	self.DrawLine(2, self.Height()-3, 2, 22)
	self.SetDrawColor(0xdd, 0xdd, 0xdd, 0xff)
	self.DrawLine(3, 23, self.Width()-4, 23)
	self.DrawLine(self.Width()-4, 23, self.Width()-4, self.Height()-4)
	self.DrawLine(self.Width()-4, self.Height()-4, 3, self.Height()-4)
	self.DrawLine(3, self.Height()-4, 3, 23)
	self.SetDrawColor(0xbb, 0xbb, 0xbb, 0xff)
	self.DrawLine(4, 24, self.Width()-5, 24)
	self.DrawLine(self.Width()-5, 24, self.Width()-5, self.Height()-5)
	self.DrawLine(self.Width()-5, self.Height()-5, 4, self.Height()-5)
	self.DrawLine(4, self.Height()-5, 4, 24)
	self.SetDrawColor(0x88, 0x88, 0x88, 0xff)
	self.DrawLine(5, 25, self.Width()-6, 25)
	self.DrawLine(self.Width()-6, 25, self.Width()-6, self.Height()-6)
	self.DrawLine(self.Width()-6, self.Height()-6, 4, self.Height()-6)
	self.DrawLine(5, self.Height()-6, 5, 25)

	if self.resizable {
		self.DrawLine(self.Width()-25, self.Height()-6, self.Width()-25, self.Height()-1)
		self.DrawLine(self.Width()-6, self.Height()-25, self.Width()-1, self.Height()-25)
	}

	if self.hasfocus {
		rectSrc := sdl.Rect{0, 0, mainlefths.W, mainlefths.H}
		rectDst := sdl.Rect{3, 3, mainlefths.W, mainlefths.H}
		if self.buttonOnClose && self.cursorInsideClose {
			if mainlefthclickeds.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		} else {
			if mainlefths.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		}
		if self.expandable {
			rectSrc = sdl.Rect{0, 0, mainrighths.W, mainrighths.H}
			rectDst = sdl.Rect{maxW - 19, 3, mainrighths.W, mainrighths.H}
			if self.buttonOnExpand && self.cursorInsideExpand {
				if mainrighthclickeds.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
				}
			} else {
				if mainrighths.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
				}
			}
		}
	} else {
		rectSrc := sdl.Rect{0, 0, mainlefts.W, mainlefts.H}
		rectDst := sdl.Rect{3, 3, mainlefts.W, mainlefts.H}
		if mainlefts.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
		}
		if self.expandable {
			rectSrc = sdl.Rect{0, 0, mainrights.W, mainrights.H}
			rectDst = sdl.Rect{maxW - 19, 3, mainrights.W, mainrights.H}
			if mainrights.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		}
	}

	if self.menubar != nil {
		self.menubar.Repaint()
		rectSrc = sdl.Rect{0, 0, self.menubar.Width(), self.menubar.Height()}
		rectDst = sdl.Rect{self.menubar.X(), self.menubar.Y(), self.menubar.Width(), self.menubar.Height()}
		self.menubar.Surface().Blit(&rectSrc, self.Surface(), &rectDst)
	}
	maxwidth:=self.subwidget.Width()
	if maxwidth>self.Width()-12 { maxwidth=self.Width()-12 }
	maxheight:=self.subwidget.Height()
	if self.menubar!=nil {
		if maxheight>self.Height()-32-self.menubar.Height() { maxheight=self.Height()-32-self.menubar.Height() }
	} else {
		if maxheight>self.Height()-32 { maxheight=self.Height()-32 }
	}
	self.subwidget.Repaint()
	rectSrc = sdl.Rect{0, 0, maxwidth, maxheight}
	rectDst = sdl.Rect{self.subwidget.X(), self.subwidget.Y(), maxwidth, maxheight}
	self.subwidget.Surface().Blit(&rectSrc, self.Surface(), &rectDst)
}

func (self *SWS_MainWidget) AddChild(child SWS_Widget) {
	self.subwidget.AddChild(child)
}

func (self *SWS_MainWidget) MousePressDown(x, y int32, button uint8) {
	wText, _, _ := self.font.SizeUTF8(self.label)
	maxW := int32(wText) + 40
	if maxW > self.Width() {
		maxW = self.Width()
	}
	if x > 2 && x < 18 && y > 2 && y < 18 {
		self.buttonOnClose = true
		self.cursorInsideClose = true
		PostUpdate()
		return
	}
	if self.expandable && x > maxW-19 && x < maxW-3 && y > 2 && y < 18 {
		self.buttonOnExpand = true
		self.cursorInsideExpand = true
		PostUpdate()
		return
	}
	if x < 40+int32(wText) && y < 20 {
		self.inmove = true
	}
	if (x >= self.Width()-25 && y >= self.Height()-6) || (x >= self.Width()-6 && y >= self.Height()-25) {
		self.onResize = true
	}
}

func (self *SWS_MainWidget) MousePressUp(x, y int32, button uint8) {
	self.onResize = false
	self.inmove = false
	if self.buttonOnClose == true {
		self.buttonOnClose = false
		PostUpdate()
	}
	if self.buttonOnExpand == true {
		self.buttonOnExpand = false
		PostUpdate()
	}
}

func (self *SWS_MainWidget) MouseMove(x, y, xrel, yrel int32) {

	if self.inmove {
		self.x += xrel
		self.y += yrel
		PostUpdate()
		return
	}
	if self.onResize {
		self.Resize(x, y)
		return
	}
	wText, _, _ := self.font.SizeUTF8(self.label)
	maxW := int32(wText) + 40
	if maxW > self.Width() {
		maxW = self.Width()
	}

	if self.buttonOnClose {
		if x > 2 && x < 18 && y > 2 && y < 18 {
			self.cursorInsideClose = true
		} else {
			self.cursorInsideClose = false
		}
		PostUpdate()
	}
	if self.buttonOnExpand {
		if x > maxW-19 && x < maxW-3 && y > 2 && y < 18 {
			self.cursorInsideExpand = true
		} else {
			self.cursorInsideExpand = false
		}
		PostUpdate()
	}
}

func (self *SWS_MainWidget) Resize(width, height int32) {
	if width < 60 {
		width = 60
	}
	if height < 80 {
		height = 80
	}
	self.SWS_CoreWidget.Resize(width, height)
	if self.menubar == nil {
		self.subwidget.Resize(width-12, height-32)
	} else {
		self.menubar.Resize(width-12, self.menubar.Height())
		self.subwidget.Resize(width-12, height-32-self.menubar.Height())
	}
	PostUpdate()
}

func CreateMainWidget(w, h int32, s string, expandable bool, resizable bool) *SWS_MainWidget {
	corewidget := CreateCoreWidget(w, h)
	subwidget := CreateCoreWidget(w-12, h-32)
	subwidget.Move(6, 26)
	corewidget.AddChild(subwidget)
	widget := &SWS_MainWidget{SWS_CoreWidget: *corewidget,
		label:              s,
		hasfocus:           false,
		expandable:         expandable,
		resizable:          resizable,
		inmove:             false,
		buttonOnClose:      false,
		cursorInsideClose:  false,
		buttonOnExpand:     false,
		cursorInsideExpand: false,
		onResize:           false,
		subwidget:          subwidget,
		menubar:            nil}
	subwidget.SetParent(widget)
	return widget
}
