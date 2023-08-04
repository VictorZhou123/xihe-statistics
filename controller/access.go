package controller

import (
	"net/http"
	"project/xihe-statistics/app"
	"project/xihe-statistics/domain/repository"

	"github.com/gin-gonic/gin"
)

func AddRouterForAccessController(
	rg *gin.RouterGroup,
	r repository.Access,
) {
	ctl := AccessController{
		s: app.NewAccessService(r),
	}

	rg.GET("/v1/access", ctl.AddAccess)
}

type AccessController struct {
	baseController
	s app.AccessService
}

// @Summary AddAccess
// @Description add record of access
// @Tags  Access
// @Param  body  body  AccessRequest  true  "body of add access"
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/access [get]
func (ctl *AccessController) AddAccess(ctx *gin.Context) {
	req := AccessRequest{}
	cmd, err := req.toCmd(ctx.ClientIP(), ctx.GetHeader("Referer"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newResponseCodeError(
			errorBadRequestParam, err,
		))

		return
	}

	if code, err := ctl.s.AddAccess(&cmd); err != nil {
		ctl.sendCodeMessage(ctx, code, err)
	} else {
		ctl.sendRespOfPost(ctx, "success")
	}
}
