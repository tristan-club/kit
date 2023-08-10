package pubsub

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/tristan-club/kit/log"
)

var projectID = ""
var topic = ""

var ctx = context.Background()
var client *pubsub.Client
var clientErr error

func Publish(tag string, data map[string]interface{}) {
	PublishTopic(topic, tag, data)
}

func PublishTopic(topic string, tag string, data map[string]interface{}) {
	if client == nil {
		return
	}

	if data != nil {
		data["created_time"] = time.Now().Format("2006-01-02T15:04:05.000Z")
	} else {
		log.Error().Msgf("send nil pubsub message")
		return
	}

	marshal, err := json.Marshal(data)

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	t := client.Topic(topic)

	result := t.Publish(context.Background(), &pubsub.Message{
		Data: marshal,
		Attributes: map[string]string{
			"tag": tag,
		},
	})

	_, err = result.Get(context.Background())

	if err != nil {
		log.Error().Err(err).Send()
	}
}

func Init() {
	if v := os.Getenv("PUBSUB_PROJECT_ID"); v != "" {
		projectID = v
	}

	if v := os.Getenv("PUBSUB_TOPIC"); v != "" {
		topic = v
	}

	client, clientErr = pubsub.NewClient(ctx, projectID)
	if clientErr != nil {
		client = nil
		log.Error().Msgf("pubsub client error: %s", clientErr)
		return
	}
}
