package models

type ValidateAccessTokenRequest struct{
	AccessToken     string `json:"accessToken"`
}

type ValidateAuthTokenRes struct {
	Claims map[string]interface{} `json:"claims"`
}
