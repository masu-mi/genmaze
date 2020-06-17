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

var GeneratedMaze = ResultType("generated_maze", func() {
	Attributes(func() {
		Field(1, "field", String)
	})
})

var _ = Service("maze", func() {
	Description("The genmaze service performs operations on numbers.")

	Method("gen", func() {
		Payload(func() {
			Field(1, "w", Int, "field size x", func() {
				Minimum(1)
				Maximum(1001)
			})
			Field(2, "h", Int, "field size y", func() {
				Minimum(1)
				Maximum(1001)
			})
			Required("w", "h")
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
