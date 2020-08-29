package main

import (
	"log"
	"net"
	"strings"
)

type network struct {
	listener net.Listener
}

var networkInfo network

func sendNetworkRequest(conn net.Conn, buffer []byte, size int) int {
	length, err := conn.Write(buffer)
	if err != nil {
		return -1
	}
	if length != size {
		return -1
	}
	return 0
}

/* Request format:
 * <command> <key> [<valueLength>]\r\n[<value>\r\n]
 */
func processNetworkRequest(conn net.Conn) {
	var position int = -1
	addr := conn.RemoteAddr().String()
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
			if position = strings.Index(data, "\\r\\n"); position < 0 {
				position = strings.Index(data, "\\n")
			}

			if position < 0 {
				/* TODO :: change error write after define error code */
				response := "INVALID command foramt"
				_, err = conn.Write([]byte(response))
				if err != nil {
					log.Printf("Network Write failure.\n")
					goto close
				}
			} else {
				/* process command */
				ret := processCommand(data, position)
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
	log.Printf("Close connection. Client Addr = %s\n", addr)
	conn.Close()
}

func acceptNetworkProcess() {
	log.Printf("Start accept network process\n")
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
