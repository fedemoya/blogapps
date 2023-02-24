package main

type PaymentsEventVisitorImpl struct {
    withdrawalCreatedHandler WithdrawalCreatedHandler
}

func NewPaymentsEventVisitorImpl(handlersFactory HandlersFactory) *PaymentsEventVisitorImpl {
    return &PaymentsEventVisitorImpl{
        withdrawalCreatedHandler: handlersFactory.CreateWithdrawalCreatedHandler(),
    }
}

func (p PaymentsEventVisitorImpl) VisitWithdrawalCreated(withdrawalCreated WithdrawalCreated) error {
    return p.withdrawalCreatedHandler.HandleWithdrawalCreated(withdrawalCreated)
}
