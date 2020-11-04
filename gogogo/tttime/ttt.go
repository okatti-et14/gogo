package ttt

import "time"

// systemLocation holds the location of asia tokyo.
var systemLocation *time.Location

func init() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	systemLocation = loc
}

// Now returns now time in systemlocation(="Asia/Tokyo")
var Now = func() time.Time {
	return time.Now().In(systemLocation)
}

// SetStubbedNow is stubbed function for UT
func SetStubbedNow(t string) {
	layout := "2006-01-02 15:04:05"
	stubbedNow, _ := time.ParseInLocation(layout, t, systemLocation)
	Now = func() time.Time { return stubbedNow }
}
