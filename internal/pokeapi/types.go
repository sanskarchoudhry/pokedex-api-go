package pokeapi

type ResShallowGenerations struct {
	Count   int              `json:"count"`
	Results []GenerationItem `json:"results"`
}

type GenerationItem struct {
	GenName string `json:"name"`
	GenUrl  string `json:"ul"`
}
