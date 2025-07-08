package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alkeshnikalje/go-pokedox/internal/pokecache"
)

func GetLocationAreas(url string,c *pokecache.Cache) (*LocationAreaResponse, error) {
	var locationArea LocationAreaResponse
	
	val,ok := c.Get(url) 
	
	if ok {
		if err := json.Unmarshal(val, &locationArea); err != nil {
			return nil, err
		}			
		return &locationArea,nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.Add(url,data)

	if err := json.Unmarshal(data, &locationArea); err != nil {
		return nil, err
	}

	return &locationArea, nil
}


func GetArea(name string, c *pokecache.Cache) (*Area, int, error) {
	areaUrl := "https://pokeapi.co/api/v2/location-area/" + name
	var area Area
	val,ok := c.Get(name)
	if ok {
		if err := json.Unmarshal(val,&area); err != nil {
			return nil,0, err
		}
		return &area,0,nil
	}
	resp, err := http.Get(areaUrl)

	if err != nil {
		return nil, resp.StatusCode, err
	}

	defer resp.Body.Close()
	
	if resp.StatusCode == 404 {
		return nil, resp.StatusCode,fmt.Errorf("received 404")
	}
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,0, err
	}
	c.Add(name,data)

	if err := json.Unmarshal(data,&area); err != nil {
		return nil,0,err
	}
	return &area,resp.StatusCode,nil	
}

func GetPokemon (name string, c *pokecache.Cache) (*Pokemon,int, error) {
	pokemonInfoUrl := "https://pokeapi.co/api/v2/pokemon/"	+ name
	var pokemonInfo Pokemon
	val,ok := c.Get(name)
	if ok {
		if err := json.Unmarshal(val,&pokemonInfo); err != nil {
			return nil,0, err
		}
		return &pokemonInfo,0,nil
	}

	resp,err := http.Get(pokemonInfoUrl)
	if err != nil {
		return nil,0,err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, resp.StatusCode,fmt.Errorf("received 404")
	}
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,0, err
	}
	c.Add(name,data)	

	if err := json.Unmarshal(data,&pokemonInfo); err != nil {
		return nil,0,err
	}
	return &pokemonInfo,resp.StatusCode,nil
}

















































