package valueobject

import "time"

// Date type
type Date struct {
	value time.Time
}

// NewDate method
func NewDate(value string) (*Date, error) {
	if value == "" {
		return nil, newRequiredError("日付")
	}

	date, err := time.ParseInLocation("2006-01-02", value, time.Local)
	if err != nil {
		return nil, newFormatError("日付")
	}

	return &Date{value: date}, nil
}

// Get method
func (vo *Date) Get() time.Time {
	return vo.value
}
