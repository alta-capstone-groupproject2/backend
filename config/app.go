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
<<<<<<< HEAD

=======
const Status = "waiting"
>>>>>>> f06af611cbbd14d15f271ff4f64aef978094944d
