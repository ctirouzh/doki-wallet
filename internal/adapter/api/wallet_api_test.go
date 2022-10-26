package api_test

import (
	"context"
	"doki/wallet/internal/adapter/api"
	"doki/wallet/internal/adapter/repository"
	"doki/wallet/internal/app"
	"doki/wallet/internal/domain"
	"doki/wallet/pb"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetBalanceService(t *testing.T) {
	// t.Parallel()
	walletRepo := repository.NewMockWalletRepo()
	transRepo := repository.NewMockTransRepo()
	walletService := app.NewWalletService(walletRepo, transRepo)
	transService := app.NewTransactionService(transRepo)
	walletAPI := api.NewWalletAPI(walletService, transService)

	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()
	ctxExeeded, exeed := context.WithTimeout(ctx, 1*time.Millisecond)
	time.Sleep(2 * time.Millisecond)
	exeed()

	type want struct {
		res *pb.GetBalanceResponse
		err error
	}
	testCases := []struct {
		name string
		ctx  context.Context
		req  *pb.GetBalanceRequest
		want want
	}{
		{
			name: "request is cancelled",
			ctx:  ctxCanceled,
			req:  &pb.GetBalanceRequest{UserId: 1},
			want: want{res: nil, err: status.Error(codes.Canceled, api.ErrRequestIsCanceled.Error())},
		},
		{
			name: "deadline is exeeded",
			ctx:  ctxExeeded,
			req:  &pb.GetBalanceRequest{UserId: 1},
			want: want{res: nil, err: status.Error(codes.DeadlineExceeded, api.ErrDeadlineExceeded.Error())},
		},
		{
			name: "successfully get balance",
			ctx:  ctx,
			req:  &pb.GetBalanceRequest{UserId: 1},
			want: want{res: &pb.GetBalanceResponse{Balance: 2000}, err: nil},
		},
		{
			name: "internal server error",
			ctx:  ctx,
			req:  &pb.GetBalanceRequest{UserId: 2},
			want: want{res: nil, err: status.Error(codes.Internal, api.ErrInternalServer.Error())},
		},
		{
			name: "wallet not found",
			ctx:  ctx,
			req:  &pb.GetBalanceRequest{UserId: 111},
			want: want{res: nil, err: status.Error(codes.NotFound, domain.ErrWalletNotFound.Error())},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotRes, gotErr := walletAPI.GetBalance(tc.ctx, tc.req)
			assert.Equal(t, tc.want.err, gotErr)
			assert.Equal(t, tc.want.res, gotRes)
		})
	}
}
