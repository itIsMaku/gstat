package main

import (
	"fmt"
	"gstat/internal/http"
	"gstat/internal/protocol"
	"gstat/internal/storage"
	"gstat/internal/tcpudp"
	"os"
)

func printCommandsHelp() {
	fmt.Println(`gstat - simple CLI tool to learn Go
			
Usage:
	gstat check <protocol> <http/ip:port>
	gstat history
		`)
}

func main() {
	storage.CreateHistoryDirectory()

	if len(os.Args) < 2 {
		printCommandsHelp()
		os.Exit(1)
		return
	}

	switch os.Args[1] {
	case "check":
		if len(os.Args) < 4 {
			fmt.Println("Usage: gstat check <protocol> <url/ip:port>")
			os.Exit(1)
			return
		}

		targetProtocol := protocol.Protocol(os.Args[2])
		target := os.Args[3]

		fmt.Printf("Checking %s with protocol %s\n", target, targetProtocol)
		var res protocol.Result

		switch targetProtocol {
		case protocol.HTTP:
			res = http.Check(target)
		default:
			res = tcpudp.Check(targetProtocol, target)
		}

		storage.Save(res)

		fmt.Println("Result:", res)
		os.Exit(0)
		return
	case "history":

	default:
		fmt.Println("Unknown command!", os.Args[1:])
		printCommandsHelp()
		os.Exit(1)
		return
	}
}
