package dinopay

import (
    "encoding/json"
    "fmt"
    "log"
    "message-processor/internal/events"
)

type EventDeserializer struct {
}

func (e EventDeserializer) Deserialize(rawEvent []byte) (events.Event[EventsVisitor], error) {
    var event events.EventEnvelope
    err := json.Unmarshal(rawEvent, &event)
    if err != nil {
        return nil, fmt.Errorf("error processing raw event %s: %w", rawEvent, err)
    }
    switch event.Type {
    case "paymentCreated":
        var withdrawalCreated PaymentCreated
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

func NewEventDeserializer() *EventDeserializer {
    return &EventDeserializer{}
}
