package fcm

type Client interface {
	// Sends FCM message to FCM server.
	Send(notification Notification, data Data, regIds ...string) (Response, error)
}

type Notification struct {
	Title       string
	Body        string
	Icon        string
	Sound       string
	Badge       string
	Tag         string
	Color       string
	ClickAction string
}

type Data map[string]interface{}
