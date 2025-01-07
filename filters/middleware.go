package filters

import (
	"github.com/astaxie/beego/context"
)

// AuthMiddleware проверяет, есть ли пользователь в сессии, и перенаправляет на страницу входа, если нет
func AuthMiddleware(ctx *context.Context) {
	username := ctx.Input.Session("username")
	if username == nil {
		ctx.Redirect(302, "/")
	}
}
