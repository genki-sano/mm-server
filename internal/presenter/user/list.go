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
	users []*entity.User
}

// New method
func (f *listFactory) New(users []*entity.User) presenter.I {
	return &listPresenter{
		users: users,
	}
}

// Exec method
func (p *listPresenter) Exec() ([]byte, error) {
	type listResponseItem struct {
		Type       uint8  `json:"type"`
		Name       string `json:"name"`
		AuthUserID string `json:"auth_user_id"`
	}
	type listResponse struct {
		Users []*listResponseItem `json:"users"`
	}

	items := make([]*listResponseItem, 0, len(p.users))
	for _, user := range p.users {
		item := &listResponseItem{
			Type:       user.Type,
			Name:       user.Name,
			AuthUserID: *user.LineUserID,
		}
		items = append(items, item)
	}
	resp := &listResponse{
		Users: items,
	}
	return json.Marshal(resp)
}
