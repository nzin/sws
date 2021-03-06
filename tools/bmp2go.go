package main

import (
	"flag"
	"fmt"
	"os"
)

//
// After translation it can be used with:
//rwops := sdl.RWFromMem(bmp)
//surface := sdl.LoadBMP_RW(rwops,1)
//

func main() {
	bmpfile := flag.String("bmpfile", "", "bmp image to read")
	flag.Parse()

	if *bmpfile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	f, err := os.Open(*bmpfile)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 16)
	var n int
	fmt.Print("var bmp = \n []byte(\"")
	for n, err = f.Read(buffer); n > 0; n, err = f.Read(buffer) {
		for i := range buffer {
			fmt.Printf("\\x%02x", uint8(buffer[i]))
		}
	}
	fmt.Println("\")")
}
