package pokeapi

type RespShallowPokemonList struct {
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Height  int           `json:"height"`
	Weight  int           `json:"weight"`
	BaseExp int           `json:"base_experience"`
	Name    string        `json:"name"`
	Stats   []PokemonStat `json:"stats"`
	Types   []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	Stat     StatInfo `json:"stat"`
}

type StatInfo struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
}
