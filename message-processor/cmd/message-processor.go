package main

import (
    "context"
    "log"
    "message-processor/internal/messages"
    "message-processor/internal/payments"
    "os/signal"
    "syscall"
)

func main() {

    mainCtx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

    messageConsumer := messages.NewRabbitMQMessageConsumer()
    paymentsEventDeserializer := payments.NewPaymentsEventDeserializer()
    handlersFactory := payments.NewHandlersFactoryImpl()
    paymentsEventVisitor := payments.NewPaymentsEventVisitorImpl(handlersFactory)

    processor := messages.NewMessageProcessor[payments.EventsVisitor](messageConsumer, paymentsEventDeserializer, paymentsEventVisitor)

    err := processor.Start()
    if err != nil {
        log.Fatalf("failed to start message processor: %s", err.Error())
    }

    log.Printf("message processor started")

    <-mainCtx.Done()

    log.Printf("message processor stopped")
}
