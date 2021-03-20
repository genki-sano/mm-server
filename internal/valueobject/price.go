package valueobject

import "strconv"

// Price type
type Price struct {
	value uint32
}

// NewPrice method
func NewPrice(value string) (*Price, error) {
	if value == "" {
		return nil, newRequiredError("価格")
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return nil, newFormatError("価格")
	}

	return &Price{value: uint32(i)}, nil
}

// Get method
func (vo *Price) Get() uint32 {
	return vo.value
}
