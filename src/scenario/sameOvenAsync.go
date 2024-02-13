package scenario

import (
	"fmt"
	"masterchef/src/baker"
	"masterchef/src/oven"
	"sync"
	"time"
)

func BakersUsingSameOvenAsync() {

	// ***************************************************************//
	// Play around with the following variable and see what happens! //
	// ***************************************************************//
	var secondsToWaitBetweenRequests int = 13

	var wg sync.WaitGroup

	baker1 := baker.Baker{
		Wait: &wg,
		Name: "Sarah",
	}
	wg.Add(1)

	baker2 := baker.Baker{
		Wait: &wg,
		Name: "Chris",
	}
	wg.Add(1)

	go MakeShortbread4(baker1, &wg)
	time.Sleep(time.Second * time.Duration(secondsToWaitBetweenRequests))
	go MakeCake4(baker2, &wg)

	wg.Wait()
	fmt.Println("Done")
}


func MakeCake4(b baker.Baker, s *sync.WaitGroup) {
	Bake4(b, "cake", 240, s)
}

func MakeShortbread4(b baker.Baker, s *sync.WaitGroup) {
	Bake4(b, "shortbread", 180, s)
}

func Bake4(b baker.Baker, item string, temp int, s *sync.WaitGroup) {
	var wg sync.WaitGroup

	fmt.Println(b.Name, "is ready to go!")
	o := oven.GetOven()
	wg.Add(1)

	// Step 1: Heat Oven
	// Step 1.5: Asyncronously set oven temp - and continue doing other stuff
	go o.SetTemp(temp, &wg)
	fmt.Println(b.Name, "set oven to", temp, "and is letting it preheat in the background")
	
	// Step 2: Mix ingredients
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(b.Name, "is busy mixing ingredients...")
	}

	// Step 2.5: Wait for oven to preheat
	fmt.Println(b.Name, "is waiting for the oven to preheat")
	wg.Wait()
	fmt.Println(b.Name, "is ready to put ingredients in")

	// Step 3: Put into oven and wait
	o.PutIntoOven(item)
	fmt.Println(b.Name, "put", item, "into oven")

	for i := 0; i < 4; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println(b.Name, "is waiting for the", item, "to bake...")
	}

	shortbread := o.TakeOutOfOven()
	ovenTemp := o.Temp
	b.ServeItem(item, temp, shortbread, ovenTemp)

	s.Done()
}