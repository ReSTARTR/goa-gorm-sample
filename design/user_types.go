package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

// UserPayload defines the data structure used in the create user request body.
// It is also the base type for the user media type used to render users.
var UserPayload = Type("UserPayload", func() {
	Attribute("title", func() {})
	Attribute("description", func() {})
})