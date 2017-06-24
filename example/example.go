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
	//"github.com/veandco/go-sdl2/sdl"
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
	
	vbox := sws.CreateListWidget(200,10)
	vbox.AddItem(25,"Element 1","",nil)
	vbox.AddItem(25,"Element 2","",nil)
	vbox.AddItem(25,"Element 3","",nil)
	vbox.AddItem(25, "longer element 4","",nil)
	
	vboxscroll := sws.CreateScrollWidget(200, 200)
	vboxscroll.ShowHorizontalScrollbar(false)
	vboxscroll.SetInnerWidget(vbox)
	sv.SetLeftWidget(vboxscroll)

	corewidget2 := sws.CreateCoreWidget(200, 300)
	sv.SetRightWidget(corewidget2)

	b2 := sws.CreateButtonWidget(100, 100, "idea")
	b2.SetImage("idea.png")
	corewidget2.AddChild(b2)

	l2 := sws.CreateLabel(100, 100, "idea2")
	l2.SetImage("idea.png")
	l2.SetCentered(true)
	corewidget2.AddChild(l2)
	l2.Move(0, 100)

	b3 := sws.CreateFlatButtonWidget(100, 100, "idea")
	b3.SetImage("idea.png")
	corewidget2.AddChild(b3)
	b3.Move(100, 0)
	
	ta := sws.CreateTextAreaWidget(100,100,"initial text another line t\tt\tt")
        corewidget2.AddChild(ta)
	ta.Move(100,100)

	ta2 := sws.CreateTextAreaWidget(100,100,"initial text another line t\tt\tt")
	ta2.SetReadonly(true)
        corewidget2.AddChild(ta2)
	ta2.Move(0,200)

	main1 := sws.CreateMainWidget(200, 100, "main1", false, true)
	main1.Move(400, 300)
	main1.SetMenuBar(menubar)
	
	main1.SetInnerWidget(sv)
	
	root.AddChild(main1)

	for sws.PoolEvent() == false {
	}

}
