package sws

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

type SWS_Label struct {
	SWS_CoreWidget
	label     string
	image     *sdl.Surface
	textcolor sdl.Color
	centered  bool
}

func (self *SWS_Label) SetTextColor(color sdl.Color) {
	self.textcolor=color
}

func (self *SWS_Label) SetCentered(centered bool) {
        self.centered=centered
	PostUpdate()
}

func (self *SWS_Label) SetText(label string) {
        self.label=label
	PostUpdate()
}

func (self *SWS_Label) SetImage(image string) {
	if img,err := img.Load(image); err==nil {
        	self.image=img
	}
	PostUpdate()
}

func (self *SWS_Label) Repaint() {
	self.SWS_CoreWidget.Repaint()
	// text
	var text *sdl.Surface
	var err error
	if (self.label!="") {
		if text, err = self.Font().RenderUTF8_Blended(self.label, self.textcolor); err != nil {
		//      fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
		}
		defer text.Free()
        }
	if (text!=nil && self.image == nil) {
		wGap := self.Width() - text.W
		hGap := self.Height() - text.H
		if self.centered==false {
			wGap=0
		}
		rectSrc := sdl.Rect{0, 0, text.W, text.H}
		rectDst := sdl.Rect{(wGap/2), (hGap/2), self.Width()-(wGap/2), self.       Height()-(hGap/2)}
		if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
		}
	} else if (text==nil && self.image !=nil) {
		wGap := self.Width() - self.image.W
		hGap := self.Height() - self.image.H
		if self.centered==false {
			wGap=0
		}
		rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
		rectDst := sdl.Rect{(wGap/2), (hGap/2), self.Width()-(wGap/2), self.       Height()-(hGap/2)}
		if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
		}
	} else if (text!=nil && self.image !=nil) {
		wTGap := self.Width() - text.W
		wIGap := self.Width() - self.image.W
		hGap := self.Height() - self.image.H - text.H
		if self.centered==false {
			wTGap=0
			wIGap=0
		}
		rectSrc := sdl.Rect{0, 0, self.image.W, self.image.H}
		rectDst := sdl.Rect{(wIGap/2), (hGap/2), self.Width()-(wIGap/2), self.     Height()-(hGap/2)}
		if err = self.image.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
		}
		rectSrc = sdl.Rect{0, 0, text.W, text.H}
		rectDst = sdl.Rect{(wTGap/2), (hGap/2)+self.image.H, self.Width()-(wTGap/  2), self.Height()-(hGap/2)-self.image.H}
		if err = text.Blit(&rectSrc, self.Surface(), &rectDst); err != nil {
		}
	}
}

func CreateLabel(w, h int32, s string) *SWS_Label {
	corewidget := CreateCoreWidget(w, h)
	widget := &SWS_Label{SWS_CoreWidget: *corewidget,
		label: s,
		textcolor:    sdl.Color{0, 0, 0, 255}}
	return widget
}
