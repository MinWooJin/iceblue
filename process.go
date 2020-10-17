package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

const maxTokenCount int = 10
const operationToken int = 0
const keyToken int = 1

// Operation Code
const (
	SET = iota
	UPDATE
	DELETE
	GET
	STATS
	QUIT
	UNKNOWN = -1
)

func tokenizeCommand(data string, position int) (int, [maxTokenCount]string) {
	var tokenCount int = 0
	var tokens [maxTokenCount]string
	var startingIndex int = 0

	for i := 0; i < position; i++ {
		if data[i] == ' ' {
			tokens[tokenCount] = data[startingIndex:i]
			startingIndex = i + 1
			tokenCount++
		}
	}
	tokens[tokenCount] = data[startingIndex:position]
	tokenCount++

	return tokenCount, tokens
}

func tryReadNextline(conn net.Conn, data string, position int, vlen int) (string, int) {
	var nextline string = ""
	var ret int = -1

	/* postion is end index of \r\n in first line */
	remainData := data[position:len(data)]
	remainDataLength := len(remainData)
	if remainDataLength <= vlen {
		var buffer bytes.Buffer
		var additionalData []byte
		buffer.WriteString(remainData)
		additionalData, ret = readNetworkRequest(conn, vlen-remainDataLength)
		if ret < 0 {
			return nextline, ret
		}
		buffer.Write(additionalData)
		remainData = buffer.String()
	}

	/* check protocol.. CRLF(\r\n) or LF(\n) */
	var nextlinePosition int
	if nextlinePosition = strings.Index(remainData, "\\r\\n"); nextlinePosition < 0 {
		nextlinePosition = strings.Index(remainData, "\\n")
	}
	if nextlinePosition == vlen {
		nextline = remainData[:nextlinePosition]
		ret = 0
	}

	return nextline, ret
}

func getOperationCode(operationToken string, tokenCount int) int {
	if operationToken == "set" && tokenCount == 3 {
		return SET
	} else if operationToken == "update" && tokenCount == 3 {
		return UPDATE
	} else if operationToken == "delete" && tokenCount == 2 {
		return DELETE
	} else if operationToken == "get" && tokenCount == 2 {
		return GET
	} else if operationToken == "stats" && tokenCount == 1 {
		return STATS
	} else if operationToken == "quit" && tokenCount == 1 {
		return QUIT
	} else {
		return UNKNOWN
	}
}

func processStoreCommand(conn net.Conn, key string, velnStr string, data string, endPosition int) int {
	vlen, err := strconv.Atoi(velnStr)
	if err != nil {
		if sendNetworkRequest(conn, "CLIENT_ERROR bad line format") < 0 {
			return -1
		}
	}
	value, ret := tryReadNextline(conn, data, endPosition, vlen)
	if ret < 0 {
		if sendNetworkRequest(conn, "CLIENT_ERROR failed value read") < 0 {
			return -1
		}
	}
	ret = store(key, value)
	if ret == 0 {
		if sendNetworkRequest(conn, "SUCCESS") < 0 {
			return -1
		}
	} else {
		/* TODO :: error hanlding according to error code */
		if sendNetworkRequest(conn, "SERVER_ERROR store failed") < 0 {
			return -1
		}
	}
	return 0
}

func processUpdateCommand(conn net.Conn, key string, velnStr string, data string, endPosition int) int {
	vlen, err := strconv.Atoi(velnStr)
	if err != nil {
		if sendNetworkRequest(conn, "CLIENT_ERROR bad line format") < 0 {
			return -1
		}
	}
	value, ret := tryReadNextline(conn, data, endPosition, vlen)
	if ret < 0 {
		if sendNetworkRequest(conn, "CLIENT_ERROR failed value read") < 0 {
			return -1
		}
	}
	ret = update(key, value)
	if ret == 0 {
		if sendNetworkRequest(conn, "SUCCESS") < 0 {
			return -1
		}
	} else {
		/* TODO :: error hanlding according to error code */
		if sendNetworkRequest(conn, "SERVER_ERROR update failed") < 0 {
			return -1
		}
	}
	return 0
}

func processDeleteCommand(conn net.Conn, key string) int {
	delete(key)
	if sendNetworkRequest(conn, "DELETED") < 0 {
		return -1
	}
	return 0
}

func processGetCommand(conn net.Conn, key string) int {
	var buffer bytes.Buffer
	value, ret := get(key)
	if ret == 0 {
		/* separate header and body to support multiple keys */
		headerStr := fmt.Sprintf("VALUE %d\r\n", len(value))
		bodyStr := fmt.Sprintf("%s\r\n", value)
		buffer.WriteString(headerStr)
		buffer.WriteString(bodyStr)
	} else {
		/* TODO :: error hanlding according to error code */
	}
	buffer.WriteString("END\r\n")
	if sendNetworkRequest(conn, buffer.String()) < 0 {
		return -1
	}
	return 0
}

func processStatsCommand(conn net.Conn) int {
	/* TODO :: process stats command */
	return 0
}

func processUnknownCommand(conn net.Conn, operation string) int {
	log.Printf("Unknown operation = %s, addr = %s\n",
		operation, conn.RemoteAddr().String())

	if sendNetworkRequest(conn, "CLIENT_ERROR unknown operation") < 0 {
		return -1
	}
	return 0
}

func processQuitCommand(conn net.Conn) int {
	if sendNetworkRequest(conn, "Quit this connection") < 0 {
		return -1
	}
	return -2
}

func processCommand(conn net.Conn, data string, position int, endPosition int) int {
	tokenCount, tokens := tokenizeCommand(data, position)

	operationCode := getOperationCode(tokens[operationToken], tokenCount)
	switch operationCode {
	case SET:
		return processStoreCommand(conn, tokens[keyToken], tokens[keyToken+1], data, endPosition)
	case UPDATE:
		return processUpdateCommand(conn, tokens[keyToken], tokens[keyToken+1], data, endPosition)
	case DELETE:
		return processDeleteCommand(conn, tokens[keyToken])
	case GET:
		return processGetCommand(conn, tokens[keyToken])
	case STATS:
		return processStatsCommand(conn)
	case QUIT:
		return processQuitCommand(conn)
	case UNKNOWN:
		return processUnknownCommand(conn, tokens[operationToken])
	}
	return 0
}

func initializeProcessRoutine(routineCount int) {
	log.Printf("Initialize process routine\n")
}

func destroyProcessRoutine() {
	log.Printf("Destroy process routine\n")
}
