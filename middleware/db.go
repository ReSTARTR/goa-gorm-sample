package middleware

import (
	"context"
	"net/http"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// Db returns handler
// ref: http://ikawaha.hateblo.jp/entry/2016/10/11/234941
func Db(d *gorm.DB) goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, r *http.Request) error {
			ctx = context.WithValue(ctx, "DB", d)
			return h(ctx, rw, r)
		}
	}
}
