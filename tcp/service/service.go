package service

import (
	"entrytask/rpc/function"
	"entrytask/tcp/service/user"
)

var S function.RPCFunction

func init() {
	S.UserLogin = user.Login
	S.UserGetProfile = user.GetProfile
	S.UserCheckLogin = user.CheckLogin
	S.UserUpdateProfile = user.UpdateProfile
	S.UserUpdateProfileNickname = user.UpdateProfileNickname
	S.UserUpdateProfilePicture = user.UpdateProfilePicture
}
