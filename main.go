package main

import (
	"fmt"
	"neck/internals/frets"
	"neck/internals/music"
	"neck/internals/neck"
	"neck/internals/render"
)

func main() {

	neck := neck.GuitarNeck()

	stringToPrint := render.RenderFretString(frets.NewFretString(music.E), 1, 5, "3", music.G)
	fmt.Println("\n\n")
	fmt.Println(stringToPrint)

	fmt.Println(render.RenderChord(neck, "7-9-9-8-7-7", music.B))
	fmt.Println(render.RenderChord(neck, "x-x-x-x-x-x", music.G))
	fmt.Println(render.RenderChord(neck, "0-2-2-0-0-x", music.E))
}
