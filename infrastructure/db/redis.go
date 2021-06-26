package db

import (
	"github.com/InnoSoft/task/infrastructure/config"
	"gopkg.in/redis.v3"
)

type Redis struct {
}

func NewRedis() *Redis {
	return &Redis{}
}

func (db *Redis) NewClient() (*redis.Client , error){
	viper := config.NewViper()
	address := viper.Cache.Address
	//fmt.Println("viper",viper)
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "redis:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	client := redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0,
	})
	_, err := client.Ping().Result()
	return client , err
}