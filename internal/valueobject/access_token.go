package valueobject

// AccessToken type
type AccessToken struct {
	value string
}

// NewAccessToken method
func NewAccessToken(value string) (*AccessToken, error) {
	if value == "" {
		return nil, newRequiredError("アクセストークン")
	}
	if containMutibyte(value) {
		return nil, newContainMutibyteError("アクセストークン")
	}

	return &AccessToken{value: value}, nil
}

// Get method
func (vo *AccessToken) Get() string {
	return vo.value
}
