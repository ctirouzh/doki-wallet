package repository

import (
	"doki/wallet/internal/domain"
	"errors"
	"time"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

// Find retrieves transaction record by id.
func (repo *TransactionRepo) Find(id uint) (*domain.Transaction, error) {
	trans := new(domain.Transaction)
	if err := repo.db.Where("id=?", id).First(trans).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrTransactionNotFound
		}
		return nil, err
	}
	return trans, nil
}

// Create creates new transaction record using given wallet and amount.
// It internally sets the transaction's status to "DONE".
func (repo *TransactionRepo) Create(wallet_id uint, amount int64) (*domain.Transaction, error) {
	trans := &domain.Transaction{
		WalletID: wallet_id,
		Amount:   amount,
		Status:   domain.TRANS_STATUS_DONE,
		LockedAt: time.Now(),
	}
	if err := repo.db.Create(trans).Error; err != nil {
		return nil, err
	}
	return trans, nil
}

func (repo *TransactionRepo) Update(id uint, status domain.TransStatus) (*domain.Transaction, error) {
	return nil, ErrRepositoryMethodNotImplemented
}

func (repo *TransactionRepo) Sum(from, to time.Time) int64 {
	var result int64
	repo.db.Table("transactions").
		Where("status=?", domain.TRANS_STATUS_DONE).
		Where("created_at BETWEEN ? AND ?", from.String(), to.String()).
		Select("sum(amount)").Row().Scan(&result)
	return result
}
