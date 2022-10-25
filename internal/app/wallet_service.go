package app

import "doki/wallet/internal/domain"

type WalletService struct {
	walletRepo domain.WalletRepository
}

// NewWalletService is a factory which returns a pointer to new WalletService struct.
func NewWalletService(walletRepo domain.WalletRepository) *WalletService {
	return &WalletService{walletRepo: walletRepo}
}
