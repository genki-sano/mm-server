package valueobject

// Memo type
type Memo struct {
	value string
}

// NewMemo method
func NewMemo(value string) (*Memo, error) {
	// if containMutibyte(value) {
	// 	return nil, newContainMutibyteError("メモ")
	// }

	return &Memo{value: value}, nil
}

// Get method
func (vo *Memo) Get() string {
	return vo.value
}
