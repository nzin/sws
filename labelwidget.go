package sws

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type LabelWidget struct {
	CoreWidget
	label     string
	image     *sdl.Surface
	imageleft bool
	textcolor sdl.Color
	centered  bool
}

func (self *LabelWidget) AlignImageLeft(alignleft bool) {
	self.imageleft = alignleft
}

func (self *LabelWidget) SetTextColor(color sdl.Color) {
	self.textcolor = color
}

func (self *LabelWidget) SetCentered(centered bool) {
	self.centered = centered
	self.PostUpdate()
}

func (self *LabelWidget) SetText(label string) {
	self.label = label
	self.PostUpdate()
}

func (self *LabelWidget) SetImage(image string) {
	if img, err := img.Load(image); err == nil {
		self.image = img
	}
	self.PostUpdate()
}

func (self *LabelWidget) SetImageSurface(img *sdl.Surface) {
	self.image = img
	self.PostUpdate()
}

func (self *LabelWidget) Repaint() {
	self.CoreWidget.Repaint()
	// text
	var text *sdl.Surface
	var err error
	if self.label != "" {
		if text, err = self.Font().RenderUTF8_Blended(self.label, self.textcolor); err != nil {
		}
		defer text.Free()
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
			wGap := self.Width() - self.image.W - text.W
			if self.centered == false {
				wGap = 0
			}
			rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
			rectDst := sdl.Rect{2 + (wGap / 2), 2 + (hIGap / 2), self.Width() - 2 - (wGap / 2), self.Height() - 2 - (hIGap / 2)}
			if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
			rectSrc = sdl.Rect{0, 0, text.W, text.H}
			rectDst = sdl.Rect{2 + (wGap / 2) + self.image.W, 2 + (hTGap / 2), self.Width() - 2 - (wGap / 2) - self.image.W, self.Height() - 2 - (hTGap / 2)}
			if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
			}
		}
	}
}

func NewLabelWidget(w, h int32, s string) *LabelWidget {
	corewidget := NewCoreWidget(w, h)
	widget := &LabelWidget{CoreWidget: *corewidget,
		label:     s,
		textcolor: sdl.Color{0, 0, 0, 255}}
	return widget
}
