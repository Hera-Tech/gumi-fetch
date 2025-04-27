package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:8000/v1/shows" // Adjust base URL as needed

// Helper function to make HTTP requests and return response
func makeRequest(t *testing.T, method, url string, body interface{}, expectedStatusCode int) *http.Response {
	var jsonBody []byte
	if body != nil {
		var err error
		jsonBody, err = json.Marshal(body)
		assert.NoError(t, err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	// Assert the status code
	assert.Equal(t, expectedStatusCode, resp.StatusCode)

	return resp
}

type ShowListResponse struct {
	Data []map[string]interface{} `json:"data"`
}

func TestShowController(t *testing.T) {
	// Test case: Register a new show
	t.Run("register show", func(t *testing.T) {
		// Prepare the payload for registering a new show
		showPayload := map[string]interface{}{
			"id":           11,
			"title":        "New Show",
			"source":       "Source",
			"source_id":    "new_show_1",
			"main_picture": "https://cdn.myanimelist.net/images/anime/1770/97704.jpg",
		}

		// Register the show
		resp := makeRequest(t, "POST", baseURL, showPayload, http.StatusCreated)

		// Check the response body (you can modify according to your response format)
		var responseBody map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&responseBody)
		assert.NoError(t, err)
	})

	// Test case: List all shows
	t.Run("list shows", func(t *testing.T) {
		// Make a GET request to list all shows
		resp := makeRequest(t, "GET", baseURL, nil, http.StatusOK)

		// Check that the response body contains the expected list of shows
		var response ShowListResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Data)
	})

	// Test case: Unregister a show
	t.Run("unregister show", func(t *testing.T) {
		// Make a DELETE request to unregister the show with ID 1
		url := fmt.Sprintf("%s/11", baseURL)
		resp := makeRequest(t, "DELETE", url, nil, http.StatusNoContent)

		// Check that the response body is empty and status code is OK
		assert.Empty(t, resp.Body)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})

	// Test case: Try to list shows again after unregistering
	t.Run("list shows after unregistering", func(t *testing.T) {
		// Make a GET request to list all shows
		resp := makeRequest(t, "GET", baseURL, nil, http.StatusOK)

		// Check that the response body does not contain the deleted show
		var response ShowListResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.NotContains(t, response.Data, map[string]interface{}{"id": 1})
	})
}
