package main

import (
	"doki/wallet/config"
	"doki/wallet/database"
	"log"

	"fmt"
)

func main() {
	fmt.Println("[server]>>> Welcome!")

	fmt.Println("[server]>>> Parsing configs...")
	cfg := config.Parse()

	log.Println("config: ", cfg)

	fmt.Println("[server]>>> Connecting to database...")
	database.ConnectMySQL(cfg.MySQL)
	database.Migrate()
	db := database.GetDB()
	fmt.Println(db)
}
