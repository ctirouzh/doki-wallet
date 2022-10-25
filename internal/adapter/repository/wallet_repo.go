package repository

import (
	"doki/wallet/internal/domain"
	"errors"

	"gorm.io/gorm"
)

type WalletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) *WalletRepo {
	return &WalletRepo{db: db}
}

func (repo *WalletRepo) Find(user_id uint) (*domain.Wallet, error) {
	return nil, errors.New("repository interface not implemented")
}
