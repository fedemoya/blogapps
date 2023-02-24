package main

import (
    "encoding/json"
    "errors"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
    "sync"
    "testing"
)

func TestMessageProcessor(t *testing.T) {

    withdrawalCreated := WithdrawalCreated{
        WithdrawalId:       "some-unique-id",
        Amount:             100,
        SourceAccount:      "source account data",
        DestinationAccount: "destination account data",
    }

    withdrawalCreatedJSON, err := json.Marshal(withdrawalCreated)
    require.NoError(t, err)

    event := EventEnvelope{
        Type: "withdrawal.created",
        Data: withdrawalCreatedJSON,
    }

    eventJSON, err := json.Marshal(event)
    require.NoError(t, err)

    t.Run("handler returns no error, message is acked", func(t *testing.T) {

        messagesCh := make(chan Message)
        messageConsumerMock := &MockMessageConsumer{}
        messageConsumerMock.On("Consume").Return((<-chan Message)(messagesCh), nil)
        messageConsumerMock.On("Close").Return(nil)

        wg := sync.WaitGroup{}

        messageMock := &MockMessage{}
        messageMock.On("Payload").Return(eventJSON)
        messageMock.On("Ack").Return(nil).Run(func(args mock.Arguments) {
            wg.Done()
        })

        paymentsEventDeserializerMock := &MockEventsDeserializer[PaymentsEventVisitor]{}
        paymentsEventDeserializerMock.On("Deserialize", messageMock).Return(withdrawalCreated, nil)

        paymentsEventVisitorMock := &MockPaymentsEventVisitor{}
        paymentsEventVisitorMock.On("VisitWithdrawalCreated", withdrawalCreated).Return(nil)

        processor := NewMessageProcessor[PaymentsEventVisitor](
            messageConsumerMock,
            paymentsEventDeserializerMock,
            paymentsEventVisitorMock,
        )

        wg.Add(1)

        processor.Start()

        messagesCh <- messageMock

        wg.Wait()
    })

    t.Run("handler returns error, message is nacked", func(t *testing.T) {

        messagesCh := make(chan Message)
        messageConsumerMock := &MockMessageConsumer{}
        messageConsumerMock.On("Consume").Return((<-chan Message)(messagesCh), nil)
        messageConsumerMock.On("Close").Return(nil)

        wg := sync.WaitGroup{}

        messageMock := &MockMessage{}
        messageMock.On("Payload").Return(eventJSON)
        messageMock.On("Nack").Return(nil).Run(func(args mock.Arguments) {
            wg.Done()
        })

        paymentsEventDeserializerMock := &MockEventsDeserializer[PaymentsEventVisitor]{}
        paymentsEventDeserializerMock.On("Deserialize", messageMock).Return(withdrawalCreated, nil)

        paymentsEventVisitorMock := &MockPaymentsEventVisitor{}
        paymentsEventVisitorMock.On("VisitWithdrawalCreated", withdrawalCreated).Return(errors.New("boom"))

        processor := NewMessageProcessor[PaymentsEventVisitor](
            messageConsumerMock,
            paymentsEventDeserializerMock,
            paymentsEventVisitorMock,
        )

        wg.Add(1)

        processor.Start()

        messagesCh <- messageMock

        wg.Wait()
    })
}
