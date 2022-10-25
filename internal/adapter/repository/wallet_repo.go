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
	wallet := new(domain.Wallet)
	if err := repo.db.Where("user_id=?", user_id).First(wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrWalletNotFound
		}
		return nil, err
	}
	return wallet, nil
}
