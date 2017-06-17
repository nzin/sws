package sws

import (
        "github.com/veandco/go-sdl2/sdl"
        "github.com/veandco/go-sdl2/sdl_ttf"
	"unsafe"
)


var defaultFont *ttf.Font
var LatoRegular20 *ttf.Font
var LatoRegular24 *ttf.Font

func InitFonts() {
	rwops := sdl.RWFromMem(unsafe.Pointer(&latoRegular[0]), len(latoRegular))
	defaultFont,_ = ttf.OpenFontRW(rwops,1,16)
	LatoRegular20,_ = ttf.OpenFontRW(rwops,1,20)
	LatoRegular24,_ = ttf.OpenFontRW(rwops,1,24)
}

var latoRegular=
