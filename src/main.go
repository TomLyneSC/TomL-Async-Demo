package main

import (
	"masterchef/src/scenario"
)

func main() {
	scenario.Welcome()

	/*
		I will be using a baking analogy to explain the difference between
		synchronous and asynchronous code, showing the benefits, as well as
		the potential disadvantages
	*/

	// Example 1
	/*
		When baking synchronously, the baker can only go onto the next instruction
		when the current instruction has fully completed. Unfortunately, this means
		they are wasting lots of time waiting for the oven to heat.
	*/

	//scenario.BakingSynchronously()

	// Let's see if we can speed this up...
	// Example 2
	/*
		When baking asynchronously, the baker can set a task going in the background.
		This means that *while* the oven is heating up, they can carry on with other
		tasks, such as mixing the ingredients.
	*/

	//scenario.BakingAsynchronously()

	// This causes a significant speed increase! But what happens when two bakers want
	// to use the same oven?

	// Example 3
	/*
		If we put the bakers into a queue for the oven, we get no problems at all!
	*/

	//scenario.BakersUsingSameOven()

	// But if we have no orderly queue in place and the bakers are fighting for control
	// of the oven, disaster is inevitable...
	// Example 4
	/*
		When multiple bakers are asynchronously using the same oven, things can get
		messy and unpredictable. Depending on when the bakers start their bakes, the
		end results can become wildly different!!!
	*/

	//scenario.BakersUsingSameOvenAsync()
}
