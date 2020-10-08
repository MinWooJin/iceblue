package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

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

func handleSignal(signalChannel chan os.Signal, doneChannel chan bool) {
	signal := <-signalChannel
	log.Printf("Receive signal. %v", signal)
	doneChannel <- true
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

	signalChannel := make(chan os.Signal, 1)
	quit := make(chan bool, 1)

	var waitGroups sync.WaitGroup
	waitGroups.Add(1)

	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go handleSignal(signalChannel, quit)
	go acceptNetworkProcess(&waitGroups, quit)

	waitGroups.Wait()

	destroyNetworkModule()
	destroyProcessRoutine()
	destroyStore()
}
