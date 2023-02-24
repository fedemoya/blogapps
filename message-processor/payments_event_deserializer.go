package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type PaymentsEventDeserializer struct {
}

func NewPaymentsEventDeserializer() *PaymentsEventDeserializer {
    return &PaymentsEventDeserializer{}
}

func (r *PaymentsEventDeserializer) Deserialize(message Message) (Event[PaymentsEventVisitor], error) {
    var event EventEnvelope
    err := json.Unmarshal(message.Payload(), &event)
    if err != nil {
        return nil, fmt.Errorf("error processing message %s: %w", message.Payload(), err)
    }
    switch event.Type {
    case "withdrawal.created":
        var withdrawalCreated WithdrawalCreated
        err := json.Unmarshal(event.Data, &withdrawalCreated)
        if err != nil {
            return nil, fmt.Errorf("error unmarshalling withdrawal.created event data %s: %s", event.Data, err.Error())
        }
        return withdrawalCreated, nil
    default:
        log.Printf("unexpected event type: %s", event.Type)
        return EmptyEvent{}, nil
    }
}
