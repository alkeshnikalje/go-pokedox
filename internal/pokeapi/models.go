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

// structs to unmarshal pokemon information. 

type Pokemon struct {
	BaseExp int 			`json:"base_experience"`
	Weight 	int 			`json:"weight"`
	Name	string			`json:"name"`	
	Height 	int 			`json:"height"`
	Stats   []StatResponse	`json:"stats"`
	Types   []TypeResponse	`json:"types"`

}

type StatResponse struct {
	BaseStart int		`json:"base_stat"`
	Stat 	  StatName	`json:"stat"`	

}

type StatName struct {
	Name string		`json:"name"`		
}

type TypeResponse struct {
	Type TypeName 	`json:"type"`
}

type TypeName struct {
	Name string		`json:"name"`
}



































