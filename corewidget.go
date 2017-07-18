package sws

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
)

//
// main "abstract" class that implement the SWS_Widget interface
// Mostly it is used as a base class for all other widgets
//
// One comment: each SWS_CoreWidget (and derivated), keep a
// bitmap of its content (into a surface *sdl.Surface) member
// It allows to compose sub element into main element easier, and
// to provide cache when we need to refresh the content.
//
type SWS_CoreWidget struct {
	surface  *sdl.Surface
	renderer *sdl.Renderer
	children []SWS_Widget
	parent   SWS_Widget
	bgColor  uint32
	x        int32
	y        int32
	width    int32
	height   int32
	font     *ttf.Font
}

func CreateCoreWidget(w, h int32) *SWS_CoreWidget {
	surface, err := sdl.CreateRGBSurface(0, w, h, 32, 0x00ff0000, 0x0000ff00, 0x000000ff, 0xff000000)
	if err != nil {
		panic(err)
	}
	renderer, err := sdl.CreateSoftwareRenderer(surface)
	if err != nil {
		panic(err)
	}

	widget := &SWS_CoreWidget{x: 0,
		y:        0,
		width:    w,
		height:   h,
		bgColor:  0xffdddddd,
		surface:  surface,
		renderer: renderer,
		font:     defaultFont}
	return widget
}

func (self *SWS_CoreWidget) Destroy() {
	self.Parent().RemoveChild(self)
	self.surface.Free()
}

//
// when we resize we need to
// destroy the current surface, recreate it, and
// trigger a PostUpdate() to ask for a refresh
//
func (self *SWS_CoreWidget) Resize(width, height int32) {
	if width<0 { width=0}
	if height<0 { height=0}
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
	PostUpdate()
}

func (self *SWS_CoreWidget) FillRect(x, y, w, h int32, c uint32) {
	surface := self.Surface()

	rect := sdl.Rect{x, y, w, h}
	surface.FillRect(&rect, c)
}

//
// Write text into the surface
//
func (self *SWS_CoreWidget) WriteText(x, y int32, str string, color sdl.Color) (int32, int32) {
	var solid *sdl.Surface
	var err error

	if str == "" {
		return 0, 0
	}

	if solid, err = self.Font().RenderUTF8_Blended(str, color); err != nil {
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
func (self *SWS_CoreWidget) WriteTextCenter(x, y int32, str string, color sdl.Color) {
	var solid *sdl.Surface
	var err error

	if solid, err = self.Font().RenderUTF8_Blended(str, color); err != nil {
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

func (self *SWS_CoreWidget) DrawLine(x1, y1, x2, y2 int32) {
	self.Renderer().DrawLine(int(x1), int(y1), int(x2), int(y2))
}

//
// the color should be 0xrrggbbaa (r=red, g=green, b=blue, a=alpha)
//
func (self *SWS_CoreWidget) SetDrawColorHex(color uint32) {
	b := color & 0xff
	g := (color >> 8) & 0xff
	r := (color >> 16) & 0xff
	a := (color >> 24) & 0xff
	self.Renderer().SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
}

func (self *SWS_CoreWidget) SetDrawColor(r, g, b, c uint8) {
	self.Renderer().SetDrawColor(r, g, b, c)
}

func (self *SWS_CoreWidget) DrawPoint(x, y int32) {
	self.Renderer().DrawPoint(int(x), int(y))
}

func (self *SWS_CoreWidget) KeyDown(key sdl.Keycode, mod uint16) {
}

func (self *SWS_CoreWidget) KeyUp(key sdl.Keycode, mod uint16) {
}

func (self *SWS_CoreWidget) Font() *ttf.Font {
	return self.font
}

func (self *SWS_CoreWidget) SetFont(font *ttf.Font) {
	self.font=font
}

func (self *SWS_CoreWidget) HasFocus(focus bool) {
}

func (self *SWS_CoreWidget) Surface() *sdl.Surface {
	return self.surface
}

func (self *SWS_CoreWidget) Renderer() *sdl.Renderer {
	return self.renderer
}

func (self *SWS_CoreWidget) RemoveChild(child SWS_Widget) {

	//fmt.Println("todestroy:",child)
	for i, c := range self.children {
		//fmt.Println("child:",c)
		if c == child {
			//fmt.Println("found")
			if i == 0 {
				self.children = self.children[1:]
			} else {
				self.children = append(self.children[:i], self.children[i+1:]...)
			}
			return
		}
	}
}

func (self *SWS_CoreWidget) AddChild(child SWS_Widget) {
	// we don't want to add it twice
	for i,c:= range self.children {
		if c==child {
			if i == 0 {
				self.children = self.children[1:]
			} else {
				self.children = append(self.children[:i], self.children[i+1:]...)
			}
		}
	} 
	self.children = append(self.children, child)
	child.SetParent(self)
}

func (self *SWS_CoreWidget) SetParent(father SWS_Widget) {
	self.parent = father
}

func (self *SWS_CoreWidget) Parent() SWS_Widget {
	return self.parent
}

func (self *SWS_CoreWidget) GetChildren() []SWS_Widget {
	return self.children
}

//
// This function seems to be obvious, but if we want to
// create widget that are not squared (like SWS_MainWindow), the
// calculation is a bit more complicated
//
func (self *SWS_CoreWidget) IsInside(x, y int32) bool {
	return x >= 0 && y >= 0 && x < self.Width() && y < self.Height()
}

func (self *SWS_CoreWidget) TranslateXYToWidget(globalX, globalY int32) (x, y int32) {
	if self.Parent() == nil {
		return globalX - self.X(), globalY - self.Y()
	}
	return self.Parent().TranslateXYToWidget(globalX-self.X(), globalY-self.Y())
}

func (self *SWS_CoreWidget) MousePressDown(x, y int32, button uint8) {
}

func (self *SWS_CoreWidget) MousePressUp(x, y int32, button uint8) {
}

func (self *SWS_CoreWidget) MouseMove(x, y, xrel, yrel int32) {
}

//
// X() and Y() correspond to the position of this widget into its parent
// widget
//
func (self *SWS_CoreWidget) X() int32 {
	return self.x
}

//
// X() and Y() correspond to the position of this widget into its parent
// widget
//
func (self *SWS_CoreWidget) Y() int32 {
	return self.y
}

func (self *SWS_CoreWidget) Width() int32 {
	return self.width
}

func (self *SWS_CoreWidget) Height() int32 {
	return self.height
}

func (self *SWS_CoreWidget) SetColor(color uint32) {
	self.bgColor = color
	PostUpdate()
}

func (self *SWS_CoreWidget) Move(x, y int32) {
	self.x = x
	self.y = y
}

func (self *SWS_CoreWidget) DragMove(x, y int32, payload DragPayload) {
}

func (self *SWS_CoreWidget) DragEnter(x,y int32, payload DragPayload) {
}

func (self *SWS_CoreWidget) DragLeave() {
}

func (self *SWS_CoreWidget) DragDrop(x,y int32, payload DragPayload) {
}


//
// One of the main window: how do we want to write to our widget
// content
//
func (self *SWS_CoreWidget) Repaint() {
	if self.bgColor != 0 {
		self.FillRect(0, 0, self.width, self.height, self.bgColor)
	}
	for _, child := range self.children {
		// adjust the clipping to the current child
		child.Repaint()
		rectSrc := sdl.Rect{0, 0, child.Width(), child.Height()}
		rectDst := sdl.Rect{child.X(), child.Y(), child.Width(), child.Height()}
		child.Surface().Blit(&rectSrc, self.Surface(), &rectDst)
	}
}
