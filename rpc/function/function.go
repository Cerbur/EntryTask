package function

import "encoding/gob"

func init() {
	gob.Register(UserLoginArgs{})
	gob.Register(UserLoginReply{})
	gob.Register(UserGetProfileArgs{})
	gob.Register(UserGetProfileReply{})
	gob.Register(UserCheckLoginArgs{})
	gob.Register(UserCheckLoginReply{})
	gob.Register(UserUpdateProfileArgs{})
	gob.Register(UserUpdateProfileReply{})
}

type RPCFunction struct {
	UserLogin                 func(args UserLoginArgs) (UserLoginReply, error)                 // 用户登陆RPC函数
	UserCheckLogin            func(args UserCheckLoginArgs) (UserCheckLoginReply, error)       // 用户检查登陆状态
	UserGetProfile            func(args UserGetProfileArgs) (UserGetProfileReply, error)       // 用户获取个人信息函数
	UserUpdateProfile         func(args UserUpdateProfileArgs) (UserUpdateProfileReply, error) // 用户修改个人信息
	UserUpdateProfileNickname func(args UserUpdateProfileArgs) (UserUpdateProfileReply, error) // 用户修改个人信息
	UserUpdateProfilePicture  func(args UserUpdateProfileArgs) (UserUpdateProfileReply, error) // 用户修改个人信息
}

type UserLoginArgs struct {
	Name     string
	Password string
}

type UserLoginReply struct {
	Token string
}

type UserCheckLoginArgs struct {
	Token string
}

type UserCheckLoginReply struct {
	Exist bool
}

type UserGetProfileArgs struct {
	Token string
}

type UserGetProfileReply struct {
	Username string

	Nickname string

	ProfilePicture string
}

type UserUpdateProfileArgs struct {
	Token          string
	Nickname       string
	ProfilePicture string
}

type UserUpdateProfileReply struct {
	Nickname       string
	ProfilePicture string
}
