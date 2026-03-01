package ipapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPAPI struct {
	IP          string  `json:"query"`
	Hostname    string  `json:"hostname"`
	City        string  `json:"city"`
	Region      string  `json:"region"`
	CountryCode string  `json:"countryCode"`
	CountryName string  `json:"country"`
	RegionName  string  `json:"regionName"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	Org         string  `json:"org"`
	Postal      string  `json:"zip"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	ASN         string  `json:"as"`
}

type IPInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func Fetch(ip string) (*IPAPI, error) {
	fmt.Println(ip)
	resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ipapi: upstream returned status %d", resp.StatusCode)
	}

	var info IPAPI
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}
