package store

import (
	e "github.com/faerlin-developer/expense-tracker-with-gRPC/internal/errors"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/internal/expense"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/internal/logger"
)

// InMemoryStore is an in-memory implementation of the Store interface.
type InMemoryStore struct {
	expenses       map[string]expense.Expense // maps expense ID to expense data
	expensesByUser map[string][]string        // maps user ID to a list of expense IDs
	log            logger.Logger
}

// NewInMemoryStore returns an instance of InMemoryStore.
func NewInMemoryStore(log logger.Logger) *InMemoryStore {
	return &InMemoryStore{
		log:            log,
		expenses:       map[string]expense.Expense{},
		expensesByUser: map[string][]string{},
	}
}

func (s *InMemoryStore) Put(exp expense.Expense) {

	id := exp.ID
	userId := exp.UserID

	// Add to map of expenses by id
	_, ok := s.expenses[id]
	s.expenses[id] = exp

	// Add to map of expense ids by user id
	if !ok {
		s.expensesByUser[userId] = append(s.expensesByUser[userId], id)
	}
}

func (s *InMemoryStore) Get(id string) (expense.Expense, error) {

	// Check if id is present in storage
	exp, ok := s.expenses[id]
	if !ok {
		return expense.Expense{}, e.InvalidInputError{Field: id, Reason: "id not found"}
	}

	return exp, nil
}

func (s *InMemoryStore) Delete(id string) {

	// If expense id is not present in storage, return immediately
	exp, ok := s.expenses[id]
	if !ok {
		return
	}

	// Remove expense from both internal maps
	userId := exp.UserID
	delete(s.expenses, id)
	s.expensesByUser[userId] = deleteByValue(s.expensesByUser[userId], id)
}

func (s *InMemoryStore) List(userId string) []expense.Expense {

	// Check if user id is present in storage
	ids, ok := s.expensesByUser[userId]
	if !ok {
		return []expense.Expense{}
	}

	// Collect expenses for specified user
	var list []expense.Expense
	for _, id := range ids {
		list = append(list, s.expenses[id])
	}

	return list
}

func (s *InMemoryStore) GetAllUsers() []string {

	keys := make([]string, len(s.expensesByUser))

	i := 0
	for k, _ := range s.expensesByUser {
		keys[i] = k
		i++
	}

	return keys
}

// deleteByValue deletes an element from a slice by value, works for any type.
func deleteByValue[T comparable](slice []T, value T) []T {
	for i, item := range slice {
		if item == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
