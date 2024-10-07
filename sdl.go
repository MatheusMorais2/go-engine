package main

import (
	"fmt"

	sdl "github.com/veandco/go-sdl2/sdl"
)

func arroz() {
	fmt.Println("Hello, world!")

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic("Failed to initialize a window")
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		"Main", 
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		1080,
		720,
		sdl.WINDOW_SHOWN | sdl.WINDOW_RESIZABLE,
	)
	defer window.Destroy()



	if err != nil {
		panic("Failed to create the SDL window")
	}

	runWindow := true
	for runWindow {
		renderer, err := window.GetRenderer()
		
		if err != nil {
			panic("Failed to get renderer")
		}

		err = renderer.DrawLine(-1,-1,-100,-100)
		if err != nil {
			fmt.Printf("error: %+v\n", sdl.GetError())
			panic(err)
		}
	}

	
}