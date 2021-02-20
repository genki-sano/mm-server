package gateway

import (
	"context"
)

// FirebaseDataAccess type
type FirebaseDataAccess interface {
	CreateCustomToken(context.Context, string) (string, error)
}
