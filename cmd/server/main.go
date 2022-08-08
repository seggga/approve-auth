package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	config "github.com/seggga/approve-auth/internal/config"
	"github.com/seggga/approve-auth/internal/http/defaultmux"
	"github.com/seggga/approve-auth/internal/server/grpc"
	"github.com/seggga/approve-auth/internal/server/rest"
	"github.com/seggga/approve-auth/internal/storage/mongo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	conf := config.Read()

	logger := initLogger("debug")
	defer logger.Sync()

	// initialize storage
	userStore, err := mongo.New(conf.Mongo.DSN, conf.Data)
	// userStore, err := memstore.New(conf.Data)
	if err != nil {
		logger.Fatalf("error connecting to mongo db storage: %v", err)
	}

	// start REST service
	mux := defaultmux.New(userStore, conf.JWT.Secret, logger)
	restServer := rest.NewServer(mux, logger)
	restServer.Start()

	// start gRPC service
	grpcServer := grpc.NewServer(userStore, conf.JWT.Secret, logger)
	grpcServer.Start()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	<-ctx.Done()

	restServer.Stop()
	cancel()

	logger.Info("program exit")
}

func initLogger(level string) *zap.SugaredLogger {

	l, err := zapcore.ParseLevel(level)
	if err != nil {
		log.Fatalf("wrong value for level: %s, program exit", level)
	}

	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(l),
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    map[string]interface{}{"service": "auth"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "msg",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal("error parsing logger config:", err)
	}

	logger.Info("logger construction succeeded")

	return logger.Sugar()
}
