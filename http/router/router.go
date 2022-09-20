package router

import (
	"entrytask/http/controller/img"
	"entrytask/http/controller/login"
	"entrytask/http/controller/profile"
	"entrytask/http/middleware"
	"entrytask/http/server"
)

func GetHttpServer() *server.HttpServer {
	httpServer := server.NewHttpServer("127.0.0.1:8888")
	httpServer.POST("/api/login", login.POSTHandler)
	httpServer.GET("/api/img", img.GETHandler)
	httpServer.GET("/api/profile",
		middleware.AuthLoginRequired(profile.GETHandler))
	httpServer.PUT("/api/profile/nickname",
		middleware.AuthLoginRequired(profile.PUTNicknameHandler))
	httpServer.PUT("/api/profile/picture",
		middleware.AuthLoginRequired(profile.PUTProfilePictureHandler))
	return httpServer
}
