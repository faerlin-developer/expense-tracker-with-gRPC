package main

import (
	"context"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/api"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand/v2"
	"os"
	"time"
)

const HOST = "localhost"
const PORT = "5000"

func main() {

	ctx := context.Background()
	log := logger.NewStructuredLogger()

	log.Info(ctx, "starting client 1")

	// Connect to gRPC service
	conn, err := grpc.NewClient(
		HOST+":"+PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Error(ctx, "connect failed", "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Create expense client
	client := api.NewExpensesClient(conn)

	i := 0
	userIDs := []string{"A1", "B2", "C3"}

	// Call CreateExpense endpoint indefinitely
	for {

		id := userIDs[i]
		category := "Food"
		amount := randomFloat(0, 100)
		i = (i + 1) % len(userIDs)

		res, err := client.CreateExpense(context.Background(), randomRequest(id, category, amount))
		if err != nil {
			log.Error(ctx, "create expense failed", "error", err)
			os.Exit(1)
		}

		log.Info(ctx, "created expense", "id", res.Id)

		time.Sleep(time.Second)
	}
}

// randomRequest returns a random request for the CreateExpense endpoint.
func randomRequest(id string, category string, amount float64) *api.CreateExpenseRequest {
	return &api.CreateExpenseRequest{
		UserId:   id,
		Category: category,
		Amount:   amount,
	}
}

// randomFloat returns a random float in the range [min, max).
func randomFloat(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
