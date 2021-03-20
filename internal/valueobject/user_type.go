package valueobject

import (
	"strconv"

	"github.com/genki-sano/mm-server/internal/entity"
)

// UserType type
type UserType struct {
	value uint8
}

// NewUserType method
func NewUserType(value string) (*UserType, error) {
	if value == "" {
		return nil, newRequiredError("性別")
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return nil, newFormatError("性別")
	}
	ut := uint8(i)
	if ut != entity.UserTypeMan && ut != entity.UserTypeWoman {
		return nil, newContainItemsError("性別")
	}

	return &UserType{value: ut}, nil
}

// Get method
func (vo *UserType) Get() uint8 {
	return vo.value
}
