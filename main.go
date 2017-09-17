//go:generate goagen bootstrap -d github.com/ReSTARTR/goa-sample/design

package main

import (
	"github.com/ReSTARTR/goa-sample/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("celler")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "bottle" controller
	c := NewBottleController(service)
	app.MountBottleController(service, c)

	// Start service
	if err := service.ListenAndServe(":8000"); err != nil {
		service.LogError("startup", "err", err)
	}

}
