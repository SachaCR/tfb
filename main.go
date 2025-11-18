package main

import (
	"fmt"
	"neck/internals/frets"
	"neck/internals/music"
	"neck/internals/neck"
	"neck/internals/render"
)

func main() {

	// fmt.Println(myString.FretToNote(0))
	// fmt.Println(myString.FretToNote(5))
	// fmt.Println(myString.FretToNote(12))
	// fmt.Println(myString.FretToNote(13))
	// fmt.Println(myString.FretToNote(14))
	// fmt.Println(myString.FretToNote(1))

	neck := neck.GuitarNeck()

	fmt.Println(neck.Tuning())

	fmt.Println("┬─⬤ ─┬────┬────┬")
	fmt.Println("┼─⬤ ─┼────┼────┼")
	fmt.Println("┼────┼─⬤ ─┼────┼")
	fmt.Println("┼────┼────┼─⬤ ─┼")
	fmt.Println("┼─⬤ ─┼─Ab─┼─⬤ ─┼")
	fmt.Println("┴─⬤ ─┴────┴────┴")

	stringToPrint := render.RenderFretString(frets.NewFretString(music.E), 2, 5, []int{4})
	fmt.Println(stringToPrint)
}

// first: '|----',
// normal: '|----',
// last: '|----',

// first: '┬────',
// normal: '┼────',
// last: '┴────',

// render 7-9-9-8-7-7

//┼────┼─⬤ ─┼─⬤ ◯───◉───⊕───🔴───⚪───⚫───⭕───🔵
