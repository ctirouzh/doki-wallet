package domain

import "time"

type Wallet struct {
	ID        uint
	UserID    uint
	Balance   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
