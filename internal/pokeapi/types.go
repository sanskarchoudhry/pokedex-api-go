package pokeapi

type ResShallowGenerations struct {
	Count   int              `json:"count"`
	Results []GenerationItem `json:"results"`
}

type GenerationItem struct {
	GenName string `json:"name"`
	GenUrl  string `json:"ul"`
}

type GenerationDetails struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	MainRegion struct {
		Name string `json:"name"`
	} `json:"main_region"`
}
