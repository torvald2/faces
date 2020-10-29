package models

import (
	"time"
)

type Profile struct {
	Id          int       `json:"profile_id"`
	Descriptor  []float64 `json:"-"`
	Name        string    `json:"user_name"`
	ImageUrl    string    `json:"avatar_url"`
	ShopNum     int       `json:"shop_num"`
	CreatedDate time.Time `json:"created_date"`
}
