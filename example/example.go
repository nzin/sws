package main

// https://medium.com/random-go-tips/method-overriding-680cfd51ce40

/*
 * clipping: chaque widget ecrit dans sa fenetre
 * on copie la surface sur le widget pere
 */

import (
    _ "github.com/veandco/go-sdl2/sdl"
    _ "github.com/veandco/go-sdl2/sdl_ttf"
    "fmt"
    "flag"
    "nzin/sws"
    _ "runtime/pprof"
    _ "log"
)



func a() {
    fmt.Println("clicked!!!")
    sub:=sws.CreateMenuWidget()
    sub.AddItem(sws.CreateMenuItemLabel("sub bla bla",nil))
    sub.AddItem(sws.CreateMenuItemLabel("sub bla bla 2",nil))
    sub.AddItem(sws.CreateMenuItemLabel("sub bla bla 3",nil))
    sub2:=sws.CreateMenuWidget()
    sub2.AddItem(sws.CreateMenuItemLabel("sub bla bla",nil))
    sub2.AddItem(sws.CreateMenuItemLabel("sub bla bla 2",nil))
    sub2.AddItem(sws.CreateMenuItemLabel("sub bla bla 3",nil))
    m:=sws.CreateMenuWidget()
    m.AddItem(sws.CreateMenuItemLabel("bla bla",nil))
    m.AddItem(sws.CreateMenuItemLabel("bla bla 2",nil))
    m.AddItem(sws.CreateMenuItemLabel("bla bla 3",nil))
    i4:=sws.CreateMenuItemLabel("sub1",nil)
    i4.SetSubMenu(sub)
    m.AddItem(i4)
    i5:=sws.CreateMenuItemLabel("sub2",nil)
    i5.SetSubMenu(sub2)
    m.AddItem(i5)
    m.Move(400,100)
    sws.ShowMenu(m)
}


var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
    flag.Parse()
     
    root := sws.Init(800,600)
        
    //f:=sws.CreateCoreWidget(200,100)
    f:=sws.CreateMainWidget(200,100,"very looooooooooooong title ",true)
    //f.SetColor(0xffff0000)
    f.Move(100,10)
    c:=sws.CreateLabel(100,50,"Footcheball")
    f.AddChild(c)
    c.Move(-10,85)
    b:=sws.CreateButtonWidget(100,25,"click")
    b.SetClicked(a)
    f.AddChild(b)
    b.Move(10,10)
    i:=sws.CreateInputWidget(100,25,"text")
    f.AddChild(i)
    i.Move(50,50)
    root.AddChild(f)
        
    main1:=sws.CreateMainWidget(200,100,"main1",false)
    main1.Move(400,300)
    root.AddChild(main1)
        

    for sws.PoolEvent() == false {
    }

}
