package repository

import (
	"doki/wallet/internal/domain"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (repo *TransactionRepo) Find(id uint) (*domain.Transaction, error) {
	return nil, ErrRepositoryMethodNotImplemented

}

func (repo *TransactionRepo) Create(wallet_id uint, amount int64) (*domain.Transaction, error) {
	return nil, ErrRepositoryMethodNotImplemented

}
func (repo *TransactionRepo) Update(id uint, status domain.TransStatus) (*domain.Transaction, error) {
	return nil, ErrRepositoryMethodNotImplemented
}
