package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
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
const PaymentStatus = "unpaid"
const MethodPost = "POST"
const PaymentBankTransferBCA = "BANK_TRANSFER_BCA"

func JWT() string {
	SECRET_JWT := os.Getenv("SECRET_JWT")
	return SECRET_JWT
}

func EncryptKey() string {
	ENCRYPT_KEY := os.Getenv("ENCRYPT_KEY")
	return ENCRYPT_KEY
}

func MidtransServerKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file")
	}
	MIDTRANS_SERVER_KEY := os.Getenv("MIDTRANS_EVENT_SERVER_KEY")
	return MIDTRANS_SERVER_KEY
}

func MidtransOrderServerKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file")
	}
	MIDTRANS_ORDER_SERVER_KEY := os.Getenv("MIDTRANS_ORDER_SERVER_KEY")
	return MIDTRANS_ORDER_SERVER_KEY
}
