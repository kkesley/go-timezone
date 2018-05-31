package timezone

import (
	"fmt"
	"testing"
)

func TestParseTime(test *testing.T) {
	time := ParseTime("5/1/2018", "Australia/Melbourne")
	if time == nil {
		test.Fail()
	}

	if year, month, date := (*time).Date(); year != 2018 || month.String() != "May" || date != 1 {
		test.Fail()
	}
}

func TestParseTimeUTC(test *testing.T) {
	time := ParseTimeUTC("5/1/2018", "Australia/Melbourne")
	if time == nil {
		test.Fail()
	}

	//Melbourne is UTC+10:00
	if year, month, date := (*time).Date(); year != 2018 || month.String() != "April" || date != 30 {
		test.Fail()
	}

	if name, offset := (*time).Zone(); name != "UTC" && offset != 0 {
		fmt.Println(name)
		fmt.Println(offset)
		test.Fail()
	}
}
