package allergies

var allergenScores = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) []string {
	var results []string
	for allergen, score := range allergenScores {
		if allergies&score != 0 {
			results = append(results, allergen)
		}
	}
	return results
}

func AllergicTo(allergies uint, allergen string) bool {
	score, exists := allergenScores[allergen]
	if !exists {
		return false
	}
	return allergies&score != 0
}
