package handler

import (
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/admin"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/auth"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/patient"
)

type Handler struct {
	patientService *patient.Patient
	adminService   *admin.Admin
	authService    *auth.Auth
}

func NewHandler(patientService *patient.Patient, adminService *admin.Admin, authService *auth.Auth) *Handler {
	return &Handler{
		patientService: patientService,
		adminService:   adminService,
		authService:    authService,
	}
}
