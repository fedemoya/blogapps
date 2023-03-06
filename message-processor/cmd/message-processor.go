package main

import (
    "context"
    "log"
    "message-processor/internal/dinopay"
    "message-processor/internal/messages"
    "message-processor/internal/payments"
    "os/signal"
    "syscall"
)

func main() {

    mainCtx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

    rabbitMQMessageConsumer := messages.NewRabbitMQMessageConsumer()
    paymentsEventDeserializer := payments.NewEventDeserializer()
    paymentsEventsHandlersFactory := payments.NewHandlersFactoryImpl()
    paymentsEventsVisitor := payments.NewPaymentsEventVisitorImpl(paymentsEventsHandlersFactory)

    paymentMessagesProcessor := messages.NewMessageProcessor[payments.EventsVisitor](rabbitMQMessageConsumer, paymentsEventDeserializer, paymentsEventsVisitor)

    err := paymentMessagesProcessor.Start()
    if err != nil {
        log.Fatalf("failed to start message paymentMessagesProcessor: %s", err.Error())
    }

    log.Printf("message paymentMessagesProcessor started")

    webhookMessageConsumer := messages.NewWebhookMessageConsumer()
    dinoPayEventDeserializer := dinopay.NewEventDeserializer()
    dinopayEventsHandlersFactory := dinopay.NewHandlersFactoryImpl()
    dinopayEventsVisitor := dinopay.NewEventVisitorImpl(dinopayEventsHandlersFactory)

    dinopayMessagesProcessor := messages.NewMessageProcessor[dinopay.EventsVisitor](webhookMessageConsumer, dinoPayEventDeserializer, dinopayEventsVisitor)

    err = dinopayMessagesProcessor.Start()
    if err != nil {
        log.Fatalf("failed to start message dinopayMessagesProcessor: %s", err.Error())
    }

    log.Printf("message dinopayMessagesProcessor started")

    <-mainCtx.Done()
}
