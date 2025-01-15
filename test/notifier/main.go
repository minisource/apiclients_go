package main

import (
	"fmt"
	"time"

	"github.com/minisource/apiclients/notifier"
	"github.com/minisource/apiclients/notifier/models"
	"github.com/minisource/common_go/http/helper"
	"github.com/minisource/common_go/http/services"
)

func main() {
	jwtManager := services.NewJWTManager("notifier", "evuR6Oh5V5wQ_ZsKUYBvfVYIij", "http://127.0.0.1:5001")
	for {
		client := helper.APIClient{
			BaseURL:    "http://127.0.0.1:5000",
			JWTManager: jwtManager,
		}
		notifier.NewNotifireService(&client)
		service := notifier.GetNotifierService()
		err := service.SendSMS(models.SMSRequest{
			To:       "+1912",
			Body:     "123456",
			Template: "otppersianbeauti",
		})
		if err != nil {
			fmt.Printf("failed to send SMS: %v\n", err)
		}
		time.Sleep(5 * time.Second)
	}
}
