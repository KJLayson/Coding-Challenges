// Package weather : This is the weather package providing you with opportunities to do weather like things using weather like functions for weather like purposes.
package weather

// CurrentCondition : This variable represents the current condition of the weather in the form of a string. Why a single string? Couldn't afford an entire spool.
var CurrentCondition string

// CurrentLocation: This variable represent the current location, of which we are measuring the weather. Not to be confused with the weather in Nantucket.
var CurrentLocation string

// Forecast : The Forcast function returns the location and its current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
