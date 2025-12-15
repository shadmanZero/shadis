package command

import (
	"errors"
	"net"

	"github.com/shadman/shadis/internal/resp"
)

type EchoCommand struct{}

func (e *EchoCommand) Execute(conn net.Conn, args []string) error {
	if conn == nil {
		return errors.New("no connection to execute commands")
	}
	if len(args) < 2 {
		return resp.WriteError(conn, "ERR wrong number of arguments for 'echo' command")
	}
	return resp.WriteBulkString(conn, args[1])
}
