package utils

import "os"

func GetEnv(key, defaultValue string) string  {
	if envVal, exist := os.LookupEnv(key); exist {
		return envVal
	}
	return defaultValue
}
