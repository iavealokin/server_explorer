package sqlstore

import (
	"server_explorer/pkg/bus"
	"server_explorer/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetUser)
}

func GetUser(user *models.User) error {
	if _, err := x.Table("users").Get(user); /*x.Get(user)*/ err != nil {
		return err
	}
	user.Authorized = true
	return nil
}
