package timezone

import (
	"testing"
	"time"
)

func TestConvertTime(test *testing.T) {
	time := time.Now().UTC()
	convertedTime := ConvertTime(time, "Australia/Melbourne")
	if name, _ := convertedTime.Zone(); name != "AEST" {
		test.Fail()
	}
}
