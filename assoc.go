package main

import (
	"fmt"
)

type assocST struct {
	hashtable     []*item
	hashtableSize uint64
}

var assoc assocST

func hash(str string) uint32 {
	var hvalue uint32
	/* TODO :: make hash function */
	return hvalue
}

func assocGet(hvalue uint32, key string) *item {
	var getItem *item

	return getItem
}

func assocInsert(hvalue uint32, it *item) bool {
	return true
}

func initializeAssoc(hashtableSize uint64) {
	assoc.hashtableSize = hashtableSize
	assoc.hashtable = make([]*item, assoc.hashtableSize)

	fmt.Printf("initialize assoc module.\n")
}
