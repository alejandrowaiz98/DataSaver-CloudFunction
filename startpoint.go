package p

import "context"

type PubSubMessage struct {
	Data       []byte            `json:"data"`
	Attributes map[string]string `json:"attributes"`
}

func Startpoint(ctx context.Context, m PubSubMessage) error {

	return nil

}
