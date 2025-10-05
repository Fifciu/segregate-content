package services

import (
	"testing"
)

func TestParseCoordinateString(t *testing.T) {
	service := &ElevationApiService{}

	tests := []struct {
		name        string
		input       string
		expected    float64
		expectError bool
	}{
		{
			name:        "Valid latitude North",
			input:       "51 deg 24' 4.44\" N",
			expected:    51.401233,
			expectError: false,
		},
		{
			name:        "Valid longitude East",
			input:       "21 deg 5' 39.98\" E",
			expected:    21.094439,
			expectError: false,
		},
		{
			name:        "Valid latitude South",
			input:       "40 deg 42' 51.36\" S",
			expected:    -40.714267,
			expectError: false,
		},
		{
			name:        "Valid longitude West",
			input:       "74 deg 0' 21.6\" W",
			expected:    -74.006000,
			expectError: false,
		},
		{
			name:        "Zero coordinates North",
			input:       "0 deg 0' 0.0\" N",
			expected:    0.0,
			expectError: false,
		},
		{
			name:        "Zero coordinates East",
			input:       "0 deg 0' 0.0\" E",
			expected:    0.0,
			expectError: false,
		},
		{
			name:        "Maximum latitude North",
			input:       "90 deg 0' 0.0\" N",
			expected:    90.0,
			expectError: false,
		},
		{
			name:        "Maximum latitude South",
			input:       "90 deg 0' 0.0\" S",
			expected:    -90.0,
			expectError: false,
		},
		{
			name:        "Maximum longitude East",
			input:       "180 deg 0' 0.0\" E",
			expected:    180.0,
			expectError: false,
		},
		{
			name:        "Maximum longitude West",
			input:       "180 deg 0' 0.0\" W",
			expected:    -180.0,
			expectError: false,
		},
		{
			name:        "With extra whitespace",
			input:       "  51  deg  24'  4.44\"  N  ",
			expected:    51.401233,
			expectError: false,
		},
		{
			name:        "Decimal seconds",
			input:       "45 deg 30' 15.5\" N",
			expected:    45.504306,
			expectError: false,
		},
		{
			name:        "Large minutes and seconds",
			input:       "30 deg 45' 30.0\" S",
			expected:    -30.758333,
			expectError: false,
		},
		// Error cases
		{
			name:        "Invalid format - missing deg",
			input:       "51 24' 4.44\" N",
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Invalid format - missing minutes",
			input:       "51 deg 4.44\" N",
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Invalid format - missing seconds",
			input:       "51 deg 24' N",
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Invalid format - missing direction",
			input:       "51 deg 24' 4.44\"",
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Invalid format - wrong direction",
			input:       "51 deg 24' 4.44\" X",
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric degrees",
			input:       "abc deg 24' 4.44\" N",
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric minutes",
			input:       "51 deg abc' 4.44\" N",
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Invalid format - non-numeric seconds",
			input:       "51 deg 24' abc\" N",
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Empty string",
			input:       "",
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Whitespace only",
			input:       "   ",
			expected:    0.0,
			expectError: true,
		},
		{
			name:        "Invalid format - extra characters",
			input:       "51 deg 24' 4.44\" N extra",
			expected:    0.0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ParseCoordinateString(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("ParseCoordinateString(%q) expected error but got none", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("ParseCoordinateString(%q) unexpected error: %v", tt.input, err)
				return
			}

			// Use a small tolerance for floating point comparison
			tolerance := 0.000001
			if abs(result-tt.expected) > tolerance {
				t.Errorf("ParseCoordinateString(%q) = %f, expected %f", tt.input, result, tt.expected)
			}
		})
	}
}

// Helper function for absolute value
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func TestParseCoordinateString_EdgeCases(t *testing.T) {
	service := &ElevationApiService{}

	t.Run("Very precise coordinates", func(t *testing.T) {
		input := "51 deg 24' 4.444444\" N"
		result, err := service.ParseCoordinateString(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := 51.401234
		tolerance := 0.000001
		if abs(result-expected) > tolerance {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})

	t.Run("Zero minutes and seconds", func(t *testing.T) {
		input := "45 deg 0' 0.0\" N"
		result, err := service.ParseCoordinateString(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != 45.0 {
			t.Errorf("Expected 45.0, got %f", result)
		}
	})

	t.Run("Zero seconds", func(t *testing.T) {
		input := "30 deg 30' 0.0\" S"
		result, err := service.ParseCoordinateString(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := -30.5
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	})
}

func TestParseCoordinateString_RealWorldExamples(t *testing.T) {
	service := &ElevationApiService{}

	// Real world coordinates for testing
	realWorldTests := []struct {
		name     string
		input    string
		expected float64
		location string
	}{
		{
			name:     "Warsaw, Poland",
			input:    "52 deg 13' 48.0\" N",
			expected: 52.230000,
			location: "Warsaw latitude",
		},
		{
			name:     "Warsaw, Poland",
			input:    "21 deg 0' 42.0\" E",
			expected: 21.011667,
			location: "Warsaw longitude",
		},
		{
			name:     "New York City",
			input:    "40 deg 42' 46.0\" N",
			expected: 40.712778,
			location: "NYC latitude",
		},
		{
			name:     "New York City",
			input:    "74 deg 0' 21.0\" W",
			expected: -74.005833,
			location: "NYC longitude",
		},
		{
			name:     "London, UK",
			input:    "51 deg 30' 26.0\" N",
			expected: 51.507222,
			location: "London latitude",
		},
		{
			name:     "London, UK",
			input:    "0 deg 7' 39.0\" W",
			expected: -0.127500,
			location: "London longitude",
		},
	}

	for _, tt := range realWorldTests {
		t.Run(tt.name+"_"+tt.location, func(t *testing.T) {
			result, err := service.ParseCoordinateString(tt.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			tolerance := 0.000001
			if abs(result-tt.expected) > tolerance {
				t.Errorf("ParseCoordinateString(%q) = %f, expected %f", tt.input, result, tt.expected)
			}
		})
	}
}
