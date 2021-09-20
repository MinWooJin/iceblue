package main

import (
	"fmt"
	"iceblue/pkg/process"
	"iceblue/pkg/storage"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var defaultPort = "11508"

var iceblueInfo info

type info struct {
	port string
}

func stats() {
	/* FIXME change the way of print stats to operation response */
	fmt.Printf("port = %s\n", iceblueInfo.port)
}

func initializeInfo() {
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
	storage.InitializeStore()

	process.InitializeProcessRoutine()
	success := process.InitializeNetworkModule(iceblueInfo.port)
	if !success {
		log.Panic("Fail initialize network module")
	}

	signalChannel := make(chan os.Signal, 1)
	quit := make(chan bool, 1)

	var waitGroups sync.WaitGroup
	waitGroups.Add(1)

	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go handleSignal(signalChannel, quit)
	go process.AcceptNetworkProcess(&waitGroups, quit)

	waitGroups.Wait()

	process.DestroyNetworkModule()
	process.DestroyProcessRoutine()
	storage.DestroyStore()
}
