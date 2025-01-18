package test

import (
	"testing"

	"github.com/minisource/apiclients/auth/oauth"
	"github.com/minisource/apiclients/auth/oauth/models"
	"github.com/minisource/common_go/http/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAPIClient for testing purposes
type MockAPIClient struct {
	mock.Mock
}

func TestValidateToken(t *testing.T) {
	mockClient := new(MockAPIClient)
	mockResponse := []byte(`{"valid": true}`)
	mockClient.On("DoRequest", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(200, mockResponse, nil)

	client := &helper.APIClient{
		BaseURL:    "http://127.0.0.1:5000",
	}
	oauth.NewOAuthService(client)
	service := oauth.GetOAuthService()

	resp, err := service.ValidateToken(models.ValidateOAuthTokenReq{
		Token: "ory_at_xW5O48OD9fF8MqK3sDC4mm8EJ8k0axn3Lsq8HZYWJQg.xCCpqkd0HXafg_w2f2F6dPllHQ_vMgImLoQ-LeW-_7M",
	})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	mockClient.AssertExpectations(t)
}
