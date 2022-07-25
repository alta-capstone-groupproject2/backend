package business

import (
	"errors"
	"fmt"
	"strconv"

	"lami/app/features/orders"
	"lami/app/features/orders/presentation/request"
	"lami/app/helper"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
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

	userData, _ := uc.orderData.SelectUser(dataReq.UserID)

	detailMal := fmt.Sprintf("Total Order : Rp. %d", total)

	helper.SendGmailNotify(userData.Email, "Order Merchandise", detailMal)

	return 0, nil

}

// BankCore implements orders.Business
func (uc *orderUseCase) RequestChargeBank(transfer coreapi.ChargeReq, typename string) (coreapi.ChargeReq, error) {
	var transferCore coreapi.ChargeReq
	switch {
	case typename == "permata":
		transferCore = request.ToCoreMidtransPermata(transfer)
	case typename == "mandiri":
		transferCore = request.ToCoreMidtransMandiri(transfer)
	default:
		transferCore = request.ToCoreMidtransBank(transfer)
	}

	return transferCore, nil
}

// TypeBank implements orders.Business
func (uc *orderUseCase) TypeBank(grossamount int64, typename string, idOrder int) (coreapi.ChargeReq, string, error) {
	var Transfer coreapi.ChargeReq
	switch {
	case typename == "bni":
		Transfer = coreapi.ChargeReq{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder),
				GrossAmt: int64(grossamount),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBni,
			},
		}
	case typename == "bca":
		Transfer = coreapi.ChargeReq{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder),
				GrossAmt: int64(grossamount),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBca,
			},
		}
	case typename == "bri":
		Transfer = coreapi.ChargeReq{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder),
				GrossAmt: int64(grossamount),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBri,
			},
		}
	case typename == "permata":
		Transfer = coreapi.ChargeReq{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder),
				GrossAmt: int64(grossamount),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankPermata,
				Permata: &coreapi.PermataBankTransferDetail{
					RecipientName: "lamiapp",
				},
			},
		}
	case typename == "mandiri":
		Transfer = coreapi.ChargeReq{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder),
				GrossAmt: int64(grossamount),
			},
			EChannel: &coreapi.EChannelDetail{
				BillInfo1: "BillInfo1",
				BillInfo2: "BillInfo2",
			},
		}
	default:
		return Transfer, "", errors.New("failed")
	}

	return Transfer, Transfer.TransactionDetails.OrderID, nil
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

// UpdateStatusPayments implements orders.Business
func (uc *orderUseCase) UpdateStatus(idOrder int) error {
	if idOrder == 0 {
		return errors.New("failed")
	}

	err := uc.orderData.UpdateDataStatus(idOrder)
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
