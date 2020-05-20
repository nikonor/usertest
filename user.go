package usertest

import (
	"strconv"
	"time"
)

type User interface {
	Show() string
	SetYear(uint)
	CalcAge() uint
}

type usr struct {
	Name   string `json:"name"`
	BDYear uint   `json:"age"`
}

func (u *usr) SetYear(y uint) {
	u.BDYear = y
}

func (u usr) Show() string {
	return "User=" + u.Name + ", age=" + strconv.Itoa(int(u.CalcAge()))
}

func (u usr) CalcAge() uint {
	return uint(time.Now().Year()) - u.BDYear
}

func New(name string, bdyear uint) User {
	usr := usr{
		Name:   name,
		BDYear: bdyear,
	}
	return &usr
}
