package main

import (
	"log"
)

func processCommand(data string, position int) int {
	/* DEBUG code */
	log.Printf("process command. data = %s, position %d\n", data, position)

	return 0
}

func initializeProcessRoutine(routineCount int) {
	log.Printf("Initialize process routine\n")
}

func destroyProcessRoutine() {
	log.Printf("Destroy process routine\n")
}
