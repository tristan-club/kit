package pubsub

import (
	"encoding/json"
	"github.com/tristan-club/kit/log"
)

func SendEvent(appId string, uid string, event string, v interface{}) {
	marshal, err := json.Marshal(v)
	if err != nil {
		log.Error().Msgf("json encode event value error %s, value is %v", err, v)
		marshal = []byte("")
	}

	data := map[string]interface{}{
		"app_id":  appId,
		"user_id": uid,
		"event":   event,
		"value":   string(marshal),
	}

	go Publish("", data)
}

// SendTagEvent
// the first param tag is used for filter events in a pubsub topic for a specific subscription
func SendTagEvent(tag string, appId string, uid string, event string, v interface{}) {
	marshal, err := json.Marshal(v)
	if err != nil {
		log.Error().Msgf("json encode event value error %s, value is %v", err, v)
		marshal = []byte("")
	}

	data := map[string]interface{}{
		"app_id":  appId,
		"user_id": uid,
		"event":   event,
		"value":   string(marshal),
	}

	go Publish(tag, data)
}
