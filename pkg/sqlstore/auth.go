package sqlstore

import (
	"fmt"
	"server_explorer/pkg/bus"
	"server_explorer/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetUser)
}

func GetUser(user *models.User) error {
	if _, err := x.Table("users").Get(user); /*x.Get(user)*/ err != nil {
		fmt.Println("HEUO ", err)
		return err
	}
	user.Authorized = true
	return nil
}
