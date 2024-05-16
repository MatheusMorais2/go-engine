package main

import (
	"fmt"
	"time"

	sdl "github.com/veandco/go-sdl2/sdl"
)

func main() {
	fmt.Println("Hello, world!")

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic("Failed to initialize a window")
	}

	window, err := sdl.CreateWindow("Main", 
	sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
	1080, 720, sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE)

	if err != nil {
		panic("Failed to create the SDL window")
	}

	time.Sleep(10 * time.Second)

	window.Destroy()

}