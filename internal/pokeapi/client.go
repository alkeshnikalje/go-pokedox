package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetLocationAreas(url string) (*LocationAreaResponse, error) {
	var locationArea LocationAreaResponse
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &locationArea); err != nil {
		return nil, err
	}

	return &locationArea, nil
}

