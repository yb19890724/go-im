package auth

type User struct {
	ID       int    `gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}
