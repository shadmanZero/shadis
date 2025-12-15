package command

import (
	"net"

	"github.com/shadman/shadis/internal/resp"
	"github.com/shadman/shadis/internal/store"
)

type SetCommand struct {
	store *store.Store
}

func (s *SetCommand) Execute(conn net.Conn, args []string) error {
	if len(args) < 3 {
		return resp.WriteError(conn, "ERR wrong number of arguments for 'set' command")
	}

	s.store.Set(args[1], args[2])
	return resp.WriteSimpleString(conn, "OK")
}
