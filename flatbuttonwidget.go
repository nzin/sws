package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

type FlatButtonWidget struct {
	ButtonWidget
}

func (self *FlatButtonWidget) Repaint() {
	self.CoreWidget.Repaint()

	// text
	var text *sdl.Surface
	var err error
	if self.label != "" {
		if text, err = self.Font().RenderUTF8Blended(self.label, self.textcolor); err != nil {
			//	fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
		}
		defer text.Free()
	}

	if self.cursorInside == false {
		if self.bgColor != 0 {
			self.FillRect(0, 0, self.width, self.height, self.bgColor)
		}
	} else {
		self.FillRect(0, 0, self.width, self.height, 0xff8888ff)
	}
	if text != nil && self.image == nil {
		wGap := self.Width() - text.W
		hGap := self.Height() - text.H
		if self.centered == false {
			wGap = 0
		}
		rectSrc := sdl.Rect{0, 0, text.W, text.H}
		rectDst := sdl.Rect{(wGap / 2), (hGap / 2), self.Width() - (wGap / 2), self.Height() - (hGap / 2)}
		if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
		}
	} else if text == nil && self.image != nil {
		wGap := self.Width() - self.image.W
		hGap := self.Height() - self.image.H
		if self.centered == false {
			wGap = 0
		}
		rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
		rectDst := sdl.Rect{(wGap / 2), (hGap / 2), self.Width() - (wGap / 2), self.Height() - (hGap / 2)}
		if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
		}
	} else if text != nil && self.image != nil {
		if self.imageleft == false {
			wTGap := self.Width() - text.W
			wIGap := self.Width() - self.image.W
			hGap := self.Height() - self.image.H - text.H
			if self.centered == false {
				wTGap = 0
				wIGap = 0
			}
			rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
			rectDst := sdl.Rect{(wIGap / 2), (hGap / 2), self.Width() - (wIGap / 2), self.Height() - (hGap / 2)}
			if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
			rectSrc = sdl.Rect{0, 0, text.W, text.H}
			rectDst = sdl.Rect{(wTGap / 2), (hGap / 2) + self.image.H, self.Width() - (wTGap / 2), self.Height() - (hGap / 2) - self.image.H}
			if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		} else {
			hTGap := self.Height() - text.H
			hIGap := self.Height() - self.image.H
			wGap := self.Width() - self.image.W - text.W + 10
			if self.centered == false {
				wGap = 10
			}
			rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
			rectDst := sdl.Rect{(wGap / 2), (hIGap / 2), self.Width() - (wGap / 2), self.Height() - (hIGap / 2)}
			if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
			rectSrc = sdl.Rect{0, 0, text.W, text.H}
			rectDst = sdl.Rect{(wGap / 2) + self.image.W + 5, (hTGap / 2), self.Width() - (wGap / 2) - self.image.W, self.Height() - (hTGap / 2)}
			if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		}
	}
}

func NewFlatButtonWidget(w, h int32, s string) *FlatButtonWidget {
	buttonwidget := NewButtonWidget(w, h, s)
	widget := &FlatButtonWidget{ButtonWidget: *buttonwidget}
	return widget
}
