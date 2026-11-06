package worker

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const requestTimeout = 10 * time.Second

type PingResult struct {
	IP       string `json:"ip"`
	Message  string `json:"message"`
	RTT      int32  `json:"rtt_ms"`
	Success  bool   `json:"success"`
	Location string `json:"location"`
}

type ProbeResult struct {
	Success bool   `json:"success"`
	Address string `json:"address"`
	Elapsed int64  `json:"elapsed_ms"`
	Error   string `json:"error,omitempty"`
}

type Client struct {
	apiKey string
	http   *http.Client
}

func New(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		http:   &http.Client{Timeout: requestTimeout},
	}
}

func (c *Client) Ping(target, node string) (PingResult, error) {
	result, err := get[PingResult](c, node, "ping", target)
	if err != nil {
		return PingResult{}, err
	}
	if !result.Success {
		return PingResult{}, fmt.Errorf("worker: ping failed: %s", result.Message)
	}
	return result, nil
}

func (c *Client) TCP(target, node string) (ProbeResult, error) {
	return get[ProbeResult](c, node, "tcp", target)
}

func get[T any](c *Client, node, probe, target string) (T, error) {
	var zero T

	endpoint := fmt.Sprintf("http://%s/%s/%s?key=%s", node, probe, url.PathEscape(target), url.QueryEscape(c.apiKey))
	resp, err := c.http.Get(endpoint)
	if err != nil {
		return zero, withoutURL(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return zero, fmt.Errorf("worker: %s returned status %d", node, resp.StatusCode)
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return zero, err
	}
	return result, nil
}

func withoutURL(err error) error {
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		return urlErr.Err
	}
	return err
}
