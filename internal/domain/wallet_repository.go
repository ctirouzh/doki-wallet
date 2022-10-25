package domain

import "errors"

var (
	ErrWalletNotFound      = errors.New("wallet not found")
	ErrWalletAlreadyExists = errors.New("wallet already exists")
)

type WalletRepository interface {
	Find(user_id uint) (*Wallet, error)
	Update(id uint, balance Balance) error
}
