package models

import (
	"mime/multipart"
)

type Chunk struct {
	UploadID      string // unique id for the current upload.
	ChunkNumber   int32
	TotalChunks   int32
	TotalFileSize int64 // in bytes
	Filename      string
	Data          *multipart.Part
	UploadDir     string
	ByteData      []byte
}

type Message struct {
	ToUserID string
	Message  string
}
