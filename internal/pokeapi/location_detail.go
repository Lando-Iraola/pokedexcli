package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) LocationDetail(mapName *string) (RespLocationDetail, error) {
	url := baseURL + "/location-area" + *mapName

	if cacheEntry, ok := c.cache.Get(url); ok {
		locationsResp := RespLocationDetail{}
		err := json.Unmarshal(cacheEntry, &locationsResp)
		if err != nil {
			return RespLocationDetail{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationDetail{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationDetail{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationDetail{}, err
	}

	locationsResp := RespLocationDetail{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespLocationDetail{}, err
	}

	c.cache.Add(url, dat)

	return locationsResp, nil
}
