package Siege

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetUser(username string, platform string, client Client) (Profile, error) {
	return client.GetUser(username, platform)
}
func (client Client) GetUser(username string, platform string) (Profile, error) {
	uri := fmt.Sprintf("https://public-ubiservices.ubi.com/v2/profiles?nameOnPlatform=%s&platformType=%s", username, platform)
	body, err := client.makeHTTPReq(uri)
	if err != nil {
		return Profile{}, err
	} else {
		prof := make(map[string][]Profile)
		json.Unmarshal([]byte(body), &prof)
		if len(prof["profiles"]) > 0 {
			return prof["profiles"][0], nil
		} else {
			return Profile{}, errors.New("User not found")
		}
	}
}

func GetSeasonsList() (GameSeasons, error) {
	httpClient := &http.Client{}

	req, _ := http.NewRequest("GET", "https://game-rainbow6.ubi.com/assets/data/seasons.7985adeb.json", nil)

	res, err := httpClient.Do(req)
	if err != nil {
		return GameSeasons{}, err
	}

	body, _ := ioutil.ReadAll(res.Body)

	var out GameSeasons
	json.Unmarshal(body, &out)

	return out, nil

}
