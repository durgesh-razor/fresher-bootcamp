package Environment

import (
	"fmt"
	"os"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	fmt.Println(value)
	return value
}
