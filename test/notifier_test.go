package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/minisource/apiclients/notifier"
	"github.com/minisource/apiclients/notifier/models"
	"github.com/minisource/common_go/http/helper"
	"github.com/minisource/common_go/http/services"
)

func TestSendSMS(t *testing.T) {
	mockClient := new(MockAPIClient)
	jwtManager := services.NewJWTManager("notifier", "test_secret", "http://127.0.0.1:5001")
	mockResponse := []byte(`{"success": true}`)
	mockClient.On("DoRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(200, mockResponse, nil)

	client := helper.APIClient{
		BaseURL:    "http://127.0.0.1:5000",
		JWTManager: jwtManager,
	}
	notifier.NewNotifireService(&client)
	service := notifier.GetNotifierService()

	err := service.SendSMS(models.SMSRequest{
		To:       "+989011793041",
		Body:     "123456",
		Template: "otppersianbeauti",
	})

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}
