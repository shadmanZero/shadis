package resp
const (
	
	TypeSimpleString = '+'
	TypeError        = '-'
	TypeInteger      = ':'
	TypeBulkString   = '$'
	TypeArray        = '*'
)

// Common responses
var (
	PONG = SimpleString("PONG")
	OK   = SimpleString("OK")
	NULL = []byte("$-1\r\n")
)