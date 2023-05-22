package entity

type User struct {
	Id       string `json:"id,omitempty" xorm:"pk"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Address  string `json:"address,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}
