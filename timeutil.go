package timeutil

import (
	"os"
	"time"
)

// Location represents time offset in use in a geographical area
var Location *time.Location

// Local converts UTC time to local (timezone is taken from "TZ" env var)
func Local(t time.Time) time.Time {
	if Location == nil {
		return t
	}
	return t.In(Location)
}

// Init initializes timeutil
func Init() error {
	tz := os.Getenv("TZ")
	if tz == "" {
		Location = time.Local
		return nil
	}

	l, err := time.LoadLocation(tz)
	if err != nil {
		return err
	}

	Location = l
	return nil
}

// MustInit is like Init but panics in case of error
func MustInit() {
	if err := Init(); err != nil {
		panic("timeutil.MustInit(): " + err.Error())
	}
}

func init() {
	MustInit()
}

// IsDay checks whether it is day
func IsDay() bool {
	if h := time.Now().Hour(); h >= 6 && h < 23 {
		return true
	}
	return false
}
