package utils

import (
	"fmt"
	"os"
)

func GetEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		fmt.Printf("ERROR: Environment variable %s is not set\n", key)
		return defaultValue
	}
	return value
}
