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

func formatTimeDuring(s int) interface{} {
	timeString := strconv.FormatInt(int64(s), 10) + "s"
	duration, durationErr := time.ParseDuration(timeString)

	if durationErr != nil {
		return ""
	}

	timesFormat := fmt.Sprintf("%d:%d:%d", int(duration.Hours()), int(duration.Minutes()) % 60, int(duration.Seconds()) % 60)

	return strings.TrimLeft(timesFormat, "0:")
}

func formatTimeString(s string) interface{} {
	parsedTime, parseTimeErr := time.Parse(time.RFC3339, s)

	if parseTimeErr != nil {
		return ""
	}

	timesFormat := fmt.Sprintf("%d/%d/%d %d:%d:%d", parsedTime.Year(), parsedTime.Month() + 1, parsedTime.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second())

	return strings.TrimLeft(timesFormat, "0:")
}