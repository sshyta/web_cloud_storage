package filters

import (
	"github.com/astaxie/beego/context"
)

// AuthMiddleware checks if the user is logged in and redirects to the login page if not
func AuthMiddleware(ctx *context.Context) {
	username := ctx.Input.Session("username")
	if username == nil {
		ctx.Redirect(302, "/")
	}
}
