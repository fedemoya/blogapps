package main

type HandlersFactory interface {
    CreateWithdrawalCreatedHandler() WithdrawalCreatedHandler
}

type HandlersFactoryImpl struct {
}

func NewHandlersFactoryImpl() *HandlersFactoryImpl {
    return &HandlersFactoryImpl{}
}

func (f *HandlersFactoryImpl) CreateWithdrawalCreatedHandler() WithdrawalCreatedHandler {
    return &WithdrawalCreatedHandlerImpl{}
}
