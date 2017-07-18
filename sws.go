//
// This a SDL Windowing System for Go
// Other UI toolkit (or binding) exists (Nulkear, Qt, ...), but I
// didn't found one for SDL, so I am developping it for my own need.
//
// It means that this Windowing System is far to be complete, the
// most fast, low-memory that exist, but should be complete enough for me :-).
//
// - The base "class" for all widget is the SWS_Widget interface.
//
// - And the base implementation is SWS_CoreWidget.
//
// - the root widget (background widget) is the SWS_RootWidget.
//
// - and main widget that floats on top of the root widget are SWS_MainWidget.
//
package sws

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/veandco/go-sdl2/img"
	"time"
)

var needUpdate bool

//
// Currently when a widget needs to refresh its content it calls
// this PostUpdate() function, that ask all windows to repaint
// This is not smart, we should at least have in each widget a
// boolean to know if the window is dirty (and its parent)
//
// Would be nice to implement it :-)
//
func PostUpdate() {
	needUpdate = true
}

type DragPayload interface {
	GetType() int32
}

//
// "Abstract" class of all widget
//
type SWS_Widget interface {
	AddChild(child SWS_Widget)
	RemoveChild(child SWS_Widget)
	Move(x, y int32)
	Resize(width, height int32)
	Surface() *sdl.Surface
	Renderer() *sdl.Renderer
	GetChildren() []SWS_Widget
	Parent() SWS_Widget
	SetParent(SWS_Widget)
	Repaint()
	X() int32
	Y() int32
	Width() int32
	Height() int32
	MousePressDown(x, y int32, button uint8)
	MousePressUp(x, y int32, button uint8)
	MouseMove(x, y, xrel, yrel int32)
	KeyDown(key sdl.Keycode, mod uint16)
	KeyUp(key sdl.Keycode, mod uint16)
	TranslateXYToWidget(globalX, globalY int32) (x, y int32)
	IsInside(x, y int32) bool
	HasFocus(focus bool)
	Font() *ttf.Font
	Destroy()
	DragMove(x, y int32, payload DragPayload)
	DragEnter(x,y int32, payload DragPayload)
	DragLeave()
	DragDrop(x,y int32, payload DragPayload)
}

//
// special function to deal with MainWindow focus
//
func findMainWidget(x, y int32, root *SWS_RootWidget) (target SWS_Widget) {
	target = nil

	x -= root.X()
	y -= root.Y()

	// we take the closest
	for _, child := range root.GetChildren() {
		if child==dragwidget { continue }
		maxX := child.X() + child.Width()
		maxY := child.Y() + child.Height()
		if maxX > root.Width() {
			maxX = root.Width()
		}
		if maxY > root.Height() {
			maxY = root.Height()
		}
		if x >= 0 && y >= 0 && x < maxX && y < maxY {
			if child.IsInside(x-child.X(), y-child.Y()) {
				target = child
			}
		}
	}
	if target == nil {
		return nil
	}

	// we found a child
	switch target.(type) {
	case *SWS_MainWidget:
		{
			return target
		}
	}
	return nil
}

//
// function used to find the end widget where the mouse is over
//
func findWidget(x, y int32, root SWS_Widget) (target SWS_Widget, xTarget, yTarget int32) {
	target = nil

	x -= root.X()
	y -= root.Y()

	// we take the closest
	for _, child := range root.GetChildren() {
		if child==dragwidget { continue }
		maxX := child.X() + child.Width()
		maxY := child.Y() + child.Height()
		if maxX > root.Width() {
			maxX = root.Width()
		}
		if maxY > root.Height() {
			maxY = root.Height()
		}
		if x >= 0 && y >= 0 && x < maxX && y < maxY {
			if child.IsInside(x-child.X(), y-child.Y()) {
				target = child
			}
		}
	}

	// we found a child
	if target != nil {
		return findWidget(x, y, target)
	}

	//if (x>=0 && y>=0 && x<root.Width() && y<root.Height()) {
	if root.IsInside(x, y) {
		target = root
		xTarget = x
		yTarget = y
		return
	}

	return nil, -1, -1
}

var root *SWS_RootWidget
var dragwidget SWS_Widget
var dragpayload DragPayload

//
// usually in a MouseButtonDown 
//
func NewDragEvent(x,y int32, image string, payload DragPayload) {
	dragpayload=payload
	
	draglabel:=CreateLabel(25,25,"")
	draglabel.SetColor(0)
	if img,err := img.Load(image); err==nil {
                draglabel.Resize(img.W,img.H)
        }
	draglabel.SetImage(image)
	draglabel.Move(x,y)
	dragwidget=draglabel
	root.AddChild(dragwidget)
}

//
// The SWS_RootWidget is the background widget that fill up all the
// desktop window
//
type SWS_RootWidget struct {
	SWS_CoreWidget
	window        *sdl.Window
	windowsurface *sdl.Surface
	modalwidget   SWS_Widget
}

//
// to put on top of the widget stack a particular widget
// Mainly used for SWS_MainWidget
//
func (self *SWS_RootWidget) RaiseToTop(widget SWS_Widget) {
	self.RemoveChild(widget)
	self.AddChild(widget)
}

func (self *SWS_RootWidget) RemoveChild(child SWS_Widget) {
	self.SWS_CoreWidget.RemoveChild(child)
	if self.modalwidget==child {
		self.modalwidget=nil
		mainwindowfocus=nil
		previousmainwindowfocus=nil
	}
}

func (self *SWS_RootWidget) SetModal(widget *SWS_MainWidget) {
	// just to be sure to raise it
	self.SWS_CoreWidget.RemoveChild(widget)
        self.SWS_CoreWidget.AddChild(widget)
	// set it as modal
	self.modalwidget=widget

	// this is an ugly hack
	mainwindowfocus.HasFocus(false)
	self.modalwidget.HasFocus(true)
	mainwindowfocus=self.modalwidget
}

func (self *SWS_RootWidget) WindowSurface() *sdl.Surface {
	return self.windowsurface
}

func (self *SWS_RootWidget) WindowUpdateSurface() {
	self.window.UpdateSurface()
}

func CreateRootWidget(window *sdl.Window) *SWS_RootWidget {
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	w, h := window.GetSize()

	corewidget := CreateCoreWidget(int32(w), int32(h))
	rootwidget := &SWS_RootWidget{
		SWS_CoreWidget: *corewidget,
		window:         window,
		windowsurface:  surface,
		modalwidget:    nil,
	}

	return rootwidget
}

//
// When we start the program, we must call this function
// to initialize SDL and provide the resulting SWS_RootWidget
//
// The minimum program you can write is:
//    func main() {
//        root := sws.Init(800,600)
//        for sws.PoolEvent() == false {
//        }
//    }
//
func Init(width, height int32) *SWS_RootWidget {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_FULLSCREEN_DESKTOP)
	if err != nil {
		panic(err)
	}
	//defer window.Destroy()

	if err := ttf.Init(); err != nil {
		panic(err)
	}

	InitSprites()
	InitFonts()

	PostUpdate()

	root = CreateRootWidget(window)
	root.SetColor(0xff111111)
	return root
}

//
// When we need to get keyboard event (especially), we need to get
// the focus
//
func (root *SWS_RootWidget) SetFocus(widget SWS_Widget) {
	mainwidget := widget
	for mainwidget.Parent() != nil && mainwidget.Parent() != SWS_Widget(root) {
		mainwidget = mainwidget.Parent()
	}
	if mainwidget != nil {
		switch mainwidget.(type) {
		case *SWS_MainWidget:
			{
				mainwindowfocus = mainwidget
				if previousmainwindowfocus != mainwindowfocus {
					if previousmainwindowfocus != nil {
						previousmainwindowfocus.HasFocus(false)
					}
					if mainwindowfocus != nil {
						mainwindowfocus.HasFocus(true)
						root.RaiseToTop(mainwindowfocus)
					}
					previousmainwindowfocus = mainwindowfocus
				}
			}
		}
	}

	focus = widget
	if previousFocus != focus {
		if previousFocus != nil && previousFocus != mainwindowfocus {
			previousFocus.HasFocus(false)
		}
		if focus != nil {
			focus.HasFocus(true)
		}
		previousFocus = focus
	}
}

var previousFocus, focus SWS_Widget
var previousmainwindowfocus, mainwindowfocus SWS_Widget
var buttonDown = false

//
// main loop event function.
// see func Init(width,height int32) for an example
//
func PoolEvent() bool {
	var quit bool = false
	var xTarget, yTarget int32

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			quit = true
		case *sdl.MouseButtonEvent:
			fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
				t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			//fmt.Println(findWidget(t.X,t.Y,root))
			if t.Type == sdl.MOUSEBUTTONDOWN {
				buttonDown = true

				// if we click outside of a menu -> destroy the menu
				menu := findMenu(t.X, t.Y)

				if menu == nil && root.modalwidget==nil {
					// special case for main window
					mainwindowfocus = findMainWidget(t.X, t.Y, root)
					if previousmainwindowfocus != mainwindowfocus {
						if previousmainwindowfocus != nil {
							previousmainwindowfocus.HasFocus(false)
						}
						if mainwindowfocus != nil {
							mainwindowfocus.HasFocus(true)
							root.RaiseToTop(mainwindowfocus)
						}
						previousmainwindowfocus = mainwindowfocus
					}
				}

				// else find the widget
				focus, xTarget, yTarget = findWidget(t.X, t.Y, root)
				if root.modalwidget!=nil {
					focus, xTarget, yTarget = findWidget(t.X-root.X(), t.Y-root.Y(), root.modalwidget)
				}
				if previousFocus != focus {
					if previousFocus != nil && previousFocus != mainwindowfocus {
						previousFocus.HasFocus(false)
					}
					if focus != nil {
						focus.HasFocus(true)
					}
					previousFocus = focus
				}

				if menu == nil && menuInitiator == nil {
					hideMenu(nil)
				}

				if focus != nil {
					focus.MousePressDown(xTarget, yTarget, t.Button)
				}
			}
			if t.Type == sdl.MOUSEBUTTONUP {
				buttonDown = false
				// if we click outside of a menu -> destroy the menu
				menu := findMenu(t.X, t.Y)
				if menu == nil && menuInitiator == nil {
					hideMenu(nil)
				}

				if focus != nil {
					xTarget, yTarget = focus.TranslateXYToWidget(t.X, t.Y)
					focus.MousePressUp(xTarget, yTarget, t.Button)
				}
				if (dragwidget!=nil) {
					afterW, axTarget, ayTarget := findWidget(t.X, t.Y, root)
					if afterW!=nil {
						afterW.DragDrop(axTarget, ayTarget,dragpayload)
					}
					root.RemoveChild(dragwidget)
					dragwidget=nil
				}
			}
		case *sdl.MouseMotionEvent:
			//                        fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
			//                        t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)

			if t.Type == sdl.MOUSEMOTION {
				if buttonDown == false {
					beforeW, bxTarget, byTarget := findWidget(t.X-t.XRel, t.Y-t.YRel, root)
					afterW, axTarget, ayTarget := findWidget(t.X, t.Y, root)

					if beforeW != afterW {
						if beforeW != nil {
							beforeW.MouseMove(bxTarget+t.XRel, byTarget+t.YRel, t.XRel, t.YRel)
						}
					}
					if afterW != nil {
						afterW.MouseMove(axTarget, ayTarget, t.XRel, t.YRel)
					}
				} else {
					// button down
					if focus!=nil {
						xTarget, yTarget = focus.TranslateXYToWidget(t.X, t.Y)
						focus.MouseMove(xTarget, yTarget, t.XRel, t.YRel)
					}
					// specific case: button is down AND we are dragging something
					if (dragwidget!=nil) {
						beforeW, _, _ := findWidget(t.X-t.XRel, t.Y-t.YRel, root)
						afterW, axTarget, ayTarget := findWidget(t.X, t.Y, root)
						dragwidget.Move(dragwidget.X()+t.XRel,dragwidget.Y()+t.YRel)
						if (beforeW==afterW) && beforeW!=nil {
							afterW.DragMove(axTarget, ayTarget,dragpayload)
						} else {
							if (beforeW!=nil) {
								beforeW.DragLeave()
							}
							if (afterW!=nil) {
								afterW.DragEnter(axTarget, ayTarget,dragpayload)
							}
						}
						PostUpdate()
					}
				}
			}
		case *sdl.KeyDownEvent:
			fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
				t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			if focus != nil {
				focus.KeyDown(t.Keysym.Sym, t.Keysym.Mod)
			}

		case *sdl.KeyUpEvent:
			//fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
			//t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			if focus != nil {
				focus.KeyUp(t.Keysym.Sym, t.Keysym.Mod)
			}

		}
	}
	if needUpdate == true {
		needUpdate = false
		root.Repaint()
		rectSrc := sdl.Rect{0, 0, root.Width(), root.Height()}
		rectDst := sdl.Rect{root.X(), root.Y(), root.Width(), root.Height()}
		root.surface.Blit(&rectSrc, root.WindowSurface(), &rectDst)

		if menuStack != nil {
			for _, m := range menuStack {
				//                fmt.Println("menu display")
				m.Repaint()
				rectSrc := sdl.Rect{0, 0, m.Width(), m.Height()}
				rectDst := sdl.Rect{m.X(), m.Y(), m.Width(), m.Height()}
				m.surface.Blit(&rectSrc, root.WindowSurface(), &rectDst)
			}
		}
		root.WindowUpdateSurface()
	}
	time.Sleep(25 * time.Millisecond)
	TriggerEvents()
	return quit
}
