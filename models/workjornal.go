package models

import "time"

type OperationType int

const (
	Coming OperationType = iota
	Leaving
)

type JornalOperation struct {
	Id               int           `json:"operation_id"`
	UserId           int           `json:"user_id"`
	RecognizedUserID int           `json:"user_id_recognized"`
	OperationDate    time.Time     `json:"operation_date"`
	OperationType    OperationType `json:"operation_type"`
}
