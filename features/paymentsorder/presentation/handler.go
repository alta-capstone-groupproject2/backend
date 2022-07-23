package presentation

import (
	"fmt"
	"lami/app/config"
	"lami/app/features/paymentsorder"
	"lami/app/features/paymentsorder/presentation/request"
	"lami/app/features/paymentsorder/presentation/response"
	"lami/app/helper"
	"lami/app/middlewares"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentsHandler struct {
	paymentBusiness paymentsorder.Business
}

func NewPaymentHandler(business paymentsorder.Business) *PaymentsHandler {
	return &PaymentsHandler{
		paymentBusiness: business,
	}
}

var order coreapi.Client

func (h *PaymentsHandler) PostPayment(c echo.Context) error {
	midtrans.ServerKey = config.MidtransOrderServerKey()
	order.New(midtrans.ServerKey, midtrans.Sandbox)
	typeName := c.Param("type")

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	idOrder, totalprice, errOrder := h.paymentBusiness.Payments(userID_token)
	if errOrder != nil {
		return c.JSON(http.StatusInternalServerError, "failed h.paymentBusiness")
	}

	var Transfer request.ChargeRequest
	switch {
	case typeName == "bni":
		Transfer = request.ChargeRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder),
				GrossAmt: int64(totalprice),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBni,
			},
		}
	case typeName == "bca":
		Transfer = request.ChargeRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder),
				GrossAmt: int64(totalprice),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBca,
			},
		}
	case typeName == "bri":
		Transfer = request.ChargeRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder) + "bri",
				GrossAmt: int64(totalprice),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBri,
			},
		}
	case typeName == "permata":
		Transfer = request.ChargeRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder) + "permata1",
				GrossAmt: int64(totalprice),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankPermata,
				Permata: &coreapi.PermataBankTransferDetail{
					RecipientName: "lamiapp",
				},
			},
		}
	case typeName == "mandiri":
		Transfer = request.ChargeRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder),
				GrossAmt: int64(totalprice),
			},
			EChannelDetails: coreapi.EChannelDetail{
				BillInfo1: "BillInfo1",
				BillInfo2: "BillInfo2",
			},
		}
	default:
		return c.JSON(http.StatusInternalServerError, "failed param")
	}

	var TransferCore coreapi.ChargeReq
	if typeName == "bni" || typeName == "bca" || typeName == "bri" {
		TransferCore = request.ToCoreMidtransBank(Transfer)
	} else if typeName == "permata" {
		TransferCore = request.ToCoreMidtransPermata(Transfer)
	} else if typeName == "mandiri" {
		TransferCore = request.ToCoreMidtransMandiri(Transfer)
	} else {
		TransferCore = request.ToCoreMidtransEMoney(Transfer)
	}

	resp, err := coreapi.ChargeTransaction(&TransferCore)
	if err != nil {
		fmt.Println("Error coreapi api, with global config", err.GetMessage())
	}

	if typeName == "bni" || typeName == "bca" || typeName == "bri" || typeName == "permata" {
		return c.JSON(http.StatusOK, response.FromCoreChargeMidtrans(*resp))
	} else if typeName == "permata" {
		return c.JSON(http.StatusOK, response.FromCoreChargePermata(*resp))
	} else if typeName == "mandiri" {
		return c.JSON(http.StatusOK, response.FromCoreChargeMandiri(*resp))
	} else {
		return c.JSON(http.StatusOK, response.FromCoreChargeMidtrans(*resp))
	}

}

func (h *PaymentsHandler) PutPayment(c echo.Context) error {
	panic("unimplemented")
}

func Random() string {
	time.Sleep(500 * time.Millisecond)
	return strconv.FormatInt(time.Now().Unix(), 10)
}
