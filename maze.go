package mazeapi

import (
	"context"
	maze "genmaze/gen/maze"
	"genmaze/model"
	"log"
)

// maze service example implementation.
// The example methods log the requests and return zero values.
type mazesrvc struct {
	logger *log.Logger
}

// NewMaze returns the maze service implementation.
func NewMaze(logger *log.Logger) maze.Service {
	return &mazesrvc{logger}
}

// Gen implements gen.
func (s *mazesrvc) Gen(ctx context.Context, p *maze.GenPayload) (res *maze.GeneratedMaze, err error) {
	w, h := p.W, p.H
	m, _ := model.CreateMazeWithWallGrow(w, h)
	str := m.String()
	res = &maze.GeneratedMaze{}
	res.Field = &str
	s.logger.Print("maze.gen")
	return
}
