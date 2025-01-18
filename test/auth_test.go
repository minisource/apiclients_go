package test

// import (
// 	"fmt"

// 	"github.com/minisource/apiclients/auth/auth"
// 	"github.com/minisource/apiclients/auth/auth/models"
// 	"github.com/minisource/common_go/http/helper"
// 	"github.com/minisource/common_go/http/services"
// )

// func main() {
// 	jwtManager := services.NewJWTManager("auth", "2vnwyhUbf8md564wDJWXzQIaL_", "http://127.0.0.1:5001")
// 	for {
// 		client := helper.APIClient{
// 			BaseURL:    "http://127.0.0.1:5000",
// 			JWTManager: jwtManager,
// 		}
// 		auth.NewAuthService(&client)
// 		service := auth.GetAuthService()
// 		res, err := service.ValidateAccessToken(models.ValidateAccessTokenRequest{
// 			AccessToken: "test",
// 		})
// 		if err != nil {
// 			fmt.Printf("failed to send SMS: %v\n", err)
// 		}
// 		fmt.Print(res)
// 	}
// }