package main

import (
	"fmt"
	"log"
)

/* TODO :: connection control module */
/* TODO :: worker and worker control module (maybe use goroutine) */
/* TODO :: network IO module */
/* TODO :: request parsing module */
/* TODO :: processing module */

var defaultRoutineCount = 6
var defaultPort = "11508"

var iceblueInfo info

type info struct {
	routineCount int
	port         string
}

func stats() {
	/* FIXME change the way of print stats to operation response */
	fmt.Printf("routineCount = %d\n", iceblueInfo.routineCount)
	fmt.Printf("port = %s\n", iceblueInfo.port)
}

func initializeInfo() {
	iceblueInfo.routineCount = defaultRoutineCount
	iceblueInfo.port = defaultPort
}

func main() {
	log.Printf("Start IceBlue Simple Key-value in memory storage.\n")

	initializeInfo()
	initializeStore()

	initializeProcessRoutine(iceblueInfo.routineCount)
	success := initializeNetworkModule(iceblueInfo.port)
	if !success {
		log.Panic("Fail initialize network module")
	}

	acceptNetworkProcess()

	destroyNetworkModule()
	destroyProcessRoutine()
	destroyStore()
}
