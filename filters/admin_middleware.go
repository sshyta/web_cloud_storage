package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"web_cloud_storage/models"
)

// AdminMiddleware ensures that only users with roles_id = 1 (admin) can access certain pages
func AdminMiddleware(ctx *context.Context) {
	// Проверяем наличие сессии
	username := ctx.Input.Session("username")
	if username == nil {
		ctx.Redirect(302, "/")
		return
	}

	// Проверяем роль пользователя
	o := orm.NewOrm()
	user := models.Users{}
	err := o.QueryTable("users").Filter("login", username).One(&user)
	if err != nil {
		// Если пользователь не найден
		ctx.Redirect(302, "/")
		return
	}

	// Проверяем, является ли пользователь администратором
	if user.RolesID != 1 {
		ctx.Redirect(302, "/")
		return
	}
}
