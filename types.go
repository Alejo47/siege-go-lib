package Siege

type AccountStats struct {
	General   map[string]int
	Gamemodes struct {
		Bomb    map[string]int
		Secure  map[string]int
		Hostage map[string]int
	}
	Casual    map[string]int
	Ranked    map[string]int
	Weapons   map[string]int
	Operators map[string]OperatorStats
}
type RankedSeason struct {
	Wins     float64 `json:"wins"`
	Abandons float64 `json:"abandons"`
	Losses   float64 `json:"losses"`
	MaxMMR   float64 `json:"max_mmr"`
	MaxRank  float64 `json:"max_rank"`
	MMR      float64 `json:"mmr"`
	Rank     float64 `json:"rank"`
	Region   string  `json:"region"`
}
type OperatorStats struct {
	Kills       int
	Deaths      int
	Wins        int
	Losses      int
	Playtime    int
	OP_Specific map[string]int
}
type SiegeOperator struct {
	Role          string
	Code          string
	Name          string
	OperatorStats OperatorStats
}
type Client struct {
	Ticket     string `json:"ticket"`
	Username   string `json:"username"`
	Expiration string `json:"expiration"`
}
type Platform struct {
	Space   string `json:"spaceId"`
	Sandbox string `json:"sandbox"`
}
type Profile struct {
	ProfileId      string `json:"profileId"`
	UserId         string `json:"userId"`
	PlatformType   string `json:"platformType"`
	IdOnPlatform   string `json:"idOnPlatform"`
	NameOnPlatform string `json:"nameOnPlatform"`
}
type RainbowSixAccount struct {
	Username string
	Platform string
	Profile  Profile
}
