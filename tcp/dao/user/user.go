package user

import (
	"encoding/json"
	"entrytask/tcp/db"
	"entrytask/tcp/model"
	"entrytask/utils"
	"fmt"
	"log"
)

// mysql

func FindUserAllById(id uint32) model.User {
	user := model.User{}
	if err := db.Db.QueryRow("SELECT `id`,`profile_picture`,`nickname`,`username` FROM `t_user` WHERE `id` = ?", id).
		Scan(&user.ID, &user.Profile, &user.Nickname, &user.UserName); err != nil {
		log.Println("dao.user.FindUserAllById error:", err)
	}
	return user
}

func FindUserAllByUserNameAndPassword(username string, password string) model.User {
	p := utils.Md5(password)
	user := model.User{}
	if err := db.Db.QueryRow("SELECT `id`,`profile_picture`,`nickname`,`username` FROM `t_user` WHERE `username` = ? AND `password` = ?", username, p).
		Scan(&user.ID, &user.Profile, &user.Nickname, &user.UserName); err != nil {
		log.Println("dao.user.FindUserAllByUserNameAndPassword error:", err)
	}
	return user
}

func UpdateUserNicknameAndProfileById(id uint32, nickname string, profile string) error {
	_, err := db.Db.Exec("UPDATE t_user SET `profile_picture` = ?,`nickname` = ? WHERE `id` = ?", profile, nickname, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func UpdateUserNicknameById(id uint32, nickname string) error {
	_, err := db.Db.Exec("UPDATE t_user SET `nickname` = ? WHERE `id` = ?", nickname, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func UpdateUserProfileById(id uint32, profile string) error {
	_, err := db.Db.Exec("UPDATE t_user SET `profile_picture` = ? WHERE `id` = ?", profile, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

// Redis

func CheckUUIDExist(uuid string) (bool, error) {
	result, err := db.R.Exists(uuid).Result()
	if err != nil {
		return false, err
	}
	if result > 0 {
		return true, nil
	}
	return false, nil
}

func SaveUUIDWithId(id uint32, uuid string) error {
	return db.R.Set(uuid, id, 0).Err()
}

func GetUserIdByUUID(uuid string) (uint32, error) {
	i, err := db.R.Get(uuid).Int64()
	if err != nil {
		return 0, err
	}

	return uint32(i), nil
}

func SaveCacheWithUserById(user model.User, userId uint32) error {
	// user To JSON user
	marshal, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return db.R.Set(fmt.Sprint("userId:", userId), marshal, 0).Err()
}

func RemoveCacheWithUserByID(userId uint32) error {
	return db.R.Del(fmt.Sprint("userId:", userId)).Err()
}

func GetUserByUserById(userId uint32) (model.User, error) {
	user := model.User{}
	get := db.R.Get(fmt.Sprint("userId:", userId))
	bytes, err := get.Bytes()
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return user, err
	}
	return user, err
}
