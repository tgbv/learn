package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tgbv/learn/server1/internal/util"
)

func Ping(reqCtx *gin.Context) {
	reqCtx.String(http.StatusOK, "%s", "hellooo!")
}

func GetFile(reqCtx *gin.Context) {
	fcontent, error := ioutil.ReadFile(util.Pubpath(reqCtx.Param("target")))
	if error != nil {
		panic(error)
	}

	reqCtx.Header("Content-type", "text/html")
	reqCtx.String(http.StatusOK, "%s", fcontent)
}
