/*
Package space exports a single Age function which calculates a given Earth age
in year units of another planet of the solar system.
*/
package space

// Planet is an alias of string, and intends to represent the name of the planet
// we're referring to for the years conversion.
type Planet string

const earthYrPerSec float64 = 1.0 / 31557600.0

// Age takes in the number of seconds as a float64 and the planet name as a string
// and calculates the years in units of that planet's orbital period time.
// It is assumed that planet is valid solar system planet name string excluding "Pluto".
func Age(seconds float64, planet Planet) float64 {
	var scaleFactor float64

	switch planet {
	case "Mercury":
		scaleFactor = 4.1520187
	case "Venus":
		scaleFactor = 1.6254949
	case "Earth":
		scaleFactor = 1.0000000
	case "Mars":
		scaleFactor = 0.53168418
	case "Jupiter":
		scaleFactor = 0.084298445
	case "Saturn":
		scaleFactor = 0.033958742
	case "Uranus":
		scaleFactor = 0.011902375
	case "Neptune":
		scaleFactor = 0.0060682809
	}

	return seconds * earthYrPerSec * scaleFactor
}
