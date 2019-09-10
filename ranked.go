package Siege

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (ticket *UbiTicket) GetRankedInformation(season, region, platformId string, profiles ...string) map[string]RankedSeason {

	platform := PLATFORMS[platformId]

	uri := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/%s/sandboxes/%s/r6karma/players?board_id=pvp_ranked&profile_ids=%s&region_id=%s&season_id=%s",
		platform.Space,
		platform.Sandbox,
		strings.Join(profiles, ","),
		region,
		season,
	)

	body, err := ticket.makeHTTPReq(uri)

	if err != nil {
		return nil
	}

	var output map[string]map[string]RankedSeason

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil
	}

	return output["players"]

}
