package main

import (
	"fmt"
)

// gstat check <protocol> <host>

func printCommandsHelp() {
	fmt.Println(`gstat - simple CLI tool to learn Go
			
Usage:
	gstat check <protocol> <http/ip:port>
	gstat history
		`)
}

func main() {
	/*	if len(os.Args) < 2 {
		printCommandsHelp()
		os.Exit(1)
		return
	}*/

	/*switch os.Args[1] {
	case "check":

	case "history":
	default:
		fmt.Println("Unknown command!", os.Args[1:])
		printCommandsHelp()
		os.Exit(1)
		return
	}*/
}
