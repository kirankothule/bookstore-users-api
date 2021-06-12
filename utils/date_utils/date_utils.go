package date_utils

import (
	"time"
)

const (
	apiDataLayout = "2006-01-02T07:31:48Z"
	apiDBLayout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetDateNowString() string {
	return GetNow().Format(apiDataLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}
