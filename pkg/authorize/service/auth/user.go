package auth

type User struct {
	ID       uint   `json:"id,omitempty"`        // 列名为 `id`
	Username string `json:"username"`  // 列名为 `username`
	Password string `json:"password"`  // 列名为 `password`
	Avatar   string `json:"avatar,omitempty"`    // 列名为  `avatar`
	Token    string `json:"token,omitempty"`     // 列名为  `token`
	Nickname string `json:"nickname,omitempty"` // 列名为  `token`
	Created  string `json:"created,omitempty"`
	Updated  string `json:"updated"`
}
