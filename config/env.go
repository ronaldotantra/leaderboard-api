package config

import (
	"log"
	"os"
	"strconv"
)

func getString(key, def string) string {
	res := os.Getenv(key)
	if res == "" {
		return def
	}

	return res
}

func getInt(key string, def int) int {
	res := os.Getenv(key)
	if res == "" {
		return def
	}

	resInt, err := strconv.Atoi(res)
	if err != nil {
		log.Printf("failed to get %s - %v", key, err)
		return def
	}
	if resInt == 0 {
		return def
	}

	return resInt
}
