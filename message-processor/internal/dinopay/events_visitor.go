package dinopay

type EventsVisitor interface {
    VisitPaymentCreated(created PaymentCreated) error
}
