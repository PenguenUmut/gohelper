package main

import (
	"log"
	
	"./model"
	"./helper"
)




func main() {
	test_conf()
	test_redis()
}


// helper.config test
func test_conf() {
	var conf model.Config
	helper.ReadConfig(&conf, "config.json")
	log.Println("Name: ", conf.Name)
	log.Println("Version: ", conf.Version)
	log.Println("MyObject.Enabled: ", conf.MyObject.Enabled)
	log.Println("MyObject.ID: ", conf.MyObject.ID)
}

// helper.redis test
func test_redis(){
	redisClient := helper.CreateRedisClient("localhost",6379,1,"",":")
	redisClient.Set("key1", "value1")
	redisClient.Set("key3", 5)
	d := redisClient.Get("key1")
	log.Println("d",d)
	d = redisClient.Get("key2")
	log.Println("d",d)
}