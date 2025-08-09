// Package weather provides a forecast for Goblinocus.
package weather

// CurrentCondition represents the current condition.
var CurrentCondition string
// CurrentLocation represents the current location.
var CurrentLocation string

// Forecast prints the current condition 
// of a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
