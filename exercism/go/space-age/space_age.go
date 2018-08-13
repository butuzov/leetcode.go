package space

type Planet string

const (
	SECONDS = 31557600
)

const (
	YEAR_IN_SECONDS_MERCURY = SECONDS * 0.2408467
	YEAR_IN_SECONDS_VENUS   = SECONDS * 0.61519726
	YEAR_IN_SECONDS_EARTH   = SECONDS * 1.
	YEAR_IN_SECONDS_MARS    = SECONDS * 1.8808158
	YEAR_IN_SECONDS_JUPITER = SECONDS * 11.862615
	YEAR_IN_SECONDS_SATURN  = SECONDS * 29.447498
	YEAR_IN_SECONDS_URANUS  = SECONDS * 84.016846
	YEAR_IN_SECONDS_NEPTUNE = SECONDS * 164.79132
)

var Planets map[Planet]float64 = map[Planet]float64{
	Planet("Venus"):   YEAR_IN_SECONDS_VENUS,
	Planet("Mercury"): YEAR_IN_SECONDS_MERCURY,
	Planet("Earth"):   YEAR_IN_SECONDS_EARTH,
	Planet("Mars"):    YEAR_IN_SECONDS_MARS,
	Planet("Jupiter"): YEAR_IN_SECONDS_JUPITER,
	Planet("Saturn"):  YEAR_IN_SECONDS_SATURN,
	Planet("Uranus"):  YEAR_IN_SECONDS_URANUS,
	Planet("Neptune"): YEAR_IN_SECONDS_NEPTUNE,
}

func Age(ageSec float64, p Planet) float64 {

	for planet, yearSec := range Planets {
		if planet == p {
			return ageSec / yearSec
		}
	}

	return 0
}
