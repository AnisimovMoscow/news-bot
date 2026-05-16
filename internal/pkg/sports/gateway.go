package sports

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	url     = "https://www.sports.ru/gql/graphql/"
	timeout = 10 * time.Second
)

func request(query string, resp any, variables map[string]any) error {
	payload := struct {
		Query     string         `json:"query"`
		Variables map[string]any `json:"variables"`
	}{
		Query:     query,
		Variables: variables,
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := getClient()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return err
	}

	return nil
}

func getClient() *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}
