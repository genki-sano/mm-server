package spreadsheet

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/genki-sano/mm-server/internal/entity"
	"github.com/genki-sano/mm-server/internal/gateway"
)

type userRepository struct {
	srv           *service
	spreadsheetID string
	ctx           context.Context
}

// NewUserRepository method
func NewUserRepository() gateway.UserDataAccess {
	ctx := context.Background()
	srv, err := newService(ctx)
	if err != nil {
		// TODO: エラーハンドリングを悩み中
		log.Fatalf(err.Error())
	}
	id := os.Getenv("GOOGLE_SPREDSHEET_ID")
	return &userRepository{
		srv:           srv,
		spreadsheetID: id,
		ctx:           ctx,
	}
}

func (r *userRepository) GetAll() ([]*entity.User, error) {
	readRange := "users!A2:F"
	valueRange, err := r.srv.get(r.ctx, r.spreadsheetID, readRange)
	if err != nil {
		return nil, err
	}

	users := make([]*entity.User, 0)
	if len(valueRange.Values) == 0 {
		return users, nil
	}

	for _, items := range valueRange.Values {
		utype, err := strconv.Atoi(items[0].(string))
		if err != nil {
			// TODO: エラーハンドリングを悩み中
			log.Fatalf(err.Error())
		}
		lUserID := items[2].(string)
		fUserID := items[3].(string)

		user := &entity.User{
			Type:           uint8(utype),
			Name:           items[1].(string),
			LineUserID:     &lUserID,
			FirebaseUserID: &fUserID,
		}
		users = append(users, user)
	}
	return users, nil
}
