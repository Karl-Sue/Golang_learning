package raindrops
import "fmt"
func Convert(number int) string {
    result := ""
	if number % 3 == 0 {
        result = result + "Pling"
    } 
    if number % 5 == 0 {
        result = result + "Plang"
    } 
    if number % 7 == 0 {
        result = result + "Plong"
    }
    if number % 3 != 0 && number % 5 != 0 && number % 7 != 0 {
        result = result + fmt.Sprintf("%v", number)
    }    
    return result
}
