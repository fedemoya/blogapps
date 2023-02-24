package payments

type WithdrawalCreated struct {
    WithdrawalId       string  `json:"withdrawal_id"`
    Amount             float64 `json:"amount"`
    SourceAccount      string  `json:"source_account"`
    DestinationAccount string  `json:"destination_account"`
}

func (w WithdrawalCreated) Accept(visitor EventsVisitor) error {
    return visitor.VisitWithdrawalCreated(w)
}

type EmptyEvent struct {
}

func (EmptyEvent) Accept(_ EventsVisitor) error {
    return nil
}
