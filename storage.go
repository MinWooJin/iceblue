package main

import (
	"fmt"
	"log"
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
		log.Printf("Failed insert key(%s)\n", key)
	}
	log.Printf("[DEBUG] stored. key=%s, value=%s\n", key, value)
	mutex.Unlock()

	return ret
}

func get(key string) (string, int) {
	value := ""
	result := -1

	mutex.Lock()
	item := assocGet(hash(key), key)
	if item != nil {
		value = item.value
		result = 0
	}
	mutex.Unlock()
	if result == 0 {
		log.Printf("[DEBUG] get. key=%s, value=%s\n", key, item.value)
	} else {
		log.Printf("[DEBUG] get. key=%s\n", key)
	}
	return value, result
}

func delete(key string) {
	mutex.Lock()
	assocDelete(hash(key), key)
	mutex.Unlock()
	log.Printf("[DEBUG] delete. key=%s\n", key)
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
	if result == 0 {
		log.Printf("[DEBUG] update. key=%s, value=%s\n", key, item.value)
	} else {
		log.Printf("[DEBUG] update. key=%s\n", key)
	}
	return result
}

func initializeStore() {
	initializeAssoc(uint32(128), SAMPLE)

	log.Printf("Initialize storage module.\n")
}

func destroyStore() {
	destroyAssoc()
	log.Printf("Destory storage module.\n")
}
