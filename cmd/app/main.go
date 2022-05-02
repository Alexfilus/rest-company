package main

import (
	"log"
	"os"

	"rest-company/config"
	"rest-company/internal/server"
	"rest-company/pkg/logger"
	"rest-company/pkg/redis"
)

func main() {
	log.Println("Starting api server")
	log.Println(os.Getenv("CONFIG"))
	cfgFile, err := config.LoadConfig(os.Getenv("CONFIG"))
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	redisCli, err := redis.NewRedisClient(cfg)
	if err != nil {
		appLogger.Panic(err)
	}
	appLogger.Info("Redis connected")
	defer redisCli.Close()
	s := server.New(cfg, redisCli, appLogger)
	if err = s.Run(); err != nil {
		appLogger.Fatal(err)
	}
}
