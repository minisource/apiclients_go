package models

type ValidateOAuthTokenReq struct {
	Token    string   `json:"token"`
	Scop     *string  `json:"scope,omitempty"`
	Audience []string `json:"audience,omitempty"`
}

type ValidateOAuthTokenRes struct {
	Active    bool     `json:"active"`
	Aud       []string `json:"aud"`
	ClientID  string   `json:"client_id"`
	Exp       int64    `json:"exp"`
	Iat       int64    `json:"iat"`
	Iss       string   `json:"iss"`
	Nbf       int64    `json:"nbf"`
	Sub       string   `json:"sub"`
	TokenType string   `json:"token_type"`
	TokenUse  string   `json:"token_use"`
}
