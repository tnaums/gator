package config

import (
	"os"
	"log"
	"bufio"
	"io"
	"fmt"
	"encoding/json"
)

type Config struct {
	DbURL string `json:"db_url"`
}

func READ() Config {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	filePath := homePath + "/.gatorconfig.json"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024)
	var myStruct Config
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(string(buf[:n]))
		fmt.Println("\n")

	err2 := json.Unmarshal(buf[:n], &myStruct)
	if err2 != nil {
		log.Fatal(err)
	}		

	}

	return myStruct
}
