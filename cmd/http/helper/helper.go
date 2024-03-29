package helper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/genki-sano/mm-server/internal/presenter"
)

// ResposeData type
type ResposeData struct {
	Status int
	Body   []byte
}

// ErrorMessage type
type ErrorMessage struct {
	Error string `json:"error"`
}

// Response method
func Response(ctx *gin.Context, data ResposeData) {
	ctx.Data(data.Status, "application/json", data.Body)
}

// CreateSuccessResponseData method
func CreateSuccessResponseData(p presenter.I) ResposeData {
	res, err := p.Exec()
	if err != nil {
		return CreateErrorResponseData(http.StatusInternalServerError, err)
	}
	return ResposeData{
		Status: http.StatusOK,
		Body:   res,
	}
}

// CreateSuccessNoContentResponseData method
func CreateSuccessNoContentResponseData() ResposeData {
	return ResposeData{
		Status: http.StatusNoContent,
	}
}

// CreateErrorResponseData method
func CreateErrorResponseData(status int, err error) ResposeData {
	log.Printf("ERROR: %s", err.Error())
	return ResposeData{
		Status: status,
		Body:   CreateErrorMessage(err),
	}
}

// CreateErrorMessage method
func CreateErrorMessage(err error) []byte {
	res := ErrorMessage{
		Error: err.Error(),
	}
	body, err := json.Marshal(res)
	if err != nil {
		// TODO: エラーハンドリングを悩み中
		log.Fatalf(err.Error())
	}
	return body
}
