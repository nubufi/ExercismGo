package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	age := seconds/31557600
	ratio := 1.
	switch planet{
	case "Mercury":
		ratio = 0.2408467
	case "Venus":
			ratio = 0.61519726
	case "Mars":
			ratio = 1.8808158
	case "Jupiter":
			ratio = 11.862615
	case "Saturn":
			ratio = 29.447498
	case "Uranus":
			ratio = 84.016846
	case "Neptune":
			ratio = 164.79132
	case "Sun":
		ratio = -age
	}
	return age/ratio

}
