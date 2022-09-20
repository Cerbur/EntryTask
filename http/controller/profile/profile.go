package profile

import (
	"crypto/sha256"
	"encoding/hex"
	"entrytask/http/bind"
	"entrytask/http/config"
	"entrytask/http/controller"
	"entrytask/rpc/function"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type GetOutput struct {
	Username       string `json:"username"`
	Nickname       string `json:"nickname"`
	ProfilePicture string `json:"profile_picture"`
}

func GETHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	user, err := config.Fun.UserGetProfile(function.UserGetProfileArgs{Token: token})
	if err != nil {
		controller.HTTPError(w, -301, err.Error())
		return
	}

	controller.HTTPSuccess(w, GetOutput{
		Username:       user.Username,
		Nickname:       user.Nickname,
		ProfilePicture: user.ProfilePicture,
	})
}

type PutNicknameInput struct {
	Nickname string `json:"nickname" pattern:"^().{2,15}$"`
}

func PUTNicknameHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	// 参数绑定与校验
	input := PutNicknameInput{}
	if err := bind.Struct(r, &input); err != nil {
		controller.HTTPError(w, -403, err.Error())
		return
	}

	// RPC 调用处理
	reply, err := config.Fun.UserUpdateProfileNickname(function.UserUpdateProfileArgs{
		Token:    token,
		Nickname: input.Nickname,
	})
	if err != nil {
		controller.HTTPError(w, -301, err.Error())
		return
	}

	// 成功返回处理
	controller.HTTPSuccess(w, UpdateOutput{
		Nickname:       reply.Nickname,
		ProfilePicture: reply.ProfilePicture,
	})
}

func PUTProfilePictureHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	// 文件处理
	_, fileHead, err := r.FormFile("file")
	if err != nil {
		controller.HTTPError(w, -403, err.Error())
		return
	}
	file, err := fileHead.Open()
	if err != nil {
		controller.HTTPError(w, -502, err.Error())
		return
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(file)
	// 获取文件bytes
	byteFile, err := io.ReadAll(file)
	if err != nil {
		controller.HTTPError(w, -502, err.Error())
		return
	}
	// 获取文件后缀名
	split := strings.Split(fileHead.Filename, ".")
	fileType := split[len(split)-1]

	// 获取文件hash
	hash := sha256.New()
	hash.Write(byteFile)
	filename := fmt.Sprint(hex.EncodeToString(hash.Sum(nil)), ".", fileType)

	// 存储文件
	dst := fmt.Sprint(config.Path, filename)
	out, err := os.Create(dst)
	if err != nil {
		controller.HTTPError(w, -502, err.Error())
		return
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)
	_, err = out.Write(byteFile)
	if err != nil {
		return
	}

	// RPC 调用处理
	reply, err := config.Fun.UserUpdateProfilePicture(function.UserUpdateProfileArgs{
		Token:          token,
		ProfilePicture: filename,
	})
	if err != nil {
		controller.HTTPError(w, -501, err.Error())
		return
	}

	controller.HTTPSuccess(w, UpdateOutput{
		Nickname:       reply.Nickname,
		ProfilePicture: reply.ProfilePicture,
	})
}

type UpdateOutput struct {
	Nickname       string `json:"nickname"`
	ProfilePicture string `json:"profile_picture"`
}
