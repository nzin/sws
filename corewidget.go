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
    "os"
)



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



func CreateCoreWidget(w,h int32) *SWS_CoreWidget {
    surface,err := sdl.CreateRGBSurface(0,w,h,32,0x00ff0000,0x0000ff00,0x000000ff,0xff000000)
    if err!=nil {
        panic(err)
    }
    renderer,err := sdl.CreateSoftwareRenderer(surface)
    if err!=nil {
        panic(err)
    }
    
    widget := &SWS_CoreWidget{ x: 0,
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



func (self *SWS_CoreWidget) FillRect(x,y,w,h int32, c uint32) {
    surface := self.Surface()

    rect := sdl.Rect{x,y, w, h}
    surface.FillRect(&rect, c)
}



func (self *SWS_CoreWidget) WriteText(x,y int32, str string,color sdl.Color) (int32,int32) {
    var solid *sdl.Surface
    var err error

    if str == "" {
        return 0,0
    }

    if solid, err = self.Font().RenderUTF8_Blended(str, color); err != nil {
        fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
        return 0,0
    }
    defer solid.Free()
    rectSrc := sdl.Rect{0,0, solid.W,solid.H}
    rectDst := sdl.Rect{x,y, self.Width(),self.Height()}
    if err = solid.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
        //fmt.Fprint(os.Stderr, "Failed to put text on window surface: %s\n", err)
        return 0,0
    }

    return solid.W,solid.H
}



func (self *SWS_CoreWidget) WriteTextCenter(x,y int32, str string, color sdl.Color) {
    var solid *sdl.Surface
    var err error

    if solid, err = self.Font().RenderUTF8_Blended(str, color); err != nil {
        fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
        return
    }
    rectSrc := sdl.Rect{0,0, solid.W,solid.H}
    wGap:=self.Width()-solid.W
    hGap:=self.Height()-solid.H
    rectDst := sdl.Rect{x+(wGap/2),y+(hGap/2), self.Width(),self.Height()}
    if err = solid.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
        //fmt.Fprint(os.Stderr, "Failed to put text on window surface: %s\n", err)
        return
    }

    defer solid.Free()
}



func (self *SWS_CoreWidget) DrawLine(x1, y1, x2, y2 int32) {
    self.Renderer().DrawLine(int(x1), int(y1), int(x2), int(y2))
}



func (self *SWS_CoreWidget) SetDrawColor(r, g, b, c uint8) {
    self.Renderer().SetDrawColor(r,g,b,c)
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



func (self *SWS_CoreWidget) HasFocus(focus bool) {
}



func (self *SWS_CoreWidget) Surface() *sdl.Surface {
    return self.surface
}



func (self *SWS_CoreWidget) Renderer() *sdl.Renderer {
    return self.renderer
}



func (self *SWS_CoreWidget) RemoveChild(child SWS_Widget) {
    
    fmt.Println("todestroy:",child)
    for i,c := range self.children {
        fmt.Println("child:",c)
        if c == child {
            fmt.Println("found")
            if i==0 {
                self.children = self.children[1:]
            } else {
                self.children = append(self.children[:i],self.children[i+1:]...)
            }
            return
        }
    }
}



func (self *SWS_CoreWidget) AddChild(child SWS_Widget) {
    self.children=append(self.children,child)
    child.SetParent(self)
}



func (self *SWS_CoreWidget) SetParent(father SWS_Widget) {
    self.parent=father
}



func (self *SWS_CoreWidget) Parent() SWS_Widget {
    return self.parent
}



func (self *SWS_CoreWidget) getChildren() []SWS_Widget {
    return self.children
}



func (self *SWS_CoreWidget) IsInside(x,y int32) bool {
    return x>=0 && y>=0 && x<self.Width() && y<self.Height()
}



func (self *SWS_CoreWidget) TranslateXYToWidget(globalX,globalY int32) (x,y int32) {
    if self.Parent()==nil {
        return globalX-self.X(),globalY-self.Y()
    }
    return self.Parent().TranslateXYToWidget(globalX-self.X(),globalY-self.Y())
}



func (self *SWS_CoreWidget) MousePressDown(x,y int32, button uint8) {
    fmt.Println(x,y)
}



func (self *SWS_CoreWidget) MousePressUp(x,y int32, button uint8) {
    fmt.Println(x,y)
}



func (self *SWS_CoreWidget) MouseMove(x,y,xrel,yrel int32) {
}



func (self *SWS_CoreWidget) X() int32 {
    return self.x
}



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
    self.bgColor=color
}



func (self *SWS_CoreWidget) Move(x,y int32) {
    self.x=x
    self.y=y
}



func (self *SWS_CoreWidget) repaint() {
    if self.bgColor != 0 {
        self.FillRect(0,0,self.width,self.height,self.bgColor)
    }
    for _,child := range self.children {
        // adjust the clipping to the current child
        child.repaint()
        rectSrc := sdl.Rect{0,0, child.Width(),child.Height()}
        rectDst := sdl.Rect{child.X(), child.Y(), child.Width(),child.Height()}
        child.Surface().Blit(&rectSrc,self.Surface(),&rectDst)
    }
}



