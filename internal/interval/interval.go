package interval

import (
	"database/sql"
	"fmt"
	"gstat/internal/configuration"
	"gstat/internal/database"
	"gstat/internal/http"
	"gstat/internal/protocol"
	"gstat/internal/tcpudp"
	"sync"
	"time"
)

func StartInterval(config configuration.IntervalConfig, db *sql.DB) {
	seconds := config.Seconds
	targets := config.Targets

	for {
		var wg sync.WaitGroup
		wg.Add(len(targets))

		for _, target := range targets {
			go func() {
				defer wg.Done()

				prot := protocol.Protocol(target.Protocol)
				var res protocol.Result
				switch prot {
				case protocol.HTTP:
					res = http.Check(target.Target)
				default:
					res = tcpudp.Check(prot, target.Target)
				}

				err := database.InsertResult(db, res, time.Now())
				if err != nil {
					fmt.Println("Error inserting result during interval:", err)
					return
				}
			}()
		}

		wg.Wait()

		time.Sleep(time.Duration(seconds) * time.Second)
	}
}
