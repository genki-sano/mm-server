package presenter

import "github.com/genki-sano/mm-server/internal/entity"

// UserListFactory type
type UserListFactory interface {
	New(*[]entity.User) I
}
