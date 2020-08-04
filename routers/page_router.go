package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tplish-itpm/demo/models"
	"github.com/tplish-itpm/demo/pkgs/app"
	"github.com/tplish-itpm/demo/pkgs/e"
)

func InitPageRouter(r *gin.Engine) {
	r.GET("/guestPage", guestPage)
	r.GET("/userPage", userPage)
	r.GET("/adminPage", adminPage)
}

func guestPage(c *gin.Context) {
	appG := app.Gin{C: c}

	appG.ResponseSuc(nil)
}

func userPage(c *gin.Context) {
	appG := app.Gin{C: c}

	session := sessions.Default(c)
	role := session.Get("role")
	if role != models.RoleUser {
		appG.ResponseErrCode(e.ErrPermission)
		return
	}

	appG.ResponseSuc(nil)
}

func adminPage(c *gin.Context) {
	appG := app.Gin{C: c}

	session := sessions.Default(c)
	role := session.Get("role")
	if role != "ROLE_ADMIN" {
		appG.ResponseErrCode(e.ErrPermission)
		return
	}

	appG.ResponseSuc(nil)
}
