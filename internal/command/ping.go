package command

import (
	"errors"
	"net"

	"github.com/shadman/shadis/internal/resp"
)

type PingCommand struct{}

func (p *PingCommand) Execute(conn net.Conn, args []string) error {
	if conn == nil {
		return errors.New("no connection to execute commands")
	}
	return resp.WriteSimpleString(conn, "PONG")
}
