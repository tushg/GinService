# Step 10: Utils Package Setup

## Create String Utilities
Create `pkg/utils/string.go`:

```go
package utils

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strings"
	"unicode"
)

// IsEmpty checks if a string is empty or contains only whitespace
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotEmpty checks if a string is not empty and contains non-whitespace characters
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// Truncate truncates a string to the specified length
func Truncate(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength] + "..."
}

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ToTitle converts a string to title case
func ToTitle(s string) string {
	if IsEmpty(s) {
		return s
	}
	
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	if IsEmpty(s) {
		return s
	}
	
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteByte('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// ToCamelCase converts a string to camelCase
func ToCamelCase(s string) string {
	if IsEmpty(s) {
		return s
	}
	
	words := strings.Fields(strings.ReplaceAll(s, "_", " "))
	if len(words) == 0 {
		return s
	}
	
	result := strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		if len(words[i]) > 0 {
			runes := []rune(words[i])
			runes[0] = unicode.ToUpper(runes[0])
			result += string(runes)
		}
	}
	return result
}

// ToPascalCase converts a string to PascalCase
func ToPascalCase(s string) string {
	if IsEmpty(s) {
		return s
	}
	
	camel := ToCamelCase(s)
	if len(camel) > 0 {
		runes := []rune(camel)
		runes[0] = unicode.ToUpper(runes[0])
		return string(runes)
	}
	return camel
}

// GenerateRandomString generates a random string of specified length
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}

// IsValidEmail checks if a string is a valid email address
func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// IsValidUUID checks if a string is a valid UUID
func IsValidUUID(uuid string) bool {
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	return uuidRegex.MatchString(strings.ToLower(uuid))
}

// RemoveSpecialChars removes special characters from a string
func RemoveSpecialChars(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// CountWords counts the number of words in a string
func CountWords(s string) int {
	if IsEmpty(s) {
		return 0
	}
	return len(strings.Fields(s))
}

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
	if IsEmpty(s) {
		return true
	}
	
	cleaned := strings.ToLower(RemoveSpecialChars(s))
	reversed := Reverse(cleaned)
	return cleaned == reversed
}
```

## Create Time Utilities
Create `pkg/utils/time.go`:

```go
package utils

import (
	"fmt"
	"time"
)

// TimeFormat constants
const (
	TimeFormatISO8601     = "2006-01-02T15:04:05Z07:00"
	TimeFormatDate        = "2006-01-02"
	TimeFormatDateTime    = "2006-01-02 15:04:05"
	TimeFormatTime        = "15:04:05"
	TimeFormatRFC3339     = time.RFC3339
	TimeFormatRFC3339Nano = time.RFC3339Nano
)

// Now returns current UTC time
func Now() time.Time {
	return time.Now().UTC()
}

// NowUnix returns current UTC time as Unix timestamp
func NowUnix() int64 {
	return Now().Unix()
}

// NowUnixNano returns current UTC time as Unix nano timestamp
func NowUnixNano() int64 {
	return Now().UnixNano()
}

// ParseTime parses a time string using the specified format
func ParseTime(timeStr, format string) (time.Time, error) {
	return time.Parse(format, timeStr)
}

// ParseTimeISO8601 parses an ISO8601 time string
func ParseTimeISO8601(timeStr string) (time.Time, error) {
	return ParseTime(timeStr, TimeFormatISO8601)
}

// ParseTimeDate parses a date string (YYYY-MM-DD)
func ParseTimeDate(dateStr string) (time.Time, error) {
	return ParseTime(dateStr, TimeFormatDate)
}

// ParseTimeDateTime parses a date-time string (YYYY-MM-DD HH:MM:SS)
func ParseTimeDateTime(dateTimeStr string) (time.Time, error) {
	return ParseTime(dateTimeStr, TimeFormatDateTime)
}

// FormatTime formats a time using the specified format
func FormatTime(t time.Time, format string) string {
	return t.Format(format)
}

// FormatTimeISO8601 formats a time as ISO8601 string
func FormatTimeISO8601(t time.Time) string {
	return FormatTime(t, TimeFormatISO8601)
}

// FormatTimeDate formats a time as date string (YYYY-MM-DD)
func FormatTimeDate(t time.Time) string {
	return FormatTime(t, TimeFormatDate)
}

// FormatTimeDateTime formats a time as date-time string (YYYY-MM-DD HH:MM:SS)
func FormatTimeDateTime(t time.Time) string {
	return FormatTime(t, TimeFormatDateTime)
}

// FormatTimeRFC3339 formats a time as RFC3339 string
func FormatTimeRFC3339(t time.Time) string {
	return FormatTime(t, TimeFormatRFC3339)
}

// IsToday checks if a time is today
func IsToday(t time.Time) bool {
	now := Now()
	return t.Year() == now.Year() && t.YearDay() == now.YearDay()
}

// IsYesterday checks if a time is yesterday
func IsYesterday(t time.Time) bool {
	yesterday := Now().AddDate(0, 0, -1)
	return t.Year() == yesterday.Year() && t.YearDay() == yesterday.YearDay()
}

// IsThisWeek checks if a time is this week
func IsThisWeek(t time.Time) bool {
	now := Now()
	year, week := now.ISOWeek()
	tYear, tWeek := t.ISOWeek()
	return year == tYear && week == tWeek
}

// IsThisMonth checks if a time is this month
func IsThisMonth(t time.Time) bool {
	now := Now()
	return t.Year() == now.Year() && t.Month() == now.Month()
}

// IsThisYear checks if a time is this year
func IsThisYear(t time.Time) bool {
	return t.Year() == Now().Year()
}

// StartOfDay returns the start of the day for a given time
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the end of the day for a given time
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// StartOfWeek returns the start of the week (Monday) for a given time
func StartOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7 // Sunday becomes 7
	}
	daysToSubtract := weekday - 1
	return StartOfDay(t.AddDate(0, 0, -daysToSubtract))
}

// EndOfWeek returns the end of the week (Sunday) for a given time
func EndOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7 // Sunday becomes 7
	}
	daysToAdd := 7 - weekday
	return EndOfDay(t.AddDate(0, 0, daysToAdd))
}

// StartOfMonth returns the start of the month for a given time
func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns the end of the month for a given time
func EndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 999999999, t.Location())
}

// StartOfYear returns the start of the year for a given time
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear returns the end of the year for a given time
func EndOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, t.Location())
}

// AddDays adds the specified number of days to a time
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddMonths adds the specified number of months to a time
func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// AddYears adds the specified number of years to a time
func AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// DaysBetween returns the number of days between two times
func DaysBetween(t1, t2 time.Time) int {
	t1 = StartOfDay(t1)
	t2 = StartOfDay(t2)
	return int(t2.Sub(t1).Hours() / 24)
}

// HoursBetween returns the number of hours between two times
func HoursBetween(t1, t2 time.Time) int {
	return int(t2.Sub(t1).Hours())
}

// MinutesBetween returns the number of minutes between two times
func MinutesBetween(t1, t2 time.Time) int {
	return int(t2.Sub(t1).Minutes())
}

// SecondsBetween returns the number of seconds between two times
func SecondsBetween(t1, t2 time.Time) int {
	return int(t2.Sub(t1).Seconds())
}

// FormatDuration formats a duration in a human-readable format
func FormatDuration(d time.Duration) string {
	if d < time.Second {
		return fmt.Sprintf("%dms", d.Milliseconds())
	}
	if d < time.Minute {
		return fmt.Sprintf("%.1fs", d.Seconds())
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm %ds", int(d.Minutes()), int(d.Seconds())%60)
	}
	
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	
	if hours < 24 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	}
	
	days := int(d.Hours()) / 24
	hours = hours % 24
	return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
}

// Sleep sleeps for the specified duration
func Sleep(d time.Duration) {
	time.Sleep(d)
}

// SleepSeconds sleeps for the specified number of seconds
func SleepSeconds(seconds int) {
	Sleep(time.Duration(seconds) * time.Second)
}

// SleepMilliseconds sleeps for the specified number of milliseconds
func SleepMilliseconds(milliseconds int) {
	Sleep(time.Duration(milliseconds) * time.Millisecond)
}
```

## Create Utils Test Files
Create `pkg/utils/string_test.go`:

```go
package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty(""))
	assert.True(t, IsEmpty("   "))
	assert.False(t, IsEmpty("hello"))
	assert.False(t, IsEmpty(" hello "))
}

func TestTruncate(t *testing.T) {
	assert.Equal(t, "hello", Truncate("hello", 10))
	assert.Equal(t, "he...", Truncate("hello world", 2))
	assert.Equal(t, "hello...", Truncate("hello world", 5))
}

func TestReverse(t *testing.T) {
	assert.Equal(t, "olleh", Reverse("hello"))
	assert.Equal(t, "world", Reverse("dlrow"))
	assert.Equal(t, "", Reverse(""))
}

func TestToTitle(t *testing.T) {
	assert.Equal(t, "Hello World", ToTitle("hello world"))
	assert.Equal(t, "Hello", ToTitle("hello"))
	assert.Equal(t, "", ToTitle(""))
}

func TestToSnakeCase(t *testing.T) {
	assert.Equal(t, "hello_world", ToSnakeCase("HelloWorld"))
	assert.Equal(t, "user_id", ToSnakeCase("UserID"))
	assert.Equal(t, "api_key", ToSnakeCase("APIKey"))
}

func TestToCamelCase(t *testing.T) {
	assert.Equal(t, "helloWorld", ToCamelCase("hello_world"))
	assert.Equal(t, "userId", ToCamelCase("user_id"))
	assert.Equal(t, "apiKey", ToCamelCase("api_key"))
}

func TestToPascalCase(t *testing.T) {
	assert.Equal(t, "HelloWorld", ToPascalCase("hello_world"))
	assert.Equal(t, "UserId", ToCamelCase("user_id"))
	assert.Equal(t, "ApiKey", ToCamelCase("api_key"))
}

func TestIsValidEmail(t *testing.T) {
	assert.True(t, IsValidEmail("test@example.com"))
	assert.True(t, IsValidEmail("user.name@domain.co.uk"))
	assert.False(t, IsValidEmail("invalid-email"))
	assert.False(t, IsValidEmail("@domain.com"))
	assert.False(t, IsValidEmail("user@"))
}

func TestIsValidUUID(t *testing.T) {
	assert.True(t, IsValidUUID("550e8400-e29b-41d4-a716-446655440000"))
	assert.True(t, IsValidUUID("550E8400-E29B-41D4-A716-446655440000"))
	assert.False(t, IsValidUUID("invalid-uuid"))
	assert.False(t, IsValidUUID("550e8400-e29b-41d4-a716-44665544000"))
}

func TestRemoveSpecialChars(t *testing.T) {
	assert.Equal(t, "Hello World 123", RemoveSpecialChars("Hello, World! 123"))
	assert.Equal(t, "Test", RemoveSpecialChars("Test@#$%"))
	assert.Equal(t, "", RemoveSpecialChars("!@#$%"))
}

func TestCountWords(t *testing.T) {
	assert.Equal(t, 0, CountWords(""))
	assert.Equal(t, 1, CountWords("hello"))
	assert.Equal(t, 2, CountWords("hello world"))
	assert.Equal(t, 3, CountWords("hello   world   test"))
}

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome(""))
	assert.True(t, IsPalindrome("racecar"))
	assert.True(t, IsPalindrome("A man a plan a canal Panama"))
	assert.False(t, IsPalindrome("hello"))
}
```

Create `pkg/utils/time_test.go`:

```go
package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	now := Now()
	assert.NotZero(t, now)
	assert.Equal(t, time.UTC, now.Location())
}

func TestParseTimeDate(t *testing.T) {
	date, err := ParseTimeDate("2023-12-25")
	assert.NoError(t, err)
	assert.Equal(t, 2023, date.Year())
	assert.Equal(t, time.December, date.Month())
	assert.Equal(t, 25, date.Day())
}

func TestFormatTimeDate(t *testing.T) {
	t1 := time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)
	formatted := FormatTimeDate(t1)
	assert.Equal(t, "2023-12-25", formatted)
}

func TestIsToday(t *testing.T) {
	now := Now()
	assert.True(t, IsToday(now))
	
	yesterday := now.AddDate(0, 0, -1)
	assert.False(t, IsToday(yesterday))
}

func TestStartOfDay(t *testing.T) {
	t1 := time.Date(2023, 12, 25, 15, 30, 45, 123456789, time.UTC)
	start := StartOfDay(t1)
	
	assert.Equal(t, 0, start.Hour())
	assert.Equal(t, 0, start.Minute())
	assert.Equal(t, 0, start.Second())
	assert.Equal(t, 0, start.Nanosecond())
}

func TestEndOfDay(t *testing.T) {
	t1 := time.Date(2023, 12, 25, 15, 30, 45, 123456789, time.UTC)
	end := EndOfDay(t1)
	
	assert.Equal(t, 23, end.Hour())
	assert.Equal(t, 59, end.Minute())
	assert.Equal(t, 59, end.Second())
}

func TestDaysBetween(t *testing.T) {
	t1 := time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, 12, 27, 0, 0, 0, 0, time.UTC)
	
	days := DaysBetween(t1, t2)
	assert.Equal(t, 2, days)
}

func TestFormatDuration(t *testing.T) {
	d := 2*time.Hour + 30*time.Minute + 45*time.Second
	formatted := FormatDuration(d)
	assert.Equal(t, "2h 30m 45s", formatted)
}
```

## Verify Utils Package
```bash
# Check if utils files are created
ls -la pkg/utils/

# Expected output should show:
# string.go
# string_test.go
# time.go
# time_test.go
```

## Next Steps
After creating the utils package, proceed to the next file to create the business logic packages.
