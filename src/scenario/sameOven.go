package scenario

import (
	"fmt"
	"masterchef/src/baker"
	"masterchef/src/oven"
	"time"
)

func BakersUsingSameOven() {
	baker1 := baker.Baker{
		Name: "Sarah",
	}

	baker2 := baker.Baker{
		Name: "Chris",
	}

	MakeShortbread3(baker1)
	MakeCake3(baker2)

	fmt.Println("Done")
}


func MakeCake3(b baker.Baker) {
	Bake3(b, "cake", 240)
}

func MakeShortbread3(b baker.Baker) {
	Bake3(b, "shortbread", 180)
}

func Bake3(b baker.Baker, item string, temp int) {

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