package iolib

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

//MyReader 自定义实现了 Reader 接口的对象
type MyReader struct{}

func (mr *MyReader) Read(p []byte) (n int, err error) {
	copy(p, []byte("hello"))
	return 5, nil
}

func readFromReaderInstance(reader io.Reader, size int) {
	buf := make([]byte, size)
	readed, err := reader.Read(buf)
	if err != nil {
		fmt.Printf("read err: %\n", err)
	}

	fmt.Printf("%d bytes readed, content: %s => %v\n", readed, string(buf[:readed]), buf[:readed])
}

//CallReaderTest 调用 reader 接口
func CallReaderTest() {
	fileReader, err := os.Open("/tmp/test.txt")
	defer fileReader.Close() // Close 里会判断是否为 nil，不需要放在 if 后面
	if err != nil {
		fmt.Printf("open file err: %s", err)
	}
	readFromReaderInstance(fileReader, 100)

	strReader := strings.NewReader("i am string")
	readFromReaderInstance(strReader, 100)

	myReader := &MyReader{}
	readFromReaderInstance(myReader, 100)

	// readFromReaderInstance(os.Stdin, 100)
}

//AppendFileContentTest 追加文件内容
func AppendFileContentTest() {
	file, err := os.Open("/tmp/test.txt")
	defer file.Close()
	if err != nil {
		fmt.Printf("open file err: %s", err)
	}

	buf := bytes.NewBufferString("hello world, from file:")
	buf.ReadFrom(file) // bytes.Buffer 实现了 ReadFrom 接口，而 *File 实现了 Reader 接口
	fmt.Println(buf.String())
}

//WriteToStdoutTest 显示文件内容
func WriteToStdoutTest() {
	file, err := os.Open("/tmp/test.txt")
	defer file.Close()
	if err != nil {
		fmt.Printf("open file err: %s", err)
	}

	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}
