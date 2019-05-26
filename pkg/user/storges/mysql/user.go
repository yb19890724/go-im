package mysql

// 用户存储结构
type User struct {
	ID       int    `gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}
