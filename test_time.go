package main

import (
	"fmt"
	"time"
)

func main() {
	/*const shortForm = "2006-Jan-02"
	t, _ := time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	at, err := time.Parse("2006-01-02 15:04:05", "Feb 3, 2013 at 7:54pm (PST)")
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	nt := time.Now()
	fmt.Println(nt, at)
	if nt.After(at) {
		return
	}*/
	/*fmt.Println("", time.Now().Unix())

	t := time.Unix(1573700073, 0)
	// nt := t.Format("2006-01-02 11:04:05")
	nt := t.Format("2006-01-02 15:04:07")
	fmt.Println(nt)
	fmt.Println()*/

	/*const TimeFormat = "2006-01-02 15:04:05"
	t, err := time.Parse(TimeFormat, "2019-11-13 19:27:24")
	fmt.Println(t)
	fmt.Println("second:", t.Unix())

	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, err = time.ParseInLocation(TimeFormat, "2019-11-13 19:27:24", loc)
	fmt.Println(t)
	fmt.Println("second:", t.Unix())
	fmt.Println(time.Now().Format("2006-01-02"))

	_ = err*/

	/*loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now().Format("2006-01-02"))
	_ = loc*/

	/*const TimeFormat = "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(TimeFormat, "2019-11-13 19:27:24", loc)
	fmt.Println(t)

	fmt.Println(time.Now().Format("2006-01-02"))*/
	// fmt.Println(time.Now().Format("2006-01-02"))
	checkLastNotifyRebootTime("10:54:33")
}

func checkLastNotifyRebootTime(strRebootTime string) (bOfflineRebootNow bool) {
	const TimeFormat = "2006-01-02 15:04:05"

	fmt.Println(strRebootTime)
	strRebootTime = time.Now().Format("2006-01-02") + " " + strRebootTime
	fmt.Println(strRebootTime)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation(TimeFormat, strRebootTime, loc)
	fmt.Println(t.Unix())

	return false
}
