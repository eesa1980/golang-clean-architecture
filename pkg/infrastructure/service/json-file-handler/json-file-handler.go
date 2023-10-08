package jsonfilehandler

import (
	"encoding/json"
	"fmt"
	"os"
	"wire-demo-2/pkg/application/common/interfaces"
)

type JSONFileHandler struct {
}

var (
	file *os.File
)

func (c JSONFileHandler) Load() (*json.Decoder, error) {
	opened, err := os.Open("mock-user-data.json")

	file = opened

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Opening json file")
	}

	decoder := json.NewDecoder(file)

	return decoder, nil
}

func (c JSONFileHandler) Close() {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Closing json file")
		}
	}(file)
}

func New() interfaces.IJSONFileHandler {
	return JSONFileHandler{}
}
