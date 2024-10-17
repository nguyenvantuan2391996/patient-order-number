package models

type PatientInput struct {
	Channel     string `json:"channel"`
	Name        string `json:"name"`
	Sex         string `json:"sex"`
	RoomNumber  string `json:"room_number"`
	DoctorName  string `json:"doctor_name"`
	Status      string `json:"status"`
	OrderNumber int    `json:"order_number"`
	Age         int    `json:"age"`
}

type PatientSearchInput struct {
	StartDate string
	Limit     int
	Offset    int
}
