package storage

import (
	"fmt"
	"gstat/internal/protocol"
	"os"
	"time"
)

func CreateHistoryDirectory() {
	if _, err := os.Stat("history"); os.IsNotExist(err) {
		err = os.Mkdir("history", 0755)
		if err != nil {
			panic(err)
		}
	}
}

func Save(result protocol.Result) bool {
	file, err := os.Create(fmt.Sprintf("history/%s.txt", time.Now().Format("20060102150405")))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return false
	}

	defer file.Close()

	_, err = file.WriteString(result.String())
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return false
	}

	return true
}
