package linebot

import (
	"encoding/json"
	"net/http"

	"github.com/genki-sano/mm-server/internal/errs"
	"github.com/genki-sano/mm-server/internal/gateway"
)

func isSuccess(code int) bool {
	return code/100 == 2
}

func checkResponse(res *http.Response) error {
	if isSuccess(res.StatusCode) {
		return nil
	}
	decoder := json.NewDecoder(res.Body)
	result := errs.LinebotErrorResponse{}
	if err := decoder.Decode(&result); err != nil {
		return errs.NewLinebotError(res.StatusCode, nil)
	}
	return errs.NewLinebotError(res.StatusCode, &result)
}

func decodeToVerifyTokenResponse(res *http.Response) (*gateway.VerifyTokenResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := gateway.VerifyTokenResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func decodeToUserProfileResponse(res *http.Response) (*gateway.UserProfileResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	result := gateway.UserProfileResponse{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
