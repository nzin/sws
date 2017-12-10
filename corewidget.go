package sws

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//
// main "abstract" class that implement the Widget interface
// Mostly it is used as a base class for all other widgets
//
// One comment: each CoreWidget (and derivated), keep a
// bitmap of its content (into a surface *sdl.Surface) member
// It allows to compose sub element into main element easier, and
// to provide cache when we need to refresh the content.
//
type CoreWidget struct {
	surface                        *sdl.Surface
	renderer                       *sdl.Renderer
	children                       []Widget
	parent                         Widget
	bgColor                        uint32
	x                              int32
	y                              int32
	width                          int32
	height                         int32
	font                           *ttf.Font
	dirty                          bool
	focusOnNextInputWidgetCallback func(forward bool)
}

func NewCoreWidget(w, h int32) *CoreWidget {
	surface, err := sdl.CreateRGBSurface(0, w, h, 32, 0x00ff0000, 0x0000ff00, 0x000000ff, 0xff000000)
	if err != nil {
		panic(err)
	}
	renderer, err := sdl.CreateSoftwareRenderer(surface)
	if err != nil {
		panic(err)
	}

	widget := &CoreWidget{x: 0,
		y:        0,
		width:    w,
		height:   h,
		bgColor:  0xffdddddd,
		surface:  surface,
		renderer: renderer,
		font:     defaultFont,
		dirty:    true,
		parent:   nil,
	}
	return widget
}

//
// method used (internaly) to specify that this
// widget's content changed and needs to be refreshed.
//
func (self *CoreWidget) PostUpdate() {
	self.dirty = true
	if self.Parent() != nil && self.Parent() != self {
		self.Parent().PostUpdate()
	}
}

//
// to know if the widget has to be repaint (i.e.
// PostUpdate() has been called)
//
func (self *CoreWidget) IsDirty() bool {
	return self.dirty
}

func (self *CoreWidget) Destroy() {
	self.Parent().RemoveChild(self)
	self.surface.Free()
}

//
// when we resize we need to
// destroy the current surface, recreate it, and
// trigger a PostUpdate() to ask for a refresh
//
func (self *CoreWidget) Resize(width, height int32) {
	if width < 0 {
		width = 0
	}
	if height < 0 {
		height = 0
	}
	surface, err := sdl.CreateRGBSurface(0, width, height, 32, 0x00ff0000, 0x0000ff00, 0x000000ff, 0xff000000)
	if err != nil {
		panic(err)
	}
	renderer, err := sdl.CreateSoftwareRenderer(surface)
	if err != nil {
		panic(err)
	}
	self.surface.Free()
	self.surface = surface
	self.renderer.Destroy()
	self.renderer = renderer
	self.width = width
	self.height = height
	self.PostUpdate()
}

func (self *CoreWidget) FillRect(x, y, w, h int32, c uint32) {
	surface := self.Surface()

	rect := sdl.Rect{x, y, w, h}
	surface.FillRect(&rect, c)
}

//
// Write text into the surface
// return (width,height) written
//
func (self *CoreWidget) WriteText(x, y int32, str string, color sdl.Color) (int32, int32) {
	var solid *sdl.Surface
	var err error

	if str == "" {
		return 0, 0
	}

	if solid, err = self.Font().RenderUTF8Blended(str, color); err != nil {
		fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
		return 0, 0
	}
	defer solid.Free()
	rectSrc := sdl.Rect{0, 0, solid.W, solid.H}
	rectDst := sdl.Rect{x, y, self.Width(), self.Height()}
	if err = solid.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
		//fmt.Fprint(os.Stderr, "Failed to put text on window surface: %s\n", err)
		return 0, 0
	}

	return solid.W, solid.H
}

//
// Write text into the surface but centered
//
func (self *CoreWidget) WriteTextCenter(x, y int32, str string, color sdl.Color) {
	var solid *sdl.Surface
	var err error

	if solid, err = self.Font().RenderUTF8Blended(str, color); err != nil {
		fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
		return
	}
	rectSrc := sdl.Rect{0, 0, solid.W, solid.H}
	wGap := self.Width() - solid.W
	hGap := self.Height() - solid.H
	rectDst := sdl.Rect{x + (wGap / 2), y + (hGap / 2), self.Width(), self.Height()}
	if err = solid.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
		//fmt.Fprint(os.Stderr, "Failed to put text on window surface: %s\n", err)
		return
	}

	defer solid.Free()
}

func (self *CoreWidget) DrawLine(x1, y1, x2, y2 int32) {
	self.Renderer().DrawLine(x1, y1, x2, y2)
}

//
// the color should be 0xrrggbbaa (r=red, g=green, b=blue, a=alpha)
//
func (self *CoreWidget) SetDrawColorHex(color uint32) {
	b := color & 0xff
	g := (color >> 8) & 0xff
	r := (color >> 16) & 0xff
	a := (color >> 24) & 0xff
	self.Renderer().SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
}

func (self *CoreWidget) SetDrawColor(r, g, b, c uint8) {
	self.Renderer().SetDrawColor(r, g, b, c)
}

func (self *CoreWidget) DrawPoint(x, y int32) {
	self.Renderer().DrawPoint(x, y)
}

func (self *CoreWidget) InputText(text string) {

}

func (self *CoreWidget) KeyDown(key sdl.Keycode, mod uint16) {
}

func (self *CoreWidget) KeyUp(key sdl.Keycode, mod uint16) {
}

func (self *CoreWidget) Font() *ttf.Font {
	return self.font
}

func (self *CoreWidget) SetFont(font *ttf.Font) {
	self.font = font
}

func (self *CoreWidget) Surface() *sdl.Surface {
	return self.surface
}

func (self *CoreWidget) Renderer() *sdl.Renderer {
	return self.renderer
}

func (self *CoreWidget) RemoveChild(child Widget) {
	//fmt.Println("todestroy:", child)
	for i, c := range self.children {
		//fmt.Println("child:", c)
		if c == child {
			//fmt.Println("found", child)
			if i == 0 {
				self.children = self.children[1:]
			} else {
				self.children = append(self.children[:i], self.children[i+1:]...)
			}
			self.PostUpdate()
			return
		}
	}
}

func (self *CoreWidget) AddChild(child Widget) {
	// we don't want to add it twice
	for i, c := range self.children {
		if c == child {
			if i == 0 {
				self.children = self.children[1:]
			} else {
				self.children = append(self.children[:i], self.children[i+1:]...)
			}
		}
	}
	self.children = append(self.children, child)
	child.SetParent(self)
	if child.IsInputWidget() {
		child.SetCallbackFocusOnNextInputWidget(func(forward bool) {
			if forward == true {
				self.selectNextInputWidget(child)
			} else {
				self.selectPreviousInputWidget(child)
			}
		})
	}
	self.PostUpdate()
}

func (self *CoreWidget) SetParent(father Widget) {
	self.parent = father
}

func (self *CoreWidget) Parent() Widget {
	return self.parent
}

func (self *CoreWidget) GetChildren() []Widget {
	return self.children
}

//
// This function seems to be obvious, but if we want to
// create widget that are not squared (like MainWindow), the
// calculation is a bit more complicated
//
func (self *CoreWidget) IsInside(x, y int32) bool {
	return x >= 0 && y >= 0 && x < self.Width() && y < self.Height()
}

func (self *CoreWidget) TranslateXYToWidget(globalX, globalY int32) (x, y int32) {
	if self.Parent() == nil {
		return globalX - self.X(), globalY - self.Y()
	}
	return self.Parent().TranslateXYToWidget(globalX-self.X(), globalY-self.Y())
}

func (self *CoreWidget) MouseDoubleClick(x, y int32) {
	if self.Parent() != nil {
		self.Parent().MouseDoubleClick(x+self.X(), y+self.Y())
	}
}

func (self *CoreWidget) MousePressDown(x, y int32, button uint8) {
	if self.Parent() != nil {
		self.Parent().MousePressDown(x+self.X(), y+self.Y(), button)
	}
}

func (self *CoreWidget) MousePressUp(x, y int32, button uint8) {
	if self.Parent() != nil {
		self.Parent().MousePressUp(x+self.X(), y+self.Y(), button)
	}
}

func (self *CoreWidget) MouseMove(x, y, xrel, yrel int32) {
	if self.Parent() != nil {
		self.Parent().MouseMove(x+self.X(), y+self.Y(), xrel, yrel)
	}
}

//
// X() and Y() correspond to the position of this widget into its parent
// widget
//
func (self *CoreWidget) X() int32 {
	return self.x
}

//
// X() and Y() correspond to the position of this widget into its parent
// widget
//
func (self *CoreWidget) Y() int32 {
	return self.y
}

func (self *CoreWidget) Width() int32 {
	return self.width
}

func (self *CoreWidget) Height() int32 {
	return self.height
}

func (self *CoreWidget) SetColor(color uint32) {
	self.bgColor = color
	self.PostUpdate()
}

func (self *CoreWidget) Move(x, y int32) {
	self.x = x
	self.y = y
}

func (self *CoreWidget) DragMove(x, y int32, payload DragPayload) {
}

func (self *CoreWidget) DragEnter(x, y int32, payload DragPayload) {
}

func (self *CoreWidget) DragLeave(payload DragPayload) {
}

func (self *CoreWidget) DragDrop(x, y int32, payload DragPayload) bool {
	return false
}

func (self *CoreWidget) HasFocus(focus bool) {
}

func (self *CoreWidget) IsInputWidget() bool {
	return false
}

func (self *CoreWidget) SetCallbackFocusOnNextInputWidget(callback func(forward bool)) {
	self.focusOnNextInputWidgetCallback = callback
}

func (self *CoreWidget) selectNextInputWidget(current Widget) {
	nextChildren := make([]Widget, len(self.children))
	copy(nextChildren, self.children)
	for _, c := range self.children {
		if c != current {
			// we shift the array until we find 'current'
			nextChildren = append(nextChildren[1:], c)
		} else {
			// we unroll the array to the next inputwidget
			for _, n := range nextChildren[1:] {
				if n.IsInputWidget() {
					root.SetFocus(n)
					return
				}
			}
			return
		}
	}
}

func (self *CoreWidget) selectPreviousInputWidget(current Widget) {
	nextChildren := make([]Widget, len(self.children))
	copy(nextChildren, self.children)
	for _, c := range self.children {
		if c != current {
			// we shift the array until we find 'current'
			nextChildren = append(nextChildren[1:], c)
		} else {
			// we unroll the array to the next inputwidget
			for i := len(nextChildren) - 1; i >= 0; i-- {
				n := nextChildren[i]
				if n.IsInputWidget() {
					root.SetFocus(n)
					return
				}
			}
			return
		}
	}
}

//
// One of the main window: how do we want to write to our widget
// content
//
func (self *CoreWidget) Repaint() {
	if self.bgColor != 0 {
		self.FillRect(0, 0, self.width, self.height, self.bgColor)
	} else {
		surface, err := sdl.CreateRGBSurface(0, self.Surface().W, self.Surface().H, 32, 0x00ff0000, 0x0000ff00, 0x000000ff, 0xff000000)
		if err != nil {
			panic(err)
		}
		if self.surface != nil {
			self.surface.Free()
		}
		self.surface = surface
	}
	for _, child := range self.children {
		// adjust the clipping to the current child
		if child.IsDirty() {
			child.Repaint()
		}
		rectSrc := sdl.Rect{0, 0, child.Width(), child.Height()}
		rectDst := sdl.Rect{child.X(), child.Y(), child.Width(), child.Height()}
		child.Surface().Blit(&rectSrc, self.Surface(), &rectDst)
	}
	self.dirty = false
}
