package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) int {
	num_layers := len(layers)
	if avg == 0 {
		return (num_layers * 2)
	} else {
		return (num_layers * avg)
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, val := range layers {

		if val == "sauce" {
			sauce += 0.2

		} else if val == "noodles" {
			noodles += 50
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friends, mine []string) []string {
	mine[len(mine)-1] = friends[len(friends)-1]

	return mine
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(portion_size []float64, portions int) []float64 {
	var new_portions []float64
	for i := range portion_size {
		new_portions = append(new_portions, (portion_size[i]*float64(portions))/2)
	}

	return new_portions
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
