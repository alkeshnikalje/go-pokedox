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
















