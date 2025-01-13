package main

import (
	"fmt"

	"github.com/minisource/apiclients/auth"
	"github.com/minisource/apiclients/auth/models"
	"github.com/minisource/common_go/http/helper"
)

func main() {
	client := helper.APIClient{
		BaseURL:    "http://127.0.0.1:5001",
	}
	service := auth.NewAuthService(&client)
	_, err := service.ValidateToken(models.ValidateTokenReq{
		Token: "ory_at_xW5O48OD9fF8MqK3sDC4mm8EJ8k0axn3Lsq8HZYWJQg.xCCpqkd0HXafg_w2f2F6dPllHQ_vMgImLoQ-LeW-_7M",
	})
	if err != nil {
		fmt.Printf("failed to retrive token: %v\n", err)
	}
}
