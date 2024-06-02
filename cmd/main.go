package main

import (
	"gotvup/internal/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	/*pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = godotenv.Load(filepath.Join(pwd, "../.env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server.StartWebServer()
}
