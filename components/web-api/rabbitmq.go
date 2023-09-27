package main

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

const InquiriesQueue = "inquiries"

func publishInquiry(ch *amqp.Channel, inquiry Inquiry) error {
	inquiryJSON, err := json.Marshal(inquiry)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(context.Background(), "", InquiriesQueue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        inquiryJSON,
	})
	return err
}
