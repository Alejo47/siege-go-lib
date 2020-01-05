package Siege

import (
	"encoding/json"
	"errors"
	"fmt"
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

/*
var STATS_CASUAL = []string{
	"casualpvp_death",
	"casualpvp_kills",
	"casualpvp_matchlost",
	"casualpvp_matchplayed",
	"casualpvp_matchwon",
	"casualpvp_timeplayed",
}
var STATS_RANKED = []string{
	"rankedpvp_kills",
	"rankedpvp_death",
	"rankedpvp_matchlost",
	"rankedpvp_matchplayed",
	"rankedpvp_matchwon",
	"rankedpvp_timeplayed",
}
var STATS_GENERAL = []string{
	"generalpvp_bulletfired",
	"generalpvp_bullethit",
	"generalpvp_death",
	"generalpvp_headshot",
	"generalpvp_killassists",
	"generalpvp_kills",
	"generalpvp_matchlost",
	"generalpvp_matchplayed",
	"generalpvp_matchwon",
	"generalpvp_meleekills",
	"generalpvp_penetrationkills",
	"generalpvp_revive",
	"generalpvp_timeplayed",
}
var STATS_GAMEMODES = []string{
	"plantbombpvp_bestscore",
	"plantbombpvp_matchlost",
	"plantbombpvp_matchplayed",
	"plantbombpvp_matchwon",
	"rescuehostagepvp_bestscore",
	"rescuehostagepvp_matchlost",
	"rescuehostagepvp_matchplayed",
	"rescuehostagepvp_matchwon",
	"secureareapvp_bestscore",
	"secureareapvp_matchlost",
	"secureareapvp_matchplayed",
	"secureareapvp_matchwon",
}
var STATS_OPERATORS = []string{
	"operatorpvp_death",
	"operatorpvp_kills",
	"operatorpvp_roundlost",
	"operatorpvp_roundwon",
	"operatorpvp_timeplayed",
	"operatorpvp_ash_bonfirewallbreached",
	"operatorpvp_attackerdrone_diminishedrealitymode",
	"operatorpvp_bandit_batterykill",
	"operatorpvp_barrage_killswithturret",
	"operatorpvp_black_mirror_gadget_deployed",
	"operatorpvp_blackbeard_gunshieldblockdamage",
	"operatorpvp_blitz_flashedenemy",
	"operatorpvp_buck_kill",
	"operatorpvp_caltrop_enemy_affected",
	"operatorpvp_capitao_lethaldartkills",
	"operatorpvp_castle_kevlarbarricadedeployed",
	"operatorpvp_caveira_interrogations",
	"operatorpvp_cazador_assist_kill",
	"operatorpvp_clash_sloweddown",
	"operatorpvp_concussiongrenade_detonate",
	"operatorpvp_concussionmine_detonate",
	"operatorpvp_dazzler_gadget_detonate",
	"operatorpvp_deceiver_revealedattackers",
	"operatorpvp_doc_teammaterevive",
	"operatorpvp_echo_enemy_sonicburst_affected",
	"operatorpvp_frost_dbno",
	"operatorpvp_fuze_clusterchargekill",
	"operatorpvp_glaz_sniperkill",
	"operatorpvp_hibana_detonate_projectile",
	"operatorpvp_iq_gadgetspotbyef",
	"operatorpvp_jager_gadgetdestroybycatcher",
	"operatorpvp_kaid_electroclaw_hatches",
	"operatorpvp_kapkan_boobytrapkill",
	"operatorpvp_maverick_wallbreached",
	"operatorpvp_montagne_shieldblockdamage",
	"operatorpvp_mute_gadgetjammed",
	"operatorpvp_nomad_assist",
	"operatorpvp_phoneshacked",
	"operatorpvp_pulse_heartbeatspot",
	"operatorpvp_rook_armortakenteammate",
	"operatorpvp_rush_adrenalinerush",
	"operatorpvp_sledge_hammerhole",
	"operatorpvp_smoke_poisongaskill",
	"operatorpvp_tachanka_turretkill",
	"operatorpvp_tagger_tagdevice_spot",
	"operatorpvp_thatcher_gadgetdestroywithemp",
	"operatorpvp_thermite_reinforcementbreached",
	"operatorpvp_twitch_gadgetdestroybyshockdrone",
	"operatorpvp_valkyrie_camdeployed",
}
var STATS_WEAPONS = []string{
	"weapontypepvp_bulletfired",
	"weapontypepvp_bullethit",
	"weapontypepvp_headshot",
	"weapontypepvp_kills",
}

func (s AccountStats) InitializeMaps() AccountStats {
	stats := AccountStats{
		General:   make(map[string]int),
		Casual:    make(map[string]int),
		Ranked:    make(map[string]int),
		Weapons:   make(map[string]int),
		Operators: make(map[string]OperatorStats),
	}
	for _, v := range OP_Codes {
		stats.Operators[v.Name] = OperatorStats{OP_Specific: make(map[string]int)}
	}
	stats.Gamemodes.Bomb = make(map[string]int)
	stats.Gamemodes.Hostage = make(map[string]int)
	stats.Gamemodes.Secure = make(map[string]int)
	return stats
}

func (s OperatorStats) updateOperator(val string, n int) OperatorStats {
	switch strings.ToLower(val) {
	case "death":
		s.Deaths = n
		break
	case "kills":
		s.Kills = n
		break
	case "roundlost":
		s.Losses = n
		break
	case "roundwon":
		s.Wins = n
		break
	case "timeplayed":
		s.Playtime = n
		break
	default:
		s.OP_Specific[regexp.MustCompile(`^operatorpvp_`).ReplaceAllString(val, "")] = n
		break
	}
	return s
}

func GetStats(profileId string, platform Platform, stats []string, client Client) (AccountStats, error) {
	return client.GetStats(profileId, platform, stats)
}
func (client Client) GetStats(profileId string, platform Platform, stats []string) (AccountStats, error) {

	uri := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/%s/sandboxes/%s/playerstats2/statistics?populations=%s&statistics=%s",
		platform.Space,
		platform.Sandbox,
		profileId,
		strings.Join(stats, ","))
	body, err := client.makeHTTPReq(uri)
	if err != nil {
		return AccountStats{}, err
	}
	result := make(map[string]map[string]map[string]int)
	json.Unmarshal([]byte(body), &result)

	playerStats := AccountStats{}.InitializeMaps()

	for k, v := range result["results"][profileId] {
		if regexp.MustCompile(`^generalpvp`).Match([]byte(k)) {
			playerStats.General[strings.Split(strings.Split(k, "_")[1], ":")[0]] = v
		} else if regexp.MustCompile(`:([0-9A-F]*):([0-9A-F]*):`).Match([]byte(k)) {
			operator := OP_Codes[regexp.MustCompile(`:([0-9A-F]*):([0-9A-F]*):`).FindString(k)]
			if regexp.MustCompile(`^operatorpvp_`).Match([]byte(k)) {
				value := regexp.MustCompile(`(kills|death|roundwon|roundlost|timeplayed)\b`).FindString(k)
				if value != "" {
					playerStats.Operators[operator.Name] = playerStats.Operators[operator.Name].updateOperator(value, v)
				} else {
					playerStats.Operators[operator.Name] = playerStats.Operators[operator.Name].updateOperator(strings.Split(k, ":")[0], v)
				}
			}
		} else if regexp.MustCompile(`^(plantbomb|rescuehostage|securearea)pvp_`).Match([]byte(k)) {
			gamemode := regexp.MustCompile(`^(plantbomb|rescuehostage|securearea)`).FindString(k)
			value := strings.Split(strings.Split(k, "_")[1], ":")[0]
			switch gamemode {
			case "plantbomb":
				playerStats.Gamemodes.Bomb[value] = v
				break
			case "rescuehostage":
				playerStats.Gamemodes.Hostage[value] = v
				break
			case "securearea":
				playerStats.Gamemodes.Secure[value] = v
				break
			}
		} else if regexp.MustCompile(`^(casual|ranked|general)pvp`).Match([]byte(k)) {
			value := regexp.MustCompile(`(kills|death|matchwon|matchlost|matchplayed|timeplayed)\b`).FindString(k)
			switch regexp.MustCompile(`^(casual|ranked|general)`).FindString(k) {
			case "casual":
				playerStats.Casual[value] = v
				break
			case "ranked":
				playerStats.Ranked[value] = v
				break
			case "general":
				playerStats.General[value] = v
				break
			}
		}

	}
	return playerStats, nil
}

func GetRankedStats(profileId string, platform Platform, region string, season int, client Client) (RankedSeason, error) {
	return client.GetRankedStats(profileId, platform, region, season)
}
func (client Client) GetRankedStats(profileId string, platform Platform, region string, season int) (RankedSeason, error) {
	uri := fmt.Sprintf("https://public-ubiservices.ubi.com/v1/spaces/%s/sandboxes/%s/r6karma/players?board_id=pvp_ranked&profile_ids=%s&region_id=%s&season_id=%d",
		platform.Space,
		platform.Sandbox,
		profileId,
		region,
		season,
	)
	body, err := client.makeHTTPReq(uri)
	if err != nil {
		return RankedSeason{}, err
	}
	result := make(map[string]map[string]RankedSeason)
	json.Unmarshal([]byte(body), &result)

	return result["players"][strings.ToLower(profileId)], nil
}
*/
