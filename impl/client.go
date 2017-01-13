package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mokelab-go/fcm"
	"net/http"
)

const (
	fcm_url = "https://fcm.googleapis.com/fcm/send"
)

type client struct {
	apiKey string
}

func NewClient(apiKey string) *client {
	return &client{
		apiKey: apiKey,
	}
}

func (c *client) Send(notification fcm.Notification, data fcm.Data, regIds ...string) (fcm.Response, error) {
	msg := map[string]interface{}{
		"registration_ids": regIds,
		"data":             data,
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return fcm.Response{}, err
	}
	req, err := http.NewRequest("POST", fcm_url, bytes.NewBuffer(body))
	if err != nil {
		return fcm.Response{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("key=%s", c.apiKey))
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fcm.Response{}, err
	}
	defer resp.Body.Close()

	return fcm.Response{}, nil
}
