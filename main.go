package main

import (
	//"flag"
	//"fmt"
	"os"

	"github.com/navinds25/filesync/unixclient"
	"github.com/navinds25/filesync/unixserver"
	"github.com/navinds25/filesync/grpcserver"
	"github.com/navinds25/filesync/grpcclient"
)

func main () {
	arg := os.Args[1]
	if (arg == "server") {
		go func() {
			unixserver.UnixServer()
		}()
		grpcserver.GrpcServer()
	}
	if (arg == "grpcclient") {
		grpcclient.GrpcClient()
	}
	if (arg == "cli" || arg == "unixclient" || arg == "unix") {
		unixclient.UnixClient()
	}
}
