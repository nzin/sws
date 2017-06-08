package sws

import (
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
)



type SWS_DropdownWidget struct {
    SWS_CoreWidget
    Choices      []string
    ActiveChoice int32
    buttonState  bool
    cursorInside bool
    clicked      func()
    menu         *SWS_MenuWidget
    hasfocus     bool
}



func (self *SWS_DropdownWidget) HasFocus(hasfocus bool) {
    self.hasfocus=hasfocus
    if (hasfocus==false) {
        menuInitiator=nil
    } else {
        hideMenu(nil)
        menuInitiator=self
    }

    PostUpdate()
}



func (self *SWS_DropdownWidget) SetClicked(callback func()) {
    self.clicked=callback
}



func (self *SWS_DropdownWidget) MousePressDown(x,y int32, button uint8) {
    fmt.Println("Button.PressDown")
    if button == sdl.BUTTON_LEFT {
        self.buttonState=true
        self.cursorInside=true
        if (self.menu!=nil && self.menu.Parent()!=nil) {
            hideMenu(nil)
        } else {
            self.menu=CreateMenuWidget()
            for i,choice := range self.Choices {
                index:=i
                self.menu.AddItem(CreateMenuItemLabel(choice, func() {
                    self.ActiveChoice=int32(index)
                    if self.clicked!=nil {
                        self.clicked()
                    }
                 }))
            }
            var xx int32
            yy:=self.height
            var widget SWS_Widget
            widget=self
            for (widget!=nil) {
                xx+=widget.X()
                yy+=widget.Y()
                widget=widget.Parent()
            }
            self.menu.Move(xx,yy-2)
            ShowMenu(self.menu)
            PostUpdate()
        }
    }
}

    
    
func (self *SWS_DropdownWidget) MousePressUp(x,y int32, button uint8) {
    if button == sdl.BUTTON_LEFT {
        self.buttonState=false
        if self.cursorInside==true {
            /*if self.clicked != nil {
                self.clicked()
            }*/
            if self.menu!=nil && self.menu.Parent()!=nil {
                var xx int32
                yy:=self.height
                var widget SWS_Widget
                widget=self
                for (widget!=nil) {
                    xx+=widget.X()
                    yy+=widget.Y()
                    widget=widget.Parent()
                }
                self.menu.Move(xx,yy-2)
                ShowMenu(self.menu)
            }
        }
        self.cursorInside=false
        PostUpdate()
    }
}



func (self *SWS_DropdownWidget) MouseMove(x,y,xrel,yrel int32) {
    oldCursorInside:=self.cursorInside
    if self.buttonState == true {
        if (x>=0 && x< self.Width() && y>=0 && y<self.Height()) {
            self.cursorInside=true
        } else {
            self.cursorInside=false
        }
        if (oldCursorInside!=self.cursorInside) {
            PostUpdate()
        }
    }
}



func (self *SWS_DropdownWidget) Repaint() {
    label:=""
    if self.ActiveChoice>=int32(len(self.Choices)) {
        self.ActiveChoice=int32(len(self.Choices)-1)
    }
    if len(self.Choices)>0 {
        label=self.Choices[self.ActiveChoice]
    }

    self.SWS_CoreWidget.Repaint()
    self.SetDrawColor(0,0,0,255)
    self.DrawLine(0,1,0,self.Height()-2)
    self.DrawLine(self.Width()-1,1,self.Width()-1,self.Height()-2)
    self.DrawLine(1,0,self.Width()-2,0)
    self.DrawLine(1,self.Height()-1,self.Width()-2,self.Height()-1)
        self.WriteText(4,0,label,sdl.Color{0, 0, 0, 255})
        self.FillRect(self.Width()-25,2,20,self.Height()-4,0xffdddddd)
        self.SetDrawColor(0,0,0,255)
        //self.DrawLine(self.Width()-25,4,self.Width()-25,self.Height()-6)
        for i:=0;i<5;i++ {
            self.DrawLine(self.Width()-18+int32(i),6+int32(i)*2,self.Width()-9-int32(i),6+int32(i)*2)
            self.DrawLine(self.Width()-18+int32(i),7+int32(i)*2,self.Width()-9-int32(i),7+int32(i)*2)
        }
        // bright
        self.SetDrawColor(255,255,255,255)
        self.DrawLine(1,1,1,self.Height()-2)
        self.DrawLine(1,1,self.Width()-2,1)
        self.SetDrawColor(240,240,240,255)
        self.DrawLine(2,2,2,self.Height()-3)
        self.DrawLine(2,2,self.Width()-3,2)
        //dark
        self.SetDrawColor(50,50,50,255)
        self.DrawLine(self.Width()-2,1,self.Width()-2,self.Height()-2)
        self.DrawLine(1,self.Height()-2,self.Width()-2,self.Height()-2)
        self.SetDrawColor(150,150,150,255)
        self.DrawLine(self.Width()-3,2,self.Width()-3,self.Height()-3)
        self.DrawLine(2,self.Height()-3,self.Width()-3,self.Height()-3)
}



func CreateDropdownWidget(w,h int32, choices []string) *SWS_DropdownWidget {
    corewidget := CreateCoreWidget(w,h)
    widget := &SWS_DropdownWidget{ SWS_CoreWidget: *corewidget,
                     buttonState:  false,
                     cursorInside: false,
                     Choices:      choices,
                     ActiveChoice: 0,
              }
  return widget
}



