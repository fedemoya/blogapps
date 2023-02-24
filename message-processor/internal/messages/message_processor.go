package messages

import (
    "fmt"
    "log"
    "message-processor/internal/events"
)

type MessageProcessor[EventsVisitor any] struct {
    messageConsumer    MessageConsumer
    eventsDeserializer events.Deserializer[EventsVisitor]
    eventsVisitor      EventsVisitor
}

func NewMessageProcessor[EventsVisitor any](
    messageConsumer MessageConsumer,
    eventsDeserializer events.Deserializer[EventsVisitor],
    eventsVisitor EventsVisitor) *MessageProcessor[EventsVisitor] {
    return &MessageProcessor[EventsVisitor]{
        messageConsumer:    messageConsumer,
        eventsDeserializer: eventsDeserializer,
        eventsVisitor:      eventsVisitor,
    }
}

func (p *MessageProcessor[EV]) Start() error {
    msgCh, err := p.startMessageConsumer()
    if err != nil {
        return fmt.Errorf("failed starting message consumer: %w", err)
    }
    go p.processMsgs(msgCh)
    return nil
}

func (p *MessageProcessor[EV]) startMessageConsumer() (<-chan Message, error) {
    msgCh, err := p.messageConsumer.Consume()
    if err != nil {
        return nil, fmt.Errorf("failed consuming from message consumer: %w", err)
    }
    return msgCh, nil
}

func (p *MessageProcessor[EV]) processMsgs(ch <-chan Message) {
    for msg := range ch {
        go p.processMsg(msg)
    }
    // TODO handle this case. Probably the Start method should be blocking.
    log.Printf("messages channel closed")
}

func (p *MessageProcessor[EV]) processMsg(msg Message) {

    log.Printf("processing message with payload %s", msg.Payload())

    event, err := p.eventsDeserializer.Deserialize(msg.Payload())
    if err != nil {
        log.Printf("failed deserializing msg: %s", err.Error())
        p.nackMsg(msg)
        return
    }

    err = event.Accept(p.eventsVisitor)
    if err != nil {
        log.Printf("failed processing event: %s", err.Error())
        p.nackMsg(msg)
        return
    }
    p.ackMsg(msg)
}

func (p *MessageProcessor[EV]) ackMsg(msg Message) {
    err := msg.Ack()
    if err != nil {
        log.Printf("failed acking message: %s", err.Error())
    }
}

func (p *MessageProcessor[EventsVisitor]) nackMsg(msg Message) {
    err := msg.Nack()
    if err != nil {
        log.Printf("failed nacking message: %s", err.Error())
    }
}
