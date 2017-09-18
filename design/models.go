package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("GoaSampleStorageGroup", func() {
	Description("This is the global storage group")
	Store("mysql", gorma.MySQL, func() {
		Description("This is the Postgres relational store")

		// User
		Model("User", func() {
			BuildsFrom(func() {
				// Payload("user", "create")
				// Payload("user", "update")
			})
			RendersTo(User)
			Description("User Model Description")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the User Model PK field")
			})
			Field("title", gorma.String, func() {})
			Field("description", gorma.String, func() {})
		})
	})
})
