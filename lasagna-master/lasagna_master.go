package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, avgPrepTime int) int {

	if avgPrepTime == 0 {

		return len(layers) * 2
	}

	return len(layers) * avgPrepTime
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {

	noodles := 0
	sauce := 0.0

	for _, j := range layers {

		if j == "noodles" {
			noodles += 1
		}

		if j == "sauce" {

			sauce += 1
		}

	}

	return noodles * 50, sauce * 0.2

}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendsList []string, myList []string) {

	myList[len(myList)-1] = friendsList[len(friendsList)-1]

}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(amount []float64, total int) []float64 {

	var newAmount []float64

	for _, j := range amount {

		newAmount = append(newAmount, ((j / 2) * float64(total)))
	}
	return newAmount
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
