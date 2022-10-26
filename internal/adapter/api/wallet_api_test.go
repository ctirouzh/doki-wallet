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

func TestGetBalanceServer(t *testing.T) {
	walletAPI := initTestWalletAPIs()

	// Init contexts of test cases
	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()
	ctxExeeded, exeed := context.WithTimeout(ctx, 1*time.Millisecond)
	time.Sleep(2 * time.Millisecond)
	exeed()

	// Init test cases
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
			name: "wallet not found",
			ctx:  ctx,
			req:  &pb.GetBalanceRequest{UserId: 111},
			want: want{res: nil, err: status.Error(codes.NotFound, domain.ErrWalletNotFound.Error())},
		},
	}
	// Run tests...
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotRes, gotErr := walletAPI.GetBalance(tc.ctx, tc.req)
			assert.Equal(t, tc.want.err, gotErr)
			assert.Equal(t, tc.want.res, gotRes)
		})
	}
}

func TestAddMoneyServer(t *testing.T) {
	walletAPI := initTestWalletAPIs()

	// Init contexts of test cases
	ctx := context.Background()
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()
	ctxExeeded, exeed := context.WithTimeout(ctx, 1*time.Millisecond)
	time.Sleep(2 * time.Millisecond)
	exeed()

	// Init test cases
	type want struct {
		res     *pb.AddMoneyResponse
		balance int64
		err     error
	}
	testCases := []struct {
		name string
		ctx  context.Context
		req  *pb.AddMoneyRequest
		want want
	}{
		{
			name: "request is cancelled",
			ctx:  ctxCanceled,
			req:  &pb.AddMoneyRequest{UserId: 1, Amount: 200},
			want: want{res: nil, balance: 2000, err: status.Error(codes.Canceled, api.ErrRequestIsCanceled.Error())},
		},
		{
			name: "deadline is exeeded",
			ctx:  ctxExeeded,
			req:  &pb.AddMoneyRequest{UserId: 1, Amount: 200},
			want: want{res: nil, err: status.Error(codes.DeadlineExceeded, api.ErrDeadlineExceeded.Error())},
		},
		{
			name: "wallet not found",
			ctx:  ctx,
			req:  &pb.AddMoneyRequest{UserId: 111, Amount: 200},
			want: want{res: nil, err: status.Error(codes.NotFound, domain.ErrWalletNotFound.Error())},
		},
		{
			name: "not enough balance",
			ctx:  ctx,
			req:  &pb.AddMoneyRequest{UserId: 1, Amount: -2001},
			want: want{res: nil, err: status.Error(codes.FailedPrecondition, app.ErrNotEnoughBalance.Error())},
		},
		{
			name: "successfully add money",
			ctx:  ctx,
			req:  &pb.AddMoneyRequest{UserId: 1, Amount: 200},
			want: want{res: &pb.AddMoneyResponse{ReferenceId: 1001}, balance: 2200, err: nil},
		},
		{
			name: "successfully take money",
			ctx:  ctx,
			req:  &pb.AddMoneyRequest{UserId: 1, Amount: -300},
			want: want{res: &pb.AddMoneyResponse{ReferenceId: 1002}, balance: 1900, err: nil},
		},
	}
	// Run tests...
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotRes, gotErr := walletAPI.AddMoney(tc.ctx, tc.req)
			assert.Equal(t, tc.want.err, gotErr)
			assert.Equal(t, tc.want.res, gotRes)
			if gotErr == nil {
				req := &pb.GetBalanceRequest{UserId: tc.req.UserId}
				newBalance, _ := walletAPI.GetBalance(context.Background(), req)
				assert.Equal(t, tc.want.balance, newBalance.Balance)
			}
		})
	}

}

func initTestWalletAPIs() *api.WalletAPI {
	walletRepo := repository.NewMockWalletRepo()
	transRepo := repository.NewMockTransRepo()
	walletService := app.NewWalletService(walletRepo, transRepo)
	return api.NewWalletAPI(walletService)
}
