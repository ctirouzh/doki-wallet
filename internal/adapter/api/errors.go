package api

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrApiNotImplemented = errors.New("api not implemented")
	ErrInternalServer    = errors.New("internal server error")
)

// logError logs an error, and returns nil in case of no error
func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

// checkContextError checks if the context has Canceled or DeadlineExeeded error,
// logs the error, and returns the error with its status code
func checkContextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}
