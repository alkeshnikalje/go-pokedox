package pokeapi

// structs to unmarshal locations response

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int             `json:"count"`
	Next     string          `json:"next"`
	Previous *string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

// structs to unmarshal specific area's response

type Area struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonResponse `json:"pokemon"` 
}

type PokemonResponse struct {
	Name string `json:"name"`
	Url	 string `json:"url"`
}

// struct to unmarshal pokemon information (baseExperience)

type PokemonBaseExp struct {
	BaseExp int `json:"base_experience"`
}

