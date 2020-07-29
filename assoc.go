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

func hash(str string) uint32 {
	var hvalue uint32
	strlength := uint32(len(str))
	hashFunction := assoc.hashFunction
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

func assocInsert(it *Item) int {
	hvalue := it.hvalue
	if assocGet(hvalue, it.key) != nil {
		return -1
	}

	bucketIdx := hvalue % assoc.hashtableSize
	it.next = assoc.hashtable[bucketIdx]
	assoc.hashtable[bucketIdx] = it

	assoc.totalItems++
	if assoc.totalItems >= (assoc.hashtableSize*3)/2 {
		assocExpand()
	}
	return 0
}

func assocDelete(hvalue uint32, key string) {
	var prev *Item
	bucketIdx := hvalue % assoc.hashtableSize
	it := assoc.hashtable[bucketIdx]
	for it != nil {
		if it.keyLen == uint32(len(key)) && it.key == key {
			if prev != nil {
				prev.next = it.next
			} else {
				assoc.hashtable[bucketIdx] = it.next
			}
			assoc.totalItems--
			it.next = nil
		}
		prev = it
		it = it.next
	}
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
