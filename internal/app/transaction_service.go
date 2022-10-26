package app

import (
	"doki/wallet/internal/domain"
	"log"
	"time"
)

type TransactionService struct {
	transRepo domain.TransRepository
}

// NewTransactionService is a factory which returns a pointer to new TransactionService struct.
func NewTransactionService(transRepo domain.TransRepository) *TransactionService {
	return &TransactionService{transRepo: transRepo}
}

func (s *TransactionService) DailyReport() {
	to := time.Now()
	from := to.AddDate(0, 0, -1)
	log.Println("======================================")
	log.Println("       Transaction Daily Report       ")
	log.Println("======================================")
	log.Printf(" >>>>>>> From Datatime: %s\n", from.Format("2006-01-02 15:04:05"))
	log.Printf(" >>>>>>>>> To Datetime: %s\n", to.Format("2006-01-02 15:04:05"))
	log.Printf(" >>> Total Transaction: %d\n", s.GetTotalTransaction(from, to))
}

func (s *TransactionService) GetTotalTransaction(from, to time.Time) int64 {
	return s.transRepo.Sum(from, to)
}
