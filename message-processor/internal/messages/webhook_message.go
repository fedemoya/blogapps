package messages

type WebhookMessage struct {
    payload []byte
}

func NewWebhookMessage(payload []byte) *WebhookMessage {
    return &WebhookMessage{
        payload: payload,
    }
}

func (w *WebhookMessage) Payload() []byte {
    return w.payload
}

func (w *WebhookMessage) Ack() error {
    //TODO implement me
    return nil
}

func (w *WebhookMessage) Nack() error {
    //TODO implement me
    return nil
}
