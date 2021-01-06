package iolib

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

//ReadAtTest 指定位置读取
func ReadAtTest() {
	reader := strings.NewReader("NO.1中国人")
	p := make([]byte, 3)        // 这里一个中文占3个字节，如有是 4、5 会有乱码
	n, _ := reader.ReadAt(p, 4) // 从偏移位置 4 开始读取 3 个字节
	fmt.Printf("%s, %d\n", p, n)
}

//WriteToTest 打印到屏幕
func WriteToTest() {
	reader := strings.NewReader("NO.1中国人\n")
	reader.WriteTo(os.Stdout)
}

//MoveSeekPosTest 调整读取位置
func MoveSeekPosTest() {
	reader := strings.NewReader("中国人第一")
	reader.Seek(-9, io.SeekEnd) // 注意 UTF-8 编码，相当于 3*X 个字节移动
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r) // 人
	reader.Seek(6, io.SeekStart)
	r, _, _ = reader.ReadRune()
	fmt.Printf("%c\n", r) // 人
}

//CopyDataTest 打印
func CopyDataTest() {
	io.Copy(os.Stdout, strings.NewReader("天气真好\n"))
}

//TeeCopyTest 自动写
func TeeCopyTest() {
	reader := io.TeeReader(strings.NewReader("天气确实好\n"), os.Stdout)
	reader.Read(make([]byte, 20))
}

//ForMultiReadersTest 合并多个 reader
func ForMultiReadersTest() {
	readers := []io.Reader{
		strings.NewReader("part 1"),
		// bytes.NewReader([]byte("from bytes buffer")),
		bytes.NewBufferString(" & part 2"),
	}

	result := make([]byte, 256)
	buf := make([]byte, 32)
	multi := io.MultiReader(readers...)

	for {
		size, err := multi.Read(buf)
		if err == nil || err == io.EOF {
			result = append(result, buf[:size]...)
			if err == io.EOF {
				fmt.Println(string(result))
				break
			}
		} else {
			panic(err)
		}
	}
}

//MultiWriterTest 同时写入多个 writer
func MultiWriterTest() {
	file, err := os.OpenFile("/tmp/test.txt", os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	defer file.Close()
	if err != nil {
		panic(err)
	}

	writers := []io.Writer{file, os.Stdout}
	multi := io.MultiWriter(writers...)
	multi.Write([]byte("go demo test\n")) // 会同时打印出来
}
