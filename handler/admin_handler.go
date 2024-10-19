package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
	response "github.com/nguyenvantuan2391996/patient-order-number/base_common/response"
	"github.com/nguyenvantuan2391996/patient-order-number/handler/models"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateAccount(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "CreateAccount"))
	request := models.AccountRequest{}
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

	result, err := h.adminService.CreateAccount(ctx, request.ToAccountInput())
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "CreateAccount", err)
		ctx.JSON(http.StatusInternalServerError, responseAPI.ToResponse(constants.InternalServerError,
			nil, constants.ResponseMessage[constants.InternalServerError]))
		return
	}

	responseAPI.ToResponse(result.Status, result.Data, result.Message)
	ctx.JSON(result.Status, responseAPI)
}

func (h *Handler) UpdateAccount(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "UpdateAccount"))
	request := models.AccountUpdateRequest{}
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

	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		return
	}
	result, err := h.adminService.UpdateAccount(ctx, request.ToAccountUpdateInput(int64(userID)))
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "UpdateAccount", err)
		ctx.JSON(http.StatusInternalServerError, responseAPI.ToResponse(constants.InternalServerError,
			nil, constants.ResponseMessage[constants.InternalServerError]))
		return
	}

	responseAPI.ToResponse(result.Status, result.Data, result.Message)
	ctx.JSON(result.Status, responseAPI)
}

func (h *Handler) DeleteAccount(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "DeleteAccount"))
	request := models.DeleteAccountRequest{}
	responseAPI := response.NewResponse(ctx)

	err := ctx.ShouldBind(&request)
	if err != nil {
		logrus.Warnf(constants.FormatTaskErr, "ShouldBind", err)
		ctx.JSON(http.StatusBadRequest, responseAPI.InputError().Msg(response.ErrorMsgInput))
		return
	}

	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		return
	}
	result, err := h.adminService.DeleteAccount(ctx, request.ToDeleteAccountInput(int64(userID)))
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "DeleteAccount", err)
		ctx.JSON(http.StatusInternalServerError, responseAPI.ToResponse(constants.InternalServerError,
			nil, constants.ResponseMessage[constants.InternalServerError]))
		return
	}

	responseAPI.ToResponse(result.Status, result.Data, result.Message)
	ctx.JSON(result.Status, responseAPI)
}
