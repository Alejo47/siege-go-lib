package Siege

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
	"operatorpvp_roundlost",
	"operatorpvp_roundwon",
	"operatorpvp_death",
	"operatorpvp_kills",
	"operatorpvp_timeplayed",
	"operatorpvp_amaru_distancereeled",
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
	"operatorpvp_goyo_volcandetonate",
	"operatorpvp_gridlock_traxdeployed",
	"operatorpvp_hibana_detonate_projectile",
	"operatorpvp_iq_gadgetspotbyef",
	"operatorpvp_jager_gadgetdestroybycatcher",
	"operatorpvp_kaid_electroclawelectrify",
	"operatorpvp_kapkan_boobytrapkill",
	"operatorpvp_maverick_wallbreached",
	"operatorpvp_montagne_shieldblockdamage",
	"operatorpvp_mozzie_droneshacked",
	"operatorpvp_mute_gadgetjammed",
	"operatorpvp_nokk_observationtooldeceived",
	"operatorpvp_nomad_airjabdetonate",
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
	"operatorpvp_warden_killswithglasses",
}

var STATS_WEAPONS = []string{
	"weapontypepvp_bulletfired",
	"weapontypepvp_bullethit",
	"weapontypepvp_headshot",
	"weapontypepvp_kills",
}

var PLATFORMS = map[string]Platform{
	"uplay": {Space: "5172a557-50b5-4665-b7db-e3f2e8c5041d", Sandbox: "OSBOR_PC_LNCH_A"},
	"psn":   {Space: "05bfb3f7-6c21-4c42-be1f-97a33fb5cf66", Sandbox: "OSBOR_PS4_LNCH_A"},
	"xbl":   {Space: "98a601e5-ca91-4440-b1c5-753f601a2c90", Sandbox: "OSBOR_XBOXONE_LNCH_A"},
}

var GUN_CODES = map[string]string{
	":1:": "Assault Rifle",
	":2:": "SMG",
	":3:": "LMG",
	":4:": "Sniper",
	":5:": "Pistol",
	":6:": "Shotgun",
	":7:": "Machine Pistol",
	":B:": "Utility",
	":9:": "Unknown",
	":C:": "Unknown",
}

var OP_CODES = map[string]SiegeOperator{
	":1:1:":  SiegeOperator{Role: "both", Name: "SAS Recruit", Code: ":1:1:"},      //SAS
	":2:1:":  SiegeOperator{Role: "def", Name: "Smoke", Code: ":2:1:"},             //SAS
	":3:1:":  SiegeOperator{Role: "def", Name: "Mute", Code: ":3:1:"},              //SAS
	":4:1:":  SiegeOperator{Role: "atk", Name: "Sledge", Code: ":4:1:"},            //SAS
	":5:1:":  SiegeOperator{Role: "atk", Name: "Thatcher", Code: ":5:1:"},          //SAS
	":1:2:":  SiegeOperator{Role: "both", Name: "FBI Recruit", Code: ":1:2:"},      //FBI
	":2:2:":  SiegeOperator{Role: "def", Name: "Castle", Code: ":2:2:"},            //FBI
	":3:2:":  SiegeOperator{Role: "atk", Name: "Ash", Code: ":3:2:"},               //FBI
	":4:2:":  SiegeOperator{Role: "def", Name: "Pulse", Code: ":4:2:"},             //FBI
	":5:2:":  SiegeOperator{Role: "atk", Name: "Thermite", Code: ":5:2:"},          //FBI
	":1:3:":  SiegeOperator{Role: "both", Name: "GIGN Recruit", Code: ":1:3:"},     //GIGN
	":2:3:":  SiegeOperator{Role: "def", Name: "Doc", Code: ":2:3:"},               //GIGN
	":3:3:":  SiegeOperator{Role: "def", Name: "Rook", Code: ":3:3:"},              //GIGN
	":4:3:":  SiegeOperator{Role: "atk", Name: "Twitch", Code: ":4:3:"},            //GIGN
	":5:3:":  SiegeOperator{Role: "atk", Name: "Montagne", Code: ":5:3:"},          //GIGN
	":1:4:":  SiegeOperator{Role: "both", Name: "Spetsnaz Recruit", Code: ":1:4:"}, //SPETSNAZ
	":2:4:":  SiegeOperator{Role: "atk", Name: "Glaz", Code: ":2:4:"},              //SPETSNAZ
	":3:4:":  SiegeOperator{Role: "atk", Name: "Fuze", Code: ":3:4:"},              //SPETSNAZ
	":4:4:":  SiegeOperator{Role: "def", Name: "Kapkan", Code: ":4:4:"},            //SPETSNAZ
	":5:4:":  SiegeOperator{Role: "def", Name: "Tachanka", Code: ":5:4:"},          //SPETSNAZ
	":1:5:":  SiegeOperator{Role: "both", Name: "GSG9 Recruit", Code: ":1:5:"},     //GSG9
	":2:5:":  SiegeOperator{Role: "atk", Name: "Blitz", Code: ":2:5:"},             //GSG9
	":3:5:":  SiegeOperator{Role: "atk", Name: "IQ", Code: ":3:5:"},                //GSG9
	":4:5:":  SiegeOperator{Role: "def", Name: "Jager", Code: ":4:5:"},             //GSG9
	":5:5:":  SiegeOperator{Role: "def", Name: "Bandit", Code: ":5:5:"},            //GSG9
	":2:6:":  SiegeOperator{Role: "atk", Name: "Buck", Code: ":2:6:"},              //JTF2
	":3:6:":  SiegeOperator{Role: "def", Name: "Frost", Code: ":3:6:"},             //JTF2
	":2:7:":  SiegeOperator{Role: "atk", Name: "Blackbeard", Code: ":2:7:"},        //SEAL
	":3:7:":  SiegeOperator{Role: "def", Name: "Valkyrie", Code: ":3:7:"},          //SEAL
	":2:8:":  SiegeOperator{Role: "atk", Name: "Capitão", Code: ":2:8:"},           //BOPE
	":3:8:":  SiegeOperator{Role: "def", Name: "Caveira", Code: ":3:8:"},           //BOPE
	":2:9:":  SiegeOperator{Role: "atk", Name: "Hibana", Code: ":2:9:"},            //SAT
	":3:9:":  SiegeOperator{Role: "def", Name: "Echo", Code: ":3:9:"},              //SAT
	":2:A:":  SiegeOperator{Role: "atk", Name: "Jackal", Code: ":2:A:"},            //GEO
	":3:A:":  SiegeOperator{Role: "def", Name: "Mira", Code: ":3:A:"},              //GEO
	":3:B:":  SiegeOperator{Role: "def", Name: "Lesion", Code: ":3:B:"},            //SDU
	":2:B:":  SiegeOperator{Role: "atk", Name: "Ying", Code: ":2:B:"},              //SDU
	":2:C:":  SiegeOperator{Role: "def", Name: "Ela", Code: ":2:C:"},               //GROM
	":3:C:":  SiegeOperator{Role: "atk", Name: "Zofia", Code: ":3:C:"},             //GROM
	":2:D:":  SiegeOperator{Role: "atk", Name: "Dokkaebi", Code: ":2:D:"},          //SMB
	":3:D:":  SiegeOperator{Role: "def", Name: "Vigil", Code: ":3:D:"},             //SMB
	":3:E:":  SiegeOperator{Role: "atk", Name: "Lion", Code: ":3:E:"},              //CBRN
	":4:E:":  SiegeOperator{Role: "atk", Name: "Finka", Code: ":4:E:"},             //CBRN
	":2:F:":  SiegeOperator{Role: "def", Name: "Maestro", Code: ":2:F:"},           //GIS
	":3:F:":  SiegeOperator{Role: "def", Name: "Alibi", Code: ":3:F:"},             //GIS
	":3:10:": SiegeOperator{Role: "def", Name: "Clash", Code: ":3:10:"},            //GSUTR
	":2:10:": SiegeOperator{Role: "atk", Name: "Maverick", Code: ":2:10:"},         //GSUTR
	":2:11:": SiegeOperator{Role: "atk", Name: "Nomad", Code: ":2:11:"},            //GIGR
	":3:11:": SiegeOperator{Role: "def", Name: "Kaid", Code: ":3:11:"},             //GIGR
	":2:12:": SiegeOperator{Role: "def", Name: "Mozzie", Code: ":2:12:"},           //SASR
	":3:12:": SiegeOperator{Role: "atk", Name: "Gridlock", Code: ":3:12:"},         //SASR
	":2:13:": SiegeOperator{Role: "def", Name: "Nøkk", Code: ":2:13:"},             //JGK
	":2:14:": SiegeOperator{Role: "atk", Name: "Warden", Code: ":2:14:"},           //USSS
	":2:15:": SiegeOperator{Role: "def", Name: "Goyo", Code: ":2:15:"},             //Unknown
	":2:16:": SiegeOperator{Role: "atk", Name: "Amaru", Code: ":2:16:"},            //Unknown
}
