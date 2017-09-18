//go:generate goagen bootstrap -d github.com/ReSTARTR/goa-sample/design

package main

import (
	"net/http"

	"github.com/ReSTARTR/goa-sample/app"
	"github.com/ReSTARTR/goa-sample/controller"
	"github.com/ReSTARTR/goa-sample/models"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

func main() {
	db.AutoMigrate(&models.User{})

	// Create service
	service := goa.New("celler")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "bottle" controller
	bc := controller.NewBottleController(service)
	app.MountBottleController(service, bc)
	// Mount "user" controller
	uc := controller.NewUserController(service, db)
	app.MountUserController(service, uc)
	// Mount "swagger" controller
	sc := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, sc)

	mux := http.NewServeMux()
	mux.Handle("/", service.Mux)
	Admin := admin.New(&qor.Config{DB: db})
	Admin.AddResource(&models.User{})
	Admin.MountTo("/admin", mux)

	// Start service
	if err := http.ListenAndServe(":8000", mux); err != nil {
		service.LogError("startup", "err", err)
	}
}
