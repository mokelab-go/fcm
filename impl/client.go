package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mokelab-go/fcm"
	"net/http"
)

type client struct {
	apiKey string
}

func NewClient(apiKey string) *client {
	return &client{
		apiKey: apiKey,
	}
}

func (c *client) Send(notification Notification, data Data, regIds ...string) (fcm.Response, error) {
	return fcm.Response{}, nil
}
