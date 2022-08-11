package pubsub

func SendEvent(appId string, uid string, event string, v interface{}) {
	data := map[string]interface{}{
		"app_id":  appId,
		"user_id": uid,
		"event":   event,
		"value":   v,
	}

	go Publish(data)
}
