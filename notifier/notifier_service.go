package notifier

import (
	"github.com/minisource/apiclients_go/notifier/models"
	"github.com/minisource/common_go/http/helper"
)

var notifierService *NotifierService

type NotifierService struct {
	client *helper.APIClient
}

func NewNotifireService(client *helper.APIClient) {
	notifierService = &NotifierService{
		client: client,
	}
}

func GetNotifierService() *NotifierService {
	return notifierService
}

func (s *NotifierService) SendSMS(req models.SMSRequest) (error) {
	err := s.client.PostJSON(SMSRouter, req, nil)
    if err != nil {
        return err
    }

	return nil
}
