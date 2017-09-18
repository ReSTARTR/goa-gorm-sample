// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "celler": Models
//
// Command:
// $ goagen
// --design=github.com/ReSTARTR/goa-gorm-sample/design
// --out=$(GOPATH)/src/github.com/ReSTARTR/goa-gorm-sample
// --version=v1.3.0

package models

import (
	"context"
	"github.com/ReSTARTR/goa-gorm-sample/app"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"time"
)

// User Model Description
type User struct {
	ID          int `gorm:"primary_key"` // This is the User Model PK field
	CreatedAt   time.Time
	DeletedAt   *time.Time
	Description string
	Title       string
	UpdatedAt   time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m User) TableName() string {
	return "users"

}

// UserDB is the implementation of the storage interface for
// User.
type UserDB struct {
	Db *gorm.DB
}

// NewUserDB creates a new storage type.
func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{Db: db}
}

// DB returns the underlying database.
func (m *UserDB) DB() interface{} {
	return m.Db
}

// UserStorage represents the storage interface.
type UserStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, id int) (*User, error)
	Add(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error

	ListApplicationVndUser(ctx context.Context) []*app.ApplicationVndUser
	OneApplicationVndUser(ctx context.Context, id int) (*app.ApplicationVndUser, error)

	UpdateFromUserPayload(ctx context.Context, payload *app.UserPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserDB) TableName() string {
	return "users"

}

// CRUD Functions

// Get returns a single User as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserDB) Get(ctx context.Context, id int) (*User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "get"}, time.Now())

	var native User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of User
func (m *UserDB) List(ctx context.Context) ([]*User, error) {
	defer goa.MeasureSince([]string{"goa", "db", "user", "list"}, time.Now())

	var objs []*User
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *UserDB) Add(ctx context.Context, model *User) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding User", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *UserDB) Update(ctx context.Context, model *User) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating User", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "delete"}, time.Now())

	var obj User

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting User", "error", err.Error())
		return err
	}

	return nil
}

// UserFromUserPayload Converts source UserPayload to target User model
// only copying the non-nil fields from the source.
func UserFromUserPayload(payload *app.UserPayload) *User {
	user := &User{}
	if payload.Description != nil {
		user.Description = *payload.Description
	}
	if payload.ID != nil {
		user.ID = *payload.ID
	}
	if payload.Title != nil {
		user.Title = *payload.Title
	}

	return user
}

// UpdateFromUserPayload applies non-nil changes from UserPayload to the model and saves it
func (m *UserDB) UpdateFromUserPayload(ctx context.Context, payload *app.UserPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "user", "updatefromuserPayload"}, time.Now())

	var obj User
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving User", "error", err.Error())
		return err
	}
	if payload.Description != nil {
		obj.Description = *payload.Description
	}
	if payload.ID != nil {
		obj.ID = *payload.ID
	}
	if payload.Title != nil {
		obj.Title = *payload.Title
	}

	err = m.Db.Save(&obj).Error
	return err
}
