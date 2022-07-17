package config

import (
	"os"
)

func JWT() string {
	SECRET_JWT := os.Getenv("SECRET_JWT")
	return SECRET_JWT
}

const Admin = "admin"
const User = "user"
const UMKM = "umkm"
const Status = "waiting"
