package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime <= 0 {
		avgPrepTime = 2
	}

	return avgPrepTime * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scale := float64(portions) / 2.0

	var scaledQuantities []float64
	for _, q := range quantities {
		scaledQuantities = append(scaledQuantities, q*float64(scale))
	}

	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
