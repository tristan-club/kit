package pubsub

func SendEvent(appId string, uid string, event string, v interface{}) {
	data := map[string]interface{}{
		"app_id":  appId,
		"user_id": uid,
		"event":   event,
		"value":   v,
	}

	go Publish("", data)
}

// SendTagEvent
// the first param tag is used for filter events in a pubsub topic for a specific subscription
func SendTagEvent(tag string, appId string, uid string, event string, v interface{}) {
	data := map[string]interface{}{
		"app_id":  appId,
		"user_id": uid,
		"event":   event,
		"value":   v,
	}

	go Publish(tag, data)
}
