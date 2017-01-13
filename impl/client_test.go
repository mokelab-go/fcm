package impl

import (
	"github.com/mokelab-go/fcm"
	"testing"
)

const (
	testAPIKey = "INPUT_YOUR_API_KEY"
	testRegID  = "INPUT_YOUR_REG_ID"
)

func Test_Send(t *testing.T) {
	client := NewClient(testAPIKey)
	notification := fcm.Notification{
		Title: "test",
		Body:  "hello",
	}
	data := fcm.Data{
		"name": "moke",
	}
	_, err := client.Send(notification, data, testRegID)
	if err != nil {
		t.Errorf("Failed to send FCM message : %s", err)
	}
}
