package sws

import (
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
    "os"
)



type SWS_MainWidget struct {
    SWS_CoreWidget
    label     string
    hasfocus  bool
    inmove    bool
    Close     func()
    subwidget *SWS_CoreWidget
}



func (self *SWS_MainWidget) HasFocus(focus bool) {
    if (self.hasfocus!=focus) {
        self.hasfocus=focus
        PostUpdate()
    }
}



func (self *SWS_MainWidget) IsInside(x,y int32) bool {
    if (y<20) {
        wText,_,_ := self.font.SizeUTF8(self.label)
        return x>=0 && y>=0 && x<int32(wText)+40
    } else {
        return x>=0 && y>=0 && x<self.Width() && y<self.Height()
    }
}



func (self *SWS_MainWidget) repaint() {
    // paint header
    var solid *sdl.Surface
    var err error
    var headbgcolor=self.bgColor
    
    if self.hasfocus {
        headbgcolor=0xfffff10b
    }
    
    color:=sdl.Color{0,0,0,255}
    if self.label == "" {
        self.label="unnamed"
    }
    if solid, err = self.Font().RenderUTF8_Blended(self.label, color); err != nil {
        fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
    }
    defer solid.Free()

    maxW:=solid.W+40
    if (maxW>self.Width()) { maxW=self.Width() }
    self.FillRect(0,0,maxW,20,headbgcolor)
    rectSrc := sdl.Rect{0,0, maxW-40,solid.H}
    rectDst := sdl.Rect{20,0, maxW-40,20}
    if err = solid.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
    }


    // paint body
    if self.bgColor != 0 {
        self.FillRect(0,20,self.width,self.height-20,self.bgColor)
    }
    self.subwidget.repaint()
    rectSrc = sdl.Rect{0,0, self.subwidget.Width(),self.subwidget.Height()}
    rectDst = sdl.Rect{self.subwidget.X(), self.subwidget.Y(), self.subwidget.Width(),self.subwidget.Height()}
    self.subwidget.Surface().Blit(&rectSrc,self.Surface(),&rectDst)
}



func (self *SWS_MainWidget) AddChild(child SWS_Widget) {
    self.subwidget.AddChild(child)
}



func (self *SWS_MainWidget) MousePressDown(x,y int32, button uint8) {
    wText,_,_ := self.font.SizeUTF8(self.label)
    if (x>20 && x< 40+int32(wText) && y<20) {
        self.inmove=true
    }
}



func (self *SWS_MainWidget) MousePressUp(x,y int32, button uint8) {
    self.inmove=false
}



func (self *SWS_MainWidget) MouseMove(x,y,xrel,yrel int32) {
    if (self.inmove) {
        self.x+=xrel
        self.y+=yrel
        PostUpdate()
    }
}



func CreateMainWidget(w,h int32, s string) *SWS_MainWidget {
    corewidget := CreateCoreWidget(w,h)
    subwidget := CreateCoreWidget(w,h-20)
    subwidget.Move(0,20)
    corewidget.AddChild(subwidget)
    widget := &SWS_MainWidget{ SWS_CoreWidget: *corewidget,
                     label:     s,
                     hasfocus:  false,
                     inmove:    false, 
                     subwidget: subwidget}
  return widget
}




