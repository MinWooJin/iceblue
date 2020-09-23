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

func getOperationCode(operationToken string) int {
	return 0
}

func processCommand(conn net.Conn, data string, position int, endPosition int) int {
	tokenCount, tokens := tokenizeCommand(data, position)

	/* TODO :: define operation code using operationToken and refactoring to switch case */

	if (tokens[operationToken] == "set" || tokens[operationToken] == "update") && tokenCount == 3 {
		/* Input foramt : set {key} {vlen}\r\n{value}\r\n */
		/* TODO :: check if limits are needed of key, value length */
		for {
			key := tokens[keyToken]
			vlen, err := strconv.Atoi(tokens[keyToken+1])
			if err != nil {
				if sendNetworkRequest(conn, "CLIENT_ERROR bad line format") < 0 {
					return -1
				}
				break
			}
			value, ret := tryReadNextline(conn, data, endPosition, vlen)
			if ret < 0 {
				if sendNetworkRequest(conn, "CLIENT_ERROR failed value read") < 0 {
					return -1
				}
				break
			}
			/* TODO :: define operation code and use that */
			if tokens[operationToken] == "set" {
				ret = store(key, value)
			} else {
				ret = update(key, value)
			}
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
			break
		}
	} else if tokens[operationToken] == "delete" && tokenCount == 2 {
		for {
			key := tokens[keyToken]
			delete(key)
			if sendNetworkRequest(conn, "DELETED") < 0 {
				return -1
			}
			break
		}
	} else if tokens[operationToken] == "get" && tokenCount == 2 {
		for {
			var buffer bytes.Buffer
			key := tokens[keyToken]
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
			break
		}
	} else if tokens[operationToken] == "stats" {
		for {
			break
		}
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
