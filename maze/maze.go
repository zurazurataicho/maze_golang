package maze

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Road int = iota
	Wall
)

type Vector struct {
	x int
	y int
}

var dir []Vector = []Vector{
	{0, -1}, /** UP */
	{0, 1},  /** DOWN */
	{-1, 0}, /** LEFT */
	{1, 0},  /** RIGHT */
}

type Maze struct {
	Map    [][]int
	Width  int
	Height int
}

func New(w, h int) *Maze {
	rand.Seed(time.Now().UnixNano())
	m := new(Maze)
	m.setSize(w, h)
	m.allocMap(w, h)
	return m
}

func (m *Maze) allocMap(w, h int) {
	m.Map = make([][]int, w)
	for x := 0; x < w; x += 1 {
		m.Map[x] = make([]int, h)
	}
}

func (m *Maze) setSize(w, h int) {
	if w <= 0 {
		m.Width = 1
	} else {
		m.Width = w
	}
	if h <= 0 {
		m.Height = 1
	} else {
		m.Height = h
	}
}

func (m *Maze) Clear() {
	for y := 0; y < m.Height; y += 1 {
		for x := 0; x < m.Width; x += 1 {
			m.Map[x][y] = Wall
		}
	}
}

func (m *Maze) Print() {
	var c string
	for y := 0; y < m.Height; y += 1 {
		for x := 0; x < m.Width; x += 1 {
			if m.Map[x][y] == Wall {
				c = "#"
			} else {
				c = " "
			}
			fmt.Printf("%s", c)
		}
		fmt.Println("")
	}
}

func (m *Maze) Make() {
	x := randOdd(m.Width)
	y := randOdd(m.Height)
	fmt.Printf("(%d, %d)\n", x, y)
	m.makeMaze(x, y)
}

func randOdd(mod int) (r int) {
	r = 2 + rand.Intn(mod)
	if r%2 == 0 {
		r += 1
	}
	if r > mod {
		r -= 2
	}
	return
}

func (m *Maze) makeMaze(x, y int) {
	d := rand.Intn(4)
	dd := d
	for {
		px := x + dir[d].x*2
		py := y + dir[d].y*2
		if px < 0 || px >= m.Width || py < 0 || py >= m.Height || m.Map[px][py] != Wall {
			d += 1
			if d == 4 {
				d = 0
			}
			if d == dd {
				return
			}
			continue
		}
		m.Map[x+dir[d].x][y+dir[d].y] = Road
		m.Map[px][py] = Road
		m.makeMaze(px, py)
		d = rand.Intn(4)
		dd = d
	}
}
