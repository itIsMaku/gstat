package main

import (
	"database/sql"
	"fmt"
	"gstat/internal/configuration"
	"gstat/internal/database"
	"gstat/internal/http"
	"gstat/internal/httpserver"
	"gstat/internal/interval"
	"gstat/internal/protocol"
	"gstat/internal/storage"
	"gstat/internal/tcpudp"
	"os"
)

const ConfigFile = "config.json"

func printCommandsHelp() {
	fmt.Println(`gstat - simple CLI tool to learn Go
			
Usage:
	gstat check <protocol> <http/ip:port>
	gstat history
		`)
}

func main() {
	err := configuration.CreateConfig(ConfigFile)
	if err != nil {
		fmt.Println("Error creating configuration:", err)
		os.Exit(1)
		return
	}

	config, err := configuration.LoadConfig(ConfigFile)
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
		return
	}

	historyDir := config.HistoryDir

	storage.CreateHistoryDirectory(historyDir)

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

		storage.Save(historyDir, res)

		fmt.Println("Result:", res)
		os.Exit(0)
	case "history":
		storage.Read(historyDir)
	case "interval":
		fmt.Println("Starting interval checks...")

		db, err := database.Open(config.Database)
		if err != nil {
			fmt.Println("Error opening database:", err)
			os.Exit(1)
		}

		defer func(db *sql.DB) {
			err := database.Close(db)
			if err != nil {
				fmt.Println("Error closing database:", err)
				os.Exit(1)
			}
		}(db)

		fmt.Println("Successfully connected to the database.")

		err = database.EnsureTables(db)
		if err != nil {
			fmt.Println("Error ensuring database tables:", err)
			os.Exit(1)
		}

		go interval.StartInterval(config.Interval, db)

		err = httpserver.Start(config.Http, db)
		if err != nil {
			fmt.Println("Error in HTTP server:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command!", os.Args[1:])
		printCommandsHelp()
		os.Exit(1)
	}
}
