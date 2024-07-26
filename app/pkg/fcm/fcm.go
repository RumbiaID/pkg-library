package fcm

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	"path/filepath"
)

var client *messaging.Client

func SetupFCM(credentialsPath string) error {
	ctx := context.Background()

	serviceAccountKeyFilePath, err := filepath.Abs(credentialsPath)
	if err != nil {
		return err
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	//Messaging client
	client, err = app.Messaging(ctx)
	if err != nil {
		return err
	}

	return nil
}

func createMessage(title, body, imageUrl string, data map[string]string) *messaging.Message {
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title:    title,
			Body:     body,
			ImageURL: imageUrl,
		},
		Data: data,
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
		},
		Webpush: &messaging.WebpushConfig{
			Headers: map[string]string{
				"Urgency": "high",
			},
		},
	}

	return message
}

func SendNotification(ctx context.Context, title, body, imageUrl string, data map[string]string, token string) error {
	message := createMessage(title, body, imageUrl, data)
	message.Token = token

	_, err := client.Send(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

func SendNotificationToTopic(ctx context.Context, title string, body, imageUrl string, data map[string]string, topic string) error {
	message := createMessage(title, body, imageUrl, data)
	message.Topic = topic

	_, err := client.Send(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

func RegisterClientToTopic(ctx context.Context, topic, token string) error {
	_, err := client.SubscribeToTopic(ctx, []string{token}, topic)
	if err != nil {
		return err
	}

	return nil
}

func UnregisterClientToTopic(ctx context.Context, topic, token string) error {
	_, err := client.UnsubscribeFromTopic(ctx, []string{token}, topic)
	if err != nil {
		return err
	}

	return nil
}
