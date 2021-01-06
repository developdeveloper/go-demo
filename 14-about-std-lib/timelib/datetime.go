package timelib

import (
	"fmt"
	"time"
)

//DateTimeTest 日期时间
func DateTimeTest() {
	now := time.Now()
	fmt.Println(now)

	// 6-1-2-3-4-5 golang 的生日，这个可能很多人接受不了
	// 设计者认为比 Y-m-d H:i:s 好记 (1234567 => 01/02 03:04:05PM '06 -0700)
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	time.Sleep(1 * time.Second)
	now = time.Now()

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(time.Unix(now.Unix(), 0).Format("2006-01-02 15:04:05"))
	fmt.Println(now.Date())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Weekday())
	fmt.Println(now.AddDate(1, 0, 0))
	fmt.Println(now.Add(time.Duration(365 * 24 * 60 * 60 * 1000000000)))

	t1, _ := time.Parse("2006-01-02 15:04:05", "2020-12-31 23:59:59")
	d1 := now.Sub(t1)
	fmt.Println(d1) // 125h50m8.589151s
}

