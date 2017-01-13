package fcm

type Client interface {
	// Sends FCM message to FCM server.
	Send(notification Notification, data Data, regIds ...string) (Response, error)
}

type Data map[string]interface{}

type Response struct {
}
