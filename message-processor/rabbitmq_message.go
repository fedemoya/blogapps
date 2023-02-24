package main

import (
    "fmt"
    "github.com/rabbitmq/amqp091-go"
)

type RabbitMQMessage struct {
    delivery amqp091.Delivery
}

func NewRabbitMQMessage(delivery amqp091.Delivery) *RabbitMQMessage {
    return &RabbitMQMessage{delivery: delivery}
}

func (r *RabbitMQMessage) Payload() []byte {
    return r.delivery.Body
}

func (r *RabbitMQMessage) Ack() error {
    err := r.delivery.Ack(false)
    if err != nil {
        return fmt.Errorf("failed acking rabbitmq message: %w", err)
    }
    return nil
}

func (r *RabbitMQMessage) Nack() error {
    err := r.delivery.Nack(false, false)
    if err != nil {
        return fmt.Errorf("failed nacking rabbitmq message: %w", err)
    }
    return nil
}
