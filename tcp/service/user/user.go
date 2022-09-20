package user

import (
	"entrytask/rpc/function"
	userDao "entrytask/tcp/dao/user"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

func Login(args function.UserLoginArgs) (function.UserLoginReply, error) {
	// 查数据库
	now := time.Now().UnixNano()
	user := userDao.FindUserAllByUserNameAndPassword(args.Name, args.Password)
	end := time.Now().UnixNano()
	dtime := float64(end-now) / float64(time.Millisecond)
	fmt.Printf("[%f]: FindUserAllByUserNameAndPassword(%s,%s)\n", dtime, args.Name, args.Password)

	// 判断是否存在用户
	if user.ID == 0 {
		return function.UserLoginReply{
			Token: "",
		}, errors.New("id is not find")
	}

	// 生成UUID
	newUUID, err := uuid.NewUUID()
	if err != nil {
		log.Println(err)
		return function.UserLoginReply{}, err
	}
	token := strings.ReplaceAll(newUUID.String(), "-", "")
	// 写redis 绑定登录态
	if err := userDao.SaveUUIDWithId(user.ID, token); err != nil {
		log.Println(err)
		return function.UserLoginReply{}, err
	}

	// 写数据缓存忽略错误问题不大
	_ = userDao.SaveCacheWithUserById(user, user.ID)

	// 返回
	return function.UserLoginReply{
		Token: token,
	}, nil
}

func GetProfile(args function.UserGetProfileArgs) (function.UserGetProfileReply, error) {
	// redis中查询token是否存在
	id, err := userDao.GetUserIdByUUID(args.Token)
	if err != nil || id == 0 {
		return function.UserGetProfileReply{}, err
	}

	// 读取缓存
	user, err := userDao.GetUserByUserById(id)

	if err != nil || user.ID == 0 {
		// 缓存读取失败 读数据库
		user = userDao.FindUserAllById(id)
		// 读数据库也失败了
		if user.ID == 0 {
			return function.UserGetProfileReply{}, err
		}
		_ = userDao.SaveCacheWithUserById(user, user.ID)
	}
	// 从redis中获取到信息
	return function.UserGetProfileReply{
		Username:       user.UserName,
		Nickname:       user.Nickname,
		ProfilePicture: user.Profile,
	}, nil
}

func CheckLogin(args function.UserCheckLoginArgs) (function.UserCheckLoginReply, error) {
	exist, err := userDao.CheckUUIDExist(args.Token)
	if err != nil {
		return function.UserCheckLoginReply{}, err
	}
	return function.UserCheckLoginReply{
		Exist: exist,
	}, nil
}

func UpdateProfile(args function.UserUpdateProfileArgs) (function.UserUpdateProfileReply, error) {
	userId, err := userDao.GetUserIdByUUID(args.Token)
	if err != nil {
		return function.UserUpdateProfileReply{}, err
	}
	// 可能有bug删除缓存
	if userDao.RemoveCacheWithUserByID(userId) != nil {
		log.Println(err)
		return function.UserUpdateProfileReply{}, err
	}
	if err := userDao.UpdateUserNicknameAndProfileById(userId, args.Nickname, args.ProfilePicture); err != nil {
		return function.UserUpdateProfileReply{}, err
	}
	return function.UserUpdateProfileReply{
		Nickname:       args.Nickname,
		ProfilePicture: args.ProfilePicture,
	}, nil
}

func UpdateProfileNickname(args function.UserUpdateProfileArgs) (function.UserUpdateProfileReply, error) {
	userId, err := userDao.GetUserIdByUUID(args.Token)
	if err != nil {
		return function.UserUpdateProfileReply{}, err
	}

	// TODO可能有bug删除缓存
	if userDao.RemoveCacheWithUserByID(userId) != nil {
		log.Println(err)
		return function.UserUpdateProfileReply{}, err
	}

	if err := userDao.UpdateUserNicknameById(userId, args.Nickname); err != nil {
		return function.UserUpdateProfileReply{}, err
	}

	// TODO 写缓存
	return function.UserUpdateProfileReply{
		Nickname:       args.Nickname,
		ProfilePicture: args.ProfilePicture,
	}, nil
}

func UpdateProfilePicture(args function.UserUpdateProfileArgs) (function.UserUpdateProfileReply, error) {
	userId, err := userDao.GetUserIdByUUID(args.Token)
	if err != nil {
		return function.UserUpdateProfileReply{}, err
	}

	// TODO可能有bug删除缓存
	if userDao.RemoveCacheWithUserByID(userId) != nil {
		log.Println(err)
		return function.UserUpdateProfileReply{}, err
	}

	if err := userDao.UpdateUserProfileById(userId, args.ProfilePicture); err != nil {
		return function.UserUpdateProfileReply{}, err
	}

	// TODO 写缓存
	return function.UserUpdateProfileReply{
		Nickname:       args.Nickname,
		ProfilePicture: args.ProfilePicture,
	}, nil
}
