package dinopay

type HandlersFactory interface {
    CreatePaymentCreatedHandler() PaymentCreatedHandler
}

type HandlersFactoryImpl struct {
}

func NewHandlersFactoryImpl() *HandlersFactoryImpl {
    return &HandlersFactoryImpl{}
}

func (f *HandlersFactoryImpl) CreatePaymentCreatedHandler() PaymentCreatedHandler {
    return &PaymentCreatedHandlerImpl{}
}
