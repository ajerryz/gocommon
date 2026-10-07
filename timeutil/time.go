package timeutil

import "time"

func ToUtc(t time.Time) string {
	return t.Format(time.RFC3339)
}
