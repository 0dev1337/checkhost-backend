package ipapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetIPAPIInfo(ip string) (*IPAPI, error) {
	response, err := http.Get(fmt.Sprintf("https://ipapi.co/%s/json/", ip))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch IP info: status code %d", response.StatusCode)
	}
	var info IPAPI
	err = json.NewDecoder(response.Body).Decode(&info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
