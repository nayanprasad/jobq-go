package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/nayanprasad/jobQ-go/internal/server"
)

func main() {
	//logger
	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	slog.Debug("ping")

	//load config
	appConfig, err := LoadConfig("config/config.yaml")
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	// server setup
	cnf := server.Config{
		Addr: fmt.Sprintf(":%d", appConfig.Server.Port),
		DSN:  appConfig.Server.DB.DSN,
	}

	svr := server.New(cnf)

	h := svr.Mount()
	if err := svr.Run(h); err != nil {
		slog.Error("failied start the server", "error", err.Error())
		os.Exit(1)
	}
}
