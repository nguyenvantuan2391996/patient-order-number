package patient

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/comoutput"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
	"github.com/nguyenvantuan2391996/patient-order-number/base_common/database/entities"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/patient/models"
	"github.com/nguyenvantuan2391996/patient-order-number/internal/domains/repository"
	"github.com/sirupsen/logrus"
)

type Patient struct {
	mapWSConnection map[string]*websocket.Conn
	accountRepo     repository.IAccountRepositoryInterface
	patientRepo     repository.IPatientRepositoryInterface
}

func NewPatientService(
	accountRepo repository.IAccountRepositoryInterface,
	patientRepo repository.IPatientRepositoryInterface) *Patient {
	return &Patient{
		mapWSConnection: make(map[string]*websocket.Conn),
		accountRepo:     accountRepo,
		patientRepo:     patientRepo,
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
	err := ps.patientRepo.Create(ctx, &entities.Patient{
		Name:       input.Name,
		Sex:        input.Sex,
		RoomNumber: input.RoomNumber,
		DoctorName: input.DoctorName,
		Status:     input.Status,
		Age:        input.Age,
	})
	if err != nil {
		logrus.Error(fmt.Sprintf(constants.FormatCreateEntityErr, "Patient", err))
		return nil, err
	}

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

func (ps *Patient) GetListPatient(ctx context.Context, input *models.PatientSearchInput) (*comoutput.BaseOutput, error) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginTask, "GetListPatient", input))

	var (
		records []*entities.Patient
		total   int64
		wg      sync.WaitGroup
		err     error
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, scopedErr := ps.patientRepo.List(ctx, map[string]interface{}{}, input.Limit, input.Offset,
			fmt.Sprintf("created_at >= %v", input.StartDate))
		if scopedErr != nil {
			logrus.Error(fmt.Sprintf(constants.FormatGetEntityErr, "patient", scopedErr))
			err = scopedErr
			return
		}

		records = res
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, scopedErr := ps.patientRepo.Total(ctx, map[string]interface{}{},
			fmt.Sprintf("created_at >= %v", input.StartDate))
		if scopedErr != nil {
			logrus.Error(fmt.Sprintf(constants.FormatTaskErr, "Total", scopedErr))
			err = scopedErr
			return
		}

		total = res
	}()

	wg.Wait()
	if err != nil {
		return nil, err
	}

	return &comoutput.BaseOutput{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"records": records,
			"total":   total,
		},
	}, nil
}

func (ps *Patient) UpdatePatient(ctx context.Context, input *models.PatientInput, id int64) (*comoutput.BaseOutput, error) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginTask, "UpdatePatient", input))

	// get patient from database
	record, err := ps.patientRepo.GetByQueries(ctx, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		logrus.Error(fmt.Sprintf(constants.FormatGetEntityErr, "patient", err))
		return nil, err
	}

	if record == nil {
		return nil, fmt.Errorf("the record is not found")
	}

	// update into database
	err = ps.patientRepo.UpdateWithMap(ctx, record, map[string]interface{}{
		"name":        input.Name,
		"sex":         input.Sex,
		"room_number": input.RoomNumber,
		"doctor_name": input.DoctorName,
		"status":      input.Status,
		"age":         input.Age,
	})
	if err != nil {
		logrus.Error(fmt.Sprintf(constants.FormatUpdateEntityErr, "Patient", err))
		return nil, err
	}

	return &comoutput.BaseOutput{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"id": id,
		},
	}, nil
}

func (ps *Patient) DeletePatient(ctx context.Context, id int64) (*comoutput.BaseOutput, error) {
	logrus.Info(fmt.Sprintf(constants.FormatBeginTask, "DeletePatient", id))

	// get patient from database
	record, err := ps.patientRepo.GetByQueries(ctx, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		logrus.Error(fmt.Sprintf(constants.FormatGetEntityErr, "patient", err))
		return nil, err
	}

	if record == nil {
		return nil, fmt.Errorf("the record is not found")
	}

	// delete into database
	err = ps.patientRepo.Delete(ctx, record)
	if err != nil {
		logrus.Error(fmt.Sprintf(constants.FormatDeleteEntityErr, "patient", err))
		return nil, err
	}

	return &comoutput.BaseOutput{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"id": id,
		},
	}, nil
}
