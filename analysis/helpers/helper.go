package helpers

import (
	"strconv"
	"time"
)

const (
	Day  = "d"
	Hour = "h"
	Min  = "m"
)

func GetTime(originTime, format string) string {
	var fmtStr string

	switch format {
	case Day:
		fmtStr = "2006-01-02"
	case Hour:
		fmtStr = "2006-01-02 15"
	case Min:
		fmtStr = "2006-01-02 15:04"
	}
	if len(originTime) > 0 {
		originTime = originTime[:len(fmtStr)]
	}
	t, _ := time.Parse(fmtStr, originTime)
	return strconv.FormatInt(t.Unix(), 10)
}
