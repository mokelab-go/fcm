package fcm

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

func (n *Notification) ToMap() map[string]interface{} {
	m := map[string]interface{}{}
	if len(n.Title) > 0 {
		m["title"] = n.Title
	}
	if len(n.Body) > 0 {
		m["body"] = n.Body
	}
	if len(n.Icon) > 0 {
		m["icon"] = n.Icon
	}
	if len(n.Sound) > 0 {
		m["sound"] = n.Sound
	}
	if len(n.Badge) > 0 {
		m["badge"] = n.Badge
	}
	if len(n.Tag) > 0 {
		m["tag"] = n.Tag
	}
	if len(n.Color) > 0 {
		m["color"] = n.Color
	}
	if len(n.ClickAction) > 0 {
		m["click_action"] = n.ClickAction
	}
	return m
}
