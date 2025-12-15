package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/shadman/shadis/internal/command"
	"github.com/shadman/shadis/internal/config"
	"github.com/shadman/shadis/internal/logger"
	"github.com/shadman/shadis/internal/store"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg := config.New()
	if err := cfg.Validate(); err != nil {
		fmt.Println("Config validation failed:", err)
		os.Exit(1)
	}

	// Initialize logger
	if err := logger.Init(cfg); err != nil {
		fmt.Println("Logger init failed:", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Info("=== Shadis Redis Server ===")
	logger.Info("Configuration loaded",
		zap.String("host", cfg.ServerConfig.Host),
		zap.Int("port", cfg.ServerConfig.Port),
		zap.String("log_level", cfg.LoggingConfig.LogLevel),
	)

	// Initialize the store
	st := store.New()
	logger.Info("Store initialized")

	// Initialize command registry
	registry := command.NewRegistry()
	command.RegisterCommands(registry, st)
	logger.Info("Commands registered", zap.Int("count", registry.Count()))

	// Start TCP listener
	addr := fmt.Sprintf("%s:%d", cfg.ServerConfig.Host, cfg.ServerConfig.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal("Failed to start server", zap.String("addr", addr), zap.Error(err))
	}
	defer listener.Close()

	logger.Info("Server started", zap.String("addr", addr))
	logger.Info("Ready to accept connections")

	// Handle graceful shutdown
	go handleShutdown(listener)

	// Accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			// Check if listener was closed
			select {
			default:
				logger.Error("Failed to accept connection", zap.Error(err))
				continue
			}
		}
		logger.Debug("New client connected", zap.String("remote", conn.RemoteAddr().String()))
		go HandleConnection(conn, registry)
	}
}

func handleShutdown(listener net.Listener) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	logger.Info("Shutting down server...")
	listener.Close()
	os.Exit(0)
}
