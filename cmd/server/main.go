package main

import (
	"doki/wallet/config"
	"doki/wallet/database"

	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("[server]>>> Welcome!")

	fmt.Println("[server]>>> Loading environment variables...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	cfg := config.Parse()
	fmt.Println(cfg)

	fmt.Println("[server]>>> Connecting to database...")
	database.ConnectMySQL(cfg.MySQL)
	database.Migrate()
	db := database.GetDB()
	fmt.Println(db)
}
