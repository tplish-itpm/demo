package gotest

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tplish-itpm/demo/pkgs/app"
	"io"
	"net/http"
	"net/http/httptest"
)

type TP struct {
	Req    *http.Request
	Router *gin.Engine
	W      *httptest.ResponseRecorder
}

func New(router *gin.Engine, method, url string, body io.Reader) *TP {
	var tp TP
	tp.Router = router
	tp.Req = httptest.NewRequest(method, url, body)
	return &tp
}

func (tp *TP) AddCookies(cookies []*http.Cookie) *TP {
	for _, c := range cookies {
		tp.Req.AddCookie(c)
	}
	return tp
}

func (tp *TP) Send() *TP {
	tp.W = httptest.NewRecorder()
	tp.Router.ServeHTTP(tp.W, tp.Req)
	return tp
}

func (tp *TP) B() app.Response {
	var res app.Response
	_ = json.Unmarshal(tp.W.Body.Bytes(), &res)
	return res
}

func Map2JsonReader(mp map[string]interface{}) *bytes.Reader {
	body, _ := json.Marshal(mp)
	return bytes.NewReader(body)
}
