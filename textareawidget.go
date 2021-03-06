package sws

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type TextAreaWidget struct {
	CoreWidget
	text                  string
	initialCursorPosition int   // begin of selection
	endCursorPosition     int   // end of selection
	xCursor               int32 // position of the selection but in x/y pixels
	yCursor               int32 // position of the selection but in x/y pixels
	hasfocus              bool
	leftButtonDown        bool
	writeOffset           int32 // how far we scroll the text
	disabled              bool
	internalcolor         uint32
	totalHeight           int32
	heightChangeCallback  func(int32)
	yoffsetChangeCallback func(int32)
	showLineNumber        bool
	previousHeight        int32
}

const (
	GLYPH_SPACE = 0
	GLYPH_TAB   = 1
	GLYPH_WORD  = 2
	GLYPH_END   = 3
	GLYPH_ENTER = 4
)

func (self *TextAreaWidget) ShowLineNumber() {
	self.showLineNumber = true
	self.PostUpdate()
}

// SetHeightChangeCallback is used to inform when the text grows/shrink in term of number of lines
func (self *TextAreaWidget) SetHeightChangeCallback(callback func(int32)) {
	self.heightChangeCallback = callback
}

// SetHeightChangeCallback is used to inform when we scroll down/up
func (self *TextAreaWidget) SetYOffsetChangedCallback(callback func(int32)) {
	self.yoffsetChangeCallback = callback
}

func (self *TextAreaWidget) SetYOffset(offset int32) {
	if self.writeOffset != offset {
		self.writeOffset = offset
		self.PostUpdate()
		if self.yoffsetChangeCallback != nil {
			self.yoffsetChangeCallback(offset)
		}
	}
}

func (self *TextAreaWidget) Resize(w, h int32) {
	self.CoreWidget.Resize(w, h)
	if self.previousHeight != h {
		self.heightChangeCallback(self.totalHeight)
		self.previousHeight = h
	}
}

func (self *TextAreaWidget) IsInputWidget() bool {
	if self.disabled == true {
		return false
	}
	return true
}

func (self *TextAreaWidget) SetText(text string) {
	self.text = text
	if self.valueChangedCallback != nil {
		self.valueChangedCallback()
	}
	self.PostUpdate()
}

func (self *TextAreaWidget) GetText() string {
	return self.text
}

func (self *TextAreaWidget) SetDisabled(disabled bool) {
	self.disabled = disabled
	self.internalcolor = 0xffdddddd
	self.PostUpdate()
}

func (self *TextAreaWidget) HasFocus(focus bool) {
	self.hasfocus = focus
	self.PostUpdate()
}

func (self *TextAreaWidget) MousePressDown(x, y int32, button uint8) {
	if self.disabled == true {
		return
	}

	if button == sdl.BUTTON_LEFT {
		self.leftButtonDown = true
		self.UpdatePosition(x, y)
		self.endCursorPosition = self.initialCursorPosition
		self.PostUpdate()
	}
}

func (self *TextAreaWidget) MousePressUp(x, y int32, button uint8) {
	if button == sdl.BUTTON_LEFT {
		self.leftButtonDown = false
	}
}

func (self *TextAreaWidget) MouseMove(x, y, xrel, yrel int32) {
	if self.leftButtonDown == true {
		if self.writeOffset > 0 && y < 0 {
			self.writeOffset -= int32(self.Font().Height())
			if self.writeOffset < 0 {
				self.writeOffset = 0
			}
			if self.yoffsetChangeCallback != nil {
				self.yoffsetChangeCallback(self.writeOffset)
			}
		}
		if self.yCursor > self.Height()-3 {
			//self.writeOffset+=(y-(self.Height()-2)-int32(self.Font().Height()))
			//self.writeOffset+=(y-(self.Height()-2))
			self.writeOffset += int32(self.Font().Height())
			if self.yoffsetChangeCallback != nil {
				self.yoffsetChangeCallback(self.writeOffset)
			}
		}
		self.UpdatePosition(x, y)

		self.PostUpdate()
	}
}

func (self *TextAreaWidget) InputText(text string) {
	if self.disabled == true {
		return
	}
	if self.initialCursorPosition == self.endCursorPosition {
		self.text = self.text[:self.initialCursorPosition] + text + self.text[self.initialCursorPosition:]
		self.initialCursorPosition += len(text)
	} else {
		i, e := self.initialCursorPosition, self.endCursorPosition
		if i > e {
			i, e = e, i
		}
		self.text = self.text[:i] + text + self.text[e:]
		self.initialCursorPosition = i + 1
	}
	self.endCursorPosition = self.initialCursorPosition
	self.PostUpdate()
	if self.valueChangedCallback != nil {
		self.valueChangedCallback()
	}
}

func (self *TextAreaWidget) KeyDown(key sdl.Keycode, mod uint16) {
	if key == sdl.K_TAB && (mod == sdl.KMOD_CTRL || mod == sdl.KMOD_RCTRL || mod == sdl.KMOD_LCTRL || mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT) {
		if self.focusOnNextInputWidgetCallback != nil {
			self.focusOnNextInputWidgetCallback(!(mod == sdl.KMOD_LSHIFT || mod == sdl.KMOD_RSHIFT))
			return
		}
	}
	if self.disabled == true {
		return
	}
	if key == sdl.K_UP {
		self.UpdatePosition(self.xCursor, self.yCursor-int32(self.Font().Height()))
		if mod != sdl.KMOD_LSHIFT && mod != sdl.KMOD_RSHIFT {
			self.endCursorPosition = self.initialCursorPosition
		}
		self.PostUpdate()
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
		self.PostUpdate()
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
		self.PostUpdate()
	}

	if key == sdl.K_DOWN {
		self.UpdatePosition(self.xCursor, self.yCursor+int32(self.Font().Height()))
		if mod != sdl.KMOD_LSHIFT && mod != sdl.KMOD_RSHIFT {
			self.endCursorPosition = self.initialCursorPosition
		}
		self.PostUpdate()
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
		self.PostUpdate()
		if self.valueChangedCallback != nil {
			self.valueChangedCallback()
		}
	}

	if ((mod & (sdl.KMOD_CTRL | sdl.KMOD_GUI)) != 0) && key == 'x' {
		i, e := self.initialCursorPosition, self.endCursorPosition
		if i > e {
			i, e = e, i
		}
		clipboard := self.text[i:e]
		sdl.SetClipboardText(clipboard)

		self.text = self.text[:i] + self.text[e:]
		self.initialCursorPosition = i
		self.endCursorPosition = self.initialCursorPosition
		self.PostUpdate()
		if self.valueChangedCallback != nil {
			self.valueChangedCallback()
		}
	}

	if ((mod & (sdl.KMOD_CTRL | sdl.KMOD_GUI)) != 0) && key == 'v' {
		if clipboard, err := sdl.GetClipboardText(); err == nil {
			i, e := self.initialCursorPosition, self.endCursorPosition
			if i > e {
				i, e = e, i
			}
			self.text = self.text[:i] + clipboard + self.text[e:]
			self.initialCursorPosition = i + len(clipboard)
			self.endCursorPosition = self.initialCursorPosition
			self.PostUpdate()
			if self.valueChangedCallback != nil {
				self.valueChangedCallback()
			}
		}
	}

	if ((mod & (sdl.KMOD_CTRL | sdl.KMOD_GUI)) != 0) && key == 'c' {
		i, e := self.initialCursorPosition, self.endCursorPosition
		if i > e {
			i, e = e, i
		}
		clipboard := self.text[i:e]
		sdl.SetClipboardText(clipboard)
	}

	if ((mod & (sdl.KMOD_CTRL | sdl.KMOD_GUI)) != 0) && key == 'a' {
		self.endCursorPosition = 0
		self.initialCursorPosition = len(self.text)
		self.PostUpdate()
	}

	if key == sdl.K_TAB || key == sdl.K_RETURN {
		if key == sdl.K_RETURN {
			key = '\n'
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
		self.PostUpdate()
		if self.valueChangedCallback != nil {
			self.valueChangedCallback()
		}
	}

	// recompute self.xCursor, self.yCursor
	self.parseText(renderWord)

	if self.writeOffset > 0 && self.yCursor < 3 {
		self.writeOffset -= int32(self.Font().Height())
		if self.writeOffset < 0 {
			self.writeOffset = 0
		}
		if self.yoffsetChangeCallback != nil {
			self.yoffsetChangeCallback(self.writeOffset)
		}
	}
	if self.yCursor > self.Height()-3 {
		self.writeOffset += int32(self.Font().Height())
		if self.yoffsetChangeCallback != nil {
			self.yoffsetChangeCallback(self.writeOffset)
		}
	}
	self.PostUpdate()
}

type treatWord func(ta *TextAreaWidget, typeGlyph, x, y int32, word string, position int32)

func updatePositionWord(self *TextAreaWidget, typeGlyph, x, y int32, word string, position int32) {
	// we didn't reach the line
	if y+int32(self.Font().Height()) <= self.yCursor {
		self.initialCursorPosition = int(position)
		// we are on the correct line
	} else if y <= self.yCursor && y+int32(self.Font().Height()) > self.yCursor {
		if x <= self.xCursor {
			self.initialCursorPosition = int(position)
			width, _, _ := self.Font().SizeUTF8(word)
			if x+int32(width) > self.xCursor {
				for i, _ := range word {
					width, _, _ := self.Font().SizeUTF8(word[:i])
					if x+int32(width) <= self.xCursor {
						self.initialCursorPosition = int(position) + i
					}
				}
			}
		}
	}
}

func (self *TextAreaWidget) UpdatePosition(x, y int32) {
	if y < 3 && self.writeOffset == 0 {
		self.initialCursorPosition = 0
	} else {
		self.xCursor = x
		self.yCursor = y
		self.parseText(updatePositionWord)
	}
	self.PostUpdate()
}

func renderWord(self *TextAreaWidget, typeGlyph, x, y int32, word string, position int32) {
	i := self.initialCursorPosition
	e := self.endCursorPosition
	if i > e {
		i, e = e, i
	}

	if typeGlyph != GLYPH_TAB { //aka not tab
		// i < word < e
		if i <= int(position) && e >= int(position)+len(word) {
			width, _, _ := self.Font().SizeUTF8(word)
			self.FillRect(x, y, int32(width), int32(self.Font().Height()), 0xff8888ff)
		} else if i <= int(position)+len(word) && e >= int(position)+len(word) {
			widthLeft, _, _ := self.Font().SizeUTF8(word[:i-int(position)])
			widthRight, _, _ := self.Font().SizeUTF8(word[i-int(position):])
			self.FillRect(x+int32(widthLeft), y, int32(widthRight), int32(self.Font().Height()), 0xff8888ff)
		} else if i <= int(position) && e >= int(position) {
			widthLeft, _, _ := self.Font().SizeUTF8(word[:e-int(position)])
			self.FillRect(x, y, int32(widthLeft), int32(self.Font().Height()), 0xff8888ff)
		} else if i <= int(position)+len(word) && e >= int(position) {
			widthLeft, _, _ := self.Font().SizeUTF8(word[:i-int(position)])
			widthRight, _, _ := self.Font().SizeUTF8(word[i-int(position) : e-int(position)])
			self.FillRect(x+int32(widthLeft), y, int32(widthRight), int32(self.Font().Height()), 0xff8888ff)
		}
		// compute x,y cursor position
		if self.initialCursorPosition >= int(position) && self.initialCursorPosition <= int(position)+len(word) {
			width, _, _ := self.Font().SizeUTF8(word[:self.initialCursorPosition-int(position)])
			self.xCursor = x + int32(width)
			self.yCursor = y
		}
		self.WriteText(x, y, word, sdl.Color{0, 0, 0, 255})
	} else { // TAB
		if i <= int(position) && e >= int(position)+1 {
			self.FillRect(x, y, 3+(((x-3+32)>>5)<<5)-x, int32(self.Font().Height()), 0xff8888ff)
		}
		// compute x,y cursor position
		if self.initialCursorPosition == int(position) {
			self.xCursor = x
			self.yCursor = y
		}
	}

	// we print the cursor
	if self.hasfocus {
		if int32(self.endCursorPosition) >= position && int32(self.endCursorPosition) < position+int32(len(word)) {
			for i, _ := range word {
				if position+int32(i) == int32(self.endCursorPosition) {
					width, _, _ := self.Font().SizeUTF8(word[:i])
					self.SetDrawColor(0, 0, 0, 255)
					yEnd := y + int32(self.Font().Height())
					if yEnd > self.height-2 {
						yEnd = self.height - 2
					}
					self.DrawLine(x+int32(width), y, x+int32(width), yEnd)
				}
			}
		}
	}
	if self.hasfocus {
		if (typeGlyph == GLYPH_END || typeGlyph == GLYPH_ENTER) && int32(self.endCursorPosition) == position { // word is empty: enter (4) or end of text (3)
			self.SetDrawColor(0, 0, 0, 255)
			yEnd := y + int32(self.Font().Height())
			if yEnd > self.height-2 {
				yEnd = self.height - 2
			}
			self.DrawLine(x, y, x, yEnd)
		}
	}
}

func (self *TextAreaWidget) renderText(typeGlyph int32, word string, x, y, linenumber *int32, position int32, treat treatWord) {
	if typeGlyph == GLYPH_SPACE { // space
		width, _, _ := self.Font().SizeUTF8(word)
		treat(self, typeGlyph, *x, *y, word, position)
		*x = *x + int32(width)
	} else if typeGlyph == GLYPH_TAB { //tab
		if ((*x-3+32)>>5)<<5 >= self.Width()-6 {
			*x = 3
			*y += int32(self.Font().Height())
			if self.showLineNumber {
				*x += 64
			}
			treat(self, typeGlyph, *x, *y, word, position)
			*x = 32 + 3
			if self.showLineNumber {
				*x += 64
			}
		} else {
			treat(self, typeGlyph, *x, *y, word, position)
			*x = 3 + ((*x-3+32)>>5)<<5
		}
	} else if typeGlyph == GLYPH_ENTER { //enter
		treat(self, typeGlyph, *x, *y, "", position)
		*x = 3
		*y += int32(self.Font().Height())
		if self.showLineNumber {
			*linenumber++
			self.WriteText(3, *y, fmt.Sprintf("%3d", *linenumber%1000), sdl.Color{0, 0, 0, 255})
			*x += 64
		}
	} else if typeGlyph == GLYPH_WORD { // word
		width, _, _ := self.Font().SizeUTF8(word)
		if (*x+int32(width)) > self.Width()-6 && *x > 3 {
			*x = 3
			*y += int32(self.Font().Height())
			if self.showLineNumber {
				*x += 64
			}
		}
		// the word is longer than the line
		for int32(width) > self.Width()-6 && self.Width() > 0 {
			subword := ""
			for i, c := range word {
				subwidth, _, _ := self.Font().SizeUTF8(word[:i+1])
				if int32(subwidth) < self.Width()-6 || len(subword) == 0 {
					subword = subword + string(c)
				} else {
					break
				}
			}
			treat(self, typeGlyph, *x, *y, subword, position)
			*y += int32(self.Font().Height())
			word = word[len(subword):]
			position += int32(len(subword))
			width, _, _ = self.Font().SizeUTF8(word)
		}
		treat(self, typeGlyph, *x, *y, word, position)
		*x += int32(width)
	} else if typeGlyph == GLYPH_END { // end text
		treat(self, typeGlyph, *x, *y, "", position)
	}
}

func (self *TextAreaWidget) parseText(treat treatWord) {
	// cut the string into 5 types: word, tab, enter, multiple spaces, and end of string
	// for each boundary: we have to render it:
	// if word we check if we have to go to the next line
	// space we alway render on the same line
	// tab is like word
	// enter is a bit like tab
	var x, y, linenumber, typeGlyph, position int32
	linenumber = 1
	x = 3
	y = 3 - self.writeOffset
	if self.showLineNumber {
		self.WriteText(3, y, fmt.Sprintf("%3d", linenumber%1000), sdl.Color{0, 0, 0, 255})
		x += 64
	}
	typeGlyph = -1
	word := ""
	for currentpos, char := range self.text {
		if char == ' ' && typeGlyph != GLYPH_SPACE { // space
			self.renderText(typeGlyph, word, &x, &y, &linenumber, position, treat)
			typeGlyph = GLYPH_SPACE
			word = ""
			position = int32(currentpos)
		} else if char == '\t' { // tab
			self.renderText(typeGlyph, word, &x, &y, &linenumber, position, treat)
			typeGlyph = GLYPH_TAB
			word = ""
			position = int32(currentpos)
		} else if char == '\n' { // enter
			self.renderText(typeGlyph, word, &x, &y, &linenumber, position, treat)
			typeGlyph = GLYPH_ENTER
			word = ""
			position = int32(currentpos)
		} else if typeGlyph != GLYPH_WORD { // word
			self.renderText(typeGlyph, word, &x, &y, &linenumber, position, treat)
			typeGlyph = GLYPH_WORD
			word = ""
			position = int32(currentpos)
		}
		word = word + string(char)
	}
	self.renderText(typeGlyph, word, &x, &y, &linenumber, position, treat)
	// for the cursor at the end
	self.renderText(GLYPH_END, "", &x, &y, &linenumber, int32(len(self.text)), treat)
	// replace cursor (if it is at the end)
	if self.yCursor >= y {
		if self.xCursor >= x {
			self.yCursor = y
			self.xCursor = x
		} else {
			self.xCursor = x
			self.yCursor = y
		}
	}
	if self.totalHeight != y-(3-self.writeOffset)+int32(self.Font().Height()) {
		self.totalHeight = y - (3 - self.writeOffset) + int32(self.Font().Height())
		if self.heightChangeCallback != nil {
			self.heightChangeCallback(self.totalHeight)
		}
	}
}

func (self *TextAreaWidget) Repaint() {
	self.CoreWidget.Repaint()

	if self.showLineNumber {
		self.FillRect(65, 2, self.width-67, self.height-4, self.internalcolor)
	} else {
		self.FillRect(2, 2, self.width-4, self.height-4, self.internalcolor)
	}
	self.parseText(renderWord)

	if self.showLineNumber {
		self.SetDrawColor(0, 0, 0, 255)
		self.DrawLine(65, 3, 65, self.Height()-4)
	}

	// write boundaries
	if self.disabled == false {
		self.SetDrawColor(0, 0, 0, 255)
		self.DrawLine(1, 3, 1, self.Height()-4)
		self.DrawLine(self.Width()-2, 3, self.Width()-2, self.Height()-4)
		self.DrawLine(3, 1, self.Width()-4, 1)
		self.DrawLine(3, self.Height()-2, self.Width()-4, self.Height()-2)
		self.DrawPoint(2, self.Height()-3)
		self.DrawPoint(2, 2)
		self.DrawPoint(self.Width()-3, 2)
		self.DrawPoint(self.Width()-3, self.Height()-3)

		self.SetDrawColor(0xdd, 0xdd, 0xdd, 255)
		self.DrawLine(3, 2, self.Width()-4, 2)
		self.DrawLine(2, 3, 2, self.Height()-4)
		self.DrawPoint(1, 1)
		self.DrawPoint(1, 2)
		self.DrawPoint(2, 1)

		self.DrawPoint(self.Width()-2, 1)
		self.DrawPoint(self.Width()-2, 2)
		self.DrawPoint(self.Width()-3, 1)

		self.DrawPoint(1, self.Height()-2)
		self.DrawPoint(1, self.Height()-3)
		self.DrawPoint(2, self.Height()-2)

		self.DrawPoint(self.Width()-2, self.Height()-2)
		self.DrawPoint(self.Width()-2, self.Height()-3)
		self.DrawPoint(self.Width()-3, self.Height()-2)
	}

	// draw the bezel
	if self.hasfocus && self.disabled == false {
		self.SetDrawColor(0x46, 0xc8, 0xe8, 255)
	}
	self.DrawLine(0, 3, 0, self.Height()-4)
	self.DrawPoint(1, self.Height()-3)
	self.DrawPoint(2, self.Height()-2)
	self.DrawLine(3, self.Height()-1, self.Width()-4, self.Height()-1)
	self.DrawPoint(self.Width()-3, self.Height()-2)
	self.DrawPoint(self.Width()-2, self.Height()-3)
	self.DrawLine(self.Width()-1, self.Height()-4, self.Width()-1, 3)
	self.DrawPoint(self.Width()-3, 1)
	self.DrawPoint(self.Width()-2, 2)
	self.DrawLine(self.Width()-4, 0, 3, 0)
	self.DrawPoint(2, 1)
	self.DrawPoint(1, 2)
}

func NewTextAreaWidget(w, h int32, s string) *TextAreaWidget {
	corewidget := NewCoreWidget(w, h)
	widget := &TextAreaWidget{CoreWidget: *corewidget,
		text: s,
		initialCursorPosition: 0,
		hasfocus:              false,
		leftButtonDown:        false,
		writeOffset:           0,
		disabled:              false,
		internalcolor:         0xffffffff,
		showLineNumber:        false,
		previousHeight:        h,
	}
	return widget
}
