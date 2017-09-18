package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// User is the User resource media type.
var User = MediaType("application.vnd.user+json", func() {
	Description("user media type")
	Attributes(func() {
		Attribute("title", String, "title")
		Attribute("description", String, "description")
	})
	View("default", func() {
		Attribute("title")
		Attribute("description")
	})
})

// Bottle is the Bottle resource media type.
var Bottle = MediaType("application/vnd.goa.example.com+json", func() {
	Description("A bottle of wine")
	Attributes(func() {
		Attribute("id", Integer, "Unique bottle ID")
		Attribute("href", String, "API href for making requests on the bottle")
		Attribute("name", String, "Name of wine")
		Required("id", "href", "name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})
})
