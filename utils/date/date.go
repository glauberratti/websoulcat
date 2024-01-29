package date

import (
	"fmt"
	"time"
)

func NowDateAndTimeBR() string {
	n := time.Now()
	return fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d", n.Day(), n.Month(), n.Year(), n.Hour(), n.Minute(), n.Second())
}
