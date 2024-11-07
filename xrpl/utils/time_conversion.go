package utils

import (
	"math"
	"time"
)

// The rippled server and its APIs represent time as an unsigned integer.
// This number measures the number of seconds since the "Ripple Epoch" of January 1, 2000 (00:00 UTC).
// This is like the way the Unix epoch works, except the Ripple Epoch is 946684800 seconds after the Unix Epoch.
const RIPPLE_EPOCH_DIFF = 946684800

// RippleTimeToUnixTime converts a ripple timestamp to a unix timestamp.
//
// rpepoch is the number of seconds since January 1, 2000 (00:00 UTC).
//
// It returns the number of milliseconds since the Unix epoch (January 1, 1970 00:00 UTC).
func RippleTimeToUnixTime(rpepoch uint64) uint64 {
	return (rpepoch + RIPPLE_EPOCH_DIFF) * 1000
}

// UnixTimeToRippleTime converts a unix timestamp to a ripple timestamp.
//
// timestamp is the number of milliseconds since the Unix epoch (January 1, 1970 00:00 UTC).
//
// It returns the number of seconds since the Ripple epoch (January 1, 2000 00:00 UTC).
func UnixTimeToRippleTime(timestamp uint64) uint64 {
	return uint64(math.Round(float64(timestamp)/1000) - RIPPLE_EPOCH_DIFF)
}

// RippleTimeToISOTime converts a ripple timestamp to an ISO 8601 formatted time string.
//
// rpepoch is the number of seconds since January 1, 2000 (00:00 UTC).
//
// It returns the time formatted as an ISO 8601 string.
func RippleTimeToISOTime(rippleTime uint64) string {
	unixTime := RippleTimeToUnixTime(rippleTime)
	return formatTimeToISO8601WithMillis(time.Unix(int64(unixTime/1000), 0).UTC()) //.Format(time.RFC3339)
}

// IsoTimeToRippleTime converts an ISO8601 timestmap to a ripple timestamp.
//
// iso8601 is the ISO 8601 formatted string.
//
// It returns the seconds since ripple epoch (1/1/2000 GMT).
func IsoTimeToRippleTime(iso8601 string) (uint64, error) {
	//const isoDate = time. //typeof iso8601 === 'string' ? new Date(iso8601) : iso8601

	t, err := parseISO8601(iso8601)
	if err != nil {
		return 0, err
	}
	return uint64(UnixTimeToRippleTime(uint64(t.UnixMilli()))), nil
}

// ParseISO8601 parses an ISO 8601 formatted string into a time.Time object.
//
// iso8601 is the ISO 8601 formatted string.
//
// It returns the parsed time.Time object and an error if the parsing fails.
func parseISO8601(iso8601 string) (time.Time, error) {
	return time.Parse(time.RFC3339, iso8601)
}

// FormatTimeToISO8601WithMillis formats a time.Time object to an ISO 8601 string with milliseconds.
//
// t is the time.Time object to be formatted.
//
// It returns the formatted ISO 8601 string with milliseconds.
func formatTimeToISO8601WithMillis(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z")
}
