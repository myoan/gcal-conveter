package app

import (
	"context"
	"fmt"
)

func Hello() {
	fmt.Println("Hello gcal-converter!!")
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	Hello()
	return nil
}
