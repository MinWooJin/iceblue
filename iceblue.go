package main

import (
	"fmt"
	"log"
	"net"
)

/* TODO :: connection control module */
/* TODO :: worker and worker control module (maybe use goroutine) */
/* TODO :: network IO module */
/* TODO :: request parsing module */
/* TODO :: processing module */

var defaultRoutineCount = 6
var defaultPort = "11508"

type info struct {
	routineCount int
	port         string
}

type network struct {
	listener net.Listener
}

func stats(f info) {
	/* FIXME change the way of print stats to operation response */
	fmt.Printf("routineCount = %d\n", f.routineCount)
	fmt.Printf("port = %s\n", f.port)
}

func initializeInfo(f *info) {
	f.routineCount = defaultRoutineCount
	f.port = defaultPort
}

func readyProcessRoutine(routineCount int) {

}

func initializeNetworkModule(port string) (bool, net.Listener) {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Printf("Fail Listen(%v)\n", err)
		return false, nil
	}
	log.Printf("Initialize network module\n")

	/* DEBUG code */
	conn, err := l.Accept()
	if err != nil {
		log.Printf("%v\n", err)
	}
	log.Printf("asscepted")
	conn.Close()
	return true, l
}

func destroyNetworkModule(networkInfo network) {
	networkInfo.listener.Close()
	log.Printf("Destroy network module\n")
}

func main() {
	var iceblueInfo info
	var networkInfo network

	log.Printf("Start IceBlue Simple Key-value in memory storage.\n")

	initializeInfo(&iceblueInfo)
	initializeStore()

	readyProcessRoutine(iceblueInfo.routineCount)
	success, listener := initializeNetworkModule(iceblueInfo.port)
	if !success {
		log.Panic("Fail initialize network module")
	}
	networkInfo.listener = listener

	/* DEBUG (remove later) */
	stats(iceblueInfo)

	destroyNetworkModule(networkInfo)
	destroyStore()
}
