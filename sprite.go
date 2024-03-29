package sws

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var mainlefths *sdl.Surface
var mainlefth = []byte("\x42\x4d\x8a\x04\x00\x00\x00\x00\x00\x00\x8a\x00\x00\x00\x7c\x00" +
	"\x00\x00\x10\x00\x00\x00\x10\x00\x00\x00\x01\x00\x20\x00\x03\x00" +
	"\x00\x00\x00\x04\x00\x00\x25\x16\x00\x00\x25\x16\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\x00\x00\xff\x00\x00\xff" +
	"\x00\x00\xff\x00\x00\x00\x42\x47\x52\x73\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x98\xfe\xff\xff\x9a\xff\xff\xff\x99\xfe\xff\xff\x99" +
	"\xfe\xff\xff\x99\xfe\xff\xff\x99\xfe\xff\xff\x99\xfe\xff\xff\x99" +
	"\xfe\xff\xff\x99\xfe\xff\xff\x99\xfe\xff\xff\x99\xfe\xff\xff\x99" +
	"\xfe\xff\xff\x9a\xff\xff\xff\x8f\xfa\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\xa3\xff\xff\xff\x00\xa7\xb3\xff\x00\xac\xb8\xff\x00" +
	"\xac\xb8\xff\x00\xac\xb8\xff\x00\xad\xb8\xff\x00\xad\xb8\xff\x00" +
	"\xad\xb9\xff\x00\xad\xb9\xff\x00\xad\xb9\xff\x00\xae\xb9\xff\x00" +
	"\xae\xb9\xff\x00\xaa\xb6\xff\x9a\xff\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9d\xff\xff\xff\x4d\xdf\xe5\xff\x48\xdd\xe3\xff\x3f" +
	"\xda\xe0\xff\x37\xd7\xde\xff\x2e\xd4\xda\xff\x25\xd1\xd7\xff\x17" +
	"\xcd\xd4\xff\x05\xca\xd3\xff\x00\xc7\xd0\xff\x00\xc4\xcd\xff\x00" +
	"\xc1\xc9\xff\x00\xae\xb9\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9d\xff\xff\xff\x4e\xe0\xe6\xff\x4c\xde\xe4\xff\x44" +
	"\xdb\xe1\xff\x3c\xd8\xde\xff\x34\xd5\xdc\xff\x2c\xd2\xd9\xff\x24" +
	"\xcf\xd6\xff\x16\xcc\xd3\xff\x05\xc9\xd2\xff\x00\xc6\xcf\xff\x00" +
	"\xc4\xcd\xff\x00\xae\xb9\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9c\xff\xff\xff\x56\xe3\xe8\xff\x52\xe1\xe7\xff\x4c" +
	"\xde\xe4\xff\x44\xdb\xe1\xff\x3c\xd8\xde\xff\x34\xd5\xdc\xff\x2c" +
	"\xd2\xd9\xff\x24\xcf\xd6\xff\x16\xcc\xd3\xff\x05\xc9\xd2\xff\x00" +
	"\xc7\xd0\xff\x00\xad\xb9\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9c\xff\xff\xff\x5f\xe5\xeb\xff\x59\xe4\xe9\xff\x52" +
	"\xe1\xe7\xff\x4c\xde\xe4\xff\x44\xdb\xe1\xff\x3c\xd8\xde\xff\x34" +
	"\xd5\xdc\xff\x2c\xd2\xd9\xff\x24\xcf\xd6\xff\x16\xcc\xd3\xff\x05" +
	"\xca\xd3\xff\x00\xad\xb9\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9b\xff\xff\xff\x66\xe8\xee\xff\x61\xe6\xec\xff\x59" +
	"\xe4\xe9\xff\x52\xe1\xe7\xff\x4c\xde\xe4\xff\x44\xdb\xe1\xff\x3c" +
	"\xd8\xde\xff\x34\xd5\xdc\xff\x2c\xd2\xd9\xff\x24\xcf\xd6\xff\x17" +
	"\xcd\xd4\xff\x00\xad\xb9\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9b\xff\xff\xff\x6c\xeb\xf1\xff\x68\xe9\xef\xff\x61" +
	"\xe6\xec\xff\x59\xe4\xe9\xff\x52\xe1\xe7\xff\x4c\xde\xe4\xff\x44" +
	"\xdb\xe1\xff\x3c\xd8\xde\xff\x34\xd5\xdc\xff\x2c\xd2\xd9\xff\x25" +
	"\xd1\xd7\xff\x00\xad\xb8\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9b\xff\xff\xff\x74\xee\xf3\xff\x6e\xec\xf2\xff\x68" +
	"\xe9\xef\xff\x61\xe6\xec\xff\x59\xe4\xe9\xff\x52\xe1\xe7\xff\x4c" +
	"\xde\xe4\xff\x44\xdb\xe1\xff\x3c\xd8\xde\xff\x34\xd5\xdc\xff\x2e" +
	"\xd4\xda\xff\x00\xad\xb8\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9a\xff\xff\xff\x7c\xf2\xf7\xff\x75\xef\xf4\xff\x6e" +
	"\xec\xf2\xff\x68\xe9\xef\xff\x61\xe6\xec\xff\x59\xe4\xe9\xff\x52" +
	"\xe1\xe7\xff\x4c\xde\xe4\xff\x44\xdb\xe1\xff\x3c\xd8\xde\xff\x37" +
	"\xd7\xde\xff\x00\xac\xb8\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9a\xff\xff\xff\x83\xf5\xfa\xff\x7d\xf2\xf7\xff\x75" +
	"\xef\xf4\xff\x6e\xec\xf2\xff\x68\xe9\xef\xff\x61\xe6\xec\xff\x59" +
	"\xe4\xe9\xff\x52\xe1\xe7\xff\x4c\xde\xe4\xff\x44\xdb\xe1\xff\x3f" +
	"\xda\xe0\xff\x00\xac\xb8\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x99\xff\xff\xff\x89\xf8\xfd\xff\x83\xf5\xfa\xff\x7d" +
	"\xf2\xf7\xff\x75\xef\xf4\xff\x6e\xec\xf2\xff\x68\xe9\xef\xff\x61" +
	"\xe6\xec\xff\x59\xe4\xe9\xff\x52\xe1\xe7\xff\x4c\xde\xe4\xff\x48" +
	"\xdd\xe3\xff\x00\xac\xb8\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x99\xfe\xff\xff\x91\xfa\xff\xff\x89\xf8\xfd\xff\x83" +
	"\xf5\xfa\xff\x7c\xf2\xf7\xff\x74\xee\xf3\xff\x6c\xeb\xf1\xff\x66" +
	"\xe8\xee\xff\x5f\xe5\xeb\xff\x56\xe3\xe8\xff\x4e\xe0\xe6\xff\x4d" +
	"\xdf\xe5\xff\x00\xa7\xb3\xff\x9a\xff\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xae\xba\xff\xa3\xff\xff\xff\x99\xfe\xff\xff\x99\xff\xff\xff\x9a" +
	"\xff\xff\xff\x9a\xff\xff\xff\x9b\xff\xff\xff\x9b\xff\xff\xff\x9b" +
	"\xff\xff\xff\x9c\xff\xff\xff\x9c\xff\xff\xff\x9d\xff\xff\xff\x9d" +
	"\xff\xff\xff\xa3\xff\xff\xff\x98\xfe\xff\xff\x0b\xf1\xff\xff\x01" +
	"\xb3\xbe\xff\x00\xae\xba\xff\x00\xaf\xba\xff\x00\xaf\xba\xff\x00" +
	"\xaf\xba\xff\x00\xaf\xba\xff\x00\xaf\xba\xff\x00\xaf\xba\xff\x00" +
	"\xaf\xba\xff\x00\xaf\xba\xff\x00\xaf\xba\xff\x00\xaf\xba\xff\x00" +
	"\xaf\xba\xff\x00\xaf\xba\xff\x00\xaf\xba\xff\x0b\xf1\xff\xff\x0b" +
	"\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b" +
	"\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b" +
	"\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b" +
	"\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b")

var mainlefthclickeds *sdl.Surface
var mainlefthclicked = []byte("\x42\x4d\x8a\x04\x00\x00\x00\x00\x00\x00\x8a\x00\x00\x00\x7c\x00" +
	"\x00\x00\x10\x00\x00\x00\x10\x00\x00\x00\x01\x00\x20\x00\x03\x00" +
	"\x00\x00\x00\x04\x00\x00\x25\x16\x00\x00\x25\x16\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\x00\x00\xff\x00\x00\xff" +
	"\x00\x00\xff\x00\x00\x00\x42\x47\x52\x73\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x98\xfe\xff\xff\x9a\xff\xff\xff\x99\xfe\xff\xff\x99" +
	"\xfe\xff\xff\x99\xfe\xff\xff\x99\xfe\xff\xff\x99\xfe\xff\xff\x99" +
	"\xfe\xff\xff\x99\xfe\xff\xff\x99\xfe\xff\xff\x99\xfe\xff\xff\x99" +
	"\xfe\xff\xff\x9a\xff\xff\xff\x8f\xfa\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\xa3\xff\xff\xff\x00\xa7\xb3\xff\x00\xac\xb8\xff\x00" +
	"\xac\xb8\xff\x00\xac\xb8\xff\x00\xad\xb8\xff\x00\xad\xb8\xff\x00" +
	"\xad\xb9\xff\x00\xad\xb9\xff\x00\xad\xb9\xff\x00\xae\xb9\xff\x00" +
	"\xae\xb9\xff\x00\xaa\xb6\xff\x9a\xff\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9d\xff\xff\xff\x00\x00\x00\xff\x00\x00\x00\xff\x3f" +
	"\xda\xe0\xff\x37\xd7\xde\xff\x2e\xd4\xda\xff\x25\xd1\xd7\xff\x17" +
	"\xcd\xd4\xff\x05\xca\xd3\xff\x00\xc7\xd0\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xae\xb9\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9d\xff\xff\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x3c\xd8\xde\xff\x34\xd5\xdc\xff\x2c\xd2\xd9\xff\x24" +
	"\xcf\xd6\xff\x16\xcc\xd3\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xae\xb9\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9c\xff\xff\xff\x56\xe3\xe8\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x3c\xd8\xde\xff\x34\xd5\xdc\xff\x2c" +
	"\xd2\xd9\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\xc7\xd0\xff\x00\xad\xb9\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9c\xff\xff\xff\x5f\xe5\xeb\xff\x59\xe4\xe9\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x3c\xd8\xde\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x16\xcc\xd3\xff\x05" +
	"\xca\xd3\xff\x00\xad\xb9\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9b\xff\xff\xff\x66\xe8\xee\xff\x61\xe6\xec\xff\x59" +
	"\xe4\xe9\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x2c\xd2\xd9\xff\x24\xcf\xd6\xff\x17" +
	"\xcd\xd4\xff\x00\xad\xb9\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9b\xff\xff\xff\x6c\xeb\xf1\xff\x68\xe9\xef\xff\x61" +
	"\xe6\xec\xff\x59\xe4\xe9\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x3c\xd8\xde\xff\x34\xd5\xdc\xff\x2c\xd2\xd9\xff\x25" +
	"\xd1\xd7\xff\x00\xad\xb8\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9b\xff\xff\xff\x74\xee\xf3\xff\x6e\xec\xf2\xff\x68" +
	"\xe9\xef\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x3c\xd8\xde\xff\x34\xd5\xdc\xff\x2e" +
	"\xd4\xda\xff\x00\xad\xb8\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9a\xff\xff\xff\x7c\xf2\xf7\xff\x75\xef\xf4\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x59\xe4\xe9\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x3c\xd8\xde\xff\x37" +
	"\xd7\xde\xff\x00\xac\xb8\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x9a\xff\xff\xff\x83\xf5\xfa\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x68\xe9\xef\xff\x61\xe6\xec\xff\x59" +
	"\xe4\xe9\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x3f" +
	"\xda\xe0\xff\x00\xac\xb8\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x99\xff\xff\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x75\xef\xf4\xff\x6e\xec\xf2\xff\x68\xe9\xef\xff\x61" +
	"\xe6\xec\xff\x59\xe4\xe9\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xac\xb8\xff\x99\xfe\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xaf\xba\xff\x99\xfe\xff\xff\x00\x00\x00\xff\x00\x00\x00\xff\x83" +
	"\xf5\xfa\xff\x7c\xf2\xf7\xff\x74\xee\xf3\xff\x6c\xeb\xf1\xff\x66" +
	"\xe8\xee\xff\x5f\xe5\xeb\xff\x56\xe3\xe8\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xa7\xb3\xff\x9a\xff\xff\xff\x0b\xf1\xff\xff\x00" +
	"\xae\xba\xff\xa3\xff\xff\xff\x99\xfe\xff\xff\x99\xff\xff\xff\x9a" +
	"\xff\xff\xff\x9a\xff\xff\xff\x9b\xff\xff\xff\x9b\xff\xff\xff\x9b" +
	"\xff\xff\xff\x9c\xff\xff\xff\x9c\xff\xff\xff\x9d\xff\xff\xff\x9d" +
	"\xff\xff\xff\xa3\xff\xff\xff\x98\xfe\xff\xff\x0b\xf1\xff\xff\x01" +
	"\xb3\xbe\xff\x00\xae\xba\xff\x00\xaf\xba\xff\x00\xaf\xba\xff\x00" +
	"\xaf\xba\xff\x00\xaf\xba\xff\x00\xaf\xba\xff\x00\xaf\xba\xff\x00" +
	"\xaf\xba\xff\x00\xaf\xba\xff\x00\xaf\xba\xff\x00\xaf\xba\xff\x00" +
	"\xaf\xba\xff\x00\xaf\xba\xff\x00\xaf\xba\xff\x0b\xf1\xff\xff\x0b" +
	"\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b" +
	"\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b" +
	"\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b" +
	"\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b\xf1\xff\xff\x0b")

var mainrighths *sdl.Surface
var mainrighth = []byte("\x42\x4d\x8a\x04\x00\x00\x00\x00\x00\x00\x8a\x00\x00\x00\x7c\x00" +
	"\x00\x00\x10\x00\x00\x00\x10\x00\x00\x00\x01\x00\x20\x00\x03\x00" +
	"\x00\x00\x00\x04\x00\x00\x13\x0b\x00\x00\x13\x0b\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\x00\x00\xff\x00\x00\xff" +
	"\x00\x00\xff\x00\x00\x00\x42\x47\x52\x73\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x22" +
	"\xd7\xdf\xff\x11\xd4\xdd\xff\x00\xd1\xda\xff\x00\xce\xd7\xff\x00" +
	"\xcb\xd4\xff\x00\xc8\xd2\xff\x00\xc5\xcf\xff\x00\xc2\xcc\xff\x00" +
	"\xbf\xc9\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x30" +
	"\xda\xe2\xff\x22\xd7\xdf\xff\x11\xd4\xdd\xff\x00\xd1\xda\xff\x00" +
	"\xce\xd7\xff\x00\xcb\xd4\xff\x00\xc8\xd2\xff\x00\xc5\xcf\xff\x00" +
	"\xc2\xcc\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x39" +
	"\xdd\xe5\xff\x30\xda\xe2\xff\x22\xd7\xdf\xff\x11\xd4\xdd\xff\x00" +
	"\xd1\xda\xff\x00\xce\xd7\xff\x00\xcb\xd4\xff\x00\xc8\xd2\xff\x00" +
	"\xc5\xcf\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x42" +
	"\xe0\xe8\xff\x39\xdd\xe5\xff\x30\xda\xe2\xff\x22\xd7\xdf\xff\x11" +
	"\xd4\xdd\xff\x00\xd1\xda\xff\x00\xce\xd7\xff\x00\xcb\xd4\xff\x00" +
	"\xc8\xd2\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x4b" +
	"\xe3\xea\xff\x42\xe0\xe8\xff\x39\xdd\xe5\xff\x30\xda\xe2\xff\x22" +
	"\xd7\xdf\xff\x11\xd4\xdd\xff\x00\xd1\xda\xff\x00\xce\xd7\xff\x00" +
	"\xcb\xd4\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x30" +
	"\xda\xe2\xff\x22\xd7\xdf\xff\x11\xd4\xdd\xff\x00\xd1\xda\xff\x00" +
	"\xce\xd7\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x39" +
	"\xdd\xe5\xff\x30\xda\xe2\xff\x22\xd7\xdf\xff\x11\xd4\xdd\xff\x00" +
	"\xd1\xda\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x22\xd7\xdf\xff\x00\xce\xd7\xff\x00\xc5\xcf\xff\x00" +
	"\xbc\xc7\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x42" +
	"\xe0\xe8\xff\x39\xdd\xe5\xff\x30\xda\xe2\xff\x22\xd7\xdf\xff\x11" +
	"\xd4\xdd\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x42\xe0\xe8\xff\x22\xd7\xdf\xff\x00\xce\xd7\xff\x00" +
	"\xc5\xcf\xff\x00\xbc\xc7\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x4b" +
	"\xe3\xea\xff\x42\xe0\xe8\xff\x39\xdd\xe5\xff\x30\xda\xe2\xff\x22" +
	"\xd7\xdf\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x5c\xe9\xf0\xff\x42\xe0\xe8\xff\x22\xd7\xdf\xff\x00" +
	"\xce\xd7\xff\x00\xc5\xcf\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x73\xf2\xf8\xff\x5c\xe9\xf0\xff\x42\xe0\xe8\xff\x22" +
	"\xd7\xdf\xff\x00\xce\xd7\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x73\xf2\xf8\xff\x5c\xe9\xf0\xff\x42" +
	"\xe0\xe8\xff\x22\xd7\xdf\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00")

var mainrighthclickeds *sdl.Surface
var mainrighthclicked = []byte("\x42\x4d\x8a\x04\x00\x00\x00\x00\x00\x00\x8a\x00\x00\x00\x7c\x00" +
	"\x00\x00\x10\x00\x00\x00\x10\x00\x00\x00\x01\x00\x20\x00\x03\x00" +
	"\x00\x00\x00\x04\x00\x00\x13\x0b\x00\x00\x13\x0b\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\x00\x00\xff\x00\x00\xff" +
	"\x00\x00\xff\x00\x00\x00\x42\x47\x52\x73\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00\x00\x00\xff\x00" +
	"\x00\x00\xff\x00\x00\x00\xff\x00\xb3\xbe\xff\x88\xfa\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88" +
	"\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x88\xfa\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00\xb3\xbe\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00" +
	"\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00\xb3\xbe\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00" +
	"\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00\xf2\xff\xff\x00")

var mainlefts *sdl.Surface
var mainleft = []byte("\x42\x4d\x8a\x04\x00\x00\x00\x00\x00\x00\x8a\x00\x00\x00\x7c\x00" +
	"\x00\x00\x10\x00\x00\x00\x10\x00\x00\x00\x01\x00\x20\x00\x03\x00" +
	"\x00\x00\x00\x04\x00\x00\x25\x16\x00\x00\x25\x16\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\x00\x00\xff\x00\x00\xff" +
	"\x00\x00\xff\x00\x00\x00\x42\x47\x52\x73\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf7\xf7\xf7\xff\xf8\xf8\xf8\xff\xf7\xf7\xf7\xff\xf7" +
	"\xf7\xf7\xff\xf7\xf7\xf7\xff\xf7\xf7\xf7\xff\xf7\xf7\xf7\xff\xf7" +
	"\xf7\xf7\xff\xf7\xf7\xf7\xff\xf7\xf7\xf7\xff\xf7\xf7\xf7\xff\xf7" +
	"\xf7\xf7\xff\xf8\xf8\xf8\xff\xf3\xf3\xf3\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf8\xf8\xf8\xff\x9d\x9d\x9d\xff\xa2\xa2\xa2\xff\xa2" +
	"\xa2\xa2\xff\xa2\xa2\xa2\xff\xa3\xa3\xa3\xff\xa3\xa3\xa3\xff\xa3" +
	"\xa3\xa3\xff\xa3\xa3\xa3\xff\xa3\xa3\xa3\xff\xa4\xa4\xa4\xff\xa4" +
	"\xa4\xa4\xff\xa0\xa0\xa0\xff\xf8\xf8\xf8\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf8\xf8\xf8\xff\xd6\xd6\xd6\xff\xd4\xd4\xd4\xff\xd0" +
	"\xd0\xd0\xff\xcd\xcd\xcd\xff\xc9\xc9\xc9\xff\xc6\xc6\xc6\xff\xc1" +
	"\xc1\xc1\xff\xbe\xbe\xbe\xff\xbb\xbb\xbb\xff\xb8\xb8\xb8\xff\xb5" +
	"\xb5\xb5\xff\xa4\xa4\xa4\xff\xf7\xf7\xf7\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf8\xf8\xf8\xff\xd7\xd7\xd7\xff\xd5\xd5\xd5\xff\xd1" +
	"\xd1\xd1\xff\xce\xce\xce\xff\xcb\xcb\xcb\xff\xc8\xc8\xc8\xff\xc4" +
	"\xc4\xc4\xff\xc0\xc0\xc0\xff\xbd\xbd\xbd\xff\xba\xba\xba\xff\xb8" +
	"\xb8\xb8\xff\xa4\xa4\xa4\xff\xf7\xf7\xf7\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf8\xf8\xf8\xff\xda\xda\xda\xff\xd8\xd8\xd8\xff\xd5" +
	"\xd5\xd5\xff\xd1\xd1\xd1\xff\xce\xce\xce\xff\xcb\xcb\xcb\xff\xc8" +
	"\xc8\xc8\xff\xc4\xc4\xc4\xff\xc0\xc0\xc0\xff\xbd\xbd\xbd\xff\xbb" +
	"\xbb\xbb\xff\xa3\xa3\xa3\xff\xf7\xf7\xf7\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf8\xf8\xf8\xff\xdd\xdd\xdd\xff\xdb\xdb\xdb\xff\xd8" +
	"\xd8\xd8\xff\xd5\xd5\xd5\xff\xd1\xd1\xd1\xff\xce\xce\xce\xff\xcb" +
	"\xcb\xcb\xff\xc8\xc8\xc8\xff\xc4\xc4\xc4\xff\xc0\xc0\xc0\xff\xbe" +
	"\xbe\xbe\xff\xa3\xa3\xa3\xff\xf7\xf7\xf7\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf8\xf8\xf8\xff\xe0\xe0\xe0\xff\xde\xde\xde\xff\xdb" +
	"\xdb\xdb\xff\xd8\xd8\xd8\xff\xd5\xd5\xd5\xff\xd1\xd1\xd1\xff\xce" +
	"\xce\xce\xff\xcb\xcb\xcb\xff\xc8\xc8\xc8\xff\xc4\xc4\xc4\xff\xc1" +
	"\xc1\xc1\xff\xa3\xa3\xa3\xff\xf7\xf7\xf7\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf8\xf8\xf8\xff\xe3\xe3\xe3\xff\xe1\xe1\xe1\xff\xde" +
	"\xde\xde\xff\xdb\xdb\xdb\xff\xd8\xd8\xd8\xff\xd5\xd5\xd5\xff\xd1" +
	"\xd1\xd1\xff\xce\xce\xce\xff\xcb\xcb\xcb\xff\xc8\xc8\xc8\xff\xc6" +
	"\xc6\xc6\xff\xa3\xa3\xa3\xff\xf7\xf7\xf7\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf8\xf8\xf8\xff\xe6\xe6\xe6\xff\xe4\xe4\xe4\xff\xe1" +
	"\xe1\xe1\xff\xde\xde\xde\xff\xdb\xdb\xdb\xff\xd8\xd8\xd8\xff\xd5" +
	"\xd5\xd5\xff\xd1\xd1\xd1\xff\xce\xce\xce\xff\xcb\xcb\xcb\xff\xc9" +
	"\xc9\xc9\xff\xa3\xa3\xa3\xff\xf7\xf7\xf7\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf8\xf8\xf8\xff\xeb\xeb\xeb\xff\xe7\xe7\xe7\xff\xe4" +
	"\xe4\xe4\xff\xe1\xe1\xe1\xff\xde\xde\xde\xff\xdb\xdb\xdb\xff\xd8" +
	"\xd8\xd8\xff\xd5\xd5\xd5\xff\xd1\xd1\xd1\xff\xce\xce\xce\xff\xcd" +
	"\xcd\xcd\xff\xa2\xa2\xa2\xff\xf7\xf7\xf7\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf8\xf8\xf8\xff\xee\xee\xee\xff\xeb\xeb\xeb\xff\xe7" +
	"\xe7\xe7\xff\xe4\xe4\xe4\xff\xe1\xe1\xe1\xff\xde\xde\xde\xff\xdb" +
	"\xdb\xdb\xff\xd8\xd8\xd8\xff\xd5\xd5\xd5\xff\xd1\xd1\xd1\xff\xd0" +
	"\xd0\xd0\xff\xa2\xa2\xa2\xff\xf7\xf7\xf7\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf8\xf8\xf8\xff\xf1\xf1\xf1\xff\xee\xee\xee\xff\xeb" +
	"\xeb\xeb\xff\xe7\xe7\xe7\xff\xe4\xe4\xe4\xff\xe1\xe1\xe1\xff\xde" +
	"\xde\xde\xff\xdb\xdb\xdb\xff\xd8\xd8\xd8\xff\xd5\xd5\xd5\xff\xd4" +
	"\xd4\xd4\xff\xa2\xa2\xa2\xff\xf7\xf7\xf7\xff\xdd\xdd\xdd\xff\xa5" +
	"\xa5\xa5\xff\xf7\xf7\xf7\xff\xf3\xf3\xf3\xff\xf1\xf1\xf1\xff\xee" +
	"\xee\xee\xff\xeb\xeb\xeb\xff\xe6\xe6\xe6\xff\xe3\xe3\xe3\xff\xe0" +
	"\xe0\xe0\xff\xdd\xdd\xdd\xff\xda\xda\xda\xff\xd7\xd7\xd7\xff\xd6" +
	"\xd6\xd6\xff\x9d\x9d\x9d\xff\xf8\xf8\xf8\xff\xdd\xdd\xdd\xff\xa4" +
	"\xa4\xa4\xff\xf8\xf8\xf8\xff\xf7\xf7\xf7\xff\xf8\xf8\xf8\xff\xf8" +
	"\xf8\xf8\xff\xf8\xf8\xf8\xff\xf8\xf8\xf8\xff\xf8\xf8\xf8\xff\xf8" +
	"\xf8\xf8\xff\xf8\xf8\xf8\xff\xf8\xf8\xf8\xff\xf8\xf8\xf8\xff\xf8" +
	"\xf8\xf8\xff\xf8\xf8\xf8\xff\xf7\xf7\xf7\xff\xdd\xdd\xdd\xff\xa8" +
	"\xa8\xa8\xff\xa4\xa4\xa4\xff\xa5\xa5\xa5\xff\xa5\xa5\xa5\xff\xa5" +
	"\xa5\xa5\xff\xa5\xa5\xa5\xff\xa5\xa5\xa5\xff\xa5\xa5\xa5\xff\xa5" +
	"\xa5\xa5\xff\xa5\xa5\xa5\xff\xa5\xa5\xa5\xff\xa5\xa5\xa5\xff\xa5" +
	"\xa5\xa5\xff\xa5\xa5\xa5\xff\xa5\xa5\xa5\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd")

var mainrights *sdl.Surface
var mainright = []byte("\x42\x4d\x8a\x04\x00\x00\x00\x00\x00\x00\x8a\x00\x00\x00\x7c\x00" +
	"\x00\x00\x10\x00\x00\x00\x10\x00\x00\x00\x01\x00\x20\x00\x03\x00" +
	"\x00\x00\x00\x04\x00\x00\x13\x0b\x00\x00\x13\x0b\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\x00\x00\xff\x00\x00\xff" +
	"\x00\x00\xff\x00\x00\x00\x42\x47\x52\x73\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xf3" +
	"\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3" +
	"\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3" +
	"\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xa8" +
	"\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8" +
	"\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8" +
	"\xa8\xa8\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xcc" +
	"\xcc\xcc\xff\xc8\xc8\xc8\xff\xc4\xc4\xc4\xff\xc1\xc1\xc1\xff\xbe" +
	"\xbe\xbe\xff\xbc\xbc\xbc\xff\xb9\xb9\xb9\xff\xb6\xb6\xb6\xff\xb3" +
	"\xb3\xb3\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xcf" +
	"\xcf\xcf\xff\xcc\xcc\xcc\xff\xc8\xc8\xc8\xff\xc4\xc4\xc4\xff\xc1" +
	"\xc1\xc1\xff\xbe\xbe\xbe\xff\xbc\xbc\xbc\xff\xb9\xb9\xb9\xff\xb6" +
	"\xb6\xb6\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xd3" +
	"\xd3\xd3\xff\xcf\xcf\xcf\xff\xcc\xcc\xcc\xff\xc8\xc8\xc8\xff\xc4" +
	"\xc4\xc4\xff\xc1\xc1\xc1\xff\xbe\xbe\xbe\xff\xbc\xbc\xbc\xff\xb9" +
	"\xb9\xb9\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xd6" +
	"\xd6\xd6\xff\xd3\xd3\xd3\xff\xcf\xcf\xcf\xff\xcc\xcc\xcc\xff\xc8" +
	"\xc8\xc8\xff\xc4\xc4\xc4\xff\xc1\xc1\xc1\xff\xbe\xbe\xbe\xff\xbc" +
	"\xbc\xbc\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xda" +
	"\xda\xda\xff\xd6\xd6\xd6\xff\xd3\xd3\xd3\xff\xcf\xcf\xcf\xff\xcc" +
	"\xcc\xcc\xff\xc8\xc8\xc8\xff\xc4\xc4\xc4\xff\xc1\xc1\xc1\xff\xbe" +
	"\xbe\xbe\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xa8\xa8\xa8\xff\xf3" +
	"\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3" +
	"\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xcf" +
	"\xcf\xcf\xff\xcc\xcc\xcc\xff\xc8\xc8\xc8\xff\xc4\xc4\xc4\xff\xc1" +
	"\xc1\xc1\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xa8\xa8\xa8\xff\xf3" +
	"\xf3\xf3\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8" +
	"\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xd3" +
	"\xd3\xd3\xff\xcf\xcf\xcf\xff\xcc\xcc\xcc\xff\xc8\xc8\xc8\xff\xc4" +
	"\xc4\xc4\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xa8\xa8\xa8\xff\xf3" +
	"\xf3\xf3\xff\xcc\xcc\xcc\xff\xc1\xc1\xc1\xff\xb9\xb9\xb9\xff\xb1" +
	"\xb1\xb1\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xd6" +
	"\xd6\xd6\xff\xd3\xd3\xd3\xff\xcf\xcf\xcf\xff\xcc\xcc\xcc\xff\xc8" +
	"\xc8\xc8\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xa8\xa8\xa8\xff\xf3" +
	"\xf3\xf3\xff\xd6\xd6\xd6\xff\xcc\xcc\xcc\xff\xc1\xc1\xc1\xff\xb9" +
	"\xb9\xb9\xff\xb1\xb1\xb1\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xda" +
	"\xda\xda\xff\xd6\xd6\xd6\xff\xd3\xd3\xd3\xff\xcf\xcf\xcf\xff\xcc" +
	"\xcc\xcc\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xa8\xa8\xa8\xff\xf3" +
	"\xf3\xf3\xff\xe0\xe0\xe0\xff\xd6\xd6\xd6\xff\xcc\xcc\xcc\xff\xc1" +
	"\xc1\xc1\xff\xb9\xb9\xb9\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xf3" +
	"\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3" +
	"\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xa8\xa8\xa8\xff\xf3" +
	"\xf3\xf3\xff\xea\xea\xea\xff\xe0\xe0\xe0\xff\xd6\xd6\xd6\xff\xcc" +
	"\xcc\xcc\xff\xc1\xc1\xc1\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xa8" +
	"\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8" +
	"\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xf3" +
	"\xf3\xf3\xff\xf3\xf3\xf3\xff\xea\xea\xea\xff\xe0\xe0\xe0\xff\xd6" +
	"\xd6\xd6\xff\xcc\xcc\xcc\xff\xa8\xa8\xa8\xff\xf3\xf3\xf3\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xa8\xa8\xa8\xff\xf3" +
	"\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3" +
	"\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xf3\xf3\xf3\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xa8\xa8\xa8\xff\xa8" +
	"\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8" +
	"\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xa8\xa8\xa8\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd" +
	"\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd\xdd\xdd\xff\xdd")

func InitSprites() error {
	rwops, err := sdl.RWFromMem(mainlefth)
	if err != nil {
		return fmt.Errorf("Unable to read mainlefth: %v", err)
	}
	mainlefths, _ = sdl.LoadBMPRW(rwops, true)
	rwops, err = sdl.RWFromMem(mainlefthclicked)
	if err != nil {
		return fmt.Errorf("Unable to read mainlefthclicked: %v", err)
	}
	mainlefthclickeds, _ = sdl.LoadBMPRW(rwops, true)
	rwops, err = sdl.RWFromMem(mainrighth)
	if err != nil {
		return fmt.Errorf("Unable to read mainrighth: %v", err)
	}
	mainrighths, _ = sdl.LoadBMPRW(rwops, true)
	rwops, err = sdl.RWFromMem(mainrighthclicked)
	if err != nil {
		return fmt.Errorf("Unable to read mainrighthclicked: %v", err)
	}
	mainrighthclickeds, _ = sdl.LoadBMPRW(rwops, true)
	rwops, err = sdl.RWFromMem(mainleft)
	if err != nil {
		return fmt.Errorf("Unable to read mainleft: %v", err)
	}
	mainlefts, _ = sdl.LoadBMPRW(rwops, true)
	rwops, err = sdl.RWFromMem(mainright)
	if err != nil {
		return fmt.Errorf("Unable to read mainright: %v", err)
	}
	mainrights, _ = sdl.LoadBMPRW(rwops, true)
	return nil
}
