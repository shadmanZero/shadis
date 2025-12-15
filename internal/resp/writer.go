package resp

import (
	"fmt"
	"net"
	"strconv"
)

func WriteSimpleString(conn net.Conn, s string) error {
	_, err := conn.Write([]byte("+" + s + "\r\n"))
	return err
}

func WriteBulkString(conn net.Conn, s string) error {
	data := fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
	_, err := conn.Write([]byte(data))
	return err
}

func WriteError(conn net.Conn, msg string) error {
	_, err := conn.Write([]byte("-" + msg + "\r\n"))
	return err
}

func WriteInteger(conn net.Conn, n int) error {
	_, err := conn.Write([]byte(":" + strconv.Itoa(n) + "\r\n"))
	return err
}

func WriteNull(conn net.Conn) error {
	_, err := conn.Write(NULL)
	return err
}

func WriteArray(conn net.Conn, items []string) error {
	_, err := conn.Write([]byte("*" + strconv.Itoa(len(items)) + "\r\n"))
	if err != nil {
		return err
	}
	for _, item := range items {
		if err := WriteBulkString(conn, item); err != nil {
			return err
		}
	}
	return nil
}

func SimpleString(s string) []byte {
	return []byte("+" + s + "\r\n")
}

func BulkString(s string) []byte {
	return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(s), s))
}

func Error(msg string) []byte {
	return []byte("-" + msg + "\r\n")
}