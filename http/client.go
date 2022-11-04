package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

var MyClient = &http.Client{Timeout: 10 * time.Second}

// MakeGet
// Perform a http GET request to the input url
func MakeGet(url string, headers map[string]string, target interface{}) error {
	request, _ := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	response, err := MyClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		return errors.New("Error during performing HTTP request, caused by: " + err.Error())
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		bodyString := string(bodyBytes)
		return errors.New("Error caused by: " + bodyString)
	}

	err = json.Unmarshal(bodyBytes, target)
	if err != nil {
		return errors.New("Error during json unmarshall, caused by: " + err.Error())
	}
	return nil
}

// GetAuthToken
// Return Authentication token
func GetAuthToken() string {
	return BasicAuth(os.Getenv("apim_client_username"), os.Getenv("apim_client_password"))
}
