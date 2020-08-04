package main

import (
	"fmt"
	"sync"
	"time"
)

/* TODO : change print message to logging system */
/* TODO : change storage data structure map to hash table */

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
	hashVal := hash(key)
	if assocGet(hashVal, key) != nil {
		fmt.Printf("Already exist key(%s)\n", key)
		mutex.Unlock()
		return -1
	}

	storeItem := new(Item)
	storeItem.key = key
	storeItem.keyLen = uint32(len(key))
	storeItem.value = value
	storeItem.valLen = uint32(len(value))
	storeItem.hvalue = hashVal
	storeItem.time = time.Now()

	ret := assocInsert(storeItem)
	if ret == -1 {
		fmt.Printf("Failed insert key(%s)\n", key)
	}
	fmt.Printf("[DEBUG] stored. key=%s, value=%s\n", key, value)
	mutex.Unlock()

	return ret
}

func get(key string) (string, int) {
	value := ""
	result := -1

	mutex.Lock()
	item := assocGet(hash(key), key)
	fmt.Printf("[DEBUG] try get. key=%s\n", key)
	if item != nil {
		fmt.Printf("[DEBUG] get. key=%s, value=%s\n", key, item.value)
		value = item.value
		result = 0
	}
	mutex.Unlock()

	return value, result
}

func delete(key string) {
	mutex.Lock()
	assocDelete(hash(key), key)
	mutex.Unlock()
}

func update(key string, value string) int {
	result := -1

	mutex.Lock()
	item := assocGet(hash(key), key)
	if item != nil {
		item.value = value
		result = 0
	}
	mutex.Unlock()

	return result
}

func initializeStore() {
	initializeAssoc(uint32(1), SAMPLE)

	fmt.Printf("initialize storage module.\n")
}
