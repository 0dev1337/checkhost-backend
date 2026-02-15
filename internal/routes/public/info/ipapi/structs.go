package ipapi

type IPAPI struct {
	IP          string  `json:"ip"`
	Hostname    string  `json:"hostname"`
	City        string  `json:"city"`
	Region      string  `json:"region"`
	Country     string  `json:"country"`
	CountryName string  `json:"country_name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Org         string  `json:"org"`
	Postal      string  `json:"postal"`
	Timezone    string  `json:"timezone"`
}
