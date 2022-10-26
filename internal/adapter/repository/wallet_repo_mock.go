package repository

import (
	"doki/wallet/internal/domain"
	"time"

	"github.com/stretchr/testify/mock"
)

type mockWalletRepo struct {
	mock.Mock
	data map[uint]*domain.Wallet
}

func NewMockWalletRepo() *mockWalletRepo {
	data := make(map[uint]*domain.Wallet)
	data[1] = &domain.Wallet{
		ID:        1,
		UserID:    1,
		Balance:   2000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &mockWalletRepo{data: data}
}

func (repo *mockWalletRepo) Find(user_id uint) (*domain.Wallet, error) {
	user, found := repo.data[user_id]
	if !found {
		return nil, domain.ErrWalletNotFound
	}
	return user, nil
}
func (repo *mockWalletRepo) Update(id uint, balance domain.Balance) error {
	wallet, err := repo.Find(id)
	if err != nil {
		return err
	}
	wallet.Balance = balance
	repo.data[id] = wallet
	return nil
}
