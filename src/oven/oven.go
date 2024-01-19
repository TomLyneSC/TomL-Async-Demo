package oven

import (
	"time"
)

type Oven struct {
	Contents string
	Temp     int
}

var instance *ExternalOvenService

type ExternalOvenService struct {
	oven *Oven
}

func GetOven() *Oven {
	if instance == nil {
		service := ExternalOvenService{}
		oven := Oven{}
		instance = &service
		instance.oven = &oven
	}

	return instance.oven
}

func (o *Oven) SetTemp(temp int) {
	time.Sleep(time.Second * 4)
	o.Temp = temp
}

func (o *Oven) PutIntoOven(item string) {
	time.Sleep(time.Second * 2)
	o.Contents = item
}

func (o *Oven) TakeOutOfOven() string {
	time.Sleep(time.Second * 2)
	item := o.Contents
	o.Contents = ""
	return item
}

func (o *Oven) CheckOvenTemp() int {
	return o.Temp
}
