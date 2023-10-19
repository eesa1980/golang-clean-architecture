package interfaces

import (
	"os"
)

type IFileHandlerService interface {
	LoadFile(filename string) (*os.File, error)
	ToJson() (*any, error)
	Close()
}
