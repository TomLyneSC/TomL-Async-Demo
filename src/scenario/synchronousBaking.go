package scenario

import (
	"fmt"
	"masterchef/src/baker"
	"masterchef/src/oven"
	"time"
)

func BakingSynchronously() {

	timeBefore := time.Now()
	b := baker.Baker{
		Name: "Simon",
	}

	MakeCake1(b)
	timeAfter := time.Now()

	fmt.Println("Took", timeAfter.Sub(timeBefore), "to bake synchronously")
}

func MakeCake1(b baker.Baker) {
	Bake1(b, "cake", 240)
}

func Bake1(b baker.Baker, item string, temp int) {

	fmt.Println(b.Name, "is ready to go!")
	o := oven.GetOven()

	// Step 1: Heat Oven
	fmt.Println(b.Name, "set oven to", temp)
	o.SetTemp(temp, nil)
	fmt.Println(b.Name, "oven is ready", temp)

	// Step 2: Mix ingredients
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(b.Name, "is busy mixing ingredients...")
	}

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