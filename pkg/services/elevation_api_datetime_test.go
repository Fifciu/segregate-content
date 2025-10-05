package services

import (
	"testing"
	"time"
)

func TestParseDateTimeString(t *testing.T) {
	service := &ElevationApiService{}

	tests := []struct {
		name        string
		input       string
		expected    time.Time
		expectError bool
	}{
		{
			name:        "Valid datetime with 2-digit centiseconds",
			input:       "2024:06:20 11:46:13.16Z",
			expected:    time.Date(2024, 6, 20, 11, 46, 13, 160000000, time.UTC),
			expectError: false,
		},
		{
			name:        "Valid datetime with 1-digit centiseconds",
			input:       "2024:06:20 11:46:13.1Z",
			expected:    time.Date(2024, 6, 20, 11, 46, 13, 100000000, time.UTC),
			expectError: false,
		},
		{
			name:        "Valid datetime with zero centiseconds",
			input:       "2024:06:20 11:46:13.00Z",
			expected:    time.Date(2024, 6, 20, 11, 46, 13, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "Valid datetime at midnight",
			input:       "2024:01:01 00:00:00.00Z",
			expected:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "Valid datetime at end of day",
			input:       "2024:12:31 23:59:59.99Z",
			expected:    time.Date(2024, 12, 31, 23, 59, 59, 990000000, time.UTC),
			expectError: false,
		},
		{
			name:        "Valid datetime with extra whitespace",
			input:       "  2024:06:20 11:46:13.16Z  ",
			expected:    time.Date(2024, 6, 20, 11, 46, 13, 160000000, time.UTC),
			expectError: false,
		},
		{
			name:        "Valid datetime leap year",
			input:       "2024:02:29 12:00:00.00Z",
			expected:    time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC),
			expectError: false,
		},
		// Error cases
		{
			name:        "Invalid format - missing Z",
			input:       "2024:06:20 11:46:13.16",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - wrong separator",
			input:       "2024-06-20 11:46:13.16Z",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - missing space",
			input:       "2024:06:2011:46:13.16Z",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - missing fractional seconds",
			input:       "2024:06:20 11:46:13Z",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too many fractional digits",
			input:       "2024:06:20 11:46:13.123Z",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric year",
			input:       "abcd:06:20 11:46:13.16Z",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric month",
			input:       "2024:ab:20 11:46:13.16Z",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric day",
			input:       "2024:06:ab 11:46:13.16Z",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric hour",
			input:       "2024:06:20 ab:46:13.16Z",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric minute",
			input:       "2024:06:20 11:ab:13.16Z",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric second",
			input:       "2024:06:20 11:46:ab.16Z",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric fractional",
			input:       "2024:06:20 11:46:13.abZ",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Empty string",
			input:       "",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Whitespace only",
			input:       "   ",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - extra characters",
			input:       "2024:06:20 11:46:13.16Z extra",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - wrong timezone",
			input:       "2024:06:20 11:46:13.16+00:00",
			expected:    time.Time{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ParseDateTimeString(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("ParseDateTimeString(%q) expected error but got none", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("ParseDateTimeString(%q) unexpected error: %v", tt.input, err)
				return
			}

			// Compare times using Equal method for precise comparison
			if !result.Equal(tt.expected) {
				t.Errorf("ParseDateTimeString(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseDateTimeString_EdgeCases(t *testing.T) {
	service := &ElevationApiService{}

	t.Run("Maximum centiseconds", func(t *testing.T) {
		input := "2024:06:20 11:46:13.99Z"
		result, err := service.ParseDateTimeString(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 6, 20, 11, 46, 13, 990000000, time.UTC)
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Minimum centiseconds", func(t *testing.T) {
		input := "2024:06:20 11:46:13.00Z"
		result, err := service.ParseDateTimeString(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 6, 20, 11, 46, 13, 0, time.UTC)
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Single digit centiseconds padding", func(t *testing.T) {
		input := "2024:06:20 11:46:13.5Z"
		result, err := service.ParseDateTimeString(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 6, 20, 11, 46, 13, 500000000, time.UTC)
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestParseDateTimeString_RealWorldExamples(t *testing.T) {
	service := &ElevationApiService{}

	// Real world datetime examples for testing
	realWorldTests := []struct {
		name     string
		input    string
		expected time.Time
		location string
	}{
		{
			name:     "Photo taken in Warsaw",
			input:    "2024:06:20 11:46:13.16Z",
			expected: time.Date(2024, 6, 20, 11, 46, 13, 160000000, time.UTC),
			location: "Warsaw photo timestamp",
		},
		{
			name:     "Photo taken at midnight",
			input:    "2024:12:25 00:00:00.00Z",
			expected: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC),
			location: "Christmas midnight",
		},
		{
			name:     "Photo taken at noon",
			input:    "2024:07:04 12:00:00.50Z",
			expected: time.Date(2024, 7, 4, 12, 0, 0, 500000000, time.UTC),
			location: "Independence Day noon",
		},
		{
			name:     "Photo taken at end of year",
			input:    "2024:12:31 23:59:59.99Z",
			expected: time.Date(2024, 12, 31, 23, 59, 59, 990000000, time.UTC),
			location: "New Year's Eve",
		},
	}

	for _, tt := range realWorldTests {
		t.Run(tt.name+"_"+tt.location, func(t *testing.T) {
			result, err := service.ParseDateTimeString(tt.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if !result.Equal(tt.expected) {
				t.Errorf("ParseDateTimeString(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
