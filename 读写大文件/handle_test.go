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

func BufWrite(t *testing.T) error {
	filePath := "./test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString("http://c.biancheng.net/golang/ \n")
	}

	//Flush将缓存的文件真正写入到文件中
	write.Flush()

	return nil
}

buffer写入的底层逻辑
