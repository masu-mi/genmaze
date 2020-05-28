package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("maze", func() {
	Title("GenMaze Service")
	Description("Service for adding numbers, a Goa teaser")
	Server("maze", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var Position = Type("position", func() {
	Description("0-indexed position")
	Field(1, "x", Int)
	Field(2, "y", Int)
})

var GeneratedMaze = ResultType("generated_maze", func() {
	Attributes(func() {
		Field(1, "field", String)
		Field(2, "start", Position)
		Field(3, "goal", Position)
	})
})

var _ = Service("maze", func() {
	Description("The genmaze service performs operations on numbers.")

	Method("gen", func() {
		Payload(func() {
			Field(1, "x", Int, "field size x", func() {
				Minimum(1)
				Maximum(1001)
			})
			Field(2, "y", Int, "field size y", func() {
				Minimum(1)
				Maximum(1001)
			})
			Required("x", "y")
		})

		Result(GeneratedMaze)

		HTTP(func() {
			POST("/gen/")
			Response(StatusOK)
		})
		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
	Files("/openapi.yaml", "./gen/http/openapi.yaml")
	Files("/swagger-ui/{*filepath}", "./public/swagger-ui/")
})
