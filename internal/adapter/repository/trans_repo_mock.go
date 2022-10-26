package repository

import (
	"doki/wallet/internal/domain"
	"time"

	"github.com/stretchr/testify/mock"
)

var autoinc uint = 1000

type mockTransRepo struct {
	mock.Mock
	data map[uint]*domain.Transaction
}

func NewMockTransRepo() *mockTransRepo {
	return &mockTransRepo{data: make(map[uint]*domain.Transaction, 10)}
}

func (repo *mockTransRepo) Find(id uint) (*domain.Transaction, error) {

	return nil, nil
}
func (repo *mockTransRepo) Create(
	wallet_id uint, amount int64,
) (*domain.Transaction, error) {
	autoinc++
	repo.data[autoinc] = &domain.Transaction{
		ID:        autoinc,
		WalletID:  wallet_id,
		Amount:    amount,
		Status:    domain.TRANS_STATUS_DONE,
		CreatedAt: time.Now(),
		LockedAt:  time.Now(),
	}
	return repo.data[autoinc], nil
}

func (repo *mockTransRepo) Update(
	id uint, status domain.TransStatus,
) (*domain.Transaction, error) {

	return nil, nil
}
