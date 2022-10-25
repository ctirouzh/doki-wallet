package app

import (
	"doki/wallet/internal/domain"
	"errors"
)

var (
	ErrNotEnoughBalance = errors.New("not enough balance, charge your wallet")
)

type WalletService struct {
	walletRepo domain.WalletRepository
	transRepo  domain.TransRepository
}

// NewWalletService is a factory which returns a pointer to a new WalletService struct.
func NewWalletService(walletRepo domain.WalletRepository, transRepo domain.TransRepository) *WalletService {
	return &WalletService{walletRepo: walletRepo, transRepo: transRepo}
}

// GetBalance searches in wallet repository with the given user id and retrieves its balance.
func (s *WalletService) GetBalance(user_id uint) (domain.Balance, error) {
	wallet, err := s.walletRepo.Find(user_id)
	if err != nil {
		return -1, err
	}
	return wallet.Balance, nil
}

// AddMoney creates new transaction with given user's wallet_id, and returns the
// transaction reference_id.
func (s *WalletService) AddMoney(user_id uint, amount int64) (uint, error) {

	wallet, err := s.walletRepo.Find(user_id)
	if err != nil {
		return 0, err
	}

	balance := wallet.Balance + domain.Balance(amount)
	if balance < 0 {
		return 0, ErrNotEnoughBalance
	}

	trans, createErr := s.transRepo.Create(wallet.ID, amount)
	if createErr != nil {
		return 0, err
	}
	// TODO: simulate transaction using goroutine...
	// Let assume transaction is well done.
	s.walletRepo.Update(wallet.ID, balance)

	return trans.ID, nil
}
