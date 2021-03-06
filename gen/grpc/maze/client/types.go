// Code generated by goa v3.1.3, DO NOT EDIT.
//
// maze gRPC client types
//
// Command:
// $ goa gen genmaze/design

package client

import (
	mazepb "genmaze/gen/grpc/maze/pb"
	maze "genmaze/gen/maze"
	mazeviews "genmaze/gen/maze/views"
)

// NewGenRequest builds the gRPC request type from the payload of the "gen"
// endpoint of the "maze" service.
func NewGenRequest(payload *maze.GenPayload) *mazepb.GenRequest {
	message := &mazepb.GenRequest{
		W: int32(payload.W),
		H: int32(payload.H),
	}
	return message
}

// NewGenResult builds the result type of the "gen" endpoint of the "maze"
// service from the gRPC response type.
func NewGenResult(message *mazepb.GenResponse) *mazeviews.GeneratedMazeView {
	result := &mazeviews.GeneratedMazeView{}
	if message.Field != "" {
		result.Field = &message.Field
	}
	return result
}
