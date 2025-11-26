// Package weather forecast the weather of the given cities.
package weather

var (
    // CurrentCondition declared with the condition of the city.
	CurrentCondition string
    // CurrentLocation declare the city needed to be forecasted.
	CurrentLocation  string
)
// Forecast Return the forecase for the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
