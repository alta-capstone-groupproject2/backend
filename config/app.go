package config

import (
	"os"
)

const Admin = "admin"
const User = "user"
const UMKM = "umkm"

/**
* @Method Status Validasi Event
**/
const Waiting = "waiting"
const Approved = "approved"

/**
* @Method Payment with Midtrans
**/
const MethodPost = "POST"
const PaymentBankTransferBCA = "BANK_TRANSFER_BCA"

func JWT() string {
	SECRET_JWT := os.Getenv("SECRET_JWT")
	return SECRET_JWT
}

func MidtransServerKey() string {
	MIDTRANS_SERVER_KEY := os.Getenv("MIDTRANS_EVENT_SERVER_KEY")
	return MIDTRANS_SERVER_KEY
}
