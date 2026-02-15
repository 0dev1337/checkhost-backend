package ipinfo

import (
	"checkhost-backend/internal/routes/public/info/ipapi"
	"fmt"
)

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

// FromIPAPI converts an ipapi.IPAPI response to IPInfo
func FromIPAPI(apiResp *ipapi.IPAPI) *IPInfo {
	return &IPInfo{
		IP:       apiResp.IP,
		Hostname: apiResp.Hostname,
		City:     apiResp.City,
		Region:   apiResp.Region,
		Country:  apiResp.Country,
		Location: fmt.Sprintf("%.6f,%.6f", apiResp.Latitude, apiResp.Longitude),
		Org:      apiResp.Org,
		Postal:   apiResp.Postal,
		Timezone: apiResp.Timezone,
	}
}
