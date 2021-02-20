package auth

import (
	"encoding/json"

	"github.com/genki-sano/mm-server/internal/presenter"
)

type verifyFactory struct{}

// NewVerifyFactory method
func NewVerifyFactory() presenter.AuthVerifyFactory {
	return &verifyFactory{}
}

type verifyPresenter struct {
	token string
}

// New method
func (f *verifyFactory) New(token string) presenter.I {
	return &verifyPresenter{
		token: token,
	}
}

// Exec method
func (p *verifyPresenter) Exec() ([]byte, error) {
	type verifyResponse struct {
		Token string `json:"token"`
	}

	res := verifyResponse{
		Token: p.token,
	}
	return json.Marshal(res)
}
