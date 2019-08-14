package basic

import (
	"aisitar-micro/user-srv/basic/config"
	"aisitar-micro/user-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
