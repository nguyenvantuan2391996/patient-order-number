package handler

import "github.com/nguyenvantuan2391996/patient-order-number/internal/domains/patient"

type Handler struct {
	patientService *patient.Patient
}

func NewHandler(patientService *patient.Patient) *Handler {
	return &Handler{
		patientService: patientService,
	}
}
