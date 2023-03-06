package dinopay

type EventVisitorImpl struct {
    paymentCreateddHandler PaymentCreatedHandler
}

func NewEventVisitorImpl(handlersFactory HandlersFactory) *EventVisitorImpl {
    return &EventVisitorImpl{
        paymentCreateddHandler: handlersFactory.CreatePaymentCreatedHandler(),
    }
}

func (p EventVisitorImpl) VisitPaymentCreated(paymentCreatedd PaymentCreated) error {
    return p.paymentCreateddHandler.HandlePaymentCreated(paymentCreatedd)
}
