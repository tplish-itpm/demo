package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/tplish-itpm/demo/models"
)

func InitRouter() *gin.Engine {
	models.InitDB()

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("auth", store))

	InitUserRouter(r)
	InitPageRouter(r)

	return r
}
