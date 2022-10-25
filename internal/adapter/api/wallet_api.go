package api

import (
	"context"
	"doki/wallet/internal/app"
	"doki/wallet/pb"
	"errors"
)

var (
	ErrAPINotImplemented = errors.New("api not implemented")
)

type WalletAPI struct {
	pb.UnimplementedWalletServiceServer
	walletService *app.WalletService
	transService  *app.TransactionService
}

func NewWalletAPI(walletService *app.WalletService, transService *app.TransactionService) *WalletAPI {
	return &WalletAPI{walletService: walletService, transService: transService}
}

func (api *WalletAPI) GetBalance(
	ctx context.Context, req *pb.GetBalanceRequest,
) (*pb.GetBalanceResponse, error) {

	if err := checkContextError(ctx); err != nil {
		return nil, err
	}
	return nil, ErrAPINotImplemented
}

func (api *WalletAPI) AddMoney(
	ctx context.Context, req *pb.AddMoneyRequest,
) (*pb.AddMoneyResponse, error) {

	if err := checkContextError(ctx); err != nil {
		return nil, err
	}
	return nil, ErrAPINotImplemented
}
