package dinopay

type PaymentCreated struct {
    Payment Payment `json:"payment"`
}

type Payment struct {
    Id                    string         `json:"id"`
    Amount                int            `json:"amount"`
    SourceAccount         PaymentAccount `json:"sourceAccount"`
    DestinationAccount    PaymentAccount `json:"destinationAccount"`
    Status                string         `json:"status"`
    CustomerTransactionId string         `json:"customerTransactionId"`
    CreatedAt             int            `json:"createdAt"`
    UpdatedAt             int            `json:"updatedAt"`
}

type PaymentAccount struct {
    AccountHolder string `json:"accountHolder"`
    AccountNumber string `json:"accountNumber"`
}

func (p PaymentCreated) Accept(visitor EventsVisitor) error {
    return visitor.VisitPaymentCreated(p)
}

type EmptyEvent struct {
}

func (EmptyEvent) Accept(_ EventsVisitor) error {
    return nil
}
