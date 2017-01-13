package impl

import (
	"github.com/mokelab-go/fcm"
	"testing"
)

const (
	test_api_key = "INPUT_YOUR_API_KEY"
	test_reg_id  = "INPUT_YOUR_REG_ID"
)

func Test_Send(t *testing.T) {
	client := NewClient(test_api_key)
	notification := fcm.Notification{
		Title: "test",
		Body:  "hello",
	}
	data := fcm.Data{
		"name": "moke",
	}
	_, err := client.Send(notification, data, test_reg_id)
	if err != nil {
		t.Errorf("Failed to send FCM message : %s", err)
	}
}
