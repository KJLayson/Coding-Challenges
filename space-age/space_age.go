package space

type Planet string

var time_map map[Planet]float64

func Age(seconds float64, planet Planet) float64 {
	if time_map == nil {
		time_map = map[Planet]float64{
			"Mercury": 0.2408467,
			"Venus":   0.61519726,
			"Earth":   1.0,
			"Mars":    1.8808158,
			"Jupiter": 11.862615,
			"Saturn":  29.447498,
			"Uranus":  84.016846,
			"Neptune": 164.79132,
		}
	}

	const earth_seconds = 31557600.0
	planet_factor := time_map[planet]

	if planet_factor != 0 && seconds > 0 {
		return seconds / planet_factor / earth_seconds
	} else {
		return -1.0
	}

}
