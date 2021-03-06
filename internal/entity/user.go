package entity

const (
	// UserTypeWoman 女性
	UserTypeWoman uint8 = 1
	// UserTypeMan 男性
	UserTypeMan uint8 = 2
)

// User ユーザーのモデル
type User struct {
	Type           uint8
	Name           string
	LineUserID     *string
	FirebaseUserID *string
}
