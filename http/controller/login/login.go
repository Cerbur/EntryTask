package login

import (
	"entrytask/http/bind"
	"entrytask/http/config"
	"entrytask/http/controller"
	"entrytask/rpc/function"
	"net/http"
)

type UserPostInput struct {
	Name     string `json:"name" pattern:"^[a-zA-Z0-9_-]{6,15}$"`
	Password string `json:"password" pattern:"^().{6,20}$"`
}

func POSTHandler(w http.ResponseWriter, r *http.Request) {
	// 参数绑定与校验
	input := UserPostInput{}
	if err := bind.Struct(r, &input); err != nil {
		controller.HTTPError(w, -403, err.Error())
		return
	}

	// 调用RPC
	loginReply, err := config.Fun.UserLogin(function.UserLoginArgs{
		Name:     input.Name,
		Password: input.Password,
	})
	if err != nil {
		controller.HTTPError(w, -501, err.Error())
		return
	}

	controller.HTTPSuccess(w, loginReply.Token)
}
