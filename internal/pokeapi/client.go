package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func NewClient(baseURL string, timeout time.Duration) *Client {
	return &Client{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) ListGenerations() (ResShallowGenerations, error) {
	fullURL := c.baseURL + "/generation"

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return ResShallowGenerations{}, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return ResShallowGenerations{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return ResShallowGenerations{}, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResShallowGenerations{}, err
	}

	generations := ResShallowGenerations{}
	err = json.Unmarshal(data, &generations)
	if err != nil {
		return ResShallowGenerations{}, err
	}

	return generations, nil
}

func (c *Client) GetGeneration(nameOrID string) (GenerationDetails, error) {
	url := fmt.Sprintf("%s/generation/%s", c.baseURL, nameOrID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GenerationDetails{}, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return GenerationDetails{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return GenerationDetails{}, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return GenerationDetails{}, err
	}

	details := GenerationDetails{}
	err = json.Unmarshal(data, &details)
	if err != nil {
		return GenerationDetails{}, err
	}

	return details, nil
}
