package models

import (
	"encoding/json"
	"errors"
	"time"
)

type BadRequestType string

const (
	NotMe              BadRequestType = "not_me"
	MultipleRecognized BadRequestType = "multimatch"
	NotRecognized      BadRequestType = "not_recognized"
	NoFace             BadRequestType = "no_face"
	UserNotFound       BadRequestType = "user_not_found"
)

type BadRequest struct {
	Id              int            `json:"record_id"`
	UserId          int            `json:"user_id"`
	RecognizedUsers []int          `json:"recognized_users"`
	CurrentFace     []byte         `json:"-"`
	RecognizeTime   time.Time      `json:"recognize_time"`
	Shop            int            `json:"shop_id"`
	RequestId       string         `json:"request_id"`
	ErrorType       BadRequestType `json:"error_type"`
}

func (br *BadRequest) UnmarshalJSON(data []byte) error {
	type Aux BadRequest
	var a *BadRequest = (*BadRequest)(br)
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	switch br.ErrorType {
	case NotMe, MultipleRecognized, NotRecognized, NoFace, UserNotFound:
		return nil
	default:
		br.ErrorType = ""
		return errors.New("invalid value for ErrorType")
	}
	return nil
}
