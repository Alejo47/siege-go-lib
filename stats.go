package Siege

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (ticket UbiTicket) getStats(platformId string, stats []string, profiles []string) []byte {

	platform := PLATFORMS[platformId]

	uri := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/%s/sandboxes/%s/playerstats2/statistics?populations=%s&statistics=%s",
		platform.Space,
		platform.Sandbox,
		strings.Join(profiles, ","),
		strings.Join(stats, ","),
	)

	out, err := ticket.makeHTTPReq(uri)
	if err != nil {
		return []byte{}
	}

	return out
}

func (ticket *UbiTicket) GetRankedStats(platformId string, profiles ...string) map[string]map[string]interface{} {

	body := ticket.getStats(platformId, STATS_RANKED, profiles)

	var output map[string]map[string]map[string]interface{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["results"]

}

func (ticket *UbiTicket) GetCasualStats(platformId string, profiles ...string) map[string]map[string]interface{} {

	body := ticket.getStats(platformId, STATS_CASUAL, profiles)

	var output map[string]map[string]map[string]interface{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["results"]

}

func (ticket *UbiTicket) GetGeneralStats(platformId string, profiles ...string) map[string]map[string]interface{} {

	body := ticket.getStats(platformId, STATS_GENERAL, profiles)

	var output map[string]map[string]map[string]interface{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["results"]

}

func (ticket *UbiTicket) GetOperatorStats(platformId string, profiles ...string) map[string]map[string]interface{} {

	body := ticket.getStats(platformId, STATS_OPERATORS, profiles)

	var output map[string]map[string]map[string]interface{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["results"]

}

func (ticket *UbiTicket) GetGamemodesStats(platformId string, profiles ...string) map[string]map[string]interface{} {

	body := ticket.getStats(platformId, STATS_GAMEMODES, profiles)

	var output map[string]map[string]map[string]interface{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["results"]

}

func (ticket *UbiTicket) GetWeaponsStats(platformId string, profiles ...string) map[string]map[string]interface{} {

	body := ticket.getStats(platformId, STATS_WEAPONS, profiles)

	var output map[string]map[string]map[string]interface{}

	err := json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["results"]

}
