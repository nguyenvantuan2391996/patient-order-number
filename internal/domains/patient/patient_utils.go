package patient

import (
	"fmt"

	"github.com/nguyenvantuan2391996/patient-order-number/base_common/constants"
	"github.com/sirupsen/logrus"
)

func (ps *Patient) writeMessage(channel string, message interface{}) error {
	if _, ok := ps.mapWSConnection[channel]; !ok {
		logrus.Error(fmt.Sprintf("websocket channel %v has not defined", channel))
		return fmt.Errorf("websocket channel %v has not defined", channel)
	}

	if err := ps.mapWSConnection[channel].WriteJSON(message); err != nil {
		logrus.Error(fmt.Sprintf(constants.FormatTaskErr, "WriteJSON", channel))
		return err
	}

	return nil
}
