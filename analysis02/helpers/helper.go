package helpers

import (
	"strconv"
	"time"
)

const (
	Day = iota
	Hour
	Min
	Second
)

func GetTime(t string, mod int) string {
	if len(t) == 0 {
		return ""
	}

	format := ""
	switch mod {
	case Day:
		format = "2006-01-02"
	case Hour:
		format = "2006-01-02 15"
	case Min:
		format = "2006-01-02 15:04"
	case Second:
		format = "2006-01-02 15:04:05"
	}
	if len(t) > 0 && len(t) >= len(format) {
		t = t[:len(format)]
	}
	pt, _ := time.Parse(format, t)
	return strconv.FormatInt(pt.Unix(), 10)
}
