package helper

import (
	"log"
	"strconv"
	
	"gopkg.in/redis.v3"
)



type RedisClient struct {
	client *redis.Client
	prefix string
}

func CreateRedisClient(redis_address string, redis_port int, redis_db_name int64, redis_password string, prefx string) RedisClient {
	redisClient := RedisClient{ client : redis.NewClient(&redis.Options{
		Addr:     redis_address + ":" + strconv.Itoa(redis_port),
		Password: redis_password,
		DB:       redis_db_name,
	}), prefix:prefx } 
	pong, err := redisClient.client.Ping().Result()
	log.Println(pong, err)
	return redisClient
}

func (r RedisClient) Set(key string, value interface{}) {
	err := r.client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (r RedisClient) Get(key string) interface{} {
	val, err := r.client.Get(key).Result()
	if err == redis.Nil {
		log.Println(key + " does not exists")
	} else if err != nil {
		panic(err)
	} else {
		log.Println(key, val)
	}
	return val
}