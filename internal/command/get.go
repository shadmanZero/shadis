package command

import (
	"net"

	"github.com/shadman/shadis/internal/resp"
	"github.com/shadman/shadis/internal/store"
)

type GetCommand struct {
	store *store.Store
}

func (g *GetCommand) Execute(conn net.Conn, args []string) error {
	if len(args) < 2 {
		return resp.WriteError(conn, "ERR wrong number of arguments for 'get' command")
	}

	val, ok := g.store.Get(args[1])
	if !ok {
		return resp.WriteNull(conn)
	}
	return resp.WriteBulkString(conn, val)
}
