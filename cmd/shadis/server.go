package main

import (
	"bufio"
	"io"
	"net"
	"time"

	"github.com/shadman/shadis/internal/command"
	"github.com/shadman/shadis/internal/logger"
	"github.com/shadman/shadis/internal/resp"
	"go.uber.org/zap"
)

func HandleConnection(conn net.Conn, registry *command.Registry) {
	defer func() {
		conn.Close()
		logger.Debug("Connection closed", zap.String("remote", conn.RemoteAddr().String()))
	}()

	remoteAddr := conn.RemoteAddr().String()
	reader := bufio.NewReader(conn)
	commandCount := 0

	for {
		// Parse RESP command
		start := time.Now()
		args, err := resp.Parse(reader)
		
		if err == io.EOF {
			logger.Debug("Client disconnected", 
				zap.String("remote", remoteAddr),
				zap.Int("commands_processed", commandCount),
			)
			break
		}
		if err != nil {
			logger.Warn("Parse error", 
				zap.String("remote", remoteAddr),
				zap.Error(err),
			)
			break
		}

		if len(args) > 0 {
			cmdName := args[0]
			logger.Debug("Executing command",
				zap.String("remote", remoteAddr),
				zap.String("command", cmdName),
				zap.Strings("args", args[1:]),
			)

			err := registry.Execute(conn, args)
			duration := time.Since(start)
			commandCount++

			if err != nil {
				logger.Error("Command execution failed",
					zap.String("remote", remoteAddr),
					zap.String("command", cmdName),
					zap.Error(err),
					zap.Duration("duration", duration),
				)
			} else {
				logger.Debug("Command completed",
					zap.String("remote", remoteAddr),
					zap.String("command", cmdName),
					zap.Duration("duration", duration),
				)
			}
		}
	}
}
