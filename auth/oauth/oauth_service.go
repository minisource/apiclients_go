package oauth

import (
	"github.com/minisource/apiclients/auth/oauth/models"
	"github.com/minisource/common_go/http/helper"
)

var oauthService *OAuthService

type OAuthService struct {
	client *helper.APIClient
}

func NewOAuthService(client *helper.APIClient) {
	oauthService = &OAuthService{
		client: client,
	}
}

func GetOAuthService() *OAuthService {
	return oauthService
}

func (s *OAuthService) ValidateToken(req models.ValidateOAuthTokenReq) (*models.ValidateOAuthTokenRes, error) {
	data, err := s.client.MakeRequest("POST", ValidateToken, req)
	if err != nil {
		return nil, err
	}

	// Prepare a variable to hold the deserialized result (User in this case)
	var token *models.ValidateOAuthTokenRes

	// Use the GenericDeserializer to deserialize the response
	err = helper.DeserializeResponse(data, &token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
