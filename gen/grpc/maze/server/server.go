// Code generated by goa v3.1.3, DO NOT EDIT.
//
// maze gRPC server
//
// Command:
// $ goa gen genmaze/design

package server

import (
	"context"
	mazepb "genmaze/gen/grpc/maze/pb"
	maze "genmaze/gen/maze"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the mazepb.MazeServer interface.
type Server struct {
	GenH goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the maze service endpoints.
func New(e *maze.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		GenH: NewGenHandler(e.Gen, uh),
	}
}

// NewGenHandler creates a gRPC handler which serves the "maze" service "gen"
// endpoint.
func NewGenHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGenRequest, EncodeGenResponse)
	}
	return h
}

// Gen implements the "Gen" method in mazepb.MazeServer interface.
func (s *Server) Gen(ctx context.Context, message *mazepb.GenRequest) (*mazepb.GenResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "gen")
	ctx = context.WithValue(ctx, goa.ServiceKey, "maze")
	resp, err := s.GenH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*mazepb.GenResponse), nil
}
