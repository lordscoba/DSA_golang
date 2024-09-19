// Package weather provides functions to forecast the weather.
package weather

// CurrentCondition is the current weather condition.
var CurrentCondition string

// CurrentLocation is the name of the current location.
var CurrentLocation string

// Forecast returns the current weather condition in the given city
// and the name of the city and takes two string parameters city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
