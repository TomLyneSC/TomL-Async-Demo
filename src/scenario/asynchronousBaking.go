package scenario

import (
	"fmt"
	"masterchef/src/baker"
	"masterchef/src/oven"
	"sync"
	"time"
)

func BakingAsynchronously() {

	timeBefore := time.Now()
	b := baker.Baker{
		Name: "Alice",
	}

	MakeCake2(b)
	timeAfter := time.Now()

	fmt.Println("Took", timeAfter.Sub(timeBefore), "to bake asynchronously")
}

func MakeCake2(b baker.Baker) {
	Bake2(b, "cake", 240)
}

func Bake2(b baker.Baker, item string, temp int) {
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
}