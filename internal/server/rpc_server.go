package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/api"
	e "github.com/faerlin-developer/expense-tracker-with-gRPC/internal/errors"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/internal/expense"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/internal/logger"
	"github.com/faerlin-developer/expense-tracker-with-gRPC/internal/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"strconv"
	"time"
)

// RpcServer is a gRPC server.
type RpcServer struct {
	db  store.Store
	log logger.Logger
	api.UnimplementedExpensesServer
}

// NewRpcServer returns a gRPC server.
func NewRpcServer(db store.Store, log logger.Logger) *RpcServer {
	return &RpcServer{
		db:  db,
		log: log,
	}
}

func (s *RpcServer) ListenAndServe(ctx context.Context, port int) error {

	// Create network address
	host := "127.0.0.1"
	p := strconv.Itoa(port)
	networkAddr := net.JoinHostPort(host, p)

	// Create listener
	listener, err := net.Listen("tcp", networkAddr)
	if err != nil {
		return fmt.Errorf("unable to listen to tcp port %d: %w", port, err)
	}

	// Register this server to gRPC
	rpcServer := grpc.NewServer()
	api.RegisterExpensesServer(rpcServer, s)

	// Start listening
	err = rpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("error while listening: %w", err)
	}

	return nil
}

// CreateExpense handles the request to create an expense
func (s *RpcServer) CreateExpense(
	ctx context.Context,
	request *api.CreateExpenseRequest,
) (*api.CreateExpenseResponse, error) {

	exp, err := expense.New(
		request.UserId,
		request.Category,
		request.Amount,
		request.Description,
		request.Timestamp)

	if err != nil {
		// When returning an error in a gRPC endpoint,
		// error should be built from status.Error or status.Errorf
		s.log.Error(ctx, "creating expense failed", "error", err)
		var invalidErr e.InvalidInputError
		if errors.As(err, &invalidErr) {
			return nil, status.Error(codes.InvalidArgument, invalidErr.Error())
		} else {
			return nil, status.Errorf(codes.Internal, "creating expense failed: %s", err.Error())
		}
	}

	// Store expense in database
	s.db.Put(exp)

	// Set expense ID as the trace ID of this request
	ctx = context.WithValue(ctx, "traceID", exp.ID)
	s.log.Info(ctx, "create expense success")

	return &api.CreateExpenseResponse{Id: exp.ID}, nil
}

func (s *RpcServer) GetBalances(
	request *api.GetBalancesRequest,
	stream grpc.ServerStreamingServer[api.GetBalancesResponse],
) error {

	for _, userID := range s.db.GetAllUsers() {

		// Expense IDs for user
		exps := s.db.List(userID)
		numExps := len(exps)

		// Compute the balance for user
		var total float64 = 0
		for _, exp := range exps {
			total += exp.Amount
		}

		resp := &api.GetBalancesResponse{
			UserId:      userID,
			Amount:      total,
			NumExpenses: int32(numExps),
		}

		// Send the response to the client
		if err := stream.Send(resp); err != nil {
			s.log.Error(context.Background(), "stream balance failed", "error", err)
			return status.Errorf(codes.Internal, "stream balance failed: %s", err.Error())
		}

		// Simulate delay
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}
