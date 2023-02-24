package messages

type Message interface {
    Payload() []byte
    Ack() error
    Nack() error
}
