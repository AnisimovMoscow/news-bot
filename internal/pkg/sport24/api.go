package sport24

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	apiHost = "https://api.sport24.ru"

	timeout   = 10 * time.Second
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/150.0.0.0 Safari/537.36"
)

type Param struct {
	Key   string
	Value string
}

func api(url string, params []Param, resp any) error {
	client := getClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	query := req.URL.Query()
	for _, param := range params {
		query.Add(param.Key, param.Value)
	}
	req.URL.RawQuery = query.Encode()

	req.Header.Set("User-Agent", userAgent)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

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
