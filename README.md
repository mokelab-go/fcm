# fcm
Firebase Cloud Messaging(fcm) client.

# How to use

Get!

    $ go get github.com/mokelab-go/fcm

Write!

```    
package main

import (
  "fmt"
  fcm "github.com/mokelab-go/fcm/impl"
)

const (
  apiKey = "INPUT YOUR API_KEY"
  regID = "INPUT YOUR REG_ID"
)

func main() {
	client := fcm.NewClient(apiKey)
	notification := fcm.Notification{
		Title: "test",
		Body:  "hello",
	}
	data := fcm.Data{
		"name": "moke",
	}
	_, err := client.Send(notification, data, regID)
	if err != nil {
		t.Errorf("Failed to send FCM message : %s", err)
	}
} 
```
