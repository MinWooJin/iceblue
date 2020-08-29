package main

import (
	"log"
)

// enum Operation code
const (
	STORE = iota
	UPDATE
	DELETE
	GET
)

const maxTokenCount int = 10
const operationToken int = 1
const keyToken int = 2

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

func processCommand(data string, position int) int {
	tokenCount, tokens := tokenizeCommand(data, position)

	// DEBUG
	for i := 0; i < tokenCount; i++ {
		log.Printf("========DEBUG======= tokens[%d] : %s\n", i, tokens[i])
	}
	return 0
}

func initializeProcessRoutine(routineCount int) {
	log.Printf("Initialize process routine\n")
}

func destroyProcessRoutine() {
	log.Printf("Destroy process routine\n")
}
