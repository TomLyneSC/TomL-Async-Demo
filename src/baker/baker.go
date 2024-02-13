package baker

import (
	"fmt"
	"sync"
)

type Baker struct {
	Name string
	Wait *sync.WaitGroup
}

func (b *Baker) ServeItem(expectedItem string, expectedTemp int, actualItem string, actualTemp int) {
	nameStr := b.Name
	if actualItem == "" {
		fmt.Println(nameStr, "took out ...nothing?!?")
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
