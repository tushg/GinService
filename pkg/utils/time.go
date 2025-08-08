package utils

import (
	"fmt"
	"time"
)

// TimeUtils provides common time manipulation functions
type TimeUtils struct{}

// NewTimeUtils creates a new TimeUtils instance
func NewTimeUtils() *TimeUtils {
	return &TimeUtils{}
}

// FormatDateTime formats a time to a readable string
func (t *TimeUtils) FormatDateTime(tm time.Time) string {
	return tm.Format("2006-01-02 15:04:05")
}

// FormatDate formats a time to date only string
func (t *TimeUtils) FormatDate(tm time.Time) string {
	return tm.Format("2006-01-02")
}

// FormatTime formats a time to time only string
func (t *TimeUtils) FormatTime(tm time.Time) string {
	return tm.Format("15:04:05")
}

// ParseDateTime parses a datetime string
func (t *TimeUtils) ParseDateTime(dateTimeStr string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", dateTimeStr)
}

// ParseDate parses a date string
func (t *TimeUtils) ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// GetCurrentTimestamp returns current timestamp in seconds
func (t *TimeUtils) GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// GetCurrentTimestampMillis returns current timestamp in milliseconds
func (t *TimeUtils) GetCurrentTimestampMillis() int64 {
	return time.Now().UnixMilli()
}

// IsToday checks if the given time is today
func (t *TimeUtils) IsToday(tm time.Time) bool {
	now := time.Now()
	return tm.Year() == now.Year() && tm.YearDay() == now.YearDay()
}

// IsYesterday checks if the given time is yesterday
func (t *TimeUtils) IsYesterday(tm time.Time) bool {
	yesterday := time.Now().AddDate(0, 0, -1)
	return tm.Year() == yesterday.Year() && tm.YearDay() == yesterday.YearDay()
}

// IsThisWeek checks if the given time is in this week
func (t *TimeUtils) IsThisWeek(tm time.Time) bool {
	now := time.Now()
	year, week := now.ISOWeek()
	tmYear, tmWeek := tm.ISOWeek()
	return year == tmYear && week == tmWeek
}

// IsThisMonth checks if the given time is in this month
func (t *TimeUtils) IsThisMonth(tm time.Time) bool {
	now := time.Now()
	return tm.Year() == now.Year() && tm.Month() == now.Month()
}

// GetAge calculates age from birth date
func (t *TimeUtils) GetAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		age--
	}
	return age
}

// FormatDuration formats a duration to a human-readable string
func (t *TimeUtils) FormatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%.0fs", d.Seconds())
	}
	if d < time.Hour {
		return fmt.Sprintf("%.0fm", d.Minutes())
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("%.0fh", d.Hours())
	}
	return fmt.Sprintf("%.0fd", d.Hours()/24)
}

// GetStartOfDay returns the start of the day for the given time
func (t *TimeUtils) GetStartOfDay(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
}

// GetEndOfDay returns the end of the day for the given time
func (t *TimeUtils) GetEndOfDay(tm time.Time) time.Time {
	return time.Date(tm.Year(), tm.Month(), tm.Day(), 23, 59, 59, 999999999, tm.Location())
}

// GetStartOfWeek returns the start of the week (Monday) for the given time
func (t *TimeUtils) GetStartOfWeek(tm time.Time) time.Time {
	weekday := tm.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	} else {
		weekday--
	}
	return t.GetStartOfDay(tm.AddDate(0, 0, -int(weekday)))
}

// GetEndOfWeek returns the end of the week (Sunday) for the given time
func (t *TimeUtils) GetEndOfWeek(tm time.Time) time.Time {
	weekday := tm.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	} else {
		weekday--
	}
	return t.GetEndOfDay(tm.AddDate(0, 0, 7-int(weekday)))
}
