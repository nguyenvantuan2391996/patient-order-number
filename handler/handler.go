package handler

import (
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/auth"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/patient"
)

type Handler struct {
	patientService *patient.Patient
	authService    *auth.Auth
}

func NewHandler(patientService *patient.Patient, authService *auth.Auth) *Handler {
	return &Handler{
		patientService: patientService,
		authService:    authService,
	}
}
