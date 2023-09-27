package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prodigeris/Flight-searcher-go/common"
	amqp "github.com/rabbitmq/amqp091-go"
)

func publishInquiry(ch *amqp.Channel, inquiry Inquiry) error {
	inquiryJSON, err := json.Marshal(inquiry)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(context.Background(), "", common.InquiriesQueue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        inquiryJSON,
	})

	fmt.Println("Published a message", inquiry)
	return err
}
