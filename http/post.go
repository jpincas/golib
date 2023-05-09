package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func PostJSONToAPI(apiURL string, data interface{}, result interface{}) error {
	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Make HTTP POST request to API
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Decode response body to result
	if err := json.Unmarshal(body, result); err != nil {
		return err
	}

	return nil
}
