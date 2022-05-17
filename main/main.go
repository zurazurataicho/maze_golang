package main

import "zura.org/maze"

const (
	width  int = 59
	height int = 21
)

func main() {
	m := maze.New(width, height)
	for i := 0; i < 3; i += 1 {
		m.Clear()
		m.Make()
		m.Print()
	}
}
