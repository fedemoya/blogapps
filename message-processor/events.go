package main

type WithdrawalCreated struct {
    WithdrawalId       string  `json:"withdrawal_id"`
    Amount             float64 `json:"amount"`
    SourceAccount      string  `json:"source_account"`
    DestinationAccount string  `json:"destination_account"`
}

func (w WithdrawalCreated) Accept(visitor PaymentsEventVisitor) error {
    return visitor.VisitWithdrawalCreated(w)
}

type EmptyEvent struct {

}

func (EmptyEvent) Accept(_ PaymentsEventVisitor) error  {
    return nil
}