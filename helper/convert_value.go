package helper

import "time"

// SetPointerString ...
func SetPointerString(val string) *string {
	if val == "" {
		return nil
	}
	return &val
}

// SetPointerInt64 ...
func SetPointerInt64(val int64) *int64 {
	if val == 0 {
		return nil
	}
	return &val
}

// SetPointerTime ...
func SetPointerTime(val time.Time) *time.Time {
	return &val
}

// GetStringFromPointer ...
func GetStringFromPointer(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}

// GetInt64FromPointer ...
func GetInt64FromPointer(val *int64) int64 {
	if val == nil {
		return 0
	}
	return *val
}

// GetTimeFromPointer ...
func GetTimeFromPointer(val *time.Time) time.Time {
	return *val
}

// GetCurrentTimeUTC ...
func GetCurrentTimeUTC() (standartTime time.Time, unixTime int64) {
	current := time.Now().UTC()
	return current, current.Unix()
}

// FormatTimeFromString ...
func FormatTimeFromString(stringTime string) time.Time {
	location, _ := time.LoadLocation("Asia/Jakarta")
	parsedTime, _ := time.ParseInLocation("2006-01-02 15:04:05", stringTime, location)
	return parsedTime
}
