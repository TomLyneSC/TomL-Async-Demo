package oven

import (
	"sync"
	"time"
)

type Oven struct {
	Contents string
	Temp     int
}

// This oven is *special* in that it only allows one thing inside at a time,
// due to a 'developer oversight'. In this case it is intentional, but serves
// as a good case study of what happens when developers do not consider outside
// factors

// The oven acts as an external service which we have limited control over internally
// This could be a database, S3 bucket, etc
// For ease of access I have implemented it as a singleton

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

func (o *Oven) SetTemp(temp int, wg *sync.WaitGroup) {
	o.Temp = temp
	time.Sleep(time.Second * 8)

	if wg != nil {
		wg.Done()
	}
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
