package main

import (
	"context"
	"fmt"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/api"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"os"
)

const HOST = "localhost"
const PORT = "5000"

func main() {

	ctx := context.Background()
	log := logger.NewStructuredLogger()

	log.Info(ctx, "starting client 2")

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

	// Stream data from the server indefinitely
	for {

		// Initiate server streaming RPC
		stream, err := client.GetBalances(context.Background(), &api.GetBalancesRequest{})
		if err != nil {
			log.Error(ctx, "get balances failed", "error", err)
			os.Exit(1)
		}

		// Receive data until server signals EOF
		var streamErr error
		var resp *api.GetBalancesResponse
		for {
			resp, streamErr = stream.Recv()
			if streamErr != nil {
				break
			}

			log.Info(ctx, "get balance",
				"userID", resp.UserId,
				"balance", fmt.Sprintf("%.2f", resp.Amount),
				"num_expenses", resp.NumExpenses,
			)
		}

		// Terminate if streaming fails
		if streamErr != io.EOF {
			log.Error(ctx, "stream recv failed", "error", err)
			os.Exit(1)
		}
	}
}
