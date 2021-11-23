package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JokeClient struct {
	url string
}

func NewJokeClient(url string) *JokeClient {
	return &JokeClient {
		url: url,
	}
}

func (jokeClient *JokeClient) GetJoke() (*api.JokeResponse, error)  {
	url := jokeClient.url

	response, http_error := http.Get(url)

	if http_error != nil {
		return nil, http_error
	} else if response.StatusCode != http.StatusOk {
		return nil, fmt.Errorf("Api request error. Status: %v", response.StatusCode)
	}

	var data api.JokeRespose

	decoder_error := json.NewDecoder(response.Body).Decode(&data)

	if decoder_error != nil {
		return nil, decoder_error
	}

	return &data, nil
}

