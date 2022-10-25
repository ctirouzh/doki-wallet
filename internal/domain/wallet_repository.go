package domain

type WalletRepository interface {
	Find(user_id uint) (*Wallet, error)
	Create()
}
