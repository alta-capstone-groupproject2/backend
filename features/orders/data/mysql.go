package data

import (
	"errors"
	_datacart "lami/app/features/carts/data"
	"lami/app/features/orders"
	"lami/app/features/products/data"
	_mUser "lami/app/features/users/data"

	"gorm.io/gorm"
)

type mysqlOrderRepository struct {
	db *gorm.DB
}

// SelectDataHistoryOrder implements orders.Data
func (repo *mysqlOrderRepository) SelectDataHistoryOrder(idUser int) ([]orders.Core, error) {

	dataOrder := []Order{}
	dataProduct := data.Product{}

	res := repo.db.Preload("OrderDetail").Where("user_id = ?", idUser).Find(&dataOrder)
	if res.Error != nil {
		return []orders.Core{}, res.Error
	}
	length := ToCoreList(dataOrder)

	for i := 0; i < len(length); i++ {

		for j := 0; j < len(dataOrder[i].OrderDetail); j++ {

			rowsproduct, errproduct := repo.db.Model(&dataProduct).Where("id = ?", dataOrder[i].OrderDetail[j].ProductID).Select("id, name, url").Rows()
			if errproduct != nil {
				return []orders.Core{}, errors.New(errproduct.Error())
			}

			for rowsproduct.Next() {
				var dataID int
				var dataName, dataURL string
				if errrowsproduct := rowsproduct.Scan(&dataID, &dataName, &dataURL); errrowsproduct != nil {
					panic(errrowsproduct)
				}
				dataOrder[i].OrderDetail[j].Product.ID = uint(dataID)
				dataOrder[i].OrderDetail[j].Product.Name = dataName
				dataOrder[i].OrderDetail[j].Product.URL = dataURL
			}

			rowsqty, errrowsqty := repo.db.Model(&OrderDetail{}).Where("product_id = ?", dataOrder[i].OrderDetail[j].ProductID).Select("qty").Rows()
			if errrowsqty != nil {
				return []orders.Core{}, errors.New(errrowsqty.Error())
			}

			for rowsqty.Next() {
				var dataQty int
				if errqty := rowsqty.Scan(&dataQty); errqty != nil {
					panic(errqty)
				}
				dataOrder[i].OrderDetail[j].Qty = uint(dataQty)
			}

		}

	}
	return ToCoreList(dataOrder), nil
}

// UpdateStockOnProductPlusCountTotalPrice implements orders.Data
func (repo *mysqlOrderRepository) UpdateStockOnProductPlusCountTotalPrice(dataReq orders.Core, idUser int) (int, error) {
	var totalprice int
	dataCart := []_datacart.Cart{}
	dataProduct := data.Product{}

	checkQty := repo.db.Preload("Product").Where("user_id = ?", idUser).Find(&dataCart)
	if checkQty.RowsAffected >= 1 {
		for i := 0; i < int(checkQty.RowsAffected); i++ {
			if dataCart[i].Product.Stock < uint(dataCart[i].Qty) {
				return totalprice, errors.New("there is one product with a quantity that exceeds stock")
			}
		}
	}

	for i := 0; i < len(dataReq.CartID); i++ {
		dataProduct.Stock = dataCart[i].Product.Stock - uint(dataCart[i].Qty)
		errUpdateStock := repo.db.Model(dataProduct).Where("id = ?", dataCart[i].Product.ID).Updates(dataProduct)
		if errUpdateStock.Error != nil {
			return totalprice, errUpdateStock.Error
		}

		totalprice += (int(dataCart[i].Product.Price) * dataCart[i].Qty)

	}
	return totalprice, nil
}

// DeleteDataCart implements orders.Data
func (repo *mysqlOrderRepository) DeleteDataCart(dataReq orders.Core, idUser int) error {
	dataCart := []_datacart.Cart{}

	checkQty := repo.db.Preload("Product").Where("user_id = ?", idUser).Find(&dataCart)
	if checkQty.RowsAffected >= 1 {
		for i := 0; i < int(checkQty.RowsAffected); i++ {
			if dataCart[i].Product.Stock < uint(dataCart[i].Qty) {
				return errors.New("there is one product with a quantity that exceeds stock")
			}
		}
	}

	for i := 0; i < len(dataReq.CartID); i++ {
		err := repo.db.Where("user_id = ?", idUser).Delete(&dataCart, dataCart[i].ID)
		if err.Error != nil {
			return err.Error
		}
	}
	return nil
}

// AddDataOrderPlusCountRows implements orders.Data
func (repo *mysqlOrderRepository) AddDataOrder(dataReq orders.Core, idUser int, total int) (int64, error) {
	dataCart := []_datacart.Cart{}

	checkQty := repo.db.Preload("Product").Where("user_id = ?", idUser).Find(&dataCart)
	if checkQty.RowsAffected >= 1 {
		for i := 0; i < int(checkQty.RowsAffected); i++ {
			if dataCart[i].Product.Stock < uint(dataCart[i].Qty) {
				return 0, errors.New("there is one product with a quantity that exceeds stock")
			}
		}
	}

	dataReq.Status = "Pending"
	dataReq.TotalPrice = uint(total)

	dataOrder := fromCore(dataReq)
	res := repo.db.Create(&dataOrder)
	if res.Error != nil {
		return 0, res.Error
	}

	var count int64
	countOrder := repo.db.Model(&Order{}).Count(&count)
	if countOrder.Error != nil {
		return 0, countOrder.Error
	}

	return count, nil
}

// AddDataOrderDetail implements orders.Data
func (repo *mysqlOrderRepository) AddDataOrderDetail(dataReq orders.Core, row int64, idUser int) error {
	dataCart := []_datacart.Cart{}
	dataReqOrderDetail := orders.CoreDetail{}

	checkQty := repo.db.Preload("Product").Where("user_id = ?", idUser).Find(&dataCart)
	if checkQty.RowsAffected >= 1 {
		for i := 0; i < int(checkQty.RowsAffected); i++ {
			if dataCart[i].Product.Stock < uint(dataCart[i].Qty) {
				return errors.New("there is one product with a quantity that exceeds stock")
			}
		}
	}

	for i := 0; i < len(dataReq.CartID); i++ {

		dataReqOrderDetail.OrderID = int(row)

		dataReqOrderDetail.ProductID = int(dataCart[i].Product.ID)
		dataReqOrderDetail.Qty = uint(dataCart[i].Qty)

		dataOrderDetail := fromCoreDetail(dataReqOrderDetail)
		resOrderDetail := repo.db.Create(&dataOrderDetail)
		if resOrderDetail.Error != nil {
			return resOrderDetail.Error
		}

	}

	return nil
}

func (repo *mysqlOrderRepository) DataPaymentsGrossAmount(idUser int) (int, error) {
	var totalprice int
	err := repo.db.Raw("SELECT total_price FROM orders WHERE user_id = ? ORDER BY id DESC LIMIT 1", idUser).Scan(&totalprice)
	if err.Error != nil {
		return -1, err.Error
	}

	return totalprice, nil
}

func (repo *mysqlOrderRepository) DataPaymentsOrderID(idUser int) (int, error) {
	var idOrder int
	err := repo.db.Raw("SELECT id FROM orders WHERE user_id = ? ORDER BY id DESC LIMIT 1", idUser).Scan(&idOrder)
	if err.Error != nil {
		return -1, err.Error
	}

	return idOrder, nil
}

func NewOrderRepository(conn *gorm.DB) orders.Data {
	return &mysqlOrderRepository{
		db: conn,
	}
}

func (repo *mysqlOrderRepository) SelectUser(id int) (response _mUser.User, err error) {
	datauser := _mUser.User{}
	result := repo.db.Preload("Role").Find(&datauser, id)
	if result.Error != nil {
		return _mUser.User{}, result.Error
	}
	if result.RowsAffected == 0 {
		return _mUser.User{}, errors.New("user not found")
	}
	return datauser, nil
}
