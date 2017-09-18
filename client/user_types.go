// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "celler": Application User Types
//
// Command:
// $ goagen
// --design=github.com/ReSTARTR/goa-gorm-sample/design
// --out=$(GOPATH)/src/github.com/ReSTARTR/goa-gorm-sample
// --version=v1.3.0

package client

// userPayload user type.
type userPayload struct {
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	ID          *int    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Title       *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Title != nil {
		pub.Title = ut.Title
	}
	return &pub
}

// UserPayload user type.
type UserPayload struct {
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	ID          *int    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Title       *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}
