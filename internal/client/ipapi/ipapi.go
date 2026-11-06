package ipapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

type Info struct {
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

func Fetch(ip string) (*Info, error) {
	resp, err := httpClient.Get(fmt.Sprintf("http://ip-api.com/json/%s", url.PathEscape(ip)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ipapi: upstream returned status %d", resp.StatusCode)
	}

	var info Info
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}
