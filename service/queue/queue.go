package queue

import (
	"encoding/json"
	"golang.org/x/net/context"
	"log"
	"os"
)
import "cloud.google.com/go/pubsub"

var client *pubsub.Client

func Instance() (*pubsub.Client, error) {
	if client != nil {
		return client, nil
	}
	var err error
	client, err = pubsub.NewClient(context.Background(), "logicful-290101")
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Setup() {
	setupTopic("form")
	setupTopic("submissions")
	setupTopic("integration-email")
	setupTopic("integration-webhook")
	setupTopic("team-invite")
	setupSubscription("form", "form-processor")
	setupSubscription("submissions", "workflow")
	setupSubscription("integration-email", "send-submission-email")
	setupSubscription("integration-webhook", "send-submission-webhook")
	setupSubscription("team-invite", "team-invite-email")
}

func setupTopic(name string) {
	client, err := Instance()
	if err != nil {
		panic(err)
	}
	exists, err := client.Topic(name).Exists(context.Background())
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := client.CreateTopic(context.Background(), name)
		if err != nil {
			panic(err)
		}
	}
}

func setupSubscription(topic string, subscriber string) {
	client, err := Instance()
	if err != nil {
		panic(err)
	}
	exists, err := client.Subscription(prefix(subscriber)).Exists(context.Background())
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := client.CreateSubscription(context.Background(), prefix(subscriber), pubsub.SubscriptionConfig{
			Topic: getTopic(topic),
		})
		if err != nil {
			panic(err)
		}
	}
}

func getTopic(name string) *pubsub.Topic {
	return client.Topic(name)
}

func Dispatch(name string, data interface{}) error {
	client, err := Instance()
	if err != nil {
		return err
	}
	serialized, err := json.Marshal(data)
	if err != nil {
		return err
	}
	topic := client.Topic(name)
	if _, err := topic.Publish(context.Background(), &pubsub.Message{
		Data: serialized,
	}).Get(context.Background()); err != nil {
		return err
	}
	return nil
}

func Receive(subscription string, cb func([]byte) error) {
	client, err := Instance()
	if err != nil {
		log.Fatal(err.Error())
	}
	sub := client.Subscription(prefix(subscription))
	go func() {
		err = sub.Receive(context.Background(), func(ctx context.Context, message *pubsub.Message) {
			err = cb(message.Data)
			if err != nil {
				println("Queue error " + prefix(subscription) + " " + err.Error())
				message.Nack()
				return
			}
			message.Ack()
		})
		if err != nil {
			println(err.Error())
		}
	}()
}

func prefix(value string) string {
	var prefix = os.Getenv("QUEUE_PREFIX")
	return prefix + "_" + value
}
