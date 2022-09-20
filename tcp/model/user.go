package model

type User struct {
	ID       uint32 `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Profile  string `json:"profile_picture"`
	Nickname string `json:"nickname"`
}
