// Code generated by goa v3.1.3, DO NOT EDIT.
//
// maze gRPC client encoders and decoders
//
// Command:
// $ goa gen genmaze/design

package client

import (
	"context"
	mazepb "genmaze/gen/grpc/maze/pb"
	maze "genmaze/gen/maze"
	mazeviews "genmaze/gen/maze/views"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildGenFunc builds the remote method to invoke for "maze" service "gen"
// endpoint.
func BuildGenFunc(grpccli mazepb.MazeClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Gen(ctx, reqpb.(*mazepb.GenRequest), opts...)
		}
		return grpccli.Gen(ctx, &mazepb.GenRequest{}, opts...)
	}
}

// EncodeGenRequest encodes requests sent to maze gen endpoint.
func EncodeGenRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*maze.GenPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("maze", "gen", "*maze.GenPayload", v)
	}
	return NewGenRequest(payload), nil
}

// DecodeGenResponse decodes responses from the maze gen endpoint.
func DecodeGenResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*mazepb.GenResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("maze", "gen", "*mazepb.GenResponse", v)
	}
	res := NewGenResult(message)
	vres := &mazeviews.GeneratedMaze{Projected: res, View: view}
	if err := mazeviews.ValidateGeneratedMaze(vres); err != nil {
		return nil, err
	}
	return maze.NewGeneratedMaze(vres), nil
}