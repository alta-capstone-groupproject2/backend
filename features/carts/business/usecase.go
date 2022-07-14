package business

import (
	"errors"
	"lami/app/features/carts"
)

type cartUseCase struct {
	cartData cart.Data
}

// DeleteCart implements cart.Business
func (uc *cartUseCase) DeleteCart(idCart, iduser int) error {
	err := uc.cartData.DeleteDataCart(idCart, iduser)
	return err
}

// UpdateCart implements cart.Business
func (uc *cartUseCase) UpdateCart(dataReq cart.Core, idCart int) error {
	if dataReq.Qty == 0 {
		return errors.New("failed")
	}

	err := uc.cartData.UpdateDataCart(dataReq, idCart)
	if err != nil {
		return errors.New("failed to update cart")
	}

	return err
}

// SelectCart implements cart.Business
func (uc *cartUseCase) SelectCart(idUser int) ([]cart.Core, error) {
	resp, err := uc.cartData.SelectDataCart(idUser)
	return resp, err
}

// AddCart implements cart.Business
func (uc *cartUseCase) AddCart(dataReq cart.Core) (int, error) {
	// dataReq.Qty = 1
	res, err := uc.cartData.AddDataCart(dataReq)
	if err != nil {
		return -1, errors.New("failed to insert data cart")
	}

	return res, nil
}

func NewCartBusiness(crtData cart.Data) cart.Business {
	return &cartUseCase{
		cartData: crtData,
	}
}
