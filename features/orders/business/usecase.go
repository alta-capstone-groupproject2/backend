package business

import (
	"errors"
	"lami/app/features/orders"
)

type orderUseCase struct {
	orderData orders.Data
}

// SelectHistoryOrder implements orders.Business
func (uc *orderUseCase) SelectHistoryOrder(idOrder, idUser int) (orders.CoreDetail, error) {
	resp, err := uc.orderData.SelectDataHistoryOrder(idOrder, idUser)
	return resp, err
}

// AddOrder implements orders.Business
func (uc *orderUseCase) AddOrder(dataReq orders.Core, idUser int) error {
	//	Check length []cartID
	if len(dataReq.CartID) == 0 {
		return errors.New("empty cartID")
	}

	//	Check input receiver, phonenumber, address
	if dataReq.Receiver == "" || dataReq.PhoneNumber == "" || dataReq.Address == "" {
		return errors.New("all data must be filedd")
	}

	err := uc.orderData.AddDataOrder(dataReq, idUser)
	if err != nil {
		return errors.New("failed")
	}

	return nil

}

func NewOrderBusiness(odrData orders.Data) orders.Business {
	return &orderUseCase{
		orderData: odrData,
	}
}
