package maze

import (
	"../vector"
	"fmt"
	"math/rand"
	"time"
)

const (
	Road int = iota
	Wall
)

type Maze struct {
	Map    [][]int
	Width  int
	Height int
}

func New(w, h int) *Maze {
	rand.Seed(time.Now().UnixNano())
	m := new(Maze)
	m.setSize(w, h)
	m.allocMap()
	return m
}

func (m *Maze) allocMap() {
	m.Map = make([][]int, m.Width)
	for x := 0; x < m.Width; x += 1 {
		m.Map[x] = make([]int, m.Height)
	}
}

func (m *Maze) setSize(w, h int) {
	if w < 5 {
		m.Width = 5
	} else {
		m.Width = w
	}
	if h < 5 {
		m.Height = 5
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
	x := randOdd(m.Width - 2)
	y := randOdd(m.Height - 2)
	fmt.Printf("(%d, %d)\n", x, y)
	m.makeMaze(x, y)
}

func randOdd(mod int) (r int) {
	r = 1 + rand.Intn(mod)
	if r%2 == 0 {
		r += 1
	}
	if r > mod {
		r -= 2
	}
	return
}

func (m *Maze) makeMaze(x, y int) {
	v := vector.New()
	for {
		vx, vy := v.GetVector()
		px := x + vx*2
		py := y + vy*2
		if px < 0 || px >= m.Width || py < 0 || py >= m.Height || m.Map[px][py] != Wall {
			if v.Rotate() {
				return
			}
			continue
		}
		m.Map[x+vx][y+vy] = Road
		m.Map[px][py] = Road
		m.makeMaze(px, py)
		v.Reset()
	}
}
