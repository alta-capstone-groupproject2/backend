package data

import (
	"errors"
	_dataCart "lami/app/features/carts/data"
	"lami/app/features/orders"
	"lami/app/features/products"
	_dataProduct "lami/app/features/products/data"

	"gorm.io/gorm"
)

type mysqlOrderRepository struct {
	db *gorm.DB
}

// SelectDataHistoryOrder implements orders.Data
func (repo *mysqlOrderRepository) SelectDataHistoryOrder(idUser int) (orders.CoreDetail, error) {
	panic("unimplemented")
}

//	AddDataOrder implements orders.Data
func (repo *mysqlOrderRepository) AddDataOrder(dataReq orders.Core, idUser int) error {

	dataCart := []_dataCart.Cart{}
	dataProductCore := product.Core{}
	var totalPrice uint = 0
	var totalQty int = 0

	chechQty := repo.db.Preload("Product").Where("user_id = ?", idUser).Find(&dataCart)
	if chechQty.RowsAffected >= 1 {
		for i := 0; i < int(chechQty.RowsAffected); i++ {
			if dataCart[i].Product.Stock < uint(dataCart[i].Qty) {
				return errors.New("there is one product with a quantity that exceeds stock")
			}
		}
	}

	//	After successfull checking
	for i := 0; i < len(dataReq.CartID); i++ {

		//	Update stock on product
		dataProductCore.Stock = dataCart[i].Product.Stock - uint(dataCart[i].Qty)

		errUpdateStock := repo.db.Model(_dataProduct.Product{}).Where("id = ?", dataCart[i].Product.ID).Updates(dataProductCore)
		if errUpdateStock.Error != nil {
			return errUpdateStock.Error
		}

		//	Count totalprice
		totalPrice = totalPrice + dataCart[i].Product.Price

		//	Count qty
		totalQty = totalQty + dataCart[i].Qty

		//	Create new data order in database
		dataReq.TotalPrice = totalPrice
		dataReq.Status = "Approved"

		dataOrder := fromCore(dataReq)
		resOrder := repo.db.Create(&dataOrder)
		if resOrder.Error != nil {
			return resOrder.Error
		}

		//	Create new data for order detail in database
		dataReqDetail := orders.CoreDetail{
			OrderID:    int(dataOrder.ID),
			ProductID:  int(dataCart[i].Product.ID),
			TotalPrice: dataReq.TotalPrice,
			Qty:        uint(totalQty),
		}

		dataOrderDetail := fromCoreDetail(dataReqDetail)
		resOrderDetail := repo.db.Create(&dataOrderDetail)
		if resOrderDetail.Error != nil {
			return resOrderDetail.Error
		}

		//	Delete data cart in database
		errDeleteCart := repo.db.Where("user_id = ?", idUser).Delete(&dataCart, dataCart[i].ID)
		if errDeleteCart != nil {
			return errDeleteCart.Error
		}

	}

	return nil
}

func NewOrderRepository(conn *gorm.DB) orders.Data {
	return &mysqlOrderRepository{
		db: conn,
	}
}
