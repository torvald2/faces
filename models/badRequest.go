package models

import (
	"time"
)

type BadRequestType int

const (
	NotMe BadRequestType = iota
	MultipleRecognized
	NotRecognized
	NoFace
)

type BadRequest struct {
	UserId          int
	RecognizedUsers []int
	CurrentFace     []byte
	RecognizeTime   time.Time
	Shop            int
	RequestId       string
	ErrorType       BadRequestType
}
