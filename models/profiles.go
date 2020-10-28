package models

import (
	"fmt"

	fr "github.com/Kagami/go-face"
)

type Profile struct {
	Id         int           `json:"profile_id"`
	Descriptor fr.Descriptor `json:"-"`
	Name       string        `json:"user_name"`
	ImageUrl   string        `json:"avatar_url"`
	ShopNum    int           `json:"shop_num"`
}

func NewProfile(id int, points []float32, name string, imageId int, shopNum int) Profile {
	var dots fr.Descriptor
	for k, v := range points {
		dots[k] = v
	}
	return Profile{
		Id:         id,
		Descriptor: dots,
		Name:       name,
		ImageUrl:   fmt.Sprintf("/images/%v", imageId),
		ShopNum:    shopNum,
	}
}
