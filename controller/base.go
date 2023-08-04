package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Entry
)

func Init(l *logrus.Entry) {
	log = l
}

type baseController struct{}

func (ctl baseController) sendRespWithInternalError(ctx *gin.Context, data responseData) {
	log.Errorf("code: %s, err: %s", data.Code, data.Msg)

	ctx.JSON(http.StatusInternalServerError, data)
}

func (ctl baseController) sendCodeMessage(ctx *gin.Context, code string, err error) {
	if code == "" {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))
	} else {
		ctx.JSON(http.StatusBadRequest, newResponseCodeError(code, err))
	}
}

func (ctl baseController) sendRespOfPost(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, newResponseData(data))
}
