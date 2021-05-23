package date_utils

import (
	"time"
)

const (
	apiDataLayout = "2006-01-02T07:31:48Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetDateNowString() string {
	return GetNow().Format(apiDataLayout)
}
