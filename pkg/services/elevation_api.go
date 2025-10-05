package services

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"googlemaps.github.io/maps"
)

type ElevationApiService struct {
	apiKey string
	client *maps.Client
}

func NewElevationApiService(apiKey string) *ElevationApiService {
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create Google Maps client: %v", err)
	}
	return &ElevationApiService{
		apiKey: apiKey,
		client: client,
	}
}

// ParseCoordinateString converts a coordinate string in format "51 deg 24' 4.44" N" to float64
func (s *ElevationApiService) ParseCoordinateString(coordStr string) (float64, error) {
	// Remove extra whitespace and normalize the string
	coordStr = strings.TrimSpace(coordStr)

	// Regular expression to match the format: degrees deg minutes' seconds" direction
	re := regexp.MustCompile(`^(\d+)\s*deg\s*(\d+)\s*'\s*([\d.]+)\s*"\s*([NSEW])$`)
	matches := re.FindStringSubmatch(coordStr)

	if len(matches) != 5 {
		return 0, fmt.Errorf("invalid coordinate format: %s", coordStr)
	}

	// Parse degrees, minutes, and seconds
	degrees, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid degrees: %s", matches[1])
	}

	minutes, err := strconv.ParseFloat(matches[2], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid minutes: %s", matches[2])
	}

	seconds, err := strconv.ParseFloat(matches[3], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid seconds: %s", matches[3])
	}

	// Calculate decimal degrees
	decimalDegrees := degrees + minutes/60.0 + seconds/3600.0

	// Apply direction (N/E = positive, S/W = negative)
	direction := strings.ToUpper(matches[4])
	if direction == "S" || direction == "W" {
		decimalDegrees = -decimalDegrees
	}

	return decimalDegrees, nil
}

// ParseDateTimeString converts a datetime string in format "2024:06:20 11:46:13.16Z" to time.Time
func (s *ElevationApiService) ParseDateTimeString(dateTimeStr string) (time.Time, error) {
	// Remove extra whitespace and normalize the string
	dateTimeStr = strings.TrimSpace(dateTimeStr)

	// Regular expression to match the format: YYYY:MM:DD HH:MM:SS.ssZ
	re := regexp.MustCompile(`^(\d{4}):(\d{2}):(\d{2})\s+(\d{2}):(\d{2}):(\d{2})\.(\d{1,2})Z$`)
	matches := re.FindStringSubmatch(dateTimeStr)

	if len(matches) != 8 {
		return time.Time{}, fmt.Errorf("invalid datetime format: %s", dateTimeStr)
	}

	// Parse year, month, day
	year, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid year: %s", matches[1])
	}

	month, err := strconv.Atoi(matches[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid month: %s", matches[2])
	}

	day, err := strconv.Atoi(matches[3])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid day: %s", matches[3])
	}

	// Parse hour, minute, second
	hour, err := strconv.Atoi(matches[4])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid hour: %s", matches[4])
	}

	minute, err := strconv.Atoi(matches[5])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid minute: %s", matches[5])
	}

	second, err := strconv.Atoi(matches[6])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid second: %s", matches[6])
	}

	// Parse fractional seconds (centiseconds)
	fractionalStr := matches[7]
	// Pad to 2 digits if only 1 digit provided
	if len(fractionalStr) == 1 {
		fractionalStr += "0"
	}

	fractional, err := strconv.Atoi(fractionalStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid fractional seconds: %s", matches[7])
	}

	// Convert centiseconds to nanoseconds (centiseconds * 10,000,000)
	nanoseconds := fractional * 10000000

	// Create time in UTC (Z indicates UTC)
	t := time.Date(year, time.Month(month), day, hour, minute, second, nanoseconds, time.UTC)

	return t, nil
}

func (s *ElevationApiService) GetElevation(ctx context.Context, latitude, longitude float64) ([]maps.ElevationResult, error) {
	elevation, err := s.client.Elevation(ctx, &maps.ElevationRequest{
		Locations: []maps.LatLng{
			{Lat: latitude, Lng: longitude},
		},
	})
	if err != nil {
		return nil, err
	}
	return elevation, nil
}

func (s *ElevationApiService) GetTimezone(ctx context.Context, latitude, longitude float64, timestamp time.Time) (*maps.TimezoneResult, error) {
	timezone, err := s.client.Timezone(ctx, &maps.TimezoneRequest{
		Location:  &maps.LatLng{Lat: latitude, Lng: longitude},
		Timestamp: timestamp,
	})
	if err != nil {
		return nil, err
	}
	return timezone, nil
}
