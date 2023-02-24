package main

import (
    "context"
    "log"
    "os/signal"
    "syscall"
)

func main() {

    mainCtx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

    messageConsumer := NewRabbitMQMessageConsumer()
    paymentsEventDeserializer := NewPaymentsEventDeserializer()
    handlersFactory := NewHandlersFactoryImpl()
    paymentsEventVisitor := NewPaymentsEventVisitorImpl(handlersFactory)

    processor := NewMessageProcessor[PaymentsEventVisitor](messageConsumer, paymentsEventDeserializer, paymentsEventVisitor)

    err := processor.Start()
    if err != nil {
        log.Fatalf("failed to start message processor: %s", err.Error())
    }

    log.Printf("message processor started")

    <-mainCtx.Done()

    log.Printf("message processor stopped")
}
