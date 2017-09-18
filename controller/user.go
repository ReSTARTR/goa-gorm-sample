package controller

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"

	"github.com/ReSTARTR/goa-sample/app"
	"github.com/ReSTARTR/goa-sample/models"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
	db *gorm.DB
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service, db *gorm.DB) *UserController {
	return &UserController{
		Controller: service.NewController("UserController"),
		db: db,
	}
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here
	var user *app.ApplicationVndUser
	var err error
	udb := models.NewUserDB(c.db)
	user, err = udb.OneApplicationVndUser(ctx, ctx.UserID)
	if err != nil {
		return err
	}

	// UserController_Show: end_implement
	res := &app.ApplicationVndUser{}
	res = user
	return ctx.OK(res)
}
