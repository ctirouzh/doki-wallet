package domain

type TransRepository interface {
	Find(id uint) (*Transaction, error)
	Create(wallet_id uint, amount int64) (*Transaction, error)
	Update(id uint, status TransStatus) (*Transaction, error)
}
