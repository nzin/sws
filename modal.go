package sws

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func ShowModalErrorSurfaceicon(root *RootWidget, title string, iconsurface *sdl.Surface, desc string, callback func()) {
	modal := NewMainWidget(500, 200, title, false, false)
	modal.SetCloseCallback(func() {
		root.RemoveChild(modal)
		if callback != nil {
			callback()
		}
	})

	icon := NewLabelWidget(32, 32, "")
	if iconsurface != nil {
		icon.SetImageSurface(iconsurface)
	}
	icon.Move(20, 50)
	modal.AddChild(icon)

	textarea := NewTextAreaWidget(400, 70, desc)
	textarea.Move(70, 40)
	textarea.SetDisabled(true)
	modal.AddChild(textarea)

	ok := NewButtonWidget(100, 25, "Ok")
	ok.Move(370, 120)
	ok.SetClicked(func() {
		root.RemoveChild(modal)
		if callback != nil {
			callback()
		}
	})
	modal.AddChild(ok)
	modal.Move((root.Width()-500)/2, (root.Height()-200)/2)

	root.AddChild(modal)
	root.SetModal(modal)

	root.SetFocus(ok)
}

func ShowModalError(root *RootWidget, title, iconpath, desc string, callback func()) {
	img, _ := img.Load(iconpath)
	ShowModalErrorSurfaceicon(root, title, img, desc, callback)
}

func ShowModalYesNoSurfaceicon(root *RootWidget, title string, iconsurface *sdl.Surface, desc string, callbackyes func(), callbackno func()) {
	modal := NewMainWidget(500, 200, title, false, false)
	modal.SetCloseCallback(func() {
		root.RemoveChild(modal)
		if callbackno != nil {
			callbackno()
		}
	})

	icon := NewLabelWidget(32, 32, "")
	if icon != nil {
		icon.SetImageSurface(iconsurface) //"resources/icon-triangular-big.png"
	}
	icon.Move(20, 50)
	modal.AddChild(icon)

	textarea := NewTextAreaWidget(400, 70, desc)
	textarea.Move(70, 40)
	textarea.SetDisabled(true)
	modal.AddChild(textarea)

	yes := NewButtonWidget(50, 25, "Yes")
	yes.Move(310, 120)
	yes.SetClicked(func() {
		root.RemoveChild(modal)
		if callbackyes != nil {
			callbackyes()
		}
	})
	modal.AddChild(yes)

	no := NewButtonWidget(50, 25, "No")
	no.Move(370, 120)
	no.SetClicked(func() {
		root.RemoveChild(modal)
		if callbackno != nil {
			callbackno()
		}
	})
	modal.AddChild(no)

	modal.Move((root.Width()-500)/2, (root.Height()-200)/2)

	root.AddChild(modal)
	root.SetModal(modal)

	root.SetFocus(no)
}

func ShowModalYesNo(root *RootWidget, title, iconpath, desc string, callbackyes func(), callbackno func()) {
	img, _ := img.Load(iconpath)
	ShowModalYesNoSurfaceicon(root, title, img, desc, callbackyes, callbackno)
}
