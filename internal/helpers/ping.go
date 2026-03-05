package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PingResult struct {
	IP       string `json:"ip"`
	Message  string `json:"message"`
	RTT      int32  `json:"rtt_ms"`
	Success  bool   `json:"success"`
	Location string `json:"location"`
}

func PingICMP(target string) (PingResult, error) {
	response, err := http.Get(fmt.Sprintf("http://localhost:3002/ping/%s", target))
	if err != nil {
		return PingResult{}, err
	}
	defer response.Body.Close()

	var result PingResult
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return PingResult{}, err
	}

	return result, nil
}
