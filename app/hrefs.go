// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "celler": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/ReSTARTR/goa-gorm-sample/design
// --out=$(GOPATH)/src/github.com/ReSTARTR/goa-gorm-sample
// --version=v1.3.0

package app

import (
	"fmt"
	"strings"
)

// BottleHref returns the resource href.
func BottleHref(bottleID interface{}) string {
	parambottleID := strings.TrimLeftFunc(fmt.Sprintf("%v", bottleID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/bottles/%v", parambottleID)
}

// UserHref returns the resource href.
func UserHref(userID interface{}) string {
	paramuserID := strings.TrimLeftFunc(fmt.Sprintf("%v", userID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/users/%v", paramuserID)
}
