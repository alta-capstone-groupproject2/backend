package business

import (
	"errors"
	"lami/app/features/paymentsorder"
)

type paymentsUseCase struct {
	paymentsData paymentsorder.Data
}

// InsertPayments implements paymentsorder.Business
func (uc *paymentsUseCase) Payments(idUser int) (int, int, error) {
	if idUser == 0 {
		return -1, -1, errors.New("failed get idUser in usecase")
	}

	idOrder, totalprice, res := uc.paymentsData.DataPayments(idUser)
	if res != nil {
		return -1, -1, errors.New("failed to insert data")
	}

	return idOrder, totalprice, nil
}

func NewPaymentsBusiness(pymt paymentsorder.Data) paymentsorder.Business {
	return &paymentsUseCase{
		paymentsData: pymt,
	}
}
