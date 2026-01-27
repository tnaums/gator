package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
// package config

// import (
// 	"os"
// 	"log"
// 	"bufio"
// 	"io"
// 	"encoding/json"
// )

// const (
// 	configFileName = ".gatorconfig.json"
// )

// func getConfigFilePath() (string, error) {
// 	homePath, err := os.UserHomeDir()
// 	if err != nil {
// 		log.Fatal(err)
// 		return "", err
// 	}
// 	filePath := homePath + "/" + configFileName
// 	return filePath, nil
// }

// type Config struct {
// 	DbURL string `json:"db_url"`
// 	CurrentUserName string `json:"current_user_name"`
// }

// func write(cfg *Config) error{
// 	filePath, err := getConfigFilePath()
// 	if err != nil {
// 		return err
// 	}

// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 		return err
// 	}
// 	defer file.Close()

// 	j, err := json.Marshal(cfg)
// 	if err != nil {
// 		return err
// 	}
	
// 	_, err2 := file.Write(j)
// 	if err2 != nil {
// 		log.Fatal(err)
// 		return err
// 	}

// 	return nil
// }

// func (c *Config) SetUser(name string) error {
// 	c.CurrentUserName = name
// 	err := write(c)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }


// func READ() (Config, error) {
// 	var myStruct Config
	
// 	filePath, err := getConfigFilePath()
// 	if err != nil {
// 		return Config{}, err
// 	}
	
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 		return Config{}, err
// 	}
// 	defer file.Close()

// 	reader := bufio.NewReader(file)
// 	buf := make([]byte, 1024)

// 	byteSlice := make([]byte, 0)
// 	for {
// 		n, err := reader.Read(buf)
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			log.Fatal(err)
// 			return Config{}, err
// 		}
// 		byteSlice = append(byteSlice, buf[:n]...)


// 	err2 := json.Unmarshal(byteSlice, &myStruct)
// 	if err2 != nil {
// 		log.Fatal(err)
// 		return Config{}, err2
// 	}		
// 	}

// 	return myStruct, nil
// }
