package presentation

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"lami/app/features/orders"
	"lami/app/features/orders/presentation/request"
	"lami/app/features/orders/presentation/response"
	"lami/app/helper"
	"lami/app/middlewares"

	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"net/http"

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
	orderCore.Status = "Pending"

	res, err := h.orderBusiness.Order(orderCore, userID_token)
	if err != nil && res == -1 {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	//	Payments Order
	idOrder, errorderid := h.orderBusiness.PaymentsOrderID(userID_token)
	if errorderid != nil {
		return c.JSON(helper.ResponseInternalServerError("failed get order_id for payments"))
	}

	grossamount, errgrossamount := h.orderBusiness.PaymentGrossAmount(userID_token)
	if errgrossamount != nil {
		return c.JSON(helper.ResponseInternalServerError("failed get gross amount for payments"))
	}

	transfer, paymentsID, errtransfer := h.orderBusiness.TypeBank(int64(grossamount), typeName, idOrder)
	if errtransfer != nil {
		return c.JSON(helper.ResponseBadRequest("failed param for type bank"))
	}
	fmt.Println(paymentsID)

	transferCore, errcore := h.orderBusiness.RequestChargeBank(transfer, typeName)
	if errcore != nil {
		return c.JSON(helper.ResponseInternalServerError("failed data bank transfer"))
	}

	resp, errcharge := coreapi.ChargeTransaction(&transferCore)
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

func (h *OrderHandler) PostUpdatedStatusPayments(c echo.Context) error {
	idOrderParam := c.Param("idOrder")
	idOrder, errparam := strconv.Atoi(idOrderParam)
	if errparam != nil {
		return c.JSON(helper.ResponseBadRequest("failed get param"))
	}

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	url := fmt.Sprintf("https://api.sandbox.midtrans.com/v2/"+"%s"+"/status", idOrderParam)

	payload := strings.NewReader("\n\n")

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, payload)
	if err != nil {
		fmt.Println(err)
	}

	midtrans.ServerKey = os.Getenv("MIDTRANS_ORDER_SERVER_KEY") + ":"
	sEnc := base64.StdEncoding.EncodeToString([]byte(midtrans.ServerKey))

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+sEnc)

	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed get api status")
	}

	errstatus := h.orderBusiness.UpdateStatus(idOrder)
	if errstatus != nil {
		c.JSON(helper.ResponseBadRequest("fauled update status in database"))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed res.Body")
	}
	defer res.Body.Close()

	resp := coreapi.TransactionStatusResponse{}
	json.Unmarshal(body, &resp)

	return c.JSON(http.StatusOK, "Success confirm")
}
