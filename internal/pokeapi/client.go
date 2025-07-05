package pokeapi

import (
	"encoding/json"
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


func GetArea(name string, c *pokecache.Cache) (*Area, error) {
	areaUrl := "https://pokeapi.co/api/v2/location-area/" + name
	var area Area
	val,ok := c.Get(name)
	if ok {
		if err := json.Unmarshal(val,&area); err != nil {
			return nil, err
		}
		return &area,nil
	}
	resp, err := http.Get(areaUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	c.Add(name,data)

	if err := json.Unmarshal(data,&area); err != nil {
		return nil,err
	}
	return &area,nil	
}

































