package data

import (
	"errors"
	"fmt"
	_datacart "lami/app/features/carts/data"
	"lami/app/features/orders"
	"lami/app/features/products/data"

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
	result := ToCoreList(dataOrder)

	for i := 0; i < len(result); i++ {

		dataOrderDetail := []OrderDetail{}

		// idProduct := []data.Product{}
		for j := 0; j < len(dataOrder[i].OrderDetail); j++ {
			fmt.Println("dataOrder[i].OrderDetail[j].ProductID", dataOrder[i].OrderDetail[j].ProductID)
			res3 := repo.db.Find(&dataProduct).Where("id = ?", dataOrder[i].OrderDetail[j].ProductID)
			if res3.Error != nil {
				return []orders.Core{}, res3.Error
			}
			dataOrderDetail2 := OrderDetail{
				ProductID: int(dataProduct.ID),
				Name:      dataProduct.Name,
				URL:       dataProduct.URL,
			}
			// idProduct = append(idProduct, dataProduct)
			dataOrderDetail = append(dataOrderDetail, dataOrderDetail2)

			fmt.Println("dataOrderDetail ", i, " : ", dataOrderDetail)

			result[i].Product = ToCoreDetailList(dataOrderDetail)
		}

		fmt.Println("result[", i, "]", " : ", result[i].Product)
	}

	return result, nil
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
