package sws

// https://medium.com/random-go-tips/method-overriding-680cfd51ce40

/*
 * clipping: chaque widget ecrit dans sa fenetre
 * on copie la surface sur le widget pere
 */

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "fmt"
    "time"
)

var defaultFont *ttf.Font
var needUpdate bool


func PostUpdate() {
    needUpdate = true
}



/*
 * "Abstract" class of all widget
 */
type SWS_Widget interface {
    AddChild(child SWS_Widget)
    RemoveChild(child SWS_Widget)
    Move(x,y int32)
    Resize(width,height int32)
    Surface() *sdl.Surface
    Renderer() *sdl.Renderer
    getChildren() []SWS_Widget
    Parent() SWS_Widget
    SetParent(SWS_Widget)
    Repaint()
    X() int32
    Y() int32
    Width() int32
    Height() int32
    MousePressDown(x,y int32, button uint8)
    MousePressUp(x,y int32, button uint8)
    MouseMove(x,y,xrel,yrel int32)
    KeyDown(key sdl.Keycode, mod uint16)
    KeyUp(key sdl.Keycode, mod uint16)
    TranslateXYToWidget(globalX,globalY int32) (x,y int32)
    IsInside(x,y int32) bool
    HasFocus(focus bool)
    Font() *ttf.Font
    Destroy()
}



/*
 * special function to deal with MainWindow focus
 */
func findMainWidget(x,y int32, root *SWS_RootWidget) (target SWS_Widget)  {
    target=nil
    
    x-=root.X()
    y-=root.Y()

    // we take the closest
    for _,child := range root.getChildren() {
        maxX:=child.X()+child.Width()
        maxY:=child.Y()+child.Height()
        if (maxX>root.Width()) { maxX = root.Width() }
        if (maxY>root.Height()) { maxY = root.Height() }
        if (x>=0 && y>=0 && x<maxX && y<maxY) {
            if (child.IsInside(x-child.X(),y-child.Y())) {
                target=child
            }
        }
    }
    if target==nil {
        return nil
    }
    
    // we found a child
    switch target.(type) {
      case *SWS_MainWidget: {
        return target
      }
    }
    return nil
}



func findWidget(x,y int32, root SWS_Widget) (target SWS_Widget, xTarget,yTarget int32)  {
    target=nil
    
    x-=root.X()
    y-=root.Y()


    // we take the closest
    for _,child := range root.getChildren() {
        maxX:=child.X()+child.Width()
        maxY:=child.Y()+child.Height()
        if (maxX>root.Width()) { maxX = root.Width() }
        if (maxY>root.Height()) { maxY = root.Height() }
        if (x>=0 && y>=0 && x<maxX && y<maxY) {
            if (child.IsInside(x-child.X(),y-child.Y())) {
                target=child
            }
        }
    }
    
    // we found a child
    if (target!=nil) {
        return findWidget(x,y,target)
    }
    
    //if (x>=0 && y>=0 && x<root.Width() && y<root.Height()) {
    if (root.IsInside(x,y)) {
        target=root
        xTarget=x
        yTarget=y
        return
    }

    return nil,-1,-1
}



var root *SWS_RootWidget



type SWS_RootWidget struct {
    SWS_CoreWidget
    window        *sdl.Window
    windowsurface *sdl.Surface
}



func (self *SWS_RootWidget) RaiseToTop(widget SWS_Widget) {
    self.RemoveChild(widget)
    self.AddChild(widget)
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
    
    w,h := window.GetSize()

    corewidget := CreateCoreWidget(int32(w),int32(h))
    rootwidget := &SWS_RootWidget {
         SWS_CoreWidget: *corewidget,
         window:         window,
         windowsurface:  surface }
    
    return rootwidget
}



func Init(width,height int32) (*SWS_RootWidget) {
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

    if defaultFont, err = ttf.OpenFont("Lato-Regular.ttf", 16); err != nil {
    //if p.font, err = ttf.OpenFont("Arial.ttf", 16); err != nil {
        panic(err)
    }
    
    InitSprites()

    PostUpdate()
    
    root=CreateRootWidget(window)
    root.SetColor(0xff111111)
    return root
}



func (root *SWS_RootWidget) SetFocus(widget SWS_Widget) {
    mainwidget:=widget
    for mainwidget.Parent()!=nil && mainwidget.Parent()!=SWS_Widget(root) {
        mainwidget=mainwidget.Parent()
    }
    if mainwidget!=nil {
        switch mainwidget.(type) {
            case *SWS_MainWidget: {
                mainwindowfocus=mainwidget
                if (previousmainwindowfocus!=mainwindowfocus) {
                    if previousmainwindowfocus!=nil {
                        previousmainwindowfocus.HasFocus(false)
                    }
                    if mainwindowfocus!=nil {
                        mainwindowfocus.HasFocus(true)
                        root.RaiseToTop(mainwindowfocus)
                    }
                    previousmainwindowfocus=mainwindowfocus
                }
            }
        }
    }

    focus=widget
    if previousFocus!=focus {
        if previousFocus!=nil && previousFocus!=mainwindowfocus {
            previousFocus.HasFocus(false)
        }
        if focus!=nil {
            focus.HasFocus(true)
        }
        previousFocus=focus
    }
}



var previousFocus,focus SWS_Widget
var previousmainwindowfocus,mainwindowfocus SWS_Widget
var buttonDown = false

func PoolEvent() (bool) {
    var quit bool = false
    var xTarget,yTarget int32

    for event := sdl.PollEvent(); event != nil ; event = sdl.PollEvent() {
        switch t := event.(type) {
            case *sdl.QuitEvent:
                quit = true
            case *sdl.MouseButtonEvent:
                fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
                t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
                //fmt.Println(findWidget(t.X,t.Y,root))
                if t.Type == sdl.MOUSEBUTTONDOWN {
                    buttonDown=true

                    // if we click outside of a menu -> destroy the menu
                    menu := findMenu(t.X,t.Y)
                    if menu==nil && currentMenuBar==nil {
                        hideMenu(nil)
                    }
                            
                    // special case for main window
                    mainwindowfocus = findMainWidget(t.X,t.Y,root) 
                    if (previousmainwindowfocus!=mainwindowfocus) {
                        if previousmainwindowfocus!=nil {
                            previousmainwindowfocus.HasFocus(false)
                        }
			if mainwindowfocus!=nil {
                            mainwindowfocus.HasFocus(true)
                            root.RaiseToTop(mainwindowfocus)
                        }
                        previousmainwindowfocus=mainwindowfocus
                    }

                    // else find the widget
                    focus, xTarget,yTarget = findWidget(t.X,t.Y,root) 
                    if previousFocus!=focus {
                        if previousFocus!=nil && previousFocus!=mainwindowfocus {
                            previousFocus.HasFocus(false)
                        }
                        if focus!=nil {
                            focus.HasFocus(true)
                        }
                        previousFocus=focus
                    }
                    if focus != nil {
                        focus.MousePressDown(xTarget,yTarget,t.Button)
                    }
                }
                if t.Type == sdl.MOUSEBUTTONUP {
                    buttonDown=false
                    // if we click outside of a menu -> destroy the menu
                    menu := findMenu(t.X,t.Y)
                    if menu==nil && currentMenuBar==nil {
                        hideMenu(nil)
                    }
                            
                    if focus != nil {
                        xTarget,yTarget=focus.TranslateXYToWidget(t.X,t.Y)
                        focus.MousePressUp(xTarget,yTarget,t.Button)
                    }
                }
            case *sdl.MouseMotionEvent:
//                        fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
//                        t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)

                if t.Type == sdl.MOUSEMOTION {
                    if buttonDown==false {
                        beforeW, bxTarget,byTarget := findWidget(t.X-t.XRel,t.Y-t.YRel,root) 
                        afterW, axTarget,ayTarget := findWidget(t.X,t.Y,root) 
                            
                        if (beforeW!=afterW) {
                            if (beforeW!=nil) {
                                beforeW.MouseMove(bxTarget+t.XRel,byTarget+t.YRel,t.XRel,t.YRel)
                            }
                        }
                        if (afterW!=nil) {
                            afterW.MouseMove(axTarget,ayTarget,t.XRel,t.YRel)
                        }
                    } else { // button down
                        xTarget,yTarget=focus.TranslateXYToWidget(t.X,t.Y)
                        focus.MouseMove(xTarget,yTarget,t.XRel,t.YRel)
                    }
                }
            case *sdl.KeyDownEvent:
                fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
                t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
                if focus != nil {
                    focus.KeyDown(t.Keysym.Sym,t.Keysym.Mod)
                }

            case *sdl.KeyUpEvent:
                //fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
                //t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
                if focus != nil {
                    focus.KeyUp(t.Keysym.Sym,t.Keysym.Mod)
                }

        }
    }
    if needUpdate == true {
        needUpdate = false
        root.Repaint()
        rectSrc := sdl.Rect{0,0, root.Width(),root.Height()}
        rectDst := sdl.Rect{root.X(), root.Y(), root.Width(),root.Height()}
        root.surface.Blit(&rectSrc,root.WindowSurface(),&rectDst)
                
        if (menuStack!=nil) {
            for _,m := range menuStack {
//                fmt.Println("menu display")
                m.Repaint()
                rectSrc := sdl.Rect{0,0, m.Width(),m.Height()}
                rectDst := sdl.Rect{m.X(), m.Y(), m.Width(),m.Height()}
                m.surface.Blit(&rectSrc,root.WindowSurface(),&rectDst)
            }
        }
        root.WindowUpdateSurface()
    }
    time.Sleep(25 * time.Millisecond)
    TriggerEvents()
    return quit
}

