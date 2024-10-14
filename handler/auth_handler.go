package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
	response "github.com/nguyenvantuan2391996/patient-order-number/base_common/response"
	"github.com/nguyenvantuan2391996/patient-order-number/handler/models"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Login(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "Login"))
	request := models.LoginRequest{}
	responseAPI := response.NewResponse(ctx)

	err := ctx.ShouldBind(&request)
	if err != nil {
		logrus.Warnf(constants.FormatTaskErr, "ShouldBind", err)
		ctx.JSON(http.StatusBadRequest, responseAPI.InputError().Msg(response.ErrorMsgInput))
		return
	}

	if err := request.Validate(); err != nil {
		logrus.Errorf(constants.FormatTaskErr, "Validate", err)
		ctx.JSON(http.StatusBadRequest, responseAPI.InputError().Msg(err.Error()))
		return
	}

	result, err := h.authService.Login(ctx, request.ToLoginInput())
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "Login", err)
		ctx.JSON(http.StatusInternalServerError, responseAPI.ToResponse(constants.InternalServerError,
			nil, constants.ResponseMessage[constants.InternalServerError]))
		return
	}

	responseAPI.ToResponse(result.Status, result.Data, result.Message)
	ctx.JSON(result.Status, responseAPI)
}
