package utilities

import (
	"testing"
	"time"
)

func TestParseDateTimeString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    time.Time
		expectError bool
	}{
		{
			name:        "Valid datetime",
			input:       "2024:07:03 11:04:46",
			expected:    time.Date(2024, 7, 3, 11, 4, 46, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "Valid datetime at midnight",
			input:       "2024:01:01 00:00:00",
			expected:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "Valid datetime at end of day",
			input:       "2024:12:31 23:59:59",
			expected:    time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "Valid datetime with extra whitespace",
			input:       "  2024:07:03 11:04:46  ",
			expected:    time.Date(2024, 7, 3, 11, 4, 46, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "Valid datetime leap year",
			input:       "2024:02:29 12:00:00",
			expected:    time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "Valid datetime with single digit month and day",
			input:       "2024:01:01 01:01:01",
			expected:    time.Date(2024, 1, 1, 1, 1, 1, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "Valid datetime with single digit hour, minute, second",
			input:       "2024:07:03 01:02:03",
			expected:    time.Date(2024, 7, 3, 1, 2, 3, 0, time.UTC),
			expectError: false,
		},
		// Error cases
		{
			name:        "Invalid format - missing space",
			input:       "2024:07:0311:04:46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - wrong separator",
			input:       "2024-07-03 11:04:46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - wrong time separator",
			input:       "2024:07:03 11-04-46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric year",
			input:       "abcd:07:03 11:04:46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric month",
			input:       "2024:ab:03 11:04:46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric day",
			input:       "2024:07:ab 11:04:46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric hour",
			input:       "2024:07:03 ab:04:46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric minute",
			input:       "2024:07:03 11:ab:46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric second",
			input:       "2024:07:03 11:04:ab",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short year",
			input:       "24:07:03 11:04:46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short month",
			input:       "2024:7:03 11:04:46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short day",
			input:       "2024:07:3 11:04:46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short hour",
			input:       "2024:07:03 1:04:46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short minute",
			input:       "2024:07:03 11:4:46",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short second",
			input:       "2024:07:03 11:04:6",
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
			input:       "2024:07:03 11:04:46 extra",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - missing components",
			input:       "2024:07:03 11:04",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too many components",
			input:       "2024:07:03 11:04:46:00",
			expected:    time.Time{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDateTimeString(tt.input)

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
	t.Run("Maximum valid values", func(t *testing.T) {
		input := "2024:12:31 23:59:59"
		result, err := ParseDateTimeString(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Minimum valid values", func(t *testing.T) {
		input := "2024:01:01 00:00:00"
		result, err := ParseDateTimeString(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Leap year February 29th", func(t *testing.T) {
		input := "2024:02:29 12:00:00"
		result, err := ParseDateTimeString(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC)
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Non-leap year February 28th", func(t *testing.T) {
		input := "2023:02:28 12:00:00"
		result, err := ParseDateTimeString(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2023, 2, 28, 12, 0, 0, 0, time.UTC)
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestParseDateTimeString_RealWorldExamples(t *testing.T) {
	// Real world datetime examples for testing
	realWorldTests := []struct {
		name     string
		input    string
		expected time.Time
		location string
	}{
		{
			name:     "Photo taken in July",
			input:    "2024:07:03 11:04:46",
			expected: time.Date(2024, 7, 3, 11, 4, 46, 0, time.UTC),
			location: "Summer photo timestamp",
		},
		{
			name:     "Photo taken at midnight",
			input:    "2024:12:25 00:00:00",
			expected: time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC),
			location: "Christmas midnight",
		},
		{
			name:     "Photo taken at noon",
			input:    "2024:07:04 12:00:00",
			expected: time.Date(2024, 7, 4, 12, 0, 0, 0, time.UTC),
			location: "Independence Day noon",
		},
		{
			name:     "Photo taken at end of year",
			input:    "2024:12:31 23:59:59",
			expected: time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
			location: "New Year's Eve",
		},
		{
			name:     "Photo taken in morning",
			input:    "2024:06:15 08:30:15",
			expected: time.Date(2024, 6, 15, 8, 30, 15, 0, time.UTC),
			location: "Morning photo",
		},
		{
			name:     "Photo taken in evening",
			input:    "2024:09:22 18:45:30",
			expected: time.Date(2024, 9, 22, 18, 45, 30, 0, time.UTC),
			location: "Evening photo",
		},
	}

	for _, tt := range realWorldTests {
		t.Run(tt.name+"_"+tt.location, func(t *testing.T) {
			result, err := ParseDateTimeString(tt.input)
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

func TestParseDateTimeStringWithTimezone(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    time.Time
		expectError bool
	}{
		// Valid positive timezone cases
		{
			name:        "Valid datetime with positive timezone",
			input:       "2024:07:22 11:36:53+02:00",
			expected:    time.Date(2024, 7, 22, 11, 36, 53, 0, time.FixedZone("", 2*60*60)),
			expectError: false,
		},
		{
			name:        "Valid datetime with positive timezone at midnight",
			input:       "2024:01:01 00:00:00+01:00",
			expected:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 1*60*60)),
			expectError: false,
		},
		{
			name:        "Valid datetime with positive timezone at end of day",
			input:       "2024:12:31 23:59:59+05:30",
			expected:    time.Date(2024, 12, 31, 23, 59, 59, 0, time.FixedZone("", 5*60*60+30*60)),
			expectError: false,
		},
		{
			name:        "Valid datetime with maximum positive timezone",
			input:       "2024:07:22 12:00:00+14:00",
			expected:    time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("", 14*60*60)),
			expectError: false,
		},
		{
			name:        "Valid datetime with single digit timezone components",
			input:       "2024:07:22 12:00:00+01:30",
			expected:    time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("", 1*60*60+30*60)),
			expectError: false,
		},
		// Valid negative timezone cases
		{
			name:        "Valid datetime with negative timezone",
			input:       "2024:07:22 11:36:53-02:00",
			expected:    time.Date(2024, 7, 22, 11, 36, 53, 0, time.FixedZone("", -2*60*60)),
			expectError: false,
		},
		{
			name:        "Valid datetime with negative timezone at midnight",
			input:       "2024:01:01 00:00:00-05:00",
			expected:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", -5*60*60)),
			expectError: false,
		},
		{
			name:        "Valid datetime with maximum negative timezone",
			input:       "2024:07:22 12:00:00-12:00",
			expected:    time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("", -12*60*60)),
			expectError: false,
		},
		{
			name:        "Valid datetime with negative timezone with minutes",
			input:       "2024:07:22 12:00:00-05:30",
			expected:    time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("", -5*60*60-30*60)),
			expectError: false,
		},
		// UTC timezone cases
		{
			name:        "Valid datetime with UTC timezone",
			input:       "2024:07:22 12:00:00+00:00",
			expected:    time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("", 0)),
			expectError: false,
		},
		{
			name:        "Valid datetime with UTC timezone (negative zero)",
			input:       "2024:07:22 12:00:00-00:00",
			expected:    time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("", 0)),
			expectError: false,
		},
		// Edge cases with whitespace
		{
			name:        "Valid datetime with extra whitespace",
			input:       "  2024:07:22 11:36:53+02:00  ",
			expected:    time.Date(2024, 7, 22, 11, 36, 53, 0, time.FixedZone("", 2*60*60)),
			expectError: false,
		},
		// Leap year cases
		{
			name:        "Valid leap year datetime with timezone",
			input:       "2024:02:29 12:00:00+01:00",
			expected:    time.Date(2024, 2, 29, 12, 0, 0, 0, time.FixedZone("", 1*60*60)),
			expectError: false,
		},
		// Error cases - invalid format
		{
			name:        "Invalid format - missing timezone",
			input:       "2024:07:22 11:36:53",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - missing space",
			input:       "2024:07:2211:36:53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - wrong date separator",
			input:       "2024-07-22 11:36:53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - wrong time separator",
			input:       "2024:07:22 11-36-53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - wrong timezone separator",
			input:       "2024:07:22 11:36:53+02-00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - missing timezone sign",
			input:       "2024:07:22 11:36:5302:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric year",
			input:       "abcd:07:22 11:36:53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric month",
			input:       "2024:ab:22 11:36:53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric day",
			input:       "2024:07:ab 11:36:53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric hour",
			input:       "2024:07:22 ab:36:53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric minute",
			input:       "2024:07:22 11:ab:53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric second",
			input:       "2024:07:22 11:36:ab+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric timezone hour",
			input:       "2024:07:22 11:36:53+ab:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric timezone minute",
			input:       "2024:07:22 11:36:53+02:ab",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short year",
			input:       "24:07:22 11:36:53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short month",
			input:       "2024:7:22 11:36:53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short day",
			input:       "2024:07:2 11:36:53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short hour",
			input:       "2024:07:22 1:36:53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short minute",
			input:       "2024:07:22 11:6:53+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short second",
			input:       "2024:07:22 11:36:3+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short timezone hour",
			input:       "2024:07:22 11:36:53+2:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too short timezone minute",
			input:       "2024:07:22 11:36:53+02:0",
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
			input:       "2024:07:22 11:36:53+02:00 extra",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - missing timezone components",
			input:       "2024:07:22 11:36:53+02",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too many timezone components",
			input:       "2024:07:22 11:36:53+02:00:30",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - missing datetime components",
			input:       "2024:07:22 11:36+02:00",
			expected:    time.Time{},
			expectError: true,
		},
		{
			name:        "Invalid format - too many datetime components",
			input:       "2024:07:22 11:36:53:00+02:00",
			expected:    time.Time{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDateTimeStringWithTimezone(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("ParseDateTimeStringWithTimezone(%q) expected error but got none", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("ParseDateTimeStringWithTimezone(%q) unexpected error: %v", tt.input, err)
				return
			}

			// Compare times using Equal method for precise comparison
			if !result.Equal(tt.expected) {
				t.Errorf("ParseDateTimeStringWithTimezone(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseDateTimeStringWithTimezone_EdgeCases(t *testing.T) {
	t.Run("Maximum positive timezone offset", func(t *testing.T) {
		input := "2024:07:22 12:00:00+14:00"
		result, err := ParseDateTimeStringWithTimezone(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("", 14*60*60))
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Maximum negative timezone offset", func(t *testing.T) {
		input := "2024:07:22 12:00:00-12:00"
		result, err := ParseDateTimeStringWithTimezone(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("", -12*60*60))
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("UTC timezone (positive zero)", func(t *testing.T) {
		input := "2024:07:22 12:00:00+00:00"
		result, err := ParseDateTimeStringWithTimezone(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("", 0))
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("UTC timezone (negative zero)", func(t *testing.T) {
		input := "2024:07:22 12:00:00-00:00"
		result, err := ParseDateTimeStringWithTimezone(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("", 0))
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Half-hour timezone offset", func(t *testing.T) {
		input := "2024:07:22 12:00:00+05:30"
		result, err := ParseDateTimeStringWithTimezone(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("", 5*60*60+30*60))
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Negative half-hour timezone offset", func(t *testing.T) {
		input := "2024:07:22 12:00:00-05:30"
		result, err := ParseDateTimeStringWithTimezone(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("", -5*60*60-30*60))
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Leap year February 29th with timezone", func(t *testing.T) {
		input := "2024:02:29 12:00:00+01:00"
		result, err := ParseDateTimeStringWithTimezone(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2024, 2, 29, 12, 0, 0, 0, time.FixedZone("", 1*60*60))
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Non-leap year February 28th with timezone", func(t *testing.T) {
		input := "2023:02:28 12:00:00+01:00"
		result, err := ParseDateTimeStringWithTimezone(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := time.Date(2023, 2, 28, 12, 0, 0, 0, time.FixedZone("", 1*60*60))
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestParseDateTimeStringWithTimezone_RealWorldExamples(t *testing.T) {
	// Real world datetime examples with timezones for testing
	realWorldTests := []struct {
		name     string
		input    string
		expected time.Time
		location string
	}{
		{
			name:     "Photo taken in Europe (CEST)",
			input:    "2024:07:22 11:36:53+02:00",
			expected: time.Date(2024, 7, 22, 11, 36, 53, 0, time.FixedZone("", 2*60*60)),
			location: "Central European Summer Time",
		},
		{
			name:     "Photo taken in New York (EDT)",
			input:    "2024:07:22 05:36:53-04:00",
			expected: time.Date(2024, 7, 22, 5, 36, 53, 0, time.FixedZone("", -4*60*60)),
			location: "Eastern Daylight Time",
		},
		{
			name:     "Photo taken in Tokyo (JST)",
			input:    "2024:07:22 18:36:53+09:00",
			expected: time.Date(2024, 7, 22, 18, 36, 53, 0, time.FixedZone("", 9*60*60)),
			location: "Japan Standard Time",
		},
		{
			name:     "Photo taken in India (IST)",
			input:    "2024:07:22 15:06:53+05:30",
			expected: time.Date(2024, 7, 22, 15, 6, 53, 0, time.FixedZone("", 5*60*60+30*60)),
			location: "India Standard Time",
		},
		{
			name:     "Photo taken in Australia (AEST)",
			input:    "2024:07:22 19:36:53+10:00",
			expected: time.Date(2024, 7, 22, 19, 36, 53, 0, time.FixedZone("", 10*60*60)),
			location: "Australian Eastern Standard Time",
		},
		{
			name:     "Photo taken in Hawaii (HST)",
			input:    "2024:07:22 01:36:53-10:00",
			expected: time.Date(2024, 7, 22, 1, 36, 53, 0, time.FixedZone("", -10*60*60)),
			location: "Hawaii Standard Time",
		},
		{
			name:     "Photo taken in London (GMT)",
			input:    "2024:07:22 09:36:53+00:00",
			expected: time.Date(2024, 7, 22, 9, 36, 53, 0, time.FixedZone("", 0)),
			location: "Greenwich Mean Time",
		},
		{
			name:     "Photo taken in Brazil (BRT)",
			input:    "2024:07:22 06:36:53-03:00",
			expected: time.Date(2024, 7, 22, 6, 36, 53, 0, time.FixedZone("", -3*60*60)),
			location: "Brasilia Time",
		},
		{
			name:     "Photo taken in Nepal (NPT)",
			input:    "2024:07:22 15:21:53+05:45",
			expected: time.Date(2024, 7, 22, 15, 21, 53, 0, time.FixedZone("", 5*60*60+45*60)),
			location: "Nepal Time",
		},
		{
			name:     "Photo taken in Iran (IRST)",
			input:    "2024:07:22 13:06:53+03:30",
			expected: time.Date(2024, 7, 22, 13, 6, 53, 0, time.FixedZone("", 3*60*60+30*60)),
			location: "Iran Standard Time",
		},
	}

	for _, tt := range realWorldTests {
		t.Run(tt.name+"_"+tt.location, func(t *testing.T) {
			result, err := ParseDateTimeStringWithTimezone(tt.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if !result.Equal(tt.expected) {
				t.Errorf("ParseDateTimeStringWithTimezone(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseDateTimeStringWithTimezone_TimezoneValidation(t *testing.T) {
	t.Run("Verify timezone offset calculation", func(t *testing.T) {
		// Test that the timezone offset is correctly calculated
		input := "2024:07:22 12:00:00+03:30"
		result, err := ParseDateTimeStringWithTimezone(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Check that the timezone offset is 3 hours 30 minutes = 210 minutes = 12600 seconds
		expectedOffset := 3*60*60 + 30*60 // 12600 seconds
		if result.Location().String() != time.FixedZone("", expectedOffset).String() {
			t.Errorf("Expected timezone offset %d seconds, got %s", expectedOffset, result.Location().String())
		}
	})

	t.Run("Verify negative timezone offset calculation", func(t *testing.T) {
		// Test that negative timezone offset is correctly calculated
		input := "2024:07:22 12:00:00-07:45"
		result, err := ParseDateTimeStringWithTimezone(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Check that the timezone offset is -7 hours 45 minutes = -465 minutes = -27900 seconds
		expectedOffset := -7*60*60 - 45*60 // -27900 seconds
		if result.Location().String() != time.FixedZone("", expectedOffset).String() {
			t.Errorf("Expected timezone offset %d seconds, got %s", expectedOffset, result.Location().String())
		}
	})

	t.Run("Verify UTC timezone handling", func(t *testing.T) {
		// Test that UTC timezone is handled correctly
		input := "2024:07:22 12:00:00+00:00"
		result, err := ParseDateTimeStringWithTimezone(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Check that the timezone offset is 0
		expectedOffset := 0
		if result.Location().String() != time.FixedZone("", expectedOffset).String() {
			t.Errorf("Expected timezone offset %d seconds, got %s", expectedOffset, result.Location().String())
		}
	})
}
