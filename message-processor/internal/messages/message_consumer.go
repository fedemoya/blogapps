package messages

import (
    "log"
    "time"
)

type MessageConsumer interface {
    Consume() (<-chan Message, error)
    Close() error
}

type DummyMessageConsumer struct {
    msgCh chan Message
}

func NewDummyMessageConsumer() *DummyMessageConsumer {
    return &DummyMessageConsumer{
        msgCh: make(chan Message),
    }
}

func (mc *DummyMessageConsumer) SendMessage(message Message) {
    mc.msgCh <- message
}

func (mc *DummyMessageConsumer) Consume() (<-chan Message, error) {
    return mc.msgCh, nil
}

func (mc *DummyMessageConsumer) Close() error {
    log.Printf("stopping dummy message consumer (wait 2 secs)...")
    close(mc.msgCh)

    // just to simulate a slow close method
    time.Sleep(2 * time.Second)

    return nil
}
