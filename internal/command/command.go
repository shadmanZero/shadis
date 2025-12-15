package command

import (
	"net"
	"strings"

	"github.com/shadman/shadis/internal/logger"
	"github.com/shadman/shadis/internal/resp"
	"github.com/shadman/shadis/internal/store"
	"go.uber.org/zap"
)

type Handler interface {
	Execute(conn net.Conn, args []string) error
}

type Command struct {
	Name    string
	Arity   int
	Handler Handler
}

type Registry struct {
	commands map[string]*Command
}

func NewRegistry() *Registry {
	return &Registry{
		commands: make(map[string]*Command),
	}
}

func (r *Registry) Register(cmd *Command) {
	r.commands[strings.ToUpper(cmd.Name)] = cmd
	logger.Debug("Command registered", zap.String("name", cmd.Name), zap.Int("arity", cmd.Arity))
}

func (r *Registry) Get(name string) (*Command, bool) {
	cmd, ok := r.commands[strings.ToUpper(name)]
	return cmd, ok
}

func (r *Registry) Count() int {
	return len(r.commands)
}

func (r *Registry) Execute(conn net.Conn, args []string) error {
	if len(args) == 0 {
		logger.Warn("Empty command received")
		return resp.WriteError(conn, "ERR no command")
	}

	cmdName := strings.ToUpper(args[0])
	cmd, ok := r.Get(cmdName)
	if !ok {
		logger.Warn("Unknown command", zap.String("command", cmdName))
		return resp.WriteError(conn, "ERR unknown command '"+args[0]+"'")
	}

	return cmd.Handler.Execute(conn, args)
}

// RegisterCommands registers all built-in commands
func RegisterCommands(r *Registry, s *store.Store) {
	r.Register(&Command{Name: "PING", Arity: -1, Handler: &PingCommand{}})
	r.Register(&Command{Name: "ECHO", Arity: 2, Handler: &EchoCommand{}})
	r.Register(&Command{Name: "GET", Arity: 2, Handler: &GetCommand{store: s}})
	r.Register(&Command{Name: "SET", Arity: 3, Handler: &SetCommand{store: s}})
}
