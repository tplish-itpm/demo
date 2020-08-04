package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tplish-itpm/demo/pkgs/app"
	"github.com/tplish-itpm/demo/pkgs/gotest"
	"github.com/tplish-itpm/demo/routers"
	"testing"
)

var router *gin.Engine

var user map[string]interface{}
var res app.Response

func init() {
	gin.SetMode(gin.TestMode)

	user = map[string]interface{}{
		"username": "user",
		"password": "user",
	}
}

func TestRegister(t *testing.T) {
	router = routers.InitRouter()

	res = gotest.New(router, "GET", "/users", nil).Send().B()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, []interface{}{}, res.Data)

	res = gotest.New(router, "POST", "/register", gotest.Map2JsonReader(user)).Send().B()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "ok", res.Msg)

	res = gotest.New(router, "GET", "/users", nil).Send().B()
	assert.Equal(t, 200, res.Code)
	assert.NotNil(t, res.Data)
}

func TestLogin(t *testing.T) {
	router = routers.InitRouter()

	res = gotest.New(router, "GET", "/users", nil).Send().B()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, []interface{}{}, res.Data)

	res = gotest.New(router, "POST", "/register", gotest.Map2JsonReader(user)).Send().B()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "ok", res.Msg)

	uu := gotest.Map2JsonReader(map[string]interface{}{
		"username": "user",
		"password": "user",
	})
	res = gotest.New(router, "POST", "/login", uu).Send().B()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "ok", res.Msg)
}

func TestAuth(t *testing.T) {
	router = routers.InitRouter()

	res = gotest.New(router, "GET", "/guestPage", nil).Send().B()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "ok", res.Msg)

	res = gotest.New(router, "GET", "/userPage", nil).Send().B()
	assert.NotEqual(t, 200, res.Code)
	assert.NotEqual(t, "ok", res.Msg) // <==>

	res = gotest.New(router, "POST", "/register", gotest.Map2JsonReader(user)).Send().B()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "ok", res.Msg)

	w := gotest.New(router, "POST", "/login", gotest.Map2JsonReader(user)).Send().W
	assert.Equal(t, 200, w.Code)
	tp := gotest.TP{W: w}
	assert.Equal(t, "ok", tp.B().Msg)

	res = gotest.New(router, "GET", "/guestPage", nil).AddCookies(w.Result().Cookies()).Send().B()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "ok", res.Msg)

	res = gotest.New(router, "GET", "/userPage", nil).AddCookies(w.Result().Cookies()).Send().B()
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "ok", res.Msg) // <==>
}
