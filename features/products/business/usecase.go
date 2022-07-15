package business

import (
	"errors"
	"lami/app/features/products"
)

type productUseCase struct {
	productData product.Data
}

// SelectRating implements product.Business
func (uc *productUseCase) SelectRating(idProduct int) (product.CoreRating, error) {
	resp, err := uc.productData.SelectDataRating(idProduct)
	return resp, err
}

// SelectMyProduct implements product.Business
func (uc *productUseCase) SelectMyProduct(idUser int) ([]product.Core, error) {
	resp, err := uc.productData.SelectDataMyProduct(idUser)
	return resp, err
}

// SelectProductbyIDProduct implements product.Business
func (uc *productUseCase) SelectProductbyIDProduct(idProduct int) (product.Core, error) {
	resp, err := uc.productData.SelectDataProductbyIDProduct(idProduct)
	return resp, err
}

// DeleteProduct implements product.Business
func (uc *productUseCase) DeleteProduct(idProduct int, idUser int) error {
	err := uc.productData.DeleteDataProduct(idProduct, idUser)
	return err
}

// UpdateProduct implements product.Business
func (uc *productUseCase) UpdateProduct(dataReq product.Core, idProduct, idUser int) error {
	if dataReq.Price == 0 || dataReq.Stock == 0 || dataReq.Detail == "" {
		return errors.New("all data must be filled")
	}

	err := uc.productData.UpdateDataProduct(dataReq, idProduct, idUser)
	if err != nil {
		return errors.New("failed to insert data product")
	}

	return nil
}

// AddProduct implements product.Business
func (uc *productUseCase) AddProduct(dataReq product.Core) (int, error) {
	if dataReq.Name == "" || dataReq.Price == 0 || dataReq.Stock == 0 || dataReq.Detail == "" {
		return -2, errors.New("all data must be filled")
	}

	res, err := uc.productData.AddDataProduct(dataReq)
	if err != nil {
		return -1, errors.New("failed to insert data product")
	}

	return res, nil
}

// AddProductRating implements product.Business
func (uc *productUseCase) AddProductRating(dataReq product.CoreRating) (int, error) {
	res, err := uc.productData.AddProductDataRating(dataReq)
	if err != nil {
		return -1, errors.New("failed to insert data rating product")
	}

	return res, nil
}

func NewProductBusiness(prdData product.Data) product.Business {
	return &productUseCase{
		productData: prdData,
	}
}
