package main

// https://medium.com/random-go-tips/method-overriding-680cfd51ce40

/*
 * clipping: chaque widget ecrit dans sa fenetre
 * on copie la surface sur le widget pere
 */

import (
	"flag"
	"fmt"
	"github.com/nzin/sws"
	_ "github.com/veandco/go-sdl2/sdl"
	_ "github.com/veandco/go-sdl2/sdl_ttf"
	_ "log"
	_ "runtime/pprof"
)

func a() {
	fmt.Println("clicked!!!")
	sub := sws.CreateMenuWidget()
	sub.AddItem(sws.CreateMenuItemLabel("sub bla bla", nil))
	sub.AddItem(sws.CreateMenuItemLabel("sub bla bla 2", nil))
	sub.AddItem(sws.CreateMenuItemLabel("sub bla bla 3", nil))
	sub2 := sws.CreateMenuWidget()
	sub2.AddItem(sws.CreateMenuItemLabel("sub bla bla", nil))
	sub2.AddItem(sws.CreateMenuItemLabel("sub bla bla 2", nil))
	sub2.AddItem(sws.CreateMenuItemLabel("sub bla bla 3", nil))
	m := sws.CreateMenuWidget()
	m.AddItem(sws.CreateMenuItemLabel("bla bla", nil))
	m.AddItem(sws.CreateMenuItemLabel("bla bla 2", nil))
	m.AddItem(sws.CreateMenuItemLabel("bla bla 3", nil))
	i4 := sws.CreateMenuItemLabel("sub1", nil)
	i4.SetSubMenu(sub)
	m.AddItem(i4)
	i5 := sws.CreateMenuItemLabel("sub2", nil)
	i5.SetSubMenu(sub2)
	m.AddItem(i5)
	m.Move(400, 100)
	sws.ShowMenu(m)
}

var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()

	root := sws.Init(800, 600)

	//f:=sws.CreateCoreWidget(200,100)
	f := sws.CreateMainWidget(200, 100, "very looooooooooooong title ", true, true)
	scrollwidget := sws.CreateScrollWidget(300, 200)
	corewidget := sws.CreateCoreWidget(200, 300)
	scrollwidget.SetInnerWidget(corewidget)
	//f.SetColor(0xffff0000)
	f.Move(100, 10)

	c := sws.CreateLabel(100, 50, "Footcheball")
	corewidget.AddChild(c)
	c.Move(-10, 85)

	b := sws.CreateButtonWidget(100, 25, "click")
	b.SetClicked(a)
	b.Move(10, 10)
	corewidget.AddChild(b)

	i := sws.CreateInputWidget(100, 25, "text")
	corewidget.AddChild(i)
	i.Move(50, 50)

	dd := sws.CreateDropdownWidget(100, 25, []string{"text 1", "text 2"})
	dd.Move(50, 110)
	corewidget.AddChild(dd)

	sbh := sws.CreateScrollbarWidget(100, 20, true)
	sbh.SetMaximum(1000)
	corewidget.AddChild(sbh)
	sbh.Move(50, 140)

	sbv := sws.CreateScrollbarWidget(20, 100, false)
	sbv.SetMaximum(1000)
	corewidget.AddChild(sbv)
	sbv.Move(50, 170)

	f.SetInnerWidget(scrollwidget)
	root.AddChild(f)

	filemenu := sws.CreateMenuWidget()
	filemenu.AddItem(sws.CreateMenuItemLabel("file bla bla", nil))
	filemenu.AddItem(sws.CreateMenuItemLabel("file bla bla 2", nil))
	filemenu.AddItem(sws.CreateMenuItemLabel("file bla bla 3", nil))
	file := sws.CreateMenuItemLabel("File", nil)
	file.SetSubMenu(filemenu)
	menubar := sws.CreateMenuBarWidget()
	menubar.AddItem(file)
	menubar.AddItem(sws.CreateMenuItemLabel("View", nil))
	menubar.AddItem(sws.CreateMenuItemLabel("About", nil))
	
	sv := sws.CreateSplitviewWidget(200,200,true)
	
	vbox := sws.CreateVBoxWidget(200,10)
	vbox.AddChild(sws.CreateLabel(200, 25, "Element 1"))
	vbox.AddChild(sws.CreateLabel(200, 25, "Element 2"))
	vbox.AddChild(sws.CreateLabel(200, 25, "Element 3"))
	vbox.AddChild(sws.CreateLabel(200, 25, "longer element 4"))
	
	vboxscroll := sws.CreateScrollWidget(200, 200)
	vboxscroll.ShowHorizontalScrollbar(false)
	vboxscroll.SetInnerWidget(vbox)
	sv.SetLeftWidget(vboxscroll)

	main1 := sws.CreateMainWidget(200, 100, "main1", false, true)
	main1.Move(400, 300)
	main1.SetMenuBar(menubar)
	
	main1.SetInnerWidget(sv)
	
	root.AddChild(main1)

	for sws.PoolEvent() == false {
	}

}
