package main

import (
	"fmt"
	"masterchef/src/baker"
	"sync"
	"time"
)

// ***************************************************************//
// Play around with the following variable and see what happens! //
// ***************************************************************//
var secondsToWaitBetweenRequests int = 14

func main() {
	var wg sync.WaitGroup

	baker1 := baker.Baker{
		Wait: &wg,
		Name: "Chris",
	}
	wg.Add(1)

	baker2 := baker.Baker{
		Wait: &wg,
		Name: "Sarah",
	}
	wg.Add(1)

	go baker1.MakeShortbread()

	time.Sleep(time.Second * time.Duration(secondsToWaitBetweenRequests))
	go baker2.MakeCake()

	wg.Wait()
	fmt.Println("Done")
}
