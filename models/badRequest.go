package models

import (
	"time"
)

type BadRequestType int

const (
	NotMe BadRequestType = iota
	MultipleRecognized
	NotRecognized
)

type BadRequest struct {
	UserId          int
	RecognizedUsers []int
	CurrentFace     []byte
	RecognizeTime   time.Time
	ErrorType       BadRequestType
}
