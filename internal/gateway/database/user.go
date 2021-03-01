package database

import (
	"github.com/genki-sano/mm-server/internal/entity"
	"github.com/genki-sano/mm-server/internal/infarastructure/database"
)

// UserDataAccess interface
type UserDataAccess interface {
	FindAll() (*[]entity.User, error)
}

type userRepository struct {
	db database.SQLHandler
}

// NewUserRepository method
func NewUserRepository(
	db database.SQLHandler,
) UserDataAccess {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() (*[]entity.User, error) {
	db := r.db.Connect()

	var users []entity.User
	ret := db.Find(&users)
	if ret.Error != nil {
		return nil, ret.Error
	}

	return &users, nil
}
