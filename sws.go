//
// This a SDL Windowing System for Go
// Other UI toolkit (or binding) exists (Nulkear, Qt, ...), but I
// didn't found one for SDL, so I am developping it for my own need.
//
// It means that this Windowing System is far to be complete, the
// most fast, low-memory that exist, but should be complete enough for me :-).
//
// - The base "class" for all widget is the Widget interface.
//
// - And the base implementation is CoreWidget.
//
// - the root widget (background widget) is the RootWidget.
//
// - and main widget that floats on top of the root widget are MainWidget.
//
package sws

import (
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type DragPayload interface {
	GetType() int32
	PayloadAccepted(bool)
}

//
// "Abstract" class of all widget
//
type Widget interface {
	AddChild(child Widget)
	RemoveChild(child Widget)
	Move(x, y int32)
	Resize(width, height int32)
	Surface() *sdl.Surface
	SetAlphaMod(alpha uint8)
	Renderer() *sdl.Renderer
	GetChildren() []Widget
	Parent() Widget
	SetParent(Widget)
	Repaint()
	X() int32
	Y() int32
	Width() int32
	Height() int32
	MouseDoubleClick(x, y int32)
	MousePressDown(x, y int32, button uint8)
	MousePressUp(x, y int32, button uint8)
	MouseMove(x, y, xrel, yrel int32)
	KeyDown(key sdl.Keycode, mod uint16)
	KeyUp(key sdl.Keycode, mod uint16)
	InputText(string)
	TranslateXYToWidget(globalX, globalY int32) (x, y int32)
	IsInside(x, y int32) bool
	Font() *ttf.Font
	Destroy()
	DragMove(x, y int32, payload DragPayload)
	DragEnter(x, y int32, payload DragPayload)
	DragLeave(payload DragPayload)
	DragDrop(x, y int32, payload DragPayload) bool
	IsDirty() bool
	PostUpdate()
	// deal with current focus
	HasFocus(focus bool)
	// to know if this widget receive input (checkbox, input text, slider, button, ...)
	IsInputWidget() bool
	// when we press tab and want to switch to the next input widget
	SetCallbackFocusOnNextInputWidget(callback func(forward bool))
	// when the input widget see its "value" changed
	SetCallbackValueChanged(callback func())
}

//
// special function to deal with MainWindow focus
//
func findMainWidget(x, y int32, root *RootWidget) (target Widget) {
	target = nil

	x -= root.X()
	y -= root.Y()

	// we take the closest
	for _, child := range root.GetChildren() {
		if child == dragwidget {
			continue
		}
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
	case *MainWidget:
		{
			return target
		}
	}
	return nil
}

//
// function used to find the end widget where the mouse is over
//
func findWidget(x, y int32, root Widget) (target Widget, xTarget, yTarget int32) {
	target = nil

	x -= root.X()
	y -= root.Y()

	// we take the closest
	for _, child := range root.GetChildren() {
		if child == dragwidget {
			continue
		}
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

var root *RootWidget
var dragwidget Widget
var dragpayload DragPayload

//
// usually in a MouseButtonDown
//
func NewDragEventSprite(x, y int32, sprite *sdl.Surface, payload DragPayload) {
	dragpayload = payload

	draglabel := NewLabelWidget(25, 25, "")
	draglabel.SetColor(0)
	draglabel.Resize(sprite.W, sprite.H)
	draglabel.SetImageSurface(sprite)
	draglabel.Move(x-draglabel.Width()/2, y-draglabel.Height()/2)
	dragwidget = draglabel
	root.AddChild(dragwidget)

	widget, _, _ := findWidget(x, y, root)
	if widget != nil {
		localx, localy := widget.TranslateXYToWidget(x, y)
		widget.DragEnter(localx, localy, payload)
	}
}

func NewDragEvent(x, y int32, image string, payload DragPayload) {
	if img, err := img.Load(image); err == nil {
		NewDragEventSprite(x, y, img, payload)
	}
}

//
// The RootWidget is the background widget that fill up all the
// desktop window
//
type RootWidget struct {
	CoreWidget
	window        *sdl.Window
	windowsurface *sdl.Surface
	modalwidget   Widget
}

//
// to put on top of the widget stack a particular widget
// Mainly used for MainWidget
//
func (self *RootWidget) RaiseToTop(widget Widget) {
	self.RemoveChild(widget)
	self.AddChild(widget)
}

func (self *RootWidget) RemoveChild(child Widget) {
	self.CoreWidget.RemoveChild(child)
	if self.modalwidget == child {
		self.modalwidget = nil
		mainwindowfocus = nil
		previousmainwindowfocus = nil
	}
}

func (self *RootWidget) SetModal(widget Widget) {
	// just to be sure to raise it
	self.CoreWidget.RemoveChild(widget)
	self.CoreWidget.AddChild(widget)
	// set it as modal
	self.modalwidget = widget

	// this is an ugly hack
	if mainwindowfocus != nil {
		mainwindowfocus.HasFocus(false)
	}
	self.modalwidget.HasFocus(true)
	mainwindowfocus = self.modalwidget
}

func (self *RootWidget) WindowSurface() *sdl.Surface {
	return self.windowsurface
}

func (self *RootWidget) WindowUpdateSurface() {
	self.window.UpdateSurface()
}

func NewRootWidget(window *sdl.Window) *RootWidget {
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	w, h := window.GetSize()

	corewidget := NewCoreWidget(int32(w), int32(h))
	rootwidget := &RootWidget{
		CoreWidget:    *corewidget,
		window:        window,
		windowsurface: surface,
		modalwidget:   nil,
	}

	return rootwidget
}

//
// When we start the program, we must call this function
// to initialize SDL and provide the resulting RootWidget
//
// The minimum program you can write is:
//    func main() {
//        root := sws.Init(800,600)
//        for sws.PoolEvent() == false {
//        }
//    }
//
func Init(width, height int32) *RootWidget {
	sdl.Init(sdl.INIT_EVERYTHING)
	sdl.StartTextInput()
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

	root = NewRootWidget(window)
	root.SetColor(0xff111111)
	return root
}

//
// When we need to get keyboard event (especially), we need to get
// the focus
//
func (root *RootWidget) SetFocus(widget Widget) {
	mainwidget := widget
	for mainwidget.Parent() != nil && mainwidget.Parent() != Widget(&root.CoreWidget) {
		mainwidget = mainwidget.Parent()
	}
	if mainwidget != nil {
		switch mainwidget.(type) {
		case *MainWidget:
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

var previousFocus, focus Widget
var previousmainwindowfocus, mainwindowfocus Widget
var buttonDown = false
var lastleftclick time.Time

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
			//fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
			//	t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			//fmt.Println(findWidget(t.X,t.Y,root))
			if t.Type == sdl.MOUSEBUTTONDOWN {
				buttonDown = true

				// if we click outside of a menu -> destroy the menu
				menu := findMenu(t.X, t.Y)

				if menu == nil && root.modalwidget == nil {
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
				if menu == nil && root.modalwidget != nil {
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
				if dragwidget != nil {
					afterW, axTarget, ayTarget := findWidget(t.X, t.Y, root)
					if afterW != nil {
						dragpayload.PayloadAccepted(afterW.DragDrop(axTarget, ayTarget, dragpayload))
					} else {
						dragpayload.PayloadAccepted(false)
					}
					root.RemoveChild(dragwidget)
					dragwidget = nil
					root.PostUpdate()
				}

				// left double click
				if time.Since(lastleftclick).Seconds() <= 1 && t.Button == sdl.BUTTON_LEFT && focus != nil {
					focus.MouseDoubleClick(xTarget, yTarget)
				}
				if t.Button == sdl.BUTTON_LEFT {
					lastleftclick = time.Now()
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
					if focus != nil {
						xTarget, yTarget = focus.TranslateXYToWidget(t.X, t.Y)
						focus.MouseMove(xTarget, yTarget, t.XRel, t.YRel)
					}
					// specific case: button is down AND we are dragging something
					if dragwidget != nil {
						beforeW, _, _ := findWidget(t.X-t.XRel, t.Y-t.YRel, root)
						afterW, axTarget, ayTarget := findWidget(t.X, t.Y, root)
						dragwidget.Move(dragwidget.X()+t.XRel, dragwidget.Y()+t.YRel)
						if (beforeW == afterW) && beforeW != nil {
							afterW.DragMove(axTarget, ayTarget, dragpayload)
						} else {
							if beforeW != nil {
								beforeW.DragLeave(dragpayload)
							}
							if afterW != nil {
								afterW.DragEnter(axTarget, ayTarget, dragpayload)
							}
						}
						root.PostUpdate()
					}
				}
			}
		case *sdl.KeyboardEvent:
			menu := findMenuForKeyboard()
			if menu != nil {
				if t.State == sdl.PRESSED {
					menu.KeyDown(t.Keysym.Sym, t.Keysym.Mod)
				}
				if t.State == sdl.RELEASED {
					menu.KeyUp(t.Keysym.Sym, t.Keysym.Mod)
				}
				break
			}
			//			fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\tunicode:%d\n",
			//				t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat, t.Keysym.Unicode)
			if focus != nil {
				if t.State == sdl.PRESSED {
					focus.KeyDown(t.Keysym.Sym, t.Keysym.Mod)
				}
				if t.State == sdl.RELEASED {
					focus.KeyUp(t.Keysym.Sym, t.Keysym.Mod)
				}
			}

		case *sdl.TextInputEvent:
			endString := 0
			for i := range t.Text {
				if t.Text[i] == 0 {
					break
				}
				endString++
			}
			focus.InputText(string(t.Text[:endString]))
			//		case *sdl.TextEditingEvent:
			//			fmt.Println("TextEditingEvent")
			//			fmt.Println(t.Text)
		}
	}
	if root.IsDirty() == true {
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
