package auth

import (
	"github.com/minisource/apiclients/auth/auth/models"
	"github.com/minisource/common_go/http/helper"
)


var authService *AuthService

type AuthService struct {
	client *helper.APIClient
}

func NewAuthService(client *helper.APIClient) {
	authService = &AuthService{
		client: client,
	}
}

func GetAuthService() *AuthService {
	return authService
}

func (s *AuthService) ValidateAccessToken(req models.ValidateAccessTokenRequest) (*models.ValidateAuthTokenRes, error) {
	data, err := s.client.MakeRequest("POST", ValidateToken, req)
	if err != nil {
		return nil, err
	}

	// Prepare a variable to hold the deserialized result (User in this case)
	var token *models.ValidateAuthTokenRes

	// Use the GenericDeserializer to deserialize the response
	err = helper.DeserializeResponse(data, &token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
