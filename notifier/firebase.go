package notifier

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FirebaseNotifier struct {
	Firebase *firebase.App
}

func NewFirebaseNotifier(serviceAccountKeyFilePath string) (*FirebaseNotifier, error) {
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	return &FirebaseNotifier{Firebase: app}, nil
}

func (notifier *FirebaseNotifier) SendNotification(token string, notification *messaging.Notification) error {
	if client, err := notifier.Firebase.Messaging(context.Background()); err != nil {
		return err
	} else {
		if _, err := client.Send(context.Background(), &messaging.Message{Notification: notification, Token: token}); err != nil {
			return err
		}
	}
	return nil
}
