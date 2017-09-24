package controller

import (
	"context"
	"errors"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"

	"github.com/ReSTARTR/goa-gorm-sample/app"
	"github.com/ReSTARTR/goa-gorm-sample/models"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here
	udb, err := dbFromContext(ctx)
	if err != nil {
		return err
	}
	u := models.UserFromUserPayload(ctx.Payload)
	if err := udb.Add(ctx, u); err != nil {
		return err
	}

	// UserController_Create: end_implement
	res := &app.ApplicationVndUser{}
	return ctx.OK(res)
}

// Index runs the index action.
func (c *UserController) Index(ctx *app.IndexUserContext) error {
	// UserController_Index: start_implement

	// Put your logic here
	udb, err := dbFromContext(ctx)
	if err != nil {
		return err
	}
	users := udb.ListApplicationVndUser(ctx)

	// UserController_Index: end_implement
	res := app.ApplicationVndUserCollection{}
	res = users
	return ctx.OK(res)
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here
	var udb *models.UserDB
	var err error
	udb, err = dbFromContext(ctx)
	if err != nil {
		return err
	}
	var user *app.ApplicationVndUser
	user, err = udb.OneApplicationVndUser(ctx, ctx.UserID)
	if err != nil {
		return err
	}

	// UserController_Show: end_implement
	res := &app.ApplicationVndUser{}
	res = user
	return ctx.OK(res)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement

	// Put your logic here
	udb, err := dbFromContext(ctx)
	if err != nil {
		return err
	}
	if err := udb.UpdateFromUserPayload(ctx, ctx.Payload, ctx.UserID); err != nil {
		return nil
	}

	// UserController_Update: end_implement
	res := &app.ApplicationVndUser{}
	return ctx.OK(res)
}

func dbFromContext(ctx context.Context) (*models.UserDB, error) {
	db, ok := (ctx.Value("DB")).(*gorm.DB)
	if !ok {
		return nil, errors.New("DB is not exists in context")
	}
	return models.NewUserDB(db), nil
}
