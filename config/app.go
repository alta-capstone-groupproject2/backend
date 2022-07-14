package config

import (
	"os"
)

func JWT() string {
	SECRET_JWT := os.Getenv("SECRET_JWT")
	return SECRET_JWT
}
