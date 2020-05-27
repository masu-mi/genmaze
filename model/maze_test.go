package model

import (
	"testing"
)

func TestMazeString(t *testing.T) {
	m, _ := NewMaze(1, 1)
	got := m.String()
	want := `###
# #
###
`
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}
