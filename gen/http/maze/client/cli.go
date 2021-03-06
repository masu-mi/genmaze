// Code generated by goa v3.1.3, DO NOT EDIT.
//
// maze HTTP client CLI support package
//
// Command:
// $ goa gen genmaze/design

package client

import (
	"encoding/json"
	"fmt"
	maze "genmaze/gen/maze"

	goa "goa.design/goa/v3/pkg"
)

// BuildGenPayload builds the payload for the maze gen endpoint from CLI flags.
func BuildGenPayload(mazeGenBody string) (*maze.GenPayload, error) {
	var err error
	var body GenRequestBody
	{
		err = json.Unmarshal([]byte(mazeGenBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"h\": 853,\n      \"w\": 273\n   }'")
		}
		if body.W < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.w", body.W, 1, true))
		}
		if body.W > 1001 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.w", body.W, 1001, false))
		}
		if body.H < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.h", body.H, 1, true))
		}
		if body.H > 1001 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.h", body.H, 1001, false))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &maze.GenPayload{
		W: body.W,
		H: body.H,
	}

	return v, nil
}
