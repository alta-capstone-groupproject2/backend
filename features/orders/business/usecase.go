package business

import "lami/app/features/orders"

type orderUseCase struct {
	orderData orders.Data
}

// AddOrder implements orders.Business
func (uc *orderUseCase) AddOrder(dataReq orders.Core, idUser int) (int, error) {
	panic("unimplemented")
}

func NewOrderBusiness(odrData orders.Data) orders.Business {
	return &orderUseCase{
		orderData: odrData,
	}
}
