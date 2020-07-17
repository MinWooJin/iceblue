package main

import "fmt"

/* TODO :: connection control module */
/* TODO :: worker and worker control module (maybe use goroutine) */
/* TODO :: network IO module */
/* TODO :: request parsing module */
/* TODO :: processing module */

var defaultRoutineCount = 6
var defaultAddressString = "127.0.0.1:11508"

type info struct {
	routineCount int
	address      string
}

func stats(f info) {
	/* FIXME change the way of print stats to operation response */
	fmt.Printf("routineCount = %d\n", f.routineCount)
	fmt.Printf("address = %s\n", f.address)
}

func initializeInfo(f *info) {
	f.routineCount = defaultRoutineCount
	f.address = defaultAddressString
}

func main() {
	var iceblueInfo info

	initializeInfo(&iceblueInfo)
	initializeStore()

	/* DEBUG (remove later) */
	stats(iceblueInfo)
}
