package patient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/comoutput"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/patient/models"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/repository"
	"github.com/sirupsen/logrus"
)

type Patient struct {
	mapWSConnection map[string]*websocket.Conn
	accountRepo     repository.IAccountRepositoryInterface
}

func NewPatientService(accountRepo repository.IAccountRepositoryInterface) *Patient {
	return &Patient{
		mapWSConnection: make(map[string]*websocket.Conn),
		accountRepo:     accountRepo,
	}
}

func (ps *Patient) InitWSPatient(channel string, conn *websocket.Conn) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginTask, "InitWSPatient", channel))
	if _, ok := ps.mapWSConnection[channel]; ok {
		err := ps.mapWSConnection[channel].Close()
		if err != nil {
			logrus.Error(fmt.Sprintf("close channel websocket %v is failed with err %v", channel, err))
		}
	}

	ps.mapWSConnection[channel] = conn
}

func (ps *Patient) CreatePatient(ctx context.Context, input *models.PatientInput) (*comoutput.BaseOutput, error) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginTask, "CreatePatient", input))

	// insert database
	// write message to socket
	go func() {
		_ = ps.writeMessage(input.Channel, &comoutput.BaseOutput{
			Status:  http.StatusBadRequest,
			Message: "OK",
			Data:    input,
		})
	}()

	return &comoutput.BaseOutput{
		Status: http.StatusOK,
		Data:   input,
	}, nil
}
