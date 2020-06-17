// Code generated by goa v3.1.3, DO NOT EDIT.
//
// maze HTTP server encoders and decoders
//
// Command:
// $ goa gen genmaze/design

package server

import (
	"context"
	mazeviews "genmaze/gen/maze/views"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGenResponse returns an encoder for responses returned by the maze gen
// endpoint.
func EncodeGenResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*mazeviews.GeneratedMaze)
		enc := encoder(ctx, w)
		body := NewGenResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGenRequest returns a decoder for requests sent to the maze gen
// endpoint.
func DecodeGenRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body GenRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateGenRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewGenPayload(&body)

		return payload, nil
	}
}
