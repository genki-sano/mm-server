package valueobject

import "time"

// Month type
type Month struct {
	value time.Time
}

// NewMonth method
func NewMonth(value string) (*Month, error) {
	if value == "" {
		return nil, newRequiredError("日付")
	}

	date, err := time.ParseInLocation("2006-01", value, time.Local)
	if err != nil {
		return nil, newFormatError("日付")
	}

	return &Month{value: date}, nil
}

// Get method
func (vo *Month) Get() time.Time {
	return vo.value
}
