package domain

import (
	"time"
)

type TransStatus int8

const (
	TRANS_STATUS_TODO TransStatus = iota
	TRANS_STATUS_DOING
	TRANS_STATUS_FAILED
	TRANS_STATUS_DONE
)

func (status TransStatus) String() string {
	return [...]string{"TODO", "DOING", "FAILED", "DONE"}[status]
}

func (status TransStatus) Index() int {
	return int(status)
}

type Transaction struct {
	ID        uint
	WalletID  uint
	Amount    int64
	Status    TransStatus
	CreatedAt time.Time
	LockedAt  time.Time
}
