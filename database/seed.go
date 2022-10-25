package database

import (
	"doki/wallet/internal/domain"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Seed simply adds n new wallet record with auto incremented user_id to the database.
// It only seeds if wallets table has no records.
func Seed(n int) error {
	if DB.Migrator().HasTable(&domain.Wallet{}) {
		if err := DB.First(&domain.Wallet{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			for i := 1; i <= n; i++ {
				wallet := domain.Wallet{
					UserID:  uint(i),
					Balance: domain.Balance(0),
				}
				if createErr := DB.Create(&wallet).Error; createErr != nil {
					return createErr
				}
			}
		}
	}
	if DB.Migrator().HasTable(&domain.Transaction{}) {
		if err := DB.First(&domain.Transaction{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			trans := &domain.Transaction{
				ID:       10000000000, // set @autoinc to min 11 digit number...
				WalletID: 1,
				Amount:   2000,
				Status:   domain.TRANS_STATUS_DONE,
				LockedAt: time.Now(),
			}
			if createErr := DB.Create(trans).Error; createErr != nil {
				return createErr
			}
			if updateErr := DB.Model(&domain.Wallet{}).
				Where("id=?", 1).
				Update("balance", 2000).Error; updateErr != nil {
				return updateErr
			}
		}
	}
	return nil
}
