package storage

import (
	"fmt"
	"gstat/internal/protocol"
	"os"
	"sync"
	"time"
)

func CreateHistoryDirectory(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func Save(dir string, result protocol.Result) bool {
	file, err := os.Create(fmt.Sprintf("%s/%s.txt", dir, time.Now().Format("20060102150405")))
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

func Read(dir string) {
	var wg sync.WaitGroup
	reads := make(chan struct{}, 3)

	var readFiles func(dir string)
	readFiles = func(dir string) {
		defer wg.Done()

		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("Error reading history file:", err)
			return
		}

		for _, entry := range entries {
			var fileInfo os.FileInfo
			fileInfo, err = entry.Info()
			if err != nil {
				fmt.Println("Error reading file:", entry.Name(), err)
				continue
			}

			if fileInfo.IsDir() {
				wg.Add(1)
				select {
				case reads <- struct{}{}:
					go func() {
						readFiles(fmt.Sprintf("%s/%s", dir, entry.Name()))
						<-reads
					}()
				default:
					fmt.Println("Channel full, processing sequentially")
					readFiles(fmt.Sprintf("%s/%s", dir, entry.Name()))
				}
				continue
			}

			var file []byte
			file, err = os.ReadFile(fmt.Sprintf("%s/%s", dir, entry.Name()))
			if err != nil {
				fmt.Println("Error reading content of file:", fileInfo.Name(), err)
				continue
			}

			fmt.Println(fileInfo.Name())
			fmt.Println(string(file) + "\n")
		}
	}

	wg.Add(1)
	readFiles(dir)
	wg.Wait()
}
