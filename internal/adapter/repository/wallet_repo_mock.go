package repository

import (
	"doki/wallet/internal/domain"
	"errors"
	"time"

	"github.com/stretchr/testify/mock"
)

type mockWalletRepo struct {
	mock.Mock
}

func NewMockWalletRepo() *mockWalletRepo {
	return &mockWalletRepo{}
}

func (repo *mockWalletRepo) Find(user_id uint) (*domain.Wallet, error) {
	switch user_id {
	case 1:
		return &domain.Wallet{
			ID:        1,
			UserID:    1,
			Balance:   2000,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	case 2:
		return nil, errors.New("any internal server error")
	default:
		return nil, domain.ErrWalletNotFound
	}
}
func (repo *mockWalletRepo) Update(id uint, balance domain.Balance) error {
	return nil
}
