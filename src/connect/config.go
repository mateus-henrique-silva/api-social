package connect

import (
	"log"

	"github.com/joho/godotenv"
)

var (
	Port     string
	MongoUri string
)

func Load() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal(err)
	}

}
