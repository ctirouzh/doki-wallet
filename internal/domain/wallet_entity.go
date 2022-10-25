package domain

import "time"

type Balance int64

type Wallet struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"uniqueIndex;not null"`
	Balance   Balance `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
