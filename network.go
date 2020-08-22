package main

import (
	"log"
	"net"
)

type network struct {
	listener net.Listener
}

var networkInfo network

func acceptNetworkProcess() {
	log.Printf("Start accept network process\n")
	for networkInfo.listener != nil {
		conn, err := networkInfo.listener.Accept()
		if err != nil {
			log.Panic(err)
		}
		/* DEBUG code */
		log.Printf("asscepted")
		conn.Close()
		break
	}
	log.Printf("Stop accept network process\n")
}

func initializeNetworkModule(port string) bool {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Printf("Fail Listen(%v)\n", err)
		return false
	}
	log.Printf("Initialize network module\n")

	networkInfo.listener = l
	return true
}

func destroyNetworkModule() {
	if networkInfo.listener != nil {
		networkInfo.listener.Close()
		networkInfo.listener = nil
	}
	log.Printf("Destroy network module\n")
}
