package main

import (
	"fmt"
)

// HASH functions
const (
	SAMPLE = iota
)

type assocST struct {
	hashtable     []*item
	hashtableSize uint32
	hashFunction  int
}

var assoc assocST

func hash(str string, hashFunction int) uint32 {
	var hvalue uint32
	strlength := uint32(len(str))
	/* TODO :: make hash function */
	switch hashFunction {
	case SAMPLE:
		hvalue = strlength % assoc.hashtableSize
		break
	}
	return hvalue
}

func assocGet(hvalue uint32, key string) *item {
	var getItem *item

	return getItem
}

func assocInsert(hvalue uint32, it *item) bool {
	return true
}

func assocDelete(hvalue uint32, key string) bool {
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
	assoc.hashtable = make([]*item, assoc.hashtableSize)
	assoc.hashFunction = hashFunction

	fmt.Printf("initialize assoc module.\n")
}
