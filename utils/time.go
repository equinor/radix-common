package utils

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const defaultLayout = time.RFC3339

// ParseTimestamp Converts timestamp to time by RFC3339 layout
func ParseTimestamp(timestamp string) (time.Time, error) {
	return ParseTimestampBy(defaultLayout, timestamp)
}

// ParseTimestampBy Converts timestamp to time by custom layout
func ParseTimestampBy(layout, timestamp string) (time.Time, error) {
	if len(layout) == len(timestamp) {
		return time.Parse(layout, timestamp)
	}
	return time.Parse(defaultLayout, timestamp)
}

// ParseTimestampNillable Converts timestamp to time or nil
func ParseTimestampNillable(timestamp string) *time.Time {
	t, err := ParseTimestamp(timestamp)
	if err != nil {
		return nil
	}
	return &t
}

// FormatTimestamp Converts time to formatted timestamp
func FormatTimestamp(timestamp time.Time) string {
	emptyTime := time.Time{}

	if timestamp != emptyTime {
		return timestamp.Format(time.RFC3339)
	}

	return ""
}

// FormatTime Converts kubernetes time to formatted timestamp
func FormatTime(time *metav1.Time) string {
	if time != nil {
		return FormatTimestamp(time.Time)
	}

	return ""
}

type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (c RealClock) Now() time.Time {
	return time.Now()
}

type FakeClock struct {
	now time.Time
}

func NewFakeClock(now time.Time) *FakeClock {
	return &FakeClock{now: now}
}

func (c *FakeClock) Now() time.Time {
	return c.now
}
