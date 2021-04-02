package payment

import (
	"context"
	"errors"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/genki-sano/mm-server/internal/entity"
	"github.com/genki-sano/mm-server/internal/gateway"
	"github.com/genki-sano/mm-server/internal/usecase"
	"github.com/genki-sano/mm-server/internal/valueobject"
)

type pushMessageInteractor struct {
	userRepos gateway.UserDataAccess
}

type createInteractor struct {
	pushMessageInteractor
	paymentRepos gateway.PaymentDataAccess
	linebot      gateway.LineDataAccess
}

type formatUser struct {
	Name  string
	Price string
}

// NewCreateUsecase method
func NewCreateUsecase(
	userRepos gateway.UserDataAccess,
	paymentRepos gateway.PaymentDataAccess,
	linebot gateway.LineDataAccess,
) usecase.PaytmentCreateUsecase {
	return &createInteractor{
		pushMessageInteractor: pushMessageInteractor{
			userRepos: userRepos,
		},
		paymentRepos: paymentRepos,
		linebot:      linebot,
	}
}

// Handle method
func (i *createInteractor) Handle(
	ctx context.Context,
	userType *valueobject.UserType,
	category *valueobject.Category,
	price *valueobject.Price,
	date *valueobject.Date,
	memo *valueobject.Memo,
) error {
	payment := &entity.Payment{
		ID:       0, // 仮の値を入れておく
		UserType: userType.Get(),
		Category: category.Get(),
		Price:    price.Get(),
		Date:     date.Get(),
		Memo:     memo.Get(),
	}
	if err := i.paymentRepos.Insert(payment); err != nil {
		return err
	}

	payments, err := i.paymentRepos.GetByDate(date.Get())
	if err != nil {
		return err
	}

	var womanTotal uint32
	var manTotal uint32

	for _, payment := range payments {
		if payment.UserType == entity.UserTypeWoman {
			womanTotal += payment.Price
		}
		if payment.UserType == entity.UserTypeMan {
			manTotal += payment.Price
		}
	}

	users, err := i.pushMessageInteractor.getAllUsers()
	if err != nil {
		return err
	}

	var woman *formatUser
	var man *formatUser

	var womanUid string
	var manUid string

	for _, user := range users {
		if user.Type == entity.UserTypeWoman {
			woman = &formatUser{
				Name:  user.Name,
				Price: humanize.Comma(int64(womanTotal)) + "円",
			}
			womanUid = *user.LineUserID
		}
		if user.Type == entity.UserTypeMan {
			man = &formatUser{
				Name:  user.Name,
				Price: humanize.Comma(int64(manTotal)) + "円",
			}
			manUid = *user.LineUserID
		}
	}

	alt := "支払い登録が完了したよ！"
	data := i.pushMessageInteractor.getFlexContainer(price.Get(), date.Get(), category.Get(), memo.Get(), *woman, *man)

	if err := i.linebot.PushFlexMessage(womanUid, alt, data); err != nil {
		return err
	}
	if err := i.linebot.PushFlexMessage(manUid, alt, data); err != nil {
		return err
	}

	return nil
}

func (i *pushMessageInteractor) getAllUsers() ([]*entity.User, error) {
	users, err := i.userRepos.GetAll()
	if err != nil {
		return nil, err
	}
	if len(users) != 2 {
		return nil, errors.New("不正なユーザーが登録されています。")
	}
	isErr := false
	for index, user := range users {
		if user.Type != uint8(index+1) {
			isErr = true
		}
	}
	if isErr {
		return nil, errors.New("必要なユーザーが登録されていません。")
	}
	return users, nil
}

func (i *pushMessageInteractor) getFlexContainer(price uint32, date time.Time, category, memo string, woman, man formatUser) []byte {
	return []byte(`
    {
      "type": "bubble",
      "body": {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": "支出を入力したよ！",
            "size": "sm",
            "margin": "md",
            "color": "#aaaaaa"
          },
          {
            "type": "text",
            "text": "` + humanize.Comma(int64(price)) + "円" + `",
            "size": "3xl",
            "weight": "bold"
          },
          {
            "type": "box",
            "layout": "vertical",
            "contents": [
              {
                "type": "box",
                "layout": "horizontal",
                "contents": [
                  {
                    "type": "text",
                    "text": "日付",
                    "color": "#aaaaaa",
                    "flex": 2,
                    "size": "sm"
                  },
                  {
                    "type": "text",
                    "text": "` + date.Format("2006/01/02") + `",
                    "flex": 5,
                    "size": "sm"
                  }
                ]
              },
              {
                "type": "box",
                "layout": "horizontal",
                "contents": [
                  {
                    "type": "text",
                    "text": "カテゴリ",
                    "color": "#aaaaaa",
                    "flex": 2,
                    "size": "sm"
                  },
                  {
                    "type": "text",
                    "text": "` + category + `",
                    "flex": 5,
                    "size": "sm"
                  }
                ]
              },
              {
                "type": "box",
                "layout": "horizontal",
                "contents": [
                  {
                    "type": "text",
                    "text": "メモ",
                    "color": "#aaaaaa",
                    "flex": 2,
                    "size": "sm"
                  },
                  {
                    "type": "text",
                    "text": "` + memo + `",
                    "flex": 5,
                    "size": "sm"
                  }
                ]
              }
            ],
            "margin": "lg",
            "spacing": "xs"
          },
          {
            "type": "separator",
            "margin": "xxl"
          },
          {
            "type": "box",
            "layout": "vertical",
            "contents": [
              {
                "type": "box",
                "layout": "horizontal",
                "contents": [
                  {
                    "type": "text",
                    "text": "` + woman.Name + `",
                    "flex": 0,
                    "color": "#aaaaaa"
                  },
                  {
                    "type": "text",
                    "text": "` + woman.Price + `",
                    "align": "end"
                  }
                ]
              },
              {
                "type": "box",
                "layout": "horizontal",
                "contents": [
                  {
                    "type": "text",
                    "text": "` + man.Name + `",
                    "flex": 0,
                    "color": "#aaaaaa"
                  },
                  {
                    "type": "text",
                    "text": "` + man.Price + `",
                    "align": "end"
                  }
                ]
              }
            ],
            "margin": "xxl",
            "spacing": "sm"
          }
        ]
      }
    }
  `)
}
