package main

import (
	"log"
	"os"
	"path/filepath"
	"encoding/json"
	"strconv"
)

import "github.com/PenguenUmut/gohelper/lib"
import "github.com/go-redis/redis"



var conf gohelper.Config
var redisClient *redis.Client

func main() {
	test_config()
	test_redis()
}




func test_config(){
	readConfig(&conf)

	log.Printf("config.json Name: ", conf.Name)
	log.Printf("config.json Version: ", conf.Version)
	log.Printf("config.json MyObject.Enabled: ", conf.MyObject.Enabled)
	log.Printf("config.json MyObject.ID: ", conf.MyObject.ID)
}

func readConfig(conf *gohelper.Config) {
	configFileName := "config.json"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	}
	configFileName, _ = filepath.Abs(configFileName)
	log.Printf("Loading config: %v", configFileName)

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal("File error: ", err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&conf); err != nil {
		log.Fatal("Config error: ", err.Error())
	}
}




func test_redis(){
	redisClient = createNewRedisClient("localhost",6379,0,"")
	sample_set_Client(redisClient, "key1", "value1")
	sample_set_Client(redisClient, "key3", 5)
	d := sample_get_Client(redisClient, "key1")
	log.Println("d",d)
	d = sample_get_Client(redisClient, "key2")
	log.Println("d",d)
}


func createNewRedisClient(redis_address string, redis_port int, redis_db_name int, redis_password string) *redis.Client {
	newRedisClient := redis.NewClient(&redis.Options{
		Addr:     redis_address + ":" + strconv.Itoa(redis_port),
		Password: redis_password, // no password set
		DB:       redis_db_name,  // use default DB
	})
	pong, err := newRedisClient.Ping().Result()
	log.Println(pong, err)
	return newRedisClient
}

func sample_set_Client(redisClient *redis.Client, key string, value interface{}) {
	err := redisClient.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func sample_get_Client(redisClient *redis.Client, key string) interface{} {
	val, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		log.Println(key + " does not exists")
	} else if err != nil {
		panic(err)
	} else {
		log.Println(key, val)
	}
	return val
}
