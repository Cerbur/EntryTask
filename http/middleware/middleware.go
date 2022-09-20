package middleware

import (
	"entrytask/http/config"
	"entrytask/http/controller"
	"entrytask/rpc/function"
	"net/http"
)

func AuthLoginRequired(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			controller.HTTPError(w, -301, "token 丢失")
			return
		}

		exist, err := config.Fun.UserCheckLogin(function.UserCheckLoginArgs{Token: token})
		if err != nil {
			controller.HTTPError(w, -301, err.Error())
			return
		}
		if !exist.Exist {
			controller.HTTPError(w, -301, "token 不存在")
			return
		}
		f(w, r)
	}
}
