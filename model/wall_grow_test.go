package model

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestCreateMazeWithWallGrow(t *testing.T) {
	m, _ := CreateMazeWithWallGrow(39, 29)
	fmt.Println(m.String())
	m, _ = CreateMazeWithWallGrow(19, 29)
	fmt.Println(m.String())
}
