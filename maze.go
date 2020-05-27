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
func (s *mazesrvc) Gen(ctx context.Context, p *maze.GenPayload) (res *maze.GenResult, err error) {
	x, y := p.X, p.Y
	m, _ := model.CreateMazeWithWallGrow(x, y)
	res = &maze.GenResult{}
	str := m.String()
	res.Field = &str
	s.logger.Print("maze.gen")
	return
}
