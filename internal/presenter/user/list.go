package user

import (
	"encoding/json"

	"github.com/genki-sano/mm-server/internal/entity"
	"github.com/genki-sano/mm-server/internal/presenter"
)

type listFactory struct{}

// NewListFactory method
func NewListFactory() presenter.UserListFactory {
	return &listFactory{}
}

type listPresenter struct {
	users *[]entity.User
}

// New method
func (f *listFactory) New(users *[]entity.User) presenter.I {
	return &listPresenter{
		users: users,
	}
}

// Exec method
func (p *listPresenter) Exec() ([]byte, error) {
	type userElement struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}
	type listResponse struct {
		Users []userElement `json:"users"`
	}

	items := []userElement{}
	for _, user := range *p.users {
		sex := "woman"
		if user.Type == entity.UserTypeMan {
			sex = "wan"
		}
		items = append(items, userElement{
			Type: sex,
			Name: user.Name,
		})
	}

	res := listResponse{
		Users: items,
	}
	return json.Marshal(res)
}
