package main

type PaymentsEventVisitor interface {
    VisitWithdrawalCreated(withdrawalCreated WithdrawalCreated) error
}
