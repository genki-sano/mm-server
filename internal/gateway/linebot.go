package gateway

import (
	"context"
)

// LineDataAccess type
type LineDataAccess interface {
	VerifyToken(context.Context, string) (*VerifyTokenResponse, error)
	GetProfile(context.Context, string) (*UserProfileResponse, error)
}

// VerifyTokenResponse type
type VerifyTokenResponse struct {
	Scope     string `json:"scope"`
	ClientID  string `json:"client_id"`
	ExpiresIn int    `json:"expores_in"`
}

// UserProfileResponse type
type UserProfileResponse struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}
