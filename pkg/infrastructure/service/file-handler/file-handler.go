package filehandlerservice

import (
	"fmt"
	"os"
	"wire-demo-2/pkg/application/common/interfaces"
)

type FileHandlerService struct {
}

var (
	file *os.File
)

func (f FileHandlerService) Load(filename string) (*os.File, error) {
	opened, err := os.Open(filename)

	file = opened

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Opening json file")
	}

	return file, err
}

func (f FileHandlerService) Close() {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Closing json file")
		}
	}(file)
}

func New() interfaces.IFileHandlerService {
	return FileHandlerService{}
}
