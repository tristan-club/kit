package pubsub

func SendEvent(uid string, event string, v interface{}) {
	data := map[string]interface{}{
		"user_id": uid,
		"event":   event,
		"value":   v,
	}

	go Publish(data)
}
