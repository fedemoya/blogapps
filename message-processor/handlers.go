package main

import (
    "log"
)

type WithdrawalCreatedHandler interface {
    HandleWithdrawalCreated(withdrawalCreated WithdrawalCreated) error
}

type WithdrawalCreatedHandlerImpl struct {
}

func (h *WithdrawalCreatedHandlerImpl) HandleWithdrawalCreated(withdrawalCreated WithdrawalCreated) error {
    log.Printf("handling WithdrawalCreated event: %+v", withdrawalCreated)
    return nil
}
