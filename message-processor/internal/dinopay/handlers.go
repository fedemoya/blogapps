package dinopay

import (
    "log"
)

type PaymentCreatedHandler interface {
    HandlePaymentCreated(withdrawalCreated PaymentCreated) error
}

type PaymentCreatedHandlerImpl struct {
}

func (h *PaymentCreatedHandlerImpl) HandlePaymentCreated(withdrawalCreated PaymentCreated) error {
    log.Printf("handling PaymentCreated event: %+v", withdrawalCreated)
    return nil
}
