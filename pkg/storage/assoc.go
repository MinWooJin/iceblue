package storage

import (
	"hash/fnv"
	"log"
)

// HASH functions
const (
	FNV32 = iota
	FNV32A
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

func Hash(str string) uint32 {
	var hvalue uint32
	hashFunction := assoc.hashFunction
	switch hashFunction {
	case FNV32:
		/* New32 returns a new 32-bit FNV-1 hash.Hash.
		 * Its Sum method will lay the value out in big-endian byte order.
		 */
		algorithm := fnv.New32()
		algorithm.Write([]byte(str))
		hvalue = algorithm.Sum32()
	case FNV32A:
		/* New32a returns a new 32-bit FNV-1a hash.Hash.
		 * Its Sum method will lay the value out in big-endian byte order.
		 */
		algorithm := fnv.New32a()
		algorithm.Write([]byte(str))
		hvalue = algorithm.Sum32()
	}
	/* TODO :: support more hash function */
	return hvalue
}

func AssocGet(hvalue uint32, key string) *Item {
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
		if it.KeyLen == uint32(len(key)) && it.Key == key {
			break
		}
		it = it.Next
	}
	return it
}

func AssocInsert(it *Item) int {
	hvalue := it.Hvalue
	if AssocGet(hvalue, it.Key) != nil {
		return -1
	}

	if assoc.expanding {
		expandTable()
	}

	var bucketIdx uint32
	if expandBucket(hvalue) {
		bucketIdx = hvalue % assoc.oldHashtableSize
		it.Next = assoc.oldHashtable[bucketIdx]
		assoc.oldHashtable[bucketIdx] = it
	} else {
		bucketIdx = hvalue % assoc.hashtableSize
		it.Next = assoc.hashtable[bucketIdx]
		assoc.hashtable[bucketIdx] = it
	}

	assoc.totalItems++
	if !assoc.expanding && assoc.totalItems >= (assoc.hashtableSize*3)/2 {
		assoc.expanding = true
		expandTable()
	}
	return 0
}

func AssocDelete(hvalue uint32, key string) {
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
		if it.KeyLen == uint32(len(key)) && it.Key == key {
			if prev != nil {
				prev.Next = it.Next
			} else {
				if oldHashtableDelete {
					assoc.oldHashtable[bucketIdx] = it.Next
				} else {
					assoc.hashtable[bucketIdx] = it.Next
				}
			}
			assoc.totalItems--
			it.Next = nil
		}
		prev = it
		it = it.Next
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
		log.Printf("Start hashtable expand. [size=%d]\n", assoc.hashtableSize)
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
			nextItem = it.Next
			nextBucket := it.Hvalue % assoc.hashtableSize
			it.Next = assoc.hashtable[nextBucket]
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
		log.Printf("End hashtable expand. [size=%d]\n", assoc.hashtableSize)
	}
}

func InitializeAssoc(hashtableSize uint32, hashFunction int) {
	assoc.hashtableSize = hashtableSize
	/* TODO: add failure handling */
	assoc.hashtable = make([]*Item, assoc.hashtableSize)
	assoc.hashFunction = hashFunction

	log.Printf("Initialize assoc module.[size=%d]\n", assoc.hashtableSize)
}

func allItemReleaseInBucket(startBucket uint32, endBucket uint32) {
	var bucket uint32
	for bucket = startBucket; bucket < endBucket; bucket++ {
		for it := assoc.hashtable[bucket]; it != nil; {
			temp := it
			it = it.Next
			temp.Next = nil
		}
		assoc.hashtable[bucket] = nil
	}
}

func DestroyAssoc() {
	if assoc.expanding {
		allItemReleaseInBucket(assoc.expandingBucket, assoc.oldHashtableSize)
		allItemReleaseInBucket(0, assoc.hashtableSize)
	} else {
		allItemReleaseInBucket(0, assoc.hashtableSize)
	}

	log.Printf("Destroy assoc module.\n")
}
