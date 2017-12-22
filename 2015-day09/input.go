package main

type city string
type distancesByCity map[city]int

var testInput = map[city]distancesByCity{
	"London": {
		"Dublin":  464,
		"Belfast": 518,
	},
	"Dublin": {
		"Belfast": 141,
	},
}

var input = map[city]distancesByCity{
	"Tristram": {
		"AlphaCentauri": 34,
		"Snowdin":       100,
		"Tambi":         63,
		"Faerun":        108,
		"Norrath":       111,
		"Straylight":    89,
		"Arbre":         132,
	},
	"AlphaCentauri": {
		"Snowdin":    4,
		"Tambi":      79,
		"Faerun":     44,
		"Norrath":    147,
		"Straylight": 133,
		"Arbre":      74,
	},
	"Snowdin": {
		"Tambi":      105,
		"Faerun":     95,
		"Norrath":    48,
		"Straylight": 88,
		"Arbre":      7,
	},
	"Tambi": {
		"Faerun":     68,
		"Norrath":    134,
		"Straylight": 107,
		"Arbre":      40,
	},
	"Faerun": {
		"Norrath":    11,
		"Straylight": 66,
		"Arbre":      144,
	},
	"Norrath": {
		"Straylight": 115,
		"Arbre":      135,
	},
	"Straylight": {
		"Arbre": 127,
	},
}
