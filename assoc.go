package main

import (
	"fmt"
)

// HASH functions
const (
	SAMPLE = iota
)

type assocST struct {
	hashtable        []*Item
	oldHashtable     []*Item
	hashtableSize    uint32
	oldHashtableSize uint32
	expandingBucket  uint32
	hashFunction     int
	totalItems       uint32
	expanding        bool
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

func assocGet(hvalue uint32, key string) *Item {
	var bucketIdx uint32
	var it *Item

	if expandBucket(hvalue) {
		bucketIdx = hvalue % assoc.oldHashtableSize
		it = assoc.oldHashtable[bucketIdx]
	} else {
		bucketIdx = hvalue % assoc.hashtableSize
		it = assoc.hashtable[bucketIdx]
	}

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

	if assoc.expanding {
		expandTable()
	}

	var bucketIdx uint32
	if expandBucket(hvalue) {
		bucketIdx = hvalue % assoc.oldHashtableSize
		it.next = assoc.oldHashtable[bucketIdx]
		assoc.oldHashtable[bucketIdx] = it
	} else {
		bucketIdx = hvalue % assoc.hashtableSize
		it.next = assoc.hashtable[bucketIdx]
		assoc.hashtable[bucketIdx] = it
	}

	assoc.totalItems++
	if !assoc.expanding && assoc.totalItems >= (assoc.hashtableSize*3)/2 {
		assoc.expanding = true
		expandTable()
	}
	return 0
}

func assocDelete(hvalue uint32, key string) {
	var prev *Item
	var it *Item
	var bucketIdx uint32
	var oldHashtableDelete bool

	if expandBucket(hvalue) {
		bucketIdx = hvalue % assoc.oldHashtableSize
		it = assoc.oldHashtable[bucketIdx]
		oldHashtableDelete = true
	} else {
		bucketIdx = hvalue % assoc.hashtableSize
		it = assoc.hashtable[bucketIdx]
	}

	for it != nil {
		if it.keyLen == uint32(len(key)) && it.key == key {
			if prev != nil {
				prev.next = it.next
			} else {
				if oldHashtableDelete {
					assoc.oldHashtable[bucketIdx] = it.next
				} else {
					assoc.hashtable[bucketIdx] = it.next
				}
			}
			assoc.totalItems--
			it.next = nil
		}
		prev = it
		it = it.next
	}
}

func expandBucket(hvalue uint32) bool {
	if assoc.oldHashtable != nil {
		bucket := hvalue % assoc.oldHashtableSize
		if bucket >= assoc.expandingBucket {
			return true
		}
	}
	return false
}

func expandTable() {
	maxExpandCount := 4

	if assoc.oldHashtable == nil {
		fmt.Printf("Start hashtable expand. [size=%d]\n", assoc.hashtableSize)
		assoc.oldHashtableSize = assoc.hashtableSize
		assoc.hashtableSize = assoc.hashtableSize * 2
		assoc.oldHashtable = assoc.hashtable
		assoc.hashtable = make([]*Item, assoc.hashtableSize)
		assoc.expandingBucket = 0
	}

	var bucket uint32
	var nextItem *Item
	moveCount := 0
	for bucket = assoc.expandingBucket; bucket < assoc.oldHashtableSize; bucket++ {
		it := assoc.oldHashtable[bucket]
		for it != nil {
			nextItem = it.next
			nextBucket := it.hvalue % assoc.hashtableSize
			it.next = assoc.hashtable[nextBucket]
			assoc.hashtable[nextBucket] = it
			it = nextItem
		}
		assoc.oldHashtable[bucket] = nil
		assoc.expandingBucket++

		moveCount++
		if moveCount == maxExpandCount {
			break
		}
	}
	if assoc.expandingBucket == assoc.oldHashtableSize {
		assoc.oldHashtable = nil
		assoc.oldHashtableSize = 0
		assoc.expandingBucket = 0
		assoc.expanding = false
		fmt.Printf("End hashtable expand. [size=%d]\n", assoc.hashtableSize)
	}
}

func initializeAssoc(hashtableSize uint32, hashFunction int) {
	assoc.hashtableSize = hashtableSize
	/* TODO: add failure handling */
	assoc.hashtable = make([]*Item, assoc.hashtableSize)
	assoc.hashFunction = hashFunction

	fmt.Printf("initialize assoc module.[size=%d]\n", assoc.hashtableSize)
}
