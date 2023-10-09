package interfaces

import (
	"os"
)

type IFileHandlerService interface {
	Load(filename string) (*os.File, error)
	Close()
}
