package spreadsheet

import "time"

// @see https://github.com/tealeg/xlsx/blob/d8b3cff7d2c81b9ab460f33382cafc4203a70f47/date.go

const (
	MJD_0      float64 = 2400000.5
	MJD_JD2000 float64 = 51544.5

	secondsInADay = float64((24 * time.Hour) / time.Second)
	nanosInADay   = float64((24 * time.Hour) / time.Nanosecond)
)

var (
	timeLocationUTC, _ = time.LoadLocation("UTC")

	unixEpoc = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)

	excel1900Epoc = time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	excel1904Epoc = time.Date(1904, time.January, 1, 0, 0, 0, 0, time.UTC)

	daysBetween1970And1900 = float64(unixEpoc.Sub(excel1900Epoc) / (24 * time.Hour))
	daysBetween1970And1904 = float64(unixEpoc.Sub(excel1904Epoc) / (24 * time.Hour))
)

func TimeToExcelTime(t time.Time, date1904 bool) float64 {
	daysSinceUnixEpoc := float64(t.Unix()) / secondsInADay
	nanosPart := float64(t.Nanosecond()) / nanosInADay
	var offsetDays float64
	if date1904 {
		offsetDays = daysBetween1970And1904
	} else {
		offsetDays = daysBetween1970And1900
	}
	daysSinceExcelEpoc := daysSinceUnixEpoc + offsetDays + nanosPart
	return daysSinceExcelEpoc
}
