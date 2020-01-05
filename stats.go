package Siege

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (client Client) getStats(platformId string, stats []string, profiles []string) []byte {

	platform := PLATFORMS[platformId]

	uri := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/%s/sandboxes/%s/playerstats2/statistics?populations=%s&statistics=%s",
		platform.Space,
		platform.Sandbox,
		strings.Join(profiles, ","),
		strings.Join(stats, ","),
	)

	out, err := client.makeHTTPReq(uri)
	if err != nil {
		return []byte{}
	}

	return out
}

func (client *Client) GetRankedStats(platformId string, profiles ...string) map[string]map[string]interface{} {

	body := client.getStats(platformId, STATS_RANKED, profiles)

	var output map[string]map[string]map[string]interface{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["results"]

}

func (client *Client) GetCasualStats(platformId string, profiles ...string) map[string]map[string]interface{} {

	body := client.getStats(platformId, STATS_CASUAL, profiles)

	var output map[string]map[string]map[string]interface{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["results"]

}

func (client *Client) GetGeneralStats(platformId string, profiles ...string) map[string]map[string]interface{} {

	body := client.getStats(platformId, STATS_GENERAL, profiles)

	var output map[string]map[string]map[string]interface{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["results"]

}

func (client *Client) GetOperatorStats(platformId string, profiles ...string) map[string]map[string]interface{} {

	body := client.getStats(platformId, STATS_OPERATORS, profiles)

	var output map[string]map[string]map[string]interface{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["results"]

}

func (client *Client) GetGamemodesStats(platformId string, profiles ...string) map[string]map[string]interface{} {

	body := client.getStats(platformId, STATS_GAMEMODES, profiles)

	var output map[string]map[string]map[string]interface{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["results"]

}

func (client *Client) GetWeaponsStats(platformId string, profiles ...string) map[string]map[string]interface{} {

	body := client.getStats(platformId, STATS_WEAPONS, profiles)

	var output map[string]map[string]map[string]interface{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["results"]

}
