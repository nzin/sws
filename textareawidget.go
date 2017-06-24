package sws

import (
	"github.com/veandco/go-sdl2/sdl"
)

type SWS_TextAreaWidget struct {
	SWS_CoreWidget
	text                  string
	initialCursorPosition int // begin of selection
	endCursorPosition     int // end of selection
	xCursor               int32 // position of the selection but in x/y pixels
	yCursor               int32 // position of the selection but in x/y pixels
	hasfocus              bool
	leftButtonDown        bool
	writeOffset           int32 // how far we scroll the text
	readonly              bool
}

const (
	GLYPH_SPACE = 0
	GLYPH_TAB =   1
	GLYPH_WORD =  2
	GLYPH_END =   3
	GLYPH_ENTER = 4
)

func (self *SWS_TextAreaWidget) SetText(text string) {
	self.text=text
	PostUpdate()
}

func (self *SWS_TextAreaWidget) SetReadonly(readonly bool) {
	self.readonly=readonly
	self.SetColor(0xffdddddd)
	PostUpdate()
}

func (self *SWS_TextAreaWidget) HasFocus(focus bool) {
	self.hasfocus = focus
	PostUpdate()
}

func (self *SWS_TextAreaWidget) MousePressDown(x, y int32, button uint8) {
	if self.readonly==true {
		return
	}

	if button == sdl.BUTTON_LEFT {
		self.leftButtonDown = true
		self.UpdatePosition(x,y)
		self.endCursorPosition = self.initialCursorPosition
		PostUpdate()
	}
}

func (self *SWS_TextAreaWidget) MousePressUp(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.leftButtonDown = false
	}
}

func (self *SWS_TextAreaWidget) MouseMove(x, y, xrel, yrel int32) {
	if self.leftButtonDown == true {
		if (self.writeOffset>0 && y<0) {
			self.writeOffset-=int32(self.Font().Height())
			if self.writeOffset<0 {
				self.writeOffset=0
			}
		}
		if (self.yCursor>self.Height()-2) {
			//self.writeOffset+=(y-(self.Height()-2)-int32(self.Font().Height()))
			//self.writeOffset+=(y-(self.Height()-2))
			self.writeOffset+=int32(self.Font().Height())
		}
		self.UpdatePosition(x,y)
		
		PostUpdate()
	}
}

func (self *SWS_TextAreaWidget) KeyDown(key sdl.Keycode, mod uint16) {
	if self.readonly==true {
		return
	}
	if key == sdl.K_UP {
		self.UpdatePosition(self.xCursor,self.yCursor-int32(self.Font().Height()))
		if mod != sdl.KMOD_LSHIFT && mod != sdl.KMOD_RSHIFT {
			self.endCursorPosition = self.initialCursorPosition
		}
		PostUpdate()
	}
	if key == sdl.K_LEFT {
		if mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT {
			if self.initialCursorPosition > 0 {
				self.initialCursorPosition--
			}
		} else {
			if self.initialCursorPosition > 0 {
				self.initialCursorPosition--
			}
			self.endCursorPosition = self.initialCursorPosition
		}
		PostUpdate()
	}
	if key == sdl.K_RIGHT {
		if mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT {
			if self.initialCursorPosition < len(self.text) {
				self.initialCursorPosition++
			}
		} else {
			if self.initialCursorPosition < len(self.text) {
				self.initialCursorPosition++
			}
			self.endCursorPosition = self.initialCursorPosition
		}
		PostUpdate()
	}

	if key == sdl.K_DOWN {
		self.UpdatePosition(self.xCursor,self.yCursor+int32(self.Font().Height()))
		if mod != sdl.KMOD_LSHIFT && mod != sdl.KMOD_RSHIFT {
			self.endCursorPosition = self.initialCursorPosition
		}
		PostUpdate()
	}

	if key == sdl.K_BACKSPACE {
		if self.initialCursorPosition == self.endCursorPosition {
			if self.initialCursorPosition > 0 {
				self.initialCursorPosition--
				self.text = self.text[:self.initialCursorPosition] + self.text[self.initialCursorPosition+1:]
			}
		} else {
			i, e := self.initialCursorPosition, self.endCursorPosition
			if i > e {
				i, e = e, i
			}
			self.text = self.text[:i] + self.text[e:]
			self.initialCursorPosition = i
		}
		self.endCursorPosition = self.initialCursorPosition
		PostUpdate()
	}

	if mod == sdl.KMOD_LCTRL || mod == sdl.KMOD_RCTRL {
		if key=='a' {
			self.endCursorPosition = 0
			self.initialCursorPosition = len(self.text)
			PostUpdate()
		}
	} else if (key >= 'a' && key <= 'z') || (key >= '0' && key <= '9') || key == ' ' || key == sdl.K_TAB || key == sdl.K_RETURN {
		if key >= 'a' && key <= 'z' {
			if mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT {
				key += 'A' - 'a'
			}
		}
		if key == sdl.K_RETURN {
			key='\n'
		}
		if self.initialCursorPosition == self.endCursorPosition {
			self.text = self.text[:self.initialCursorPosition] + string(key) + self.text[self.initialCursorPosition:]
			self.initialCursorPosition++
		} else {
			i, e := self.initialCursorPosition, self.endCursorPosition
			if i > e {
				i, e = e, i
			}
			self.text = self.text[:i] + string(key) + self.text[e:]
			self.initialCursorPosition = i + 1
		}
		self.endCursorPosition = self.initialCursorPosition
		PostUpdate()
	}

	// recompute self.xCursor, self.yCursor
	self.parseText(renderWord)

	if self.writeOffset>0 && self.yCursor<2 {
		self.writeOffset-=int32(self.Font().Height())
		if self.writeOffset<0 {
			self.writeOffset=0
		}
	}
	if self.yCursor>self.Height()-2 {
		self.writeOffset+=int32(self.Font().Height())
	}
	PostUpdate()
}

type treatWord func(ta *SWS_TextAreaWidget, typeGlyph,x,y int32, word string, position int32)

func updatePositionWord(self *SWS_TextAreaWidget, typeGlyph, x,y int32, word string,position int32) {
	// we didn't reach the line
	if y+int32(self.Font().Height())<=self.yCursor {
		self.initialCursorPosition=int(position)
	// we are on the correct line
	} else if (y<=self.yCursor && y+int32(self.Font().Height())>self.yCursor) {
		if (x<=self.xCursor) {
			self.initialCursorPosition=int(position)
			width, _, _ := self.Font().SizeUTF8(word)
			if (x+int32(width)>self.xCursor) {
				for i,_:=range(word) {
					width, _, _ := self.Font().SizeUTF8(word[:i])
					if (x+int32(width)<=self.xCursor) {
						self.initialCursorPosition=int(position)+i
					}
				}
			}
		}
	}
}

func (self *SWS_TextAreaWidget) UpdatePosition(x,y int32) {
	if (y<2 && self.writeOffset==0) {
		self.initialCursorPosition=0
	} else {
		self.xCursor=x
		self.yCursor=y
		self.parseText(updatePositionWord)
	}
	PostUpdate()
}

func renderWord(self *SWS_TextAreaWidget, typeGlyph, x,y int32, word string,position int32) {
	i := self.initialCursorPosition
	e := self.endCursorPosition
	if i > e {
		i, e = e, i
	}

	if (typeGlyph!=GLYPH_TAB) { //aka not tab
		// i < word < e
		if (i<=int(position) && e>=int(position)+len(word)) {
			width, _, _ := self.Font().SizeUTF8(word)
			self.FillRect(x,y, int32(width), int32(self.Font().Height()), 0xff8888ff)
		} else if (i<=int(position)+len(word) && e>=int(position)+len(word)) {
			widthLeft, _, _ := self.Font().SizeUTF8(word[:i-int(position)])
			widthRight, _, _ := self.Font().SizeUTF8(word[i-int(position):])
			self.FillRect(x+int32(widthLeft),y, int32(widthRight), int32(self.Font().Height()), 0xff8888ff)
		} else if (i<=int(position) && e>=int(position)) {
			widthLeft, _, _ := self.Font().SizeUTF8(word[:e-int(position)])
			self.FillRect(x,y, int32(widthLeft), int32(self.Font().Height()), 0xff8888ff)
		} else if (i<=int(position)+len(word) && e>=int(position)) {
			widthLeft, _, _ := self.Font().SizeUTF8(word[:i-int(position)])
			widthRight, _, _ := self.Font().SizeUTF8(word[i-int(position):e-int(position)])
			self.FillRect(x+int32(widthLeft),y, int32(widthRight), int32(self.Font().Height()), 0xff8888ff)
		}
		// compute x,y cursor position
		if (self.initialCursorPosition>=int(position) && self.initialCursorPosition<=int(position)+len(word)) {
			width,_,_ := self.Font().SizeUTF8(word[:self.initialCursorPosition-int(position)])
			self.xCursor=x+int32(width)
			self.yCursor=y
		}
		self.WriteText(x,y, word, sdl.Color{0, 0, 0, 255})
	} else { // TAB
		if (i<=int(position) && e>=int(position)+1) {
			self.FillRect(x,y, 2+(((x-2+32)>>5)<<5)-x, int32(self.Font().Height()), 0xff8888ff)
		}
		// compute x,y cursor position
		if (self.initialCursorPosition==int(position)) {
			self.xCursor=x
			self.yCursor=y
		}
	}

	// we print the cursor
	if (self.hasfocus) {
		if int32(self.endCursorPosition)>=position && int32(self.endCursorPosition)<position+int32(len(word)) {
			for i,_ := range(word) {
				if position+int32(i) == int32(self.endCursorPosition) {
					width, _, _ := self.Font().SizeUTF8(word[:i])
					self.SetDrawColor(0, 0, 0, 255)
					self.DrawLine(x+int32(width),y,x+int32(width),y+int32(self.Font().Height()))
				}
			}
		}
	}
	if (self.hasfocus) {
		if ((typeGlyph==GLYPH_END || typeGlyph==GLYPH_ENTER) && int32(self.endCursorPosition)==position) { // word is empty: enter (4) or end of text (3)
			self.SetDrawColor(0, 0, 0, 255)
			self.DrawLine(x,y,x,y+int32(self.Font().Height()))
		}
	}
}

func (self *SWS_TextAreaWidget) renderText(typeGlyph int32, word string, x,y *int32,position int32, treat treatWord) {
	if (typeGlyph==GLYPH_SPACE) { // space
		width, _, _ := self.Font().SizeUTF8(word)
		treat(self,typeGlyph,*x,*y, word, position)
		*x = *x+int32(width)
	} else if (typeGlyph==GLYPH_TAB) { //tab
		if ((*x-2+32)>>5)<<5 >= self.Width()-4 {
			*x=2
			*y+=int32(self.Font().Height())
			treat(self,typeGlyph,*x,*y, word, position)
			*x=34
		} else {
			treat(self,typeGlyph,*x,*y, word, position)
			*x=2+((*x-2+32)>>5)<<5
		}
	} else if (typeGlyph==GLYPH_ENTER) { //enter
		treat(self,typeGlyph,*x,*y, "", position)
		*x=2
		*y+=int32(self.Font().Height())
	} else if (typeGlyph==GLYPH_WORD) { // word
		width, _, _ := self.Font().SizeUTF8(word)
		if (*x+int32(width)) > self.Width()-4 && *x>2 {
			*x=2
			*y+=int32(self.Font().Height())
		}
		// the word is longer than the line
		for int32(width) > self.Width()-4 {
			subword:=""
			for i,c := range (word) {
				subwidth,_,_:=self.Font().SizeUTF8(word[:i+1])
				if (int32(subwidth)<self.Width()-4) {
					subword=subword+string(c)
				} else {
					break
				}
			}
			treat(self,typeGlyph,*x,*y, subword, position)
			*y+=int32(self.Font().Height())
			word=word[len(subword):]
			position+=int32(len(subword))
			width, _, _ = self.Font().SizeUTF8(word)
		}
		treat(self,typeGlyph,*x,*y, word, position)
		*x+=int32(width)
	} else if (typeGlyph==GLYPH_END) { // end text
		treat(self,typeGlyph,*x,*y, "", position)
	}
}

func (self *SWS_TextAreaWidget) parseText(treat treatWord) {
	// cut the string into 5 types: word, tab, enter, multiple spaces, and end of string
	// for each boundary: we have to render it:
	// if word we check if we have to go to the next line
	// space we alway render on the same line 
	// tab is like word
	// enter is a bit like tab
	var x,y,typeGlyph,position int32
	x=2
	y=2-self.writeOffset
	typeGlyph=-1
	word:=""
	for currentpos,char := range (self.text) {
		if char==' ' && typeGlyph!=GLYPH_SPACE { // space
			self.renderText(typeGlyph,word,&x,&y,position,treat)
			typeGlyph=GLYPH_SPACE
			word=""
			position=int32(currentpos)
		} else if char=='\t' { // tab
			self.renderText(typeGlyph,word,&x,&y,position,treat)
			typeGlyph=GLYPH_TAB
			word=""
			position=int32(currentpos)
		} else if char=='\n' { // enter
			self.renderText(typeGlyph,word,&x,&y,position,treat)
			typeGlyph=GLYPH_ENTER
			word=""
			position=int32(currentpos)
		} else if typeGlyph !=GLYPH_WORD { // word
			self.renderText(typeGlyph,word,&x,&y,position,treat)
			typeGlyph=GLYPH_WORD
			word=""
			position=int32(currentpos)
		}
		word=word+string(char)
	}
	self.renderText(typeGlyph,word,&x,&y,position,treat)
	// for the cursor at the end
	self.renderText(GLYPH_END,"",&x,&y,int32(len(self.text)),treat)
	// replace cursor (if it is at the end)
	if (self.yCursor>=y) {
		if (self.xCursor>=x) {
			self.yCursor=y
			self.xCursor=x
		} else {
			self.xCursor=x
			self.yCursor=y
		}
	}
}


func (self *SWS_TextAreaWidget) Repaint() {
	self.SWS_CoreWidget.Repaint()

	self.parseText(renderWord)

	// write boundaries
	if self.readonly==false {
		self.SetDrawColor(0, 0, 0, 255)
		self.DrawLine(0, 2, 0, self.Height()-3)
		self.DrawLine(self.Width()-1, 2, self.Width()-1, self.Height()-3)
		self.DrawLine(2, 0, self.Width()-3, 0)
		self.DrawLine(2, self.Height()-1, self.Width()-3, self.Height()-1)
		self.DrawPoint(1, self.Height()-2)
		self.DrawPoint(1, 1)
		self.DrawPoint(self.Width()-2, 1)
		self.DrawPoint(self.Width()-2, self.Height()-2)

		self.SetDrawColor(0xdd, 0xdd, 0xdd, 255)
		self.DrawLine(2, 1, self.Width()-3, 1)
		self.DrawLine(1, 2, 1, self.Height()-3)
		self.DrawPoint(0, 0)
		self.DrawPoint(0, 1)
		self.DrawPoint(1, 0)
	
		self.DrawPoint(self.Width()-1, 0)
		self.DrawPoint(self.Width()-1, 1)
		self.DrawPoint(self.Width()-2, 0)
	
		self.DrawPoint(0, self.Height()-1)
		self.DrawPoint(0, self.Height()-2)
		self.DrawPoint(1, self.Height()-1)
	
		self.DrawPoint(self.Width()-1, self.Height()-1)
		self.DrawPoint(self.Width()-1, self.Height()-2)
		self.DrawPoint(self.Width()-2, self.Height()-1)
	}

}

func CreateTextAreaWidget(w, h int32, s string) *SWS_TextAreaWidget {
	corewidget := CreateCoreWidget(w, h)
	widget := &SWS_TextAreaWidget{SWS_CoreWidget: *corewidget,
		text: s,
		initialCursorPosition: 0,
		hasfocus:              false,
		leftButtonDown:        false,
		writeOffset:           0,
		readonly:              false}
	widget.SetColor(0xffffffff)
	return widget
}
