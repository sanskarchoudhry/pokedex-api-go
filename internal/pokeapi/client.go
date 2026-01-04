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
	var resp ResShallowGenerations
	err := c.fetchJSON("/generation", &resp)
	return resp, err
}

func (c *Client) GetGeneration(nameOrID string) (GenerationDetails, error) {
	var resp GenerationDetails
	err := c.fetchJSON("/generation/"+nameOrID, &resp)
	return resp, err
}

func (c *Client) ListTypes() (ResShallowTypes, error) {
	var resp ResShallowTypes
	err := c.fetchJSON("/type", &resp)
	return resp, err
}

func (c *Client) GetType(name string) (TypeDetails, error) {
	var resp TypeDetails
	err := c.fetchJSON("/type/"+name, &resp)
	return resp, err
}

func (c *Client) fetchJSON(endpoint string, target interface{}) error {
	fullURL := c.baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, target)
}
