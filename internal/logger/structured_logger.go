package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// StructuredLogger provides structured and leveled logging.
// It utilizes zap.SugaredLogger for dynamic fields and enhanced structured logging.
type StructuredLogger struct {
	log *zap.SugaredLogger
}

// NewStructuredLogger initializes and returns a new Logger instance with production-level configuration.
// It sets up the logger to output to stdout, disables stack traces, and uses ISO8601 time encoding.
// If logger creation fails, the function panics.
func NewStructuredLogger() *StructuredLogger {

	// Configure logger
	logConfig := zap.NewProductionConfig()
	logConfig.DisableStacktrace = true
	// logConfig.DisableCaller = true
	logConfig.OutputPaths = []string{"stdout"}
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Build logger
	logger, err := logConfig.Build()
	if err != nil {
		panic("failed to build logger")
	}

	return &StructuredLogger{logger.Sugar()}
}

// Info logs at the info level.
func (logger *StructuredLogger) Info(ctx context.Context, msg string, keysAndValues ...any) {

	// Check if the context contains a traceID
	traceID, ok := ctx.Value("traceID").(string)
	if ok && traceID != "" {
		keysAndValues = append(keysAndValues, "traceID", traceID)
	}

	logger.log.WithOptions(zap.AddCallerSkip(1)).Infow(msg, keysAndValues...)
}

// Error logs at the error level
func (logger *StructuredLogger) Error(ctx context.Context, msg string, keysAndValues ...any) {

	// Check if the context contains a traceID
	traceID, ok := ctx.Value("traceID").(string)
	if ok && traceID != "" {
		keysAndValues = append(keysAndValues, "traceID", traceID)
	}

	logger.log.WithOptions(zap.AddCallerSkip(1)).Errorw(msg, keysAndValues...)
}
