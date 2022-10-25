package main

import (
	"doki/wallet/config"
	"doki/wallet/database"
	"doki/wallet/internal/adapter/api"
	"doki/wallet/internal/adapter/repository"
	"doki/wallet/internal/app"
	"doki/wallet/pb"
	"log"
	"net"

	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	walletRepo := repository.NewWalletRepo(db)
	walletService := app.NewWalletService(walletRepo)

	transRepo := repository.NewTransactionRepo(db)
	transService := app.NewTransactionService(transRepo)

	walletAPI := api.NewWalletAPI(walletService, transService)

	grpcServer := grpc.NewServer()

	// Register All API servers
	pb.RegisterWalletServiceServer(grpcServer, walletAPI)

	// Enable reflection feature of gRPC Server
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", cfg.Server.Address())
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	log.Printf("server started on port %s\n", cfg.Server.Port)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
