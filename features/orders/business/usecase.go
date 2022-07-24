package business

import (
	"errors"
	"lami/app/features/orders"
)

type orderUseCase struct {
	orderData orders.Data
}

// SelectHistoryOrder implements orders.Business
func (uc *orderUseCase) SelectHistoryOrder(idUser int) ([]orders.Core, error) {
	resp, err := uc.orderData.SelectDataHistoryOrder(idUser)
	return resp, err
}

// AddOrder implements orders.Business
func (uc *orderUseCase) Order(dataReq orders.Core, idUser int) (int, error) {

	//	Check length []cartID
	if len(dataReq.CartID) == 0 {
		return -1, errors.New("failed")
	}

	//	Update stock on product plus count total price
	total, err := uc.orderData.UpdateStockOnProductPlusCountTotalPrice(dataReq, idUser)
	if err != nil {
		return -1, errors.New("failed")
	}

	//	Add data order plus count rows
	rows, err2 := uc.orderData.AddDataOrder(dataReq, idUser, total)
	if err2 != nil {
		return -1, errors.New("failed")
	}

	//	Add data order detail
	err3 := uc.orderData.AddDataOrderDetail(dataReq, rows, idUser)
	if err3 != nil {
		return -1, errors.New("failed")
	}

	//	Delete data on cart database
	err4 := uc.orderData.DeleteDataCart(dataReq, idUser)
	if err4 != nil {
		return -1, errors.New("failed")
	}

	return 0, nil

}

// PaymentGrossAmount implements paymentsorder.Business
func (uc *orderUseCase) PaymentGrossAmount(idUser int) (int, error) {
	if idUser == 0 {
		return -1, errors.New("failed get idUser in usecase")
	}

	grossamount, res := uc.orderData.DataPaymentsGrossAmount(idUser)
	if res != nil {
		return -1, errors.New("failed to get gross amount")
	}

	return grossamount, nil
}

// PaymentsOrderID implements paymentsorder.Business
func (uc *orderUseCase) PaymentsOrderID(idUser int) (int, error) {
	if idUser == 0 {
		return -1, errors.New("failed get idUser in usecase")
	}

	grossamount, res := uc.orderData.DataPaymentsOrderID(idUser)
	if res != nil {
		return -1, errors.New("failed to get gross amount")
	}

	return grossamount, nil
}

func NewOrderBusiness(odrData orders.Data) orders.Business {
	return &orderUseCase{
		orderData: odrData,
	}
}
