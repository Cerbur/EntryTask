package bind

import (
	"fmt"
	"testing"
)

type User struct {
	Name     string `json:"name" pattern:"^[a-zA-Z0-9_-]{6,15}$"`
	Password string `json:"password" pattern:"^().{6,}$"`
}

func TestBindStruct(t *testing.T) {
	user := User{}
	//err := Struct(&user)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	fmt.Println(user)
}
