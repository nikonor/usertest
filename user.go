package usertest

import (
	"strconv"
	"time"
)

type User interface {
	Show() string
	SetYear(uint)
}

type Usr struct {
	Name   string `json:"name"`
	BDYear uint   `json:"age"`
}

func (u *Usr) SetYear(y uint) {
	u.BDYear = y
}

func (u Usr) Show() string {
	return "User=" + u.Name + ", age=" + strconv.Itoa(time.Now().Year()-int(u.BDYear))
}

func New(name string, bdyear uint) User {
	usr := Usr{
		Name:   name,
		BDYear: bdyear,
	}
	return &usr
}
