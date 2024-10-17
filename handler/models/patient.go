package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nguyenvantuan2391996/patient-order-number/handler/constants"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/patient/models"
)

type PatientRequest struct {
	Channel     string `form:"channel"`
	Name        string `form:"name"`
	Sex         string `form:"sex"`
	RoomNumber  string `form:"room_number"`
	DoctorName  string `form:"doctor_name"`
	Status      string `form:"status"`
	OrderNumber int    `form:"order_number"`
	Age         int    `form:"age"`
}

type PatientSearchRequest struct {
	StartDate string `form:"start_date"`
	Limit     int    `form:"limit"`
	Page      int    `form:"page"`
}

func (r *PatientRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Sex, validation.In(constants.Male, constants.Female)),
		validation.Field(&r.Status, validation.In(constants.WaitingStatus, constants.DoingStatus, constants.DoneStatus)),
	)
}

func (r *PatientSearchRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.StartDate, validation.Required),
	)
}

func (r *PatientRequest) ToPatientInput() *models.PatientInput {
	out := &models.PatientInput{}
	if r == nil {
		return out
	}

	out.Channel = r.Channel
	out.Name = r.Name
	out.Sex = r.Sex
	out.RoomNumber = r.RoomNumber
	out.DoctorName = r.DoctorName
	out.Status = r.Status
	out.OrderNumber = r.OrderNumber
	out.Age = r.Age

	return out
}

func (r *PatientSearchRequest) ToPatientSearchInput() *models.PatientSearchInput {
	out := &models.PatientSearchInput{}
	if r == nil {
		return out
	}

	out.StartDate = r.StartDate
	if r.Limit <= 0 {
		r.Limit = 10
	}

	if r.Page <= 0 {
		r.Page = 1
	}

	out.Limit = r.Limit
	out.Offset = (r.Page - 1) * r.Limit

	return out
}
