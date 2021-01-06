package main

import (
	"14-about-std-lib/alglib"
	"14-about-std-lib/cryptlib"
	"14-about-std-lib/iolib"
	"14-about-std-lib/mathlib"
	"14-about-std-lib/oslib"
	"14-about-std-lib/strlib"
	"14-about-std-lib/syslib"
	"14-about-std-lib/timelib"
	"fmt"
	"strings"
	"time"
)

func main() {
	iolib.CallReaderTest()
	iolib.ReadAtTest()
	iolib.WriteToTest()
	iolib.MoveSeekPosTest()
	iolib.AppendFileContentTest()
	iolib.WriteToStdoutTest()
	iolib.CopyDataTest()
	iolib.TeeCopyTest()
	iolib.ForMultiReadersTest()
	iolib.MultiWriterTest()
	iolib.CountWordsTest("this is a line.\n another line goes here") // 8
	oslib.OsFunctionTest()

	timelib.DateTimeTest()
	timelib.TimerAfterTest()
	timelib.TickerTest()
	timelib.TickTest()

	fmt.Println(strings.Index("都是NO.1中国人", "中国"))    // 10
	fmt.Println(strlib.Utf8Index("都是NO.1中国人", "中国")) // 6
	fmt.Println(cryptlib.GetStrMD5("123456"))

	syslib.GetCurrentPidTest()
	mathlib.CalcCosTest()

	alglib.SortNumberTest()
	alglib.SortPersonTest()

	time.Sleep(3 * time.Second)
}
