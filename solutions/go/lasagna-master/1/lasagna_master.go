package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(ingredients []string, time int) int {
    if time == 0 {
        return len(ingredients) * 2
    }
    return len(ingredients) * time
}  

// TODO: define the 'Quantities()' function
func Quantities(ingredients []string) (int,float64) {
    noodles := 0
    sauce := 0.0
    for _, val := range ingredients {
        switch val {
            case "noodles": noodles += 50
            case "sauce": sauce += 0.2
        }
    }
    return noodles, sauce
} 

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
    myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64 {
    result := []float64{}
    for _, val := range quantities {
        result = append(result, val*(float64(scale)/2.0))
    }
    return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
