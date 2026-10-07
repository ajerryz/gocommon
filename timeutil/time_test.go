package timeutil_test

import (
	"gocommon/timeutil"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	utc := timeutil.ToUtc(now)
	t.Logf("utc: %v", utc)
}
