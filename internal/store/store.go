package store

import (
	"github.com/faerlin-developer/expense-tracker-with-gRPC/internal/expense"
)

type Store interface {
	Put(exp expense.Expense)
	Get(id string) (expense.Expense, error)
	Delete(id string)
	List(userId string) []expense.Expense
	GetAllUsers() []string
}
