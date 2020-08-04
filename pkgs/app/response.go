package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tplish-itpm/demo/pkgs/e"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg"  example:"ok"`
	Data interface{} `json:"data" example:"null"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}

func (g *Gin) ResponseSuc(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: e.Success,
		Msg:  e.GetMsg(e.Success),
		Data: data,
	})
}

func (g *Gin) ResponseFail() {
	g.C.JSON(http.StatusOK, Response{
		Code: e.Fail,
		Msg:  e.GetMsg(e.Fail),
		Data: nil,
	})
}

func (g *Gin) ResponseErr(err e.Error) {
	g.C.JSON(http.StatusOK, Response{
		Code: err.Code(),
		Msg:  err.Error(),
		Data: nil,
	})
}

func (g *Gin) ResponseErrCode(code int) {
	g.C.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: nil,
	})
}
