package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"unicode"
)

// StringUtils provides common string manipulation functions
type StringUtils struct{}

// NewStringUtils creates a new StringUtils instance
func NewStringUtils() *StringUtils {
	return &StringUtils{}
}

// GenerateUUID generates a random UUID string
func (s *StringUtils) GenerateUUID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// GenerateRandomString generates a random string of specified length
func (s *StringUtils) GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	rand.Read(bytes)
	for i := range bytes {
		bytes[i] = charset[bytes[i]%byte(len(charset))]
	}
	return string(bytes)
}

// IsEmpty checks if a string is empty or contains only whitespace
func (s *StringUtils) IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

// ToTitleCase converts a string to title case
func (s *StringUtils) ToTitleCase(str string) string {
	if str == "" {
		return str
	}
	
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		if unicode.IsSpace(runes[i-1]) {
			runes[i] = unicode.ToUpper(runes[i])
		} else {
			runes[i] = unicode.ToLower(runes[i])
		}
	}
	return string(runes)
}

// Truncate truncates a string to the specified length
func (s *StringUtils) Truncate(str string, maxLength int) string {
	if len(str) <= maxLength {
		return str
	}
	return str[:maxLength] + "..."
}

// RemoveSpecialChars removes special characters from a string
func (s *StringUtils) RemoveSpecialChars(str string) string {
	var result strings.Builder
	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsNumber(char) || unicode.IsSpace(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

// Slugify converts a string to a URL-friendly slug
func (s *StringUtils) Slugify(str string) string {
	// Convert to lowercase
	str = strings.ToLower(str)
	
	// Replace spaces with hyphens
	str = strings.ReplaceAll(str, " ", "-")
	
	// Remove special characters
	str = s.RemoveSpecialChars(str)
	
	// Remove multiple consecutive hyphens
	for strings.Contains(str, "--") {
		str = strings.ReplaceAll(str, "--", "-")
	}
	
	// Remove leading and trailing hyphens
	str = strings.Trim(str, "-")
	
	return str
}
