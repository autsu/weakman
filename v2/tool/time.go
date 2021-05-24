package tool

import "time"

// UtoB UTC to BJT
func UtoB(utcTime time.Time) time.Time {
	format := utcTime.Format("2006-01-02 15:04:05")
	location, _ := time.LoadLocation("Local")
	parseInLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", format, location)
	return parseInLocation
}
