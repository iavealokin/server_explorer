package models

type User struct {
	Name       string
	Password   string
	Authorized bool `xorm:"-"`
}
