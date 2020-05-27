package model

import (
	"errors"
	"strings"
)

type Maze struct {
	x, y  int
	filed [][]byte
}

var ErrInvalidInput = errors.New("invalid input")

func NewMaze(x, y int) (*Maze, error) {
	if x%2 == 0 || y%2 == 0 {
		return nil, ErrInvalidInput
	}
	m := &Maze{
		x: x, y: y,
		filed: make([][]byte, x+2),
	}
	m.filed[0] = wallLine(y + 2)
	m.filed[x+1] = wallLine(y + 2)
	for i := 1; i < x+1; i++ {
		m.filed[i] = make([]byte, y+2)
		m.filed[i][0] = wall
		for j := 1; j < y+1; j++ {
			m.filed[i][j] = space
		}
		m.filed[i][y+1] = wall
	}
	return m, nil
}

func (m *Maze) String() string {
	b := &strings.Builder{}
	for _, l := range m.filed {
		b.Write(l)
		b.WriteByte('\n')
	}
	return b.String()
}

func wallLine(l int) []byte {
	line := make([]byte, l)
	for i := 0; i < len(line); i++ {
		line[i] = wall
	}
	return line
}

const (
	space = ' '
	wall  = '#'
)
