package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetNowString() string {

	return GetNow().Format(apiDateLayout)
}
func GetNowForMySQL() string {
	return time.Now().Format(time.RFC3339)
}
func GetNow() time.Time {
	return time.Now().UTC()
}
