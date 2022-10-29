package sqlstore

import (
	"fmt"
	"server_explorer/pkg/bus"
	"server_explorer/pkg/configs"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var x *xorm.Engine

type SQLStore struct {
	engine *xorm.Engine
	Cfg    *configs.Config `inject:""`
	Bus    bus.Bus         `inject:""`
}

func (ss *SQLStore) readConfig() {
	var err error
	database := ss.Cfg.ConfigFile.Section("database")
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		database.Key("user").MustString("user"),
		database.Key("password").MustString("password"),
		database.Key("host").MustString("host"),
		database.Key("port").MustInt(5432),
		database.Key("dbname").MustString("dbname"),
		database.Key("sslmode").MustString("disable"))
	ss.engine, err = xorm.NewEngine("postgres", connectionString)
	if err != nil {
		fmt.Println("ERROR IS ", err)
	}
	x = ss.engine
}
func New(cfg *configs.Config) *SQLStore {
	ss := &SQLStore{Cfg: cfg}
	ss.readConfig()
	return ss
}
