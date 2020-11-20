package models

import (
	"encoding/json"
	"errors"
	"time"
)

type OperationType string

const (
	Coming  OperationType = "coming"
	Leaving OperationType = "leaving"
)

type JornalOperation struct {
	Id               int           `json:"operation_id"`
	UserId           int           `json:"user_id"`
	RecognizedUserID int           `json:"user_id_recognized"`
	OperationDate    time.Time     `json:"operation_date"`
	OperationType    OperationType `json:"operation_type"`
	RequestId        string        `json:"request_id"`
	UserName         string        `json:"-"`
	ShopNum          string        `json:"-"`
}

type JornalOperationDB struct {
	UserName      string
	ShopNum       string
	OperationDete time.Time
	OperationType OperationType
}

func (op *JornalOperation) UnmarshalJSON(data []byte) error {
	type Aux JornalOperation
	var a *Aux = (*Aux)(op)
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	switch op.OperationType {
	case Coming, Leaving:
		return nil
	default:
		op.OperationType = ""
		return errors.New("invalid value for OperationType")
	}
	return nil
}
