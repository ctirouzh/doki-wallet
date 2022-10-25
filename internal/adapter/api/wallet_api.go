package api

import (
	"context"
	"doki/wallet/internal/app"
	"doki/wallet/internal/domain"
	"doki/wallet/pb"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// WalletAPI implements the WalletServiceServer interface of the pb package.
type WalletAPI struct {
	pb.UnimplementedWalletServiceServer
	walletService *app.WalletService
	transService  *app.TransactionService
}

// NewWalletAPI is a factory which returns a pointer to a new WalletAPI struct.
func NewWalletAPI(walletService *app.WalletService, transService *app.TransactionService) *WalletAPI {
	return &WalletAPI{walletService: walletService, transService: transService}
}

// GetBalance retrieves the wallet balance of the given user.
func (api *WalletAPI) GetBalance(
	ctx context.Context, req *pb.GetBalanceRequest,
) (*pb.GetBalanceResponse, error) {

	log.Printf("[api]>>> received a \"get balance\" request with user_id: %d", req.UserId)
	// Check and log if context has Canceled or DeadLineExceeded error
	if err := checkContextError(ctx); err != nil {
		return nil, err
	}

	balance, err := api.walletService.GetBalance(uint(req.UserId))
	// log.Println("[api]>>> Balance: ", balance)
	if err != nil {
		switch err {
		case domain.ErrWalletNotFound:
			return nil, status.Error(codes.NotFound, domain.ErrWalletNotFound.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer.Error())
		}
	}
	return &pb.GetBalanceResponse{Balance: int64(balance)}, nil
}

func (api *WalletAPI) AddMoney(
	ctx context.Context, req *pb.AddMoneyRequest,
) (*pb.AddMoneyResponse, error) {

	log.Printf("[api]>>> received a \"add money\" request with user_id: %d, amount: %d", req.UserId, req.Amount)
	// Check and log if context has Canceled or DeadLineExceeded error
	if err := checkContextError(ctx); err != nil {
		return nil, err
	}

	refrenceID, err := api.walletService.AddMoney(uint(req.UserId), req.Amount)
	if err != nil {
		switch err {
		case domain.ErrWalletNotFound:
			return nil, status.Error(codes.NotFound, domain.ErrWalletNotFound.Error())
		case app.ErrNotEnoughBalance:
			return nil, status.Error(codes.FailedPrecondition, app.ErrNotEnoughBalance.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer.Error())
		}
	}
	return &pb.AddMoneyResponse{ReferenceId: int64(refrenceID)}, nil
}
