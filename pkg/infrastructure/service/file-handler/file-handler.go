package filehandlerservice

import (
	"clean-architecture/pkg/application/common/interfaces"
	"encoding/json"
	"fmt"
	"os"
)

type FileHandlerService struct {
}

var (
	file *os.File
)

func (f *FileHandlerService) LoadFile(filename string) (*os.File, error) {
	opened, err := os.Open(filename)

	file = opened

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Opening json file")
	}

	return file, err
}

func (f *FileHandlerService) Close() {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Closing json file")
		}
	}(file)
}

func (f *FileHandlerService) ToJson() (*any, error) {
	return func(file *os.File) (*any, error) {
		decoder := json.NewDecoder(file)

		var j *any
		err := decoder.Decode(&j)

		if err != nil {
			panic(err)
		}

		return j, err

	}(file)
}

func New() interfaces.IFileHandlerService {
	return &FileHandlerService{}
}
