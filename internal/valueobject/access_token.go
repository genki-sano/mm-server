package valueobject

import (
	"errors"
)

// AccessToken type
type AccessToken struct {
	value string
}

// NewAccessToken method
func NewAccessToken(value string) (*AccessToken, error) {
	if value == "" {
		return nil, errors.New("アクセストークンは必須です。")
	}
	vo := &AccessToken{
		value: value,
	}
	return vo, nil
}

// Get method
func (vo *AccessToken) Get() string {
	return vo.value
}
