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
		})
	})
})

var Position = Type("position", func() {
	Description("0-indexed position")
	Attribute("x", Int, func() {
	})
	Attribute("y", Int)
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

		Result(func() {
			Attribute("field", String)
			Attribute("start", Position)
			Attribute("goal", Position)
		})

		HTTP(func() {
			POST("/gen/")
			Response(StatusOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
	Files("/openapi.yaml", "./gen/http/openapi.yaml")
	Files("/swagger-ui/{*filepath}", "./public/swagger-ui/")
})
