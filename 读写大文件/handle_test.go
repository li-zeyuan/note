package handle_test

import (
	"bufio"
	"errors"
	"io"
	"os"
	"testing"
)

func TestLineHandle(t *testing.T) {
	handleFun := func([]byte) {}
	filePath := ""
	err := LineHandle(filePath, handleFun)
	if err != nil {
		panic(err)
	}
}

func LineHandle(filePath string, handle func([]byte)) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := bufio.NewReader(f)

	for {
		line, isPrefix, err := buf.ReadLine()
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			return nil
		}
		// 行太长，buf溢出
		if isPrefix {
			return errors.New("line too long")
		}

		handle(line)
	}
}

func TestBufHandle(t *testing.T) {
	handleFun := func([]byte) {}
	filePath := ""
	err := BufHandle(filePath, handleFun)
	if err != nil {
		panic(err)
	}
}

func BufHandle(filePath string, handleFun func([]byte)) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			return nil
		}

		handleFun(buf[:n])
	}
}
