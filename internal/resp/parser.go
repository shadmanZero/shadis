package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

)

func readLine(reader *bufio.Reader) (string, error) {

	line, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(line, "\r\n"), nil
}

func parseBulkString(reader *bufio.Reader) (string, error) {
	lenStr, err := readLine(reader)
	if err != nil {
		return "", err
	}

	length, err := strconv.Atoi(lenStr)
	if err != nil {
		return "", err
	}

	if length == -1 {
		return "", nil
	}

	data := make([]byte, length+2)
	_, err = io.ReadFull(reader, data)
	if err != nil {
		return "", err
	}

	return string(data[:length]), nil
}
func parseArray(reader *bufio.Reader) ([]string, error) {
	countStr, err := readLine(reader)
	if err != nil {
		return nil, err
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return nil, err
	}

	if count == -1 {
		return nil, nil
	}
	args := make([]string, count)

	for i := 0 ; i < count ; i++ {

		prefix , err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		if prefix != '$' {
			return nil, fmt.Errorf("expected '$', got '%c'", prefix)
		}

		args[i], err = parseBulkString(reader)
		if err != nil {
			return nil, err
		}
	}
	return args, nil
}
func parseInline(reader *bufio.Reader) ([]string, error ) {

	line , err := readLine(reader)
	if err != nil {
		return nil,err
	}

	return strings.Fields(line) , nil
}

func Parse(reader *bufio.Reader) ([]string,error) {

	typeByte , err := reader.ReadByte()
	if err != nil {
		return nil,err
	}
	switch typeByte {

	case '*':
		return parseArray(reader)
	case '$':
		s , err := parseBulkString(reader)
		if err != nil {
			return nil,err
		}
		return []string{s} , nil
	default:
		reader.UnreadByte()
		return parseInline(reader)
	}
}