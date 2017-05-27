package sws

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/sdl_ttf"
    "fmt"
    "os"
)

/*
 * in the main loop
 * if menutrigger widget exists, then:
 * - if MouseMove -> find the corresponding menutrigger/menu and
 *    - on the widget: destroy childs and treat the event (eventually create child)
 * - if MouseDown and we are not on menutrigger/menu -> destroy everything
 * - if MouseUp
 *   - if we are not on the menutrigger/menu -> destroy everything
 *   - else we send a MouseUp on the menutrigger/menu
 */
type MenuItem interface {
    Repaint(selected bool) *sdl.Surface
    WidthHeight() (int32,int32)
    SubMenu() *SWS_MenuWidget
    Clicked()
    Destroy()
}

type MenuItemLabel struct {
    font          *ttf.Font
    surface       *sdl.Surface
    Label         string
    ClickCallback func()
    subMenu *SWS_MenuWidget
}



func (self *MenuItemLabel) SubMenu() *SWS_MenuWidget {
    return self.subMenu
}



func (self *MenuItemLabel) SetSubMenu(sub *SWS_MenuWidget) {
    self.subMenu=sub
}



func (self *MenuItemLabel) Destroy() {
    self.surface.Free()
}



func (self *MenuItemLabel) Repaint(selected bool) *sdl.Surface {
    
    if self.Label != "" {
        var err error
        var solid *sdl.Surface
        color := sdl.Color{0, 0, 0, 255}
        if (selected) {
            color = sdl.Color{255, 255, 255, 255}
        }
        
        if solid, err = self.font.RenderUTF8_Blended(self.Label, color); err != nil {
            fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
            return self.surface
        }
        self.surface.Free()
        self.surface=solid
    }

    return self.surface
}



func (self *MenuItemLabel) WidthHeight() (int32,int32) {
    w,h,_ := self.font.SizeUTF8(self.Label)
    return int32(w),int32(h)
}



func (self *MenuItemLabel) Clicked() {
    if self.ClickCallback!=nil {
        self.ClickCallback()
    }
    hideMenu(nil)
}



func CreateMenuItemLabel(label string, callback func()) *MenuItemLabel {
    w,h,_ := defaultFont.SizeUTF8(label)
    surface,err := sdl.CreateRGBSurface(0,int32(w),int32(h),32,0x00ff0000,0x0000ff00,0x000000ff, 0xff000000)
    if err!=nil {
        panic(err)
    }
    menuitem := &MenuItemLabel{ 
                     Label:         label,
                     ClickCallback: callback,
                     subMenu:       nil,
                     surface:       surface,
                     //renderer:      renderer,
                     font:          defaultFont}
  return menuitem
}



type SWS_MenuWidget struct {
    SWS_CoreWidget
    items         []MenuItem
    activeItem    int
    lastSubActive int
}



func (self *SWS_MenuWidget) Destroy() {
    fmt.Println("SWS_MenuWidget::Destroy()")
    // cannot do self.SWS_CoreWidget.RemoveChild(self) <- it cast it into SWS_CoreWidget (which is a WSW_Widget interface, but the data/struct are not the same)
    self.Parent().RemoveChild(self)
    self.surface.Free()
    for _,i := range self.items {
        i.Destroy()
    }
}



func (self *SWS_MenuWidget) AddItem(item MenuItem) {
    self.items=append(self.items,item)
    w,h :=item.WidthHeight()
    w=w+30+4 // add space for margins and UI border
    if w>self.width {
        self.width=w
    }
    self.height+=h

    // recreate the surface
    var err error
    self.surface.Free()
    self.surface,err = sdl.CreateRGBSurface(0,self.width,self.height,32,0x00ff0000,0x0000ff00,0x000000ff, 0xff000000)
    if err!=nil {
        panic(err)
    }
}



func (self *SWS_MenuWidget) MousePressDown(x,y int32, button uint8) {
}



func (self *SWS_MenuWidget) MousePressUp(x,y int32, button uint8) {
    if (self.activeItem!=-1) {
        submenu:=self.items[self.activeItem].SubMenu()
        if (submenu==nil) {
            self.items[self.activeItem].Clicked()
        }
    }
}



func (self *SWS_MenuWidget) MouseMove(x,y,xrel,yrel int32) {
    previousActiveItem:=self.activeItem
    x-=2 // UI border
    y-=2 // UI border
    self.activeItem=-1
    var yy int32
    yy=0
    if (x>=0 && x<self.width) {
        for i, item := range self.items {
            _,h :=item.WidthHeight()
            if (yy<=y && yy+h>y) {
                self.activeItem=i
                break
            }
            yy+=h
        }
    }
    if (previousActiveItem!=self.activeItem) {
        //hideMenu(self)
        if (previousActiveItem!=-1 && self.activeItem!=-1) {
            submenu:=self.items[previousActiveItem].SubMenu()
            if (submenu!=nil) {
                hideMenu(submenu)
            }
        }
        if (self.lastSubActive!=-1 && self.activeItem!=-1) {
            submenu:=self.items[self.lastSubActive].SubMenu()
            if (submenu!=nil) {
                hideMenu(submenu)
            }
        }
        if (self.activeItem!=-1) {
            submenu:=self.items[self.activeItem].SubMenu()
            if (submenu!=nil) {
                self.lastSubActive=self.activeItem
                submenu.Move(self.X()+self.Width()-2,self.Y()+yy)
                ShowMenu(submenu)
            } else {
                self.lastSubActive=-1
            }
        }
        PostUpdate()
    }
}



func (self *SWS_MenuWidget) Repaint() {
    var y int32
    rect := sdl.Rect{0,0, self.width, self.height}
    self.surface.FillRect(&rect, 0xffdddddd)

    renderer,err := sdl.CreateSoftwareRenderer(self.surface)
    if err!=nil {
        panic(err)
    }
    renderer.SetDrawColor(0, 0, 0, 255)
    renderer.DrawRect(&rect)
    renderer.SetDrawColor(255, 255, 255, 255)
    renderer.DrawLine(1,1,int(self.width-2),1)
    renderer.DrawLine(1,1,1,int(self.height)-2)
    renderer.SetDrawColor(0x88, 0x88, 0x88, 255)
    renderer.DrawLine(int(self.width)-2,2,int(self.width)-2,int(self.height)-2)
    renderer.DrawLine(2,int(self.height)-2,int(self.width)-2,int(self.height)-2)

    for i, item := range self.items {
        w,h :=item.WidthHeight()
        if (i==self.activeItem || i==self.lastSubActive) {
            rect := sdl.Rect{2,y+2, self.width-4, h}
            self.surface.FillRect(&rect, 0xff8888ff)
        }
        surface := item.Repaint(i==self.activeItem || i==self.lastSubActive)
        rectSrc := sdl.Rect{0,0, w,h}
        rectDst := sdl.Rect{5+2,y+2, w,h}
        surface.Blit(&rectSrc, self.surface, &rectDst)
        
        if (item.SubMenu()!=nil) {
            var i int32
            for i=0;i<5;i++ {
                rect := sdl.Rect{self.width-15+2,y+(h/2)-i+2, (5-i)*2, 2*i+1}
                self.surface.FillRect(&rect, 0xff000000)
            }
        }

        y+=h
    }
}



func CreateMenuWidget() *SWS_MenuWidget {
    corewidget := CreateCoreWidget(4,4)
    widget := &SWS_MenuWidget{ SWS_CoreWidget: *corewidget,
                     items:         make([]MenuItem, 0, 0),
                     activeItem:    -1,
                     lastSubActive: -1 }
  return widget
}



//var rootMenu SWS_Widget
var currentMenuBar *SWS_MenuBarWidget
var menuStack []*SWS_MenuWidget

func ShowMenu(menu *SWS_MenuWidget) {
    if menuStack==nil {
        menuStack=append(make([]*SWS_MenuWidget,0,0),menu)
    } else {
        menuStack=append(menuStack,menu)
    }
    fmt.Println("menuStack=",menuStack)
    root.AddChild(menu)
    PostUpdate()
}



func findMenu(x int32, y int32) *SWS_MenuWidget{
    if menuStack == nil {
        return nil
    }
    for i:=len(menuStack)-1;i>=0;i-- {
        menu := menuStack[i]
        if x>=menu.X() && x<menu.X()+menu.Width() && y>=menu.Y() && y<menu.Y()+menu.Height() {
            return menu
        }
    }
    return nil
}



func hideMenu(menu *SWS_MenuWidget) {
    if menuStack == nil {
        return
    }
    // destroy all menus
    if menu == nil {
        fmt.Println("***",menuStack)
        for _,m := range menuStack {
        //        m.Destroy()
            fmt.Println("***remove menu",m)
            root.RemoveChild(m)
        }
        menuStack=nil
        PostUpdate()
        return
    }
    // destroy submenu
    fmt.Println("destroy submenu")
    for i,m := range menuStack {
        if m==menu {
            for _,s := range menuStack[i:] {
            //    s.Destroy()
                root.RemoveChild(s)
            }
            menuStack=menuStack[:i]
            PostUpdate()
            return
        }
    }
}



type SWS_MenuBarWidget struct {
    SWS_MenuWidget
    hasfocus      bool
    clickonmenu   bool
}



func (self *SWS_MenuBarWidget) HasFocus(hasfocus bool) {
    self.hasfocus=hasfocus
    if (hasfocus==false) {
        currentMenuBar=nil
        self.activeItem=-1
        self.lastSubActive=-1
    } else {
        currentMenuBar=self
    }
    
    //menuStack=append(make([]*SWS_MenuWidget,0,0),&(self.SWS_MenuWidget))

    PostUpdate()
}



func (self *SWS_MenuBarWidget) Repaint() {
    var x int32
    rect := sdl.Rect{0,0, self.width, self.height}
    self.surface.FillRect(&rect, 0xffdddddd)
    
    renderer,err := sdl.CreateSoftwareRenderer(self.surface)
    if err!=nil {
        panic(err)
    }   
    
    for i, item := range self.items {
        w,h :=item.WidthHeight()
        w+=10
        if (i==self.activeItem || i==self.lastSubActive) {
            rect := sdl.Rect{x,0, w, self.height}
            self.surface.FillRect(&rect, 0xff8888ff)
        }   
        surface := item.Repaint(i==self.activeItem || i==self.lastSubActive)
        rectSrc := sdl.Rect{0,0, w,h}
        rectDst := sdl.Rect{x+5,0, w,h}
        surface.Blit(&rectSrc, self.surface, &rectDst)
        
        x+=w
    }
    renderer.SetDrawColor(255, 255, 255, 255)
    renderer.DrawLine(0,int(self.height-1),int(self.width-1),int(self.height-1))
}   



func (self *SWS_MenuBarWidget) MousePressDown(x,y int32, button uint8) {
    self.hasfocus=true
    self.activeItem=-1 // because we close the menu in the main sws.go loop
    self.clickonmenu=false
    var xx int32
    xx=0
    if (y>=0 && y<self.height) {
        for _, item := range self.items {
            w,_ :=item.WidthHeight()
            w+=10
            if (xx<=x && xx+w>x) {
                self.clickonmenu=true
                break
            }
            xx+=w
        }
    }

    self.MouseMove(x,y,0,0)
}



func (self *SWS_MenuBarWidget) MousePressUp(x,y int32, button uint8) {
    if (self.activeItem!=-1) {
        submenu:=self.items[self.activeItem].SubMenu()
        if (submenu==nil) {
            self.items[self.activeItem].Clicked()
        }
    }
    PostUpdate()
}



func (self *SWS_MenuBarWidget) MouseMove(x,y,xrel,yrel int32) {
    if self.hasfocus==false || self.clickonmenu==false {
        return
    }
    previousActiveItem:=self.activeItem
    self.activeItem=-1
    var xx int32
    xx=0
    if (y>=0 && y<self.height) {
        for i, item := range self.items {
            w,_ :=item.WidthHeight()
            w+=10
            if (xx<=x && xx+w>x) {
                self.activeItem=i
                break
            }
            xx+=w
        }
    }
    if (previousActiveItem!=self.activeItem) {
        //hideMenu(self)
        if (previousActiveItem!=-1 && self.activeItem!=-1) {
            submenu:=self.items[previousActiveItem].SubMenu()
            if (submenu!=nil) {
                hideMenu(submenu)
            }
        }
        if (self.lastSubActive!=-1 && self.activeItem!=-1) {
            submenu:=self.items[self.lastSubActive].SubMenu()
            if (submenu!=nil) {
                hideMenu(submenu)
            }
        }
        if (self.activeItem!=-1) {
            submenu:=self.items[self.activeItem].SubMenu()
            if (submenu!=nil) {
                self.lastSubActive=self.activeItem

                yy:=self.height
                var widget SWS_Widget
                widget=self
                for (widget!=nil) {
                    xx+=widget.X()
                    yy+=widget.Y()
                    widget=widget.Parent()
                }
                submenu.Move(xx,yy-2)
                ShowMenu(submenu)
            } else {
                self.lastSubActive=-1
            }
        }
        PostUpdate()
    }
}



func (self *SWS_MenuBarWidget) AddItem(item MenuItem) {
    self.items=append(self.items,item)
}



func CreateMenuBarWidget() *SWS_MenuBarWidget {
    menuwidget := CreateMenuWidget()
    widget := &SWS_MenuBarWidget{ SWS_MenuWidget: *menuwidget,
                     clickonmenu:   false,
                     hasfocus:      false }
    widget.height=25
  return widget
}




