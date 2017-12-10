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

	sub := sws.NewMenuWidget()
	sub.AddItem(sws.NewMenuItemLabel("sub bla bla", nil))
	sub.AddItem(sws.NewMenuItemLabel("sub bla bla 2", nil))
	sub.AddItem(sws.NewMenuItemLabel("sub bla bla 3", nil))
	sub2 := sws.NewMenuWidget()
	sub2.AddItem(sws.NewMenuItemLabel("sub bla bla", nil))
	sub2.AddItem(sws.NewMenuItemLabel("sub bla bla 2", nil))
	sub2.AddItem(sws.NewMenuItemLabel("sub bla bla 3", nil))
	m := sws.NewMenuWidget()
	m.AddItem(sws.NewMenuItemLabel("bla bla", nil))
	m.AddItem(sws.NewMenuItemLabel("bla bla 2", nil))
	m.AddItem(sws.NewMenuItemLabel("bla bla 3", nil))
	i4 := sws.NewMenuItemLabel("sub1", nil)
	i4.SetSubMenu(sub)
	m.AddItem(i4)
	i5 := sws.NewMenuItemLabel("sub2", nil)
	i5.SetSubMenu(sub2)
	m.AddItem(i5)
	m.Move(400, 100)
	sws.ShowMenu(m)
}

var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()

	root := sws.Init(800, 600)

	//f:=sws.NewCoreWidget(200,100)
	f := sws.NewMainWidget(200, 100, "very looooooooooooong title ", true, true)
	scrollwidget := sws.NewScrollWidget(300, 200)
	corewidget := sws.NewCoreWidget(200, 300)
	scrollwidget.SetInnerWidget(corewidget)
	//f.SetColor(0xffff0000)
	f.Move(100, 10)

	c := sws.NewLabelWidget(100, 25, "Footcheball")
	c.Move(25, 70)
	corewidget.AddChild(c)

	cb := sws.NewCheckboxWidget()
	cb.Move(0, 70)
	corewidget.AddChild(cb)

	b := sws.NewButtonWidget(100, 25, "click")
	b.SetClicked(func() {
		modal := sws.NewMainWidget(200, 100, "modal", false, true)
		modal.SetCloseCallback(func() {
			root.RemoveChild(modal)
		})
		root.AddChild(modal)
		root.SetModal(modal)
	})
	b.Move(10, 10)
	corewidget.AddChild(b)

	i := sws.NewInputWidget(100, 25, "text")
	corewidget.AddChild(i)
	i.Move(50, 50)

	dd := sws.NewDropdownWidget(100, 25, []string{"text 1", "text 2"})
	dd.Move(50, 110)
	corewidget.AddChild(dd)

	sbh := sws.NewScrollbarWidget(100, 20, true)
	sbh.SetMaximum(1000)
	corewidget.AddChild(sbh)
	sbh.Move(50, 140)

	sbv := sws.NewScrollbarWidget(20, 100, false)
	sbv.SetMaximum(1000)
	corewidget.AddChild(sbv)
	sbv.Move(50, 170)

	f.SetInnerWidget(scrollwidget)
	root.AddChild(f)

	filemenu := sws.NewMenuWidget()
	filemenu.AddItem(sws.NewMenuItemLabel("file bla bla", nil))
	filemenu.AddItem(sws.NewMenuItemLabel("file bla bla 2", nil))
	filemenu.AddItem(sws.NewMenuItemLabel("file bla bla 3", nil))
	file := sws.NewMenuItemLabel("File", nil)
	file.SetSubMenu(filemenu)
	menubar := sws.NewMenuBarWidget()
	menubar.AddItem(file)
	menubar.AddItem(sws.NewMenuItemLabel("View", nil))
	menubar.AddItem(sws.NewMenuItemLabel("About", nil))

	sv := sws.NewSplitviewWidget(200, 200, true)

	vbox := sws.NewListWidget(200, 10)
	vbox.AddItem(25, "Element 1", "", nil)
	vbox.AddItem(25, "Element 2", "", nil)
	vbox.AddItem(25, "Element 3", "", nil)
	vbox.AddItem(25, "longer element 4", "", nil)

	vboxscroll := sws.NewScrollWidget(200, 200)
	vboxscroll.ShowHorizontalScrollbar(false)
	vboxscroll.SetInnerWidget(vbox)
	sv.SetLeftWidget(vboxscroll)

	corewidget2 := sws.NewCoreWidget(200, 300)
	sv.SetRightWidget(corewidget2)

	b2 := sws.NewButtonWidget(100, 100, "idea")
	b2.SetImage("idea.png")
	corewidget2.AddChild(b2)

	l2 := sws.NewLabelWidget(100, 100, "idea2")
	l2.SetImage("idea.png")
	l2.SetCentered(true)
	corewidget2.AddChild(l2)
	l2.Move(0, 100)

	b3 := sws.NewFlatButtonWidget(100, 100, "idea")
	b3.SetImage("idea.png")
	corewidget2.AddChild(b3)
	b3.Move(100, 0)

	ta := sws.NewTextAreaWidget(100, 100, "initial text another line t\tt\tt")
	corewidget2.AddChild(ta)
	ta.Move(100, 100)

	ta2 := sws.NewTextAreaWidget(100, 100, "initial text another line t\tt\tt")
	ta2.SetDisabled(true)
	corewidget2.AddChild(ta2)
	ta2.Move(0, 200)

	main1 := sws.NewMainWidget(200, 100, "main1", false, true)
	main1.Move(400, 300)
	main1.SetMenuBar(menubar)

	main1.SetInnerWidget(sv)

	root.AddChild(main1)

	tabs := sws.NewTabWidget(100, 100)
	ta3 := sws.NewTextAreaWidget(100, 100, "text area3 ")
	ta4 := sws.NewTextAreaWidget(100, 100, "text area4 ")
	tabs.AddTab("text area 3", ta3)
	tabs.AddTab("text area 4", ta4)

	tree := sws.NewTreeViewWidget()
	ti1 := sws.NewTreeViewItem("tree view item 1", "", nil)
	ti2 := sws.NewTreeViewItem("tree view item 2", "", nil)
	ti3 := sws.NewTreeViewItem("tree view item 3", "", nil)
	sti1 := sws.NewTreeViewItem("sub item 1", "", nil)
	sti2 := sws.NewTreeViewItem("sub item 2", "", nil)
	ti1.AddSubItem(sti1)
	ti1.AddSubItem(sti2)
	tree.AddItem(ti1)
	tree.AddItem(ti2)
	tree.AddItem(ti3)

	tabs.AddTab("tree", tree)
	main2 := sws.NewMainWidget(300, 200, "tab", false, true)
	main2.SetInnerWidget(tabs)
	main2.Move(500, 400)
	root.AddChild(main2)

	for sws.PoolEvent() == false {
	}

}
