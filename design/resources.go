package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("bottle", func() {
	BasePath("/bottles")
	DefaultMedia(Bottle)

	Action("show", func() {
		Description("Get bottle by id")
		Routing(GET("/:bottleID"))
		Params(func() {
			Param("bottleID", Integer, "Bottle ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

var _ = Resource("user", func() {
	BasePath("/users")
	DefaultMedia(User)

	Action("index", func() {
		Description("List users")
		Routing(GET("/"))
		Response(OK, func() {
			Media(CollectionOf(User, func() {
				View("default")
			}))
		})
		Response(NotFound)
	})
	Action("show", func() {
		Description("Get user by id")
		Routing(GET("/:userID"))
		Params(func() {
			Param("userID", Integer, "User ID")
		})
		Response(OK)
		Response(NotFound)
	})
	Action("create", func() {
		Description("Create user")
		Routing(POST("/"))
		Payload(UserPayload)
		Response(OK)
		Response(NotFound)
	})
	Action("update", func() {
		Description("Update user")
		Routing(PUT("/:userID"))
		Params(func() {
			Param("userID", Integer, "User ID")
		})
		Payload(UserPayload)
		Response(OK)
		Response(NotFound)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swaggerui/*filepath", "swaggerui/dist")
})
