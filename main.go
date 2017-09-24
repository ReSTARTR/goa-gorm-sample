//go:generate goagen bootstrap -d github.com/ReSTARTR/goa-gorm-sample/design

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"

	"github.com/ReSTARTR/goa-gorm-sample/app"
	"github.com/ReSTARTR/goa-gorm-sample/controller"
	mymiddleware "github.com/ReSTARTR/goa-gorm-sample/middleware"
	"github.com/ReSTARTR/goa-gorm-sample/models"
)

var db *gorm.DB

func init() {
	config := &mysql.Config{
		User:      "root",
		Addr:      "127.0.0.1",
		DBName:    "goa_sample",
		ParseTime: true, // ref: https://github.com/Go-SQL-Driver/MySQL/issues/9
	}
	var err error
	db, err = gorm.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
	db.LogMode(true)
}

func main() {
	db.AutoMigrate(&models.User{})

	// Create service
	service := goa.New("celler")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	service.Use(mymiddleware.Db(db))

	// Mount "bottle" controller
	bc := controller.NewBottleController(service)
	app.MountBottleController(service, bc)
	// Mount "user" controller
	uc := controller.NewUserController(service)
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
