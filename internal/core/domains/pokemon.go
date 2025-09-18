package domains

type Pokemon struct {
	Name      string   `json:"name"`
	Height    int      `json:"height"`
	Weight    int      `json:"weight"`
	Types     []string `json:"types"`
	Abilities []string `json:"abilities"`
}

type PokemonResponse struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Height int    `json:"height"`
	Types  []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
}
