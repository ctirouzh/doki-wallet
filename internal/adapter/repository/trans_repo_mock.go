package repository

import (
	"doki/wallet/internal/domain"

	"github.com/stretchr/testify/mock"
)

type mockTransRepo struct {
	mock.Mock
}

func NewMockTransRepo() *mockTransRepo {
	return &mockTransRepo{}
}

func (repo *mockTransRepo) Find(id uint) (*domain.Transaction, error) {

	return nil, nil
}
func (repo *mockTransRepo) Create(
	wallet_id uint, amount int64,
) (*domain.Transaction, error) {

	return nil, nil
}

func (repo *mockTransRepo) Update(
	id uint, status domain.TransStatus,
) (*domain.Transaction, error) {

	return nil, nil
}
