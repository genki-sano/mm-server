package entity

const (
	// UserTypeWoman 女性
	UserTypeWoman uint8 = 0
	// UserTypeMan 男性
	UserTypeMan uint8 = 1
)

// User ユーザーのモデル
type User struct {
	ID             uint32
	Type           uint8
	Name           string
	LineUserID     *string
	FirebaseUserID *string
	Ats
}
