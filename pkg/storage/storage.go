package storage

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/* TODO : change print message to logging system */
/* TODO : change storage data structure map to hash table */

var mutex sync.Mutex

/* TODO :: add Value length argument for optimize */
func Store(key string, value string) int {
	mutex.Lock()
	hashVal := Hash(key)
	if AssocGet(hashVal, key) != nil {
		fmt.Printf("Already exist Key(%s)\n", key)
		mutex.Unlock()
		return -1
	}

	storeItem := new(Item)
	storeItem.Key = key
	storeItem.KeyLen = uint32(len(key))
	storeItem.Value = value
	storeItem.ValLen = uint32(len(value))
	storeItem.Hvalue = hashVal
	storeItem.Time = time.Now()

	ret := AssocInsert(storeItem)
	if ret == -1 {
		log.Printf("Failed insert Key(%s)\n", key)
	}
	log.Printf("[DEBUG] stored. Key=%s, Value=%s\n", key, value)
	mutex.Unlock()

	return ret
}

func Get(key string) (string, int) {
	value := ""
	result := -1

	mutex.Lock()
	item := AssocGet(Hash(key), key)
	if item != nil {
		value = item.Value
		result = 0
	}
	mutex.Unlock()
	if result == 0 {
		log.Printf("[DEBUG] Get. Key=%s, Value=%s\n", key, item.Value)
	} else {
		log.Printf("[DEBUG] Get. Key=%s\n", key)
	}
	return value, result
}

func Delete(key string) {
	mutex.Lock()
	AssocDelete(Hash(key), key)
	mutex.Unlock()
	log.Printf("[DEBUG] Delete. Key=%s\n", key)
}

func Update(key string, value string) int {
	result := -1

	mutex.Lock()
	item := AssocGet(Hash(key), key)
	if item != nil {
		item.Value = value
		result = 0
	}
	mutex.Unlock()
	if result == 0 {
		log.Printf("[DEBUG] Update. Key=%s, Value=%s\n", key, item.Value)
	} else {
		log.Printf("[DEBUG] Update. Key=%s\n", key)
	}
	return result
}

func InitializeStore() {
	InitializeAssoc(uint32(128), FNV32)

	log.Printf("Initialize storage module.\n")
}

func DestroyStore() {
	DestroyAssoc()
	log.Printf("Destory storage module.\n")
}
