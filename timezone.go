package gocommon

import "time"

func SetTimeZoneUTC() {
	time.Local = time.UTC
}
