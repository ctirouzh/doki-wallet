package app

import (
	"doki/wallet/internal/domain"
)

type WalletService struct {
	walletRepo domain.WalletRepository
}

// NewWalletService is a factory which returns a pointer to a new WalletService struct.
func NewWalletService(walletRepo domain.WalletRepository) *WalletService {
	return &WalletService{walletRepo: walletRepo}
}

// GetBalance searches in wallet repository with the given user id and retrieves its balance.
func (s *WalletService) GetBalance(user_id uint) (domain.Balance, error) {
	wallet, err := s.walletRepo.Find(user_id)
	if err != nil {
		return -1, err
	}
	return wallet.Balance, nil
}
