package connect

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SecretKey []byte
)

func Load() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
