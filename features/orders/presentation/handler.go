package presentation

import (
	"fmt"
	"lami/app/features/orders"
	"lami/app/features/orders/presentation/request"
	"lami/app/features/orders/presentation/response"
	"lami/app/helper"
	"lami/app/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type OrderHandler struct {
	orderBusiness orders.Business
}

func NewOrderHandler(business orders.Business) *OrderHandler {
	return &OrderHandler{
		orderBusiness: business,
	}
}

func (h *OrderHandler) PostOrder(c echo.Context) error {

	typeName := c.Param("type")

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	order := request.Order{}
	err_bind := c.Bind(&order)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("failed to bind data order"))
	}

	orderCore := request.ToCore(order)
	orderCore.UserID = userID_token

	res, err := h.orderBusiness.Order(orderCore, userID_token)
	if err != nil && res == -1 {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	//	Payments
	idOrder, errorderid := h.orderBusiness.PaymentsOrderID(userID_token)
	if errorderid != nil {
		return c.JSON(helper.ResponseInternalServerError("failed get order_id for payments"))
	}

	grossamount, errgrossamount := h.orderBusiness.PaymentGrossAmount(userID_token)
	if errgrossamount != nil {
		return c.JSON(helper.ResponseInternalServerError("failed get gross amount for payments"))
	}

	var Transfer request.ChargeRequest
	switch {
	case typeName == "bni":
		Transfer = request.ChargeRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder),
				GrossAmt: int64(grossamount),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBni,
			},
		}
	case typeName == "bca":
		Transfer = request.ChargeRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder),
				GrossAmt: int64(grossamount),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBca,
			},
		}
	case typeName == "bri":
		Transfer = request.ChargeRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder) + "bri",
				GrossAmt: int64(grossamount),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBri,
			},
		}
	case typeName == "permata":
		Transfer = request.ChargeRequest{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  strconv.Itoa(idOrder) + "permata1",
				GrossAmt: int64(grossamount),
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
				GrossAmt: int64(grossamount),
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
	if typeName == "permata" {
		TransferCore = request.ToCoreMidtransPermata(Transfer)
	} else if typeName == "mandiri" {
		TransferCore = request.ToCoreMidtransMandiri(Transfer)
	} else {
		TransferCore = request.ToCoreMidtransBank(Transfer)
	}

	resp, errcharge := coreapi.ChargeTransaction(&TransferCore)
	if errcharge != nil {
		fmt.Println("Error coreapi api, with global config", errcharge.GetMessage())
	}

	if typeName == "permata" {
		return c.JSON(http.StatusOK, response.FromCoreChargePermata(*resp))
	} else if typeName == "mandiri" {
		return c.JSON(http.StatusOK, response.FromCoreChargeMandiri(*resp))
	} else {
		return c.JSON(http.StatusOK, response.FromCoreChargeMidtrans(*resp))
	}

}

func (h *OrderHandler) GetHistoryOrder(c echo.Context) error {
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, "failed to extract token")
	}

	resp, err := h.orderBusiness.SelectHistoryOrder(userID_token)
	if err != nil {
		return c.JSON(helper.ResponseNotFound("failed get history order"))
	}

	response := response.FromCoreList(resp)
	return c.JSON(helper.ResponseStatusOkWithData("success get history order", response))
}
