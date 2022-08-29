package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"github.com/tristan-club/kit/log"
	"google.golang.org/api/option"
	"os"
	"time"
)

type PubSub struct {
	project      string
	topic        string
	subscription string
	attribute    map[string]string
}

type MessageCallback func(msg *pubsub.Message) error

func NewPubSub(proj string, topic string, subscription string) *PubSub {
	return &PubSub{
		project:      proj,
		topic:        topic,
		subscription: subscription,
		attribute:    map[string]string{},
	}
}

/*
Publish a message to the topic
Notice: this function is synchronous, so it will block until the message is published
the caller should call this function in a goroutine
*/
func (p *PubSub) Publish(data interface{}) (string, error) {
	var ctx = context.Background()
	// todo: reuse the client
	var client, err = pubsub.NewClient(ctx, p.project, option.WithEndpoint(os.Getenv("PUBSUB_END_POINT")))
	defer client.Close()

	if err != nil {
		log.Error().Err(err).Send()
		return "", nil
	}

	marshal, err := json.Marshal(data)

	if err != nil {
		log.Error().Err(err).Send()
		return "", nil
	}

	t := client.Topic(p.topic)
	t.EnableMessageOrdering = true

	log.Info().Fields(map[string]interface{}{
		"action": "publish pubsub",
		"data":   data,
	}).Send()

	msg := &pubsub.Message{
		Data:        marshal,
		Attributes:  p.attribute,
		OrderingKey: "key",
	}

	result := t.Publish(context.Background(), msg)

	id, err := result.Get(context.Background())

	if err != nil {
		log.Error().Err(err).Send()
		return "", err
	}

	return id, nil
}

/*
PullMessage will pull a message from the topic
using synchronous method, so it will block until a message is pulled
if callback function returns error, the message will not be consumed
and will be pulled again
callback 业务完成以后， 返回nil， 否则返回错误
一旦返回错误，在一段时间之内，消息会继续被重新推送
如果超过时间，还没有成功， 则消息会进入死信队列
*/
func (p *PubSub) PullMessage(callback MessageCallback, timeout time.Duration) error {
	var ctx = context.Background()
	var client, err = pubsub.NewClient(ctx, p.project, option.WithEndpoint(os.Getenv("PUBSUB_END_POINT")))
	defer client.Close()

	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	sub := client.Subscription(p.subscription)
	sub.ReceiveSettings.Synchronous = true

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// 下面两个值必须相等
	sub.ReceiveSettings.MaxOutstandingMessages = 1

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		if err := callback(msg); err != nil {
			log.Error().Msgf("[PubSub] callback error: %v, msg will be resend", err)
			msg.Nack()
		} else {
			msg.Ack()
		}
	})

	return err
}

func init() {
	if os.Getenv("PUBSUB_END_POINT") == "" {
		log.Error().Msgf("❌PUBSUB Error, Please Specify Endpoint")
	}
}
