package robot

import (
	"fmt"
)

type Robot struct {
	X int
	Y int
	Direction Direction
}

type Direction int 
const (
	N Direction = iota
	E Direction = iota
	S Direction = iota
	W Direction = iota
)

func (d Direction) ToString() string {
	// [...] means have the compiler count the array elements for you
	return [...]string{"N", "E", "S", "W"}[d]
}

func (r Robot) GetPosition() string {
	return fmt.Sprintf("%d, %d, %s", r.X, r.Y, r.Direction.ToString())
}

func (r *Robot) Turn(direction int) {
	newDirection := int(r.Direction) + direction
	switch {
	case newDirection > 3:
		newDirection = 0
	case newDirection < 0:
		newDirection = 3
	}
	r.Direction = [...]Direction{N, E, S, W}[newDirection]
}

func (r *Robot) Move(unit int) {
	switch r.Direction {
	case N:
		r.Y += unit
	case E:
		r.X += unit
	case S:
		r.Y -= unit
	case W:
		r.X -= unit
	}
}