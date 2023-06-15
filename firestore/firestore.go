package firestore

import (
	"context"
	"fmt"

	translator "alejandrowaiz.com/DataSaver/translator"
	"cloud.google.com/go/firestore"
)

type FirestoreDatabase struct {
	client     *firestore.Client
	translator translator.ITranslator
}

type IFirestoreDatabase interface {
	Save(messageData []byte, messageAttributes map[string]string) error
}

func New() (IFirestoreDatabase, error) {

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "looking-for-work-386300")

	if err != nil {
		return nil, fmt.Errorf("Err creating firestore client: %v", err)
	}

	t := translator.New()

	return &FirestoreDatabase{client: client, translator: t}, nil

}
