package ctemplate

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"html/template"
)

func unescaped (s string) interface{} {
	return template.HTML(s)
}

func formatTimeDuring(s uint) interface{} {
	timeString := strconv.FormatUint(uint64(s), 10) + "s"
	duration, durationErr := time.ParseDuration(timeString)

	if durationErr != nil {
		return ""
	}

	timesFormat := fmt.Sprintf("%d:%d:%d", int(duration.Hours()), int(duration.Minutes()) % 60, int(duration.Seconds()) % 60)

	return strings.TrimLeft(timesFormat, "0:")
}