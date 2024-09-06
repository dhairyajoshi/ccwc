package handlers


	var Handlerfactory = map[string] func(string) {
		"c": getCharacters,
		"l": getLines,
	}
