package app

import "doki/wallet/internal/domain"

type TransactionService struct {
	transRepo domain.TransRepository
}

// NewTransactionService is a factory which returns a pointer to new TransactionService struct.
func NewTransactionService(transRepo domain.TransRepository) *TransactionService {
	return &TransactionService{transRepo: transRepo}
}
