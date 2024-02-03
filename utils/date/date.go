package date

import (
	"time"
)

func NowDateAndTimeBR() string {
	n := time.Now()
	/*
	 01 - month
	 02 - day
	 03 - hour
	 04 - minute
	 05 - seconds
	 06 - year
	*/
	return n.Format("02/01/2006 03:04:05")
}
