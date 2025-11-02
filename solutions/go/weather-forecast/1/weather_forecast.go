// Package weather provides weather forecast.
package weather

var (
    // CurrentCondition represent the condition of the weather.
	CurrentCondition string 
    // CurrentLocation represent the current location (city).
	CurrentLocation  string 
)

// Forecast function gives forecast weather from city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
