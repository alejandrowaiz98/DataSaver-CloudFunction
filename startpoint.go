package p

import (
	"context"
	"fmt"
	"log"

	firestore "alejandrowaiz.com/DataSaver/firestore"
)

type PubSubMessage struct {
	Data       []byte            `json:"data"`
	Attributes map[string]string `json:"attributes"`
}

func Startpoint(ctx context.Context, m PubSubMessage) error {

	db, err := firestore.New()

	if err != nil {
		return fmt.Errorf("[ Startpoint | Firestore | New ] err: %v", err)
	}

	err = db.Save(m.Data, m.Attributes)

	if err != nil {
		return fmt.Errorf("[ Startpoint | Firestore | Save ] err: %v", err)
	}

	log.Println("Successfully saved!")

	return nil

}
