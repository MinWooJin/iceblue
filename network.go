package main

import (
	"log"
	"net"
	"strings"
	"sync"
)

type network struct {
	listener net.Listener
}

var networkInfo network

func sendNetworkRequest(conn net.Conn, data string) int {
	addr := conn.RemoteAddr().String()
	/* DEBUG log level */
	log.Printf("Send network request. Conn = %s\n", addr)

	byteData := []byte(data)
	byteLength := len(byteData)
	length, err := conn.Write(byteData)
	if err != nil || length != byteLength {
		log.Printf("Send network request failed. Conn = %s\n", addr)
		return -1
	}
	return 0
}

func readNetworkRequest(conn net.Conn, length int) ([]byte, int) {
	buffer := make([]byte, length+4 /* +4 is CRLF length*/)
	readBuffer := make([]byte, length+4)
	remainLength := length

	for {
		readLength, err := conn.Read(readBuffer)
		if err != nil {
			log.Printf("Network Read next line failure.\n")
			return nil, -1
		}
		buffer = append(buffer, readBuffer[:readLength]...)
		remainLength -= readLength
		if remainLength < 0 {
			log.Printf("Network Read next line failure.\n")
			return nil, -1
		} else if remainLength == 0 {
			break
		}
		/* reset readBuffer slice */
		readBuffer = readBuffer[:0]
	}

	return buffer, 0
}

/* Request format:
 * <command> <key> [<valueLength>]\r\n[<value>\r\n]
 */
func processNetworkRequest(conn net.Conn) {
	var position int = -1
	var endPosition int = -1
	addr := conn.RemoteAddr().String()
	/* DEBUG log leven */
	log.Printf("Read network request. Conn = %s\n", addr)

	buffer := make([]byte, 4096)
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Network Read failure.\n")
			goto close
		}
		if length > 0 {
			data := string(buffer[:length])
			/* check protocol.. CRLF(\r\n) or LF(\n) */
			position = strings.Index(data, "\\r\\n")
			if position < 0 {
				position = strings.Index(data, "\\n")
				endPosition = position + 2
			} else {
				endPosition = position + 4
			}

			if position < 0 {
				/* TODO :: change error write after define error code */
				if sendNetworkRequest(conn, "INVALID command format") < 0 {
					goto close
				}
			} else {
				/* process command */
				ret := processCommand(conn, data, position, endPosition)
				if ret < 0 {
					log.Printf("Failure process command.\n")
					goto close
				}
			}
		} else {
			log.Printf("Network Read length is zero.\n")
			goto close
		}
	}

close:
	/* DEBUG log level */
	log.Printf("Close connection. Client Addr = %s\n", addr)
	conn.Close()
}

func acceptNetworkProcess(waitGroups *sync.WaitGroup) {
	log.Printf("Start accept network process\n")
	defer waitGroups.Done()

	for networkInfo.listener != nil {
		conn, err := networkInfo.listener.Accept()
		if err != nil {
			log.Panic(err)
		}
		go processNetworkRequest(conn)
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
