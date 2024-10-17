package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
	response "github.com/nguyenvantuan2391996/patient-order-number/base_common/response"
	"github.com/nguyenvantuan2391996/patient-order-number/handler/models"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) GetPatientPage(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "GetPatientPage"))

	ctx.HTML(200, "patient.html", gin.H{
		"title": "Patient List",
	})
}

func (h *Handler) LoginPatientPage(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "LoginPatientPage"))

	ctx.HTML(200, "admin-login.html", gin.H{
		"title": "Login Patient",
	})
}

func (h *Handler) GetAdminPatientPage(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "GetAdminPatientPage"))

	ctx.HTML(200, "admin-patient.html", gin.H{
		"title": "Admin Patient",
	})
}

func (h *Handler) InitWSPatient(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "InitWSPatient"))
	responseAPI := response.NewResponse(ctx)

	protocol := "http://"
	if ctx.Request.TLS != nil {
		protocol = "https://"
	}

	if ctx.Request.Header.Get("Origin") != protocol+ctx.Request.Host {
		logrus.Error("request origin is not equal host")
		ctx.JSON(http.StatusInternalServerError, responseAPI.ToResponse(constants.InternalServerError,
			nil, constants.ResponseMessage[constants.InternalServerError]))
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Error(fmt.Sprintf(constants.FormatTaskErr, "Upgrade", err))
		ctx.JSON(http.StatusInternalServerError, responseAPI.ToResponse(constants.InternalServerError,
			nil, constants.ResponseMessage[constants.InternalServerError]))
		return
	}

	h.patientService.InitWSPatient(ctx.Param("channel"), conn)
}

func (h *Handler) CreatePatient(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "CreatePatient"))
	request := models.PatientRequest{}
	responseAPI := response.NewResponse(ctx)

	err := ctx.ShouldBind(&request)
	if err != nil {
		logrus.Warnf(constants.FormatTaskErr, "ShouldBind", err)
		ctx.JSON(http.StatusBadRequest, responseAPI.InputError().Msg(response.ErrorMsgInput))
		return
	}

	if err = request.Validate(); err != nil {
		logrus.Errorf(constants.FormatTaskErr, "Validate", err)
		ctx.JSON(http.StatusBadRequest, responseAPI.InputError().Msg(err.Error()))
		return
	}

	result, err := h.patientService.CreatePatient(ctx, request.ToPatientInput())
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "CreatePatient", err)
		ctx.JSON(http.StatusInternalServerError, responseAPI.ToResponse(constants.InternalServerError,
			nil, constants.ResponseMessage[constants.InternalServerError]))
		return
	}

	responseAPI.ToResponse(result.Status, result.Data, result.Message)
	ctx.JSON(result.Status, responseAPI)
}

func (h *Handler) GetListPatient(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "GetListPatient"))
	request := models.PatientSearchRequest{}
	responseAPI := response.NewResponse(ctx)

	err := ctx.ShouldBind(&request)
	if err != nil {
		logrus.Warnf(constants.FormatTaskErr, "ShouldBind", err)
		ctx.JSON(http.StatusBadRequest, responseAPI.InputError().Msg(response.ErrorMsgInput))
		return
	}

	if err = request.Validate(); err != nil {
		logrus.Errorf(constants.FormatTaskErr, "Validate", err)
		ctx.JSON(http.StatusBadRequest, responseAPI.InputError().Msg(err.Error()))
		return
	}

	result, err := h.patientService.GetListPatient(ctx, request.ToPatientSearchInput())
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "GetListPatient", err)
		ctx.JSON(http.StatusInternalServerError, responseAPI.ToResponse(constants.InternalServerError,
			nil, constants.ResponseMessage[constants.InternalServerError]))
		return
	}

	responseAPI.ToResponse(result.Status, result.Data, result.Message)
	ctx.JSON(result.Status, responseAPI)
}

func (h *Handler) UpdatePatient(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "UpdatePatient"))
	request := models.PatientRequest{}
	responseAPI := response.NewResponse(ctx)

	err := ctx.ShouldBind(&request)
	if err != nil {
		logrus.Warnf(constants.FormatTaskErr, "ShouldBind", err)
		ctx.JSON(http.StatusBadRequest, responseAPI.InputError().Msg(response.ErrorMsgInput))
		return
	}

	if err = request.Validate(); err != nil {
		logrus.Errorf(constants.FormatTaskErr, "Validate", err)
		ctx.JSON(http.StatusBadRequest, responseAPI.InputError().Msg(err.Error()))
		return
	}

	patientID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Warnf(constants.FormatTaskErr, "Atoi", err)
		ctx.JSON(http.StatusBadRequest, responseAPI.InputError().Msg(response.ErrorMsgInput))
		return
	}
	result, err := h.patientService.UpdatePatient(ctx, request.ToPatientInput(), int64(patientID))
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "UpdatePatient", err)
		ctx.JSON(http.StatusInternalServerError, responseAPI.ToResponse(constants.InternalServerError,
			nil, constants.ResponseMessage[constants.InternalServerError]))
		return
	}

	responseAPI.ToResponse(result.Status, result.Data, result.Message)
	ctx.JSON(result.Status, responseAPI)
}

func (h *Handler) DeletePatient(ctx *gin.Context) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginAPI, "DeletePatient"))
	responseAPI := response.NewResponse(ctx)

	patientID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Warnf(constants.FormatTaskErr, "Atoi", err)
		ctx.JSON(http.StatusBadRequest, responseAPI.InputError().Msg(response.ErrorMsgInput))
		return
	}
	result, err := h.patientService.DeletePatient(ctx, int64(patientID))
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "DeletePatient", err)
		ctx.JSON(http.StatusInternalServerError, responseAPI.ToResponse(constants.InternalServerError,
			nil, constants.ResponseMessage[constants.InternalServerError]))
		return
	}

	responseAPI.ToResponse(result.Status, result.Data, result.Message)
	ctx.JSON(result.Status, responseAPI)
}
