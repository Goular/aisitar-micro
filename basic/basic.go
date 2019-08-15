package basic

import (
	"aisitar-micro/basic/config"
	"aisitar-micro/basic/db"
	"aisitar-micro/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
