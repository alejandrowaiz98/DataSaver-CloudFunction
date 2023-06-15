package firestore

import (
	"context"
	"fmt"
	"log"
)

func (f *FirestoreDatabase) Save(messageData []byte, messageAttributes map[string]string) error {

	dataToStore := f.translator.Translate(messageData, messageAttributes)

	ctx := context.Background()

	_, err := f.client.Collection(dataToStore.DestinyCollection).Doc(dataToStore.MessageID).Set(ctx, dataToStore)

	if err != nil {
		return fmt.Errorf("[Firestore | Save] err: %v", err)
	}

	log.Printf("Message with ID %v has been saved", dataToStore.MessageID)

	return nil

}
