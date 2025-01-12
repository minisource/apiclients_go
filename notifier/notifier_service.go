package notifier

import (
	"github.com/minisource/apiclients/notifier/models"
	"github.com/minisource/common_go/http/helper"
)

type NotifierService struct {
	client *helper.APIClient
}

func NewNotifireService(client *helper.APIClient) *NotifierService {
	return &NotifierService{
		client: client,
	}
}

func (s *NotifierService) SendSMS(req models.SMSRequest) (error) {
	_, err := s.client.MakeRequestWithAuthorization("POST", SMSRouter, req)
	if err != nil {
		return err
	}

	return nil
}
