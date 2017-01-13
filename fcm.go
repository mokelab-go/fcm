package fcm

// Client is an interface for FCM server.
type Client interface {

	// Send FCM message to FCM server.
	Send(notification Notification, data Data, regIds ...string) (Response, error)
}

// Data is a "data" section in FCM message.
type Data map[string]interface{}

// Response is a response of FCM server.
type Response struct {
}
