package main

import (
	"fmt"
)

/* TODO : change print message to logging system */
/* TODO : change storage data structure map to hash table */

var table map[string]string

func store(key string, value string) int {
	_, ok := table[key]
	if ok {
		fmt.Printf("Already exist key(%s)\n", key)
		return -1
	}
	table[key] = value

	fmt.Printf("[DEBUG] stored. key=%s, value=%s\n", key, value)
	return 0
}

func get(key string) (string, int) {
	value, ok := table[key]
	fmt.Printf("[DEBUG] get. key=%s, exist=%v, value=%s\n", key, ok, value)

	if ok {
		return value, 0
	}
	return value, -1
}

func update(key string, value string) int {
	_, ok := table[key]

	if ok {
		table[key] = value
		return 0
	}
	return -1
}

func initializeStore() {
	table = make(map[string]string)

	fmt.Printf("initialize storage module.\n")
}
