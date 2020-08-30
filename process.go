package main

import (
	"log"
	"net"
)

const maxTokenCount int = 10
const operationToken int = 0
const keyToken int = 1

func tokenizeCommand(data string, positing int) (int, [maxTokenCount]string) {
	var tokenCount int = 0
	var tokens [maxTokenCount]string
	var startingIndex int = 0

	for i := 0; i < positing; i++ {
		if data[i] == ' ' {
			tokens[tokenCount] = data[startingIndex:i]
			startingIndex = i + 1
			tokenCount++
		}
	}
	tokens[tokenCount] = data[startingIndex:positing]
	tokenCount++

	return tokenCount, tokens
}

func processCommand(conn net.Conn, data string, position int) int {
	tokenCount, tokens := tokenizeCommand(data, position)

	// DEBUG
	for i := 0; i < tokenCount; i++ {
		log.Printf("========DEBUG======= tokens[%d] : %s\n", i, tokens[i])
	}

	if tokens[operationToken] == "set" {

	} else if tokens[operationToken] == "update" {

	} else if tokens[operationToken] == "delete" {

	} else if tokens[operationToken] == "get" {

	} else if tokens[operationToken] == "stats" {

	} else {
		log.Printf("Unknown operation = %s, addr = %s\n",
			tokens[operationToken], conn.RemoteAddr().String())

		if sendNetworkRequest(conn, "CLIENT_ERROR unknown operation") < 0 {
			return -1
		}
	}

	return 0
}

func initializeProcessRoutine(routineCount int) {
	log.Printf("Initialize process routine\n")
}

func destroyProcessRoutine() {
	log.Printf("Destroy process routine\n")
}
