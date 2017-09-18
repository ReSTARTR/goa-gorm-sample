package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("celler", func() {
	Title("The virtual wine celler")
	Description("A simple goa service")
	Scheme("http")
	Host("localhost:8000")
	Origin("http://localhost:8000", nil)
})
