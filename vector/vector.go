package vector

import (
	"math/rand"
	"time"
)

type Vector struct {
	current int
	origin  int
	vec     [4][2]int
}

func New() *Vector {
	rand.Seed(time.Now().UnixNano())
	v := new(Vector)
	v.Reset()
	v._init_vec_array()
	return v
}

func (v *Vector) _init_vec_array() {
	v.vec = [4][2]int{
		{0, -1}, /** UP */
		{0, 1},  /** DOWN */
		{-1, 0}, /** LEFT */
		{1, 0},  /** RIGHT */
	}
}

func (v *Vector) Reset() {
	v.current = rand.Intn(4)
	v.origin = v.current
}

func (v *Vector) GetVector() (int, int) {
	return v.vec[v.current][0], v.vec[v.current][1]
}

func (v *Vector) Rotate() bool {
	v.current += 1
	if v.current == 4 {
		v.current = 0
	}
	if v.current == v.origin {
		return true
	}
	return false
}
