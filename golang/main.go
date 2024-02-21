package main

import (
	"log"
	"machine-coding-lru-cache/internal/service"
)

func main() {
	log.Println("in main")
	cacheService := service.GetCacheService(2)

	str, err := cacheService.GetValue("2")
	if err != nil {
		log.Println("err: ", err)
	} else {
		log.Println("value: ", str)
	}

	cacheService.PutValue("1", "abcd")
	str, _ = cacheService.GetValue("1")
	log.Println("value: ", str)

	cacheService.PutValue("2", "ef")
	str, _ = cacheService.GetValue("2")
	log.Println("value: ", str)

	cacheService.PutValue("3", "gh")
	str, err = cacheService.GetValue("1")
	if err != nil {
		log.Println("err: ", err)
	} else {
		log.Println("value: ", str)
	}

}
