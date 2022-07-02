package main

import (
	"log"
)

func main() {
	redisConfig := RedisConfig{
		Address:  "localhost",
		Port:     55000,
		Password: "password",
	}

	// _ yerine client ver
	_, err := redisConfig.Connect()

	if err != nil {
		log.Fatalf("error: redisConfig.Connect, %v", err)
	}

}
