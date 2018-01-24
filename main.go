package main

import (
	"log"
	
	"./model"
	"./helper"
)




func main() {
	test_conf()
	test_redis()
	test_rpc()
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

func test_rpc() {
	rpc_client := helper.NewRPCClient("main", "http://127.0.0.1:8545", "10s")
	rpc_client.Check()
	bala,err := rpc_client.GetBalance("0xaf032fd62b7de068d66ce1ed3f14ebc6f4c25bb1")
	log.Println(bala, err)
	
	work,err2 := rpc_client.GetWork()
	log.Println(work, err2)
	log.Println("END")
}