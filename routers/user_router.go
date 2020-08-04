package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tplish-itpm/demo/models"
	"github.com/tplish-itpm/demo/pkgs/app"
	"github.com/tplish-itpm/demo/pkgs/e"
	"github.com/tplish-itpm/demo/services"
	_ "golang.org/x/crypto/bcrypt"
)

func InitUserRouter(r *gin.Engine) {
	r.POST("/register", Register)
	r.POST("/login", Login)
	r.GET("/users", ListUsers)
}

func Register(c *gin.Context) {
	appG := app.Gin{C: c}

	var user models.User

	// 绑定参数
	if err := c.ShouldBindJSON(&user); err != nil {
		appG.ResponseErr(e.ErrNewErr(err))
		return
	}

	// 是否注册成功
	if err := services.Register(&user); err != nil {
		appG.ResponseErr(err)
		return
	}

	appG.ResponseSuc(nil)
}

func Login(c *gin.Context) {
	appG := app.Gin{C: c}

	var user models.User

	// 绑定参数
	if err := c.ShouldBindJSON(&user); err != nil {
		appG.ResponseErr(e.ErrNewErr(err))
		return
	}

	// 是否登录成功
	if err := services.Login(&user); err != nil {
		appG.ResponseErr(err)
		return
	}

	session := sessions.Default(c)
	session.Set("role", user.Role.RoleName)
	if err := session.Save(); err != nil {
		appG.ResponseErr(e.ErrNewErr(err))
		return
	}

	appG.ResponseSuc(nil)
}

func ListUsers(c *gin.Context) {
	appG := app.Gin{C: c}

	users, err := models.ListUsers()
	if err != nil {
		appG.ResponseErr(err)
		return
	}

	appG.ResponseSuc(users)
}
