package timezone

import (
	"time"
)

//ConvertTime based on timezone
func ConvertTime(originalTime time.Time, zone string) time.Time {
	location, err := time.LoadLocation(zone)
	if err == nil {
		return originalTime.In(location)
	}
	return originalTime
}
