package expense

import (
	"fmt"
	e "github.com/faerlin-developer/expense-tracker-with-gRPC/internal/errors"
	"github.com/google/uuid"
	"time"
)

// Expense models a financial expense.
type Expense struct {
	ID          string    // Unique identifier for the expense
	UserID      string    // ID of the user who made the expense
	Category    string    // Category of the expense
	Amount      float64   // The amount spent
	Description string    // Description of the expense
	TimeStamp   time.Time // ISO 8601 format
}

// New creates a valid instance of Expense.
// Returns an error of type InvalidInputError if any of the arguments fails validation.
func New(userID string, category string, amount float64, description *string, timestamp *string) (Expense, error) {

	if userID == "" {
		return Expense{}, e.InvalidInputError{Field: "user_id", Reason: "must not be an empty string"}
	}

	if category == "" {
		return Expense{}, e.InvalidInputError{Field: "category", Reason: "must not be an empty string"}
	}

	if amount < 0 {
		return Expense{}, e.InvalidInputError{Field: "amount", Reason: "must not be less than zero"}
	}

	// If timestamp is not specified, assign current time.
	// Otherwise, attempt to parse it assuming ISO 8601 format.
	ts := time.Now().UTC()
	if timestamp != nil {
		var err error
		ts, err = time.Parse(time.RFC3339, *timestamp)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return Expense{}, e.InvalidInputError{Field: "timestamp", Reason: err.Error()}
		}
	}

	desc := ""
	if description != nil {
		desc = *description
	}

	// Assign this expense an internally generated ID
	id := uuid.NewString()

	return Expense{
		ID:          id,
		UserID:      userID,
		Category:    category,
		Amount:      amount,
		Description: desc,
		TimeStamp:   ts,
	}, nil
}
