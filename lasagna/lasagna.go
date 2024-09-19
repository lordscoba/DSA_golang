package lasagna

// TODO: define the 'OvenTime' constant

const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	// panic("RemainingOvenTime not implemented")
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	// panic("PreparationTime not implemented")

	// Time taken in Minutes
	var timeTaken int = 2

	return numberOfLayers * timeTaken
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	// panic("ElapsedTime not implemented")

	// define preparation time
	preparationTime := PreparationTime(numberOfLayers)

	// Total Elapsed time
	totalElapsedTime := preparationTime + actualMinutesInOven

	return totalElapsedTime
}
