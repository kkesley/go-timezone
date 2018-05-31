package timezone

import (
	"time"

	"github.com/araddon/dateparse"
	"github.com/aws/aws-sdk-go/aws"
)

//ParseTimeUTC convert string to time to UTC
func ParseTimeUTC(strTime string, zone string) *time.Time {
	time := ParseTime(strTime, zone)
	if time != nil {
		time = aws.Time(time.UTC())
	}
	return time
}

//ParseTime convert string to time
func ParseTime(strTime string, zone string) *time.Time {
	location, err := time.LoadLocation(zone)
	if err == nil {
		time, err := dateparse.ParseIn(strTime, location)
		if err != nil {
			return nil
		}
		return aws.Time(time)
	}
	time, err := dateparse.ParseLocal(strTime)
	if err != nil {
		return nil
	}
	return aws.Time(time)
}
