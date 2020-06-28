package main

import (
	"fmt"
)

func store(key string, value string) int {
	storedKey := key
	storedValue := value

	fmt.Printf("storedKey = %s, storedValue = %s\n", storedKey, storedValue)

	return 0
}
