package main

import (
	"fmt"
	"sync"
	"time"
)

/* TODO : change print message to logging system */
/* TODO : change storage data structure map to hash table */

var table map[string]Item
var mutex sync.Mutex

// Item is struct of iceblue data
type Item struct {
	key    string
	keyLen uint32
	value  string
	valLen uint32
	hvalue uint32
	time   time.Time
	next   *Item
}

func store(key string, value string) int {
	mutex.Lock()
	_, ok := table[key]
	if ok {
		fmt.Printf("Already exist key(%s)\n", key)
		mutex.Unlock()
		return -1
	}

	var storeItem Item
	storeItem.key = key
	storeItem.value = value
	storeItem.time = time.Now()
	table[key] = storeItem

	fmt.Printf("[DEBUG] stored. key=%s, value=%s\n", key, value)
	mutex.Unlock()

	return 0
}

func get(key string) (string, int) {
	value := ""
	result := -1

	mutex.Lock()
	item, ok := table[key]
	fmt.Printf("[DEBUG] get. key=%s, exist=%v, value=%s\n", key, ok, item.value)
	if ok {
		value = item.value
		result = 0
	}
	mutex.Unlock()

	return value, result
}

func update(key string, value string) int {
	result := -1

	mutex.Lock()
	item, ok := table[key]
	if ok {
		item.value = value
		table[key] = item
		result = 0
	}
	mutex.Unlock()

	return result
}

func initializeStore() {
	initializeAssoc(uint32(1024), SAMPLE)
	table = make(map[string]Item)

	fmt.Printf("initialize storage module.\n")
}
