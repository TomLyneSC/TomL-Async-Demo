package baker

import (
	"fmt"
	"masterchef/src/oven"
	"sync"
	"time"
)

type Baker struct {
	Name string
	Wait *sync.WaitGroup
}

func ServeItem(expectedItem string, expectedTemp int, actualItem string, actualTemp int, name string) {
	nameStr := name
	if actualItem == "" {
		fmt.Println(name, "took out ...nothing?!?")
		return
	}

	if actualItem != expectedItem {
		nameStr = nameStr + " expected " + expectedItem + ", but actually"
	}

	if actualTemp < expectedTemp {
		fmt.Println(nameStr, "took out an underbaked", actualItem)
	} else if actualTemp > expectedTemp {
		fmt.Println(nameStr, "took out a burnt", actualItem)
	} else {
		fmt.Println(nameStr, "took out a perfectly delicious", actualItem)
	}
}

func (b *Baker) MakeShortbread() {
	defer b.Wait.Done()

	targetItem := "shortbread"
	targetTemp := 180
	o := oven.GetOven()
	o.SetTemp(targetTemp)
	fmt.Println(b.Name, "set oven to", targetTemp)
	time.Sleep(time.Second * 1)
	fmt.Println(b.Name, "is busy mixing ingredients")
	o.PutIntoOven(targetItem)
	fmt.Println(b.Name, "put", targetItem, "into oven")

	time.Sleep(time.Second * 5)

	shortbread := o.TakeOutOfOven()
	temp := o.Temp
	ServeItem(targetItem, targetTemp, shortbread, temp, b.Name)
}

func (b *Baker) MakeCake() {
	defer b.Wait.Done()

	targetItem := "cake"
	targetTemp := 240
	o := oven.GetOven()
	o.SetTemp(targetTemp)
	fmt.Println(b.Name, "set oven to", targetTemp)
	time.Sleep(time.Second * 1)
	fmt.Println(b.Name, "is busy mixing ingredients")
	//MixIngredients()
	o.PutIntoOven(targetItem)
	fmt.Println(b.Name, "put", targetItem, "into oven")

	time.Sleep(time.Second * 5)

	cake := o.TakeOutOfOven()
	temp := o.Temp
	ServeItem(targetItem, targetTemp, cake, temp, b.Name)
}
