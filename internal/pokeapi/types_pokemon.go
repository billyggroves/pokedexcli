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
	Height  int    `json:"height"`
	Weight  int    `json:"weight"`
	BaseExp int    `json:"base_experience"`
	Name    string `json:"name"`
	Stats   []struct {
		Stat struct {
			base int `json:"base_stat"`
			Stat struct {
				Name string `json:"name"`
			} `json:"stat"`
		} `json:"height"`
	} `json:"stats"`
}
