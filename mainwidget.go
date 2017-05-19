package sws

import (
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
    "os"
)



type SWS_MainWidget struct {
    SWS_CoreWidget
    label             string
    hasfocus          bool
    expandable        bool
    inmove            bool
    Close             func()
    buttonOnClose     bool
    cursorInsideClose bool
    buttonOnExpand    bool
    cursorInsideExpand bool
    subwidget         *SWS_CoreWidget
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
    fmt.Println("SWS_MainWidget::repaint()")
    // paint header
    var solid *sdl.Surface
    var err error
    var headbgcolor=self.bgColor

/*
    var lightcolorR uint8=0xff
    var lightcolorG uint8=0xff
    var lightcolorB uint8=0xff
    var darkcolorR uint8=0x88
    var darkcolorG uint8=0x88
    var darkcolorB uint8=0x88
*/    
    if self.hasfocus {
        headbgcolor=0xfffff10b
/*        lightcolorR=0xff
        lightcolorG=0xf9
        lightcolorB=0x96
        darkcolorR=0xbd
        darkcolorG=0xb2
        darkcolorB=0x00
*/    }

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

    if self.hasfocus {
        rectSrc := sdl.Rect{0,0, mainlefts.W,mainlefts.H}
        rectDst := sdl.Rect{3,3, mainlefts.W,mainlefts.H}
        if (self.buttonOnClose && self.cursorInsideClose) {
            if mainleftclickeds.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
            }
        } else {
            if mainlefts.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
            }
        }
        if (self.expandable) {
            rectSrc = sdl.Rect{0,0, mainrights.W,mainrights.H}
            rectDst = sdl.Rect{maxW-17,3, mainrights.W,mainrights.H}
            if (self.buttonOnExpand && self.cursorInsideExpand) {
                if mainrightclickeds.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
                }
            } else {
                if mainrights.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
                }
            }
        }
    }

/*
    self.SetDrawColor(darkcolorR,darkcolorG,darkcolorB,0xff)
    self.DrawLine(maxW-14,6,maxW-14,18)
    self.DrawLine(maxW-14,6,maxW-2,6)
    self.DrawLine(maxW-13,17,maxW-3,17)
    self.DrawLine(maxW-3,7,maxW-3,17)
    self.SetDrawColor(lightcolorR,lightcolorG,lightcolorB,0xff)
    self.DrawLine(maxW-13,7,maxW-2,7)
    self.DrawLine(maxW-2,7,maxW-2,18)
    self.DrawLine(maxW-2,18,maxW-13,18)
    self.DrawLine(maxW-13,18,maxW-13,7)
    
    for i:=8; i<=16; i++ {
        for j:=8; j<=16; j++ {
            self.SetDrawColor(
               uint8(int(lightcolorR)+(int(darkcolorR)-int(lightcolorR))*(i+j-12)/24),
               uint8(int(lightcolorG)+(int(darkcolorG)-int(lightcolorG))*(i+j-12)/24),
               uint8(int(lightcolorB)+(int(darkcolorB)-int(lightcolorB))*(i+j-12)/24),
               0xff)
            self.DrawPoint(maxW-20+int32(i),int32(j))
        }
    }
 
    self.SetDrawColor(darkcolorR,darkcolorG,darkcolorB,0xff)
    self.DrawLine(maxW-18,2,maxW-18,10)
    self.DrawLine(maxW-18,2,maxW-10,2)
    self.DrawLine(maxW-16,9,maxW-11,9)
    self.DrawLine(maxW-11,4,maxW-11,9)
    self.SetDrawColor(lightcolorR,lightcolorG,lightcolorB,0xff)
    self.DrawLine(maxW-17,3,maxW-10,3)
    self.DrawLine(maxW-17,3,maxW-17,10)
    self.DrawLine(maxW-17,10,maxW-10,10)
    self.DrawLine(maxW-10,10,maxW-10,3)
    
    for i:=0; i<=4; i++ {
        for j:=0; j<=4; j++ {
            self.SetDrawColor(
               uint8(int(lightcolorR)+(int(darkcolorR)-int(lightcolorR))*(i+j)/8),
               uint8(int(lightcolorG)+(int(darkcolorG)-int(lightcolorG))*(i+j)/8),
               uint8(int(lightcolorB)+(int(darkcolorB)-int(lightcolorB))*(i+j)/8),
               0xff)
            self.DrawPoint(maxW-16+int32(i),4+int32(j))
        }
    }
      
*/

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
    maxW:=int32(wText)+40
    if (maxW>self.Width()) { maxW=self.Width() }
    if (x>2 && x<18 && y>2 && y<18) {
        self.buttonOnClose=true
        self.cursorInsideClose=true
        PostUpdate()
        return
    }
    if (self.expandable && x>maxW-18 && x<maxW-2 && y>2 && y<18) {
        self.buttonOnExpand=true
        self.cursorInsideExpand=true
        PostUpdate()
        return
    }
    if (x< 40+int32(wText) && y<20) {
        self.inmove=true
    }
}



func (self *SWS_MainWidget) MousePressUp(x,y int32, button uint8) {
    self.inmove=false
    if (self.buttonOnClose==true) {
        self.buttonOnClose=false
        PostUpdate()
    }
    if (self.buttonOnExpand==true) {
        self.buttonOnExpand=false
        PostUpdate()
    }
}



func (self *SWS_MainWidget) MouseMove(x,y,xrel,yrel int32) {

    if (self.inmove) {
        self.x+=xrel
        self.y+=yrel
        PostUpdate()
    } else {
        wText,_,_ := self.font.SizeUTF8(self.label)
        maxW:=int32(wText)+40
        if (maxW>self.Width()) { maxW=self.Width() }

        if (self.buttonOnClose) {
            if (x>2 && x<18 && y>2 && y<18) {
                self.cursorInsideClose=true
            } else {
                self.cursorInsideClose=false
            }
            PostUpdate()
        }
        if (self.buttonOnExpand) {
            if (x>maxW-18 && x<maxW-2 && y>2 && y<18) {
                self.cursorInsideExpand=true
            } else {
                self.cursorInsideExpand=false
            }
            PostUpdate()
        }
    }
}



func CreateMainWidget(w,h int32, s string,expandable bool) *SWS_MainWidget {
    corewidget := CreateCoreWidget(w,h)
    subwidget := CreateCoreWidget(w,h-20)
    subwidget.Move(0,20)
    corewidget.AddChild(subwidget)
    widget := &SWS_MainWidget{ SWS_CoreWidget: *corewidget,
                     label:             s,
                     hasfocus:          false,
                     expandable:        expandable,
                     inmove:            false, 
                     buttonOnClose:     false,
                     cursorInsideClose: false,
                     buttonOnExpand:    false,
                     cursorInsideExpand: false,
                     subwidget:         subwidget}
  return widget
}




