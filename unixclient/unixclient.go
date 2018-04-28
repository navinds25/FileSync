package unixclient

import (
	"log"
	"net"
)

//for pinging unixserver
func UnixClient() {
	arg := "ping"
	c, err := net.Dial("unix", "/tmp/example.sock")

	if err != nil {
		panic(err)
	}

	_, err = c.Write([]byte(arg))

	if err != nil {
		log.Println(err)
	}
}
