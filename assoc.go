package main

import (
	"fmt"
)

// HASH functions
const (
	SAMPLE = iota
)

type assocST struct {
	hashtable     []*Item
	hashtableSize uint32
	hashFunction  int
	totalItems    uint32
}

var assoc assocST

func hash(str string, hashFunction int) uint32 {
	var hvalue uint32
	strlength := uint32(len(str))
	/* TODO :: make hash function */
	switch hashFunction {
	case SAMPLE:
		hvalue = strlength
		break
	}
	return hvalue
}

func assocExpand() {
	/* do expand */
}

func assocGet(hvalue uint32, key string) *Item {
	bucketIdx := hvalue % assoc.hashtableSize
	it := assoc.hashtable[bucketIdx]
	for it != nil {
		if it.keyLen == uint32(len(key)) && it.key == key {
			break
		}
		it = it.next
	}
	return it
}

func assocInsert(hvalue uint32, it *Item) bool {
	if assocGet(hvalue, it.key) != nil {
		return false
	}

	bucketIdx := hvalue % assoc.hashtableSize
	it.next = assoc.hashtable[bucketIdx]
	assoc.hashtable[bucketIdx] = it

	assoc.totalItems++
	if assoc.totalItems >= (assoc.hashtableSize*3)/2 {
		assocExpand()
	}
	return true
}

func assocDelete(hvalue uint32, key string) bool {
	it := assocGet(hvalue, key)
	if it == nil {
		return false
	}
	return true
}

func checkExpandTable() bool {
	return false
}

func expandTable() bool {
	return true
}

func initializeAssoc(hashtableSize uint32, hashFunction int) {
	assoc.hashtableSize = hashtableSize
	assoc.hashtable = make([]*Item, assoc.hashtableSize)
	assoc.hashFunction = hashFunction

	fmt.Printf("initialize assoc module.\n")
}
