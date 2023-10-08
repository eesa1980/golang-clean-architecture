package interfaces

import "encoding/json"

type IJSONFileHandler interface {
	Load() (*json.Decoder, error)
	Close()
}
