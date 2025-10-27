package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) // eg. 29.10.2004
	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)  // to get the timestap "week" distance far from the time "t"
	fmt.Println(week_from_now)

	// formatting times:
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("2006 01 02")
	fmt.Println(t, "=>", s)

	/* this will be the output

	   2025-10-27 19:24:06.247853105 +0530 IST m=+0.000012795
	   27.10.2025
	   2025-10-27 13:54:06.247931001 +0000 UTC
	   2025-10-27 19:24:06.247934417 +0530 IST m=+0.000094098
	   2025-11-03 13:54:06.247931001 +0000 UTC
	   27 Oct 25 13:54 UTC
	   Mon Oct 27 13:54:06 2025
	   27 Oct 2025 13:54
	   2025-10-27 13:54:06.247931001 +0000 UTC => 2025 10 27

	*/
}
