package statsapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// consts
const (
	BaseURL = "https://statsapi.mlb.com/api"
	Ver     = "v1"
)

// Params object
type Params struct {
	TeamID           string
	LeagueID         string
	LeagueListid     string
	Season           string
	Date             string
	GameType         string
	Fields           string
	SportsID         string
	Hydrate          string
	DivisionID       string
	RosterType       string
	Timecode         string
	StandingsType    string
	LeaderCategories string
	LeaderGameTypes  string
	PlayerPool       string
	StatGroup        string
	StatType         string
}

// private method to execute http get.
func httpGet(appType string, values url.Values) (body string, err error) {
	var urlStr = BaseURL + "/" + Ver + "/" + appType

	resp, err := http.Get(urlStr + "?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	bodyBuf, _ := ioutil.ReadAll(resp.Body)
	return string(bodyBuf), nil
}

// Attendance : get attendance.
func Attendance(params *Params) (body string, err error) {
	values := url.Values{}
	values.Set("teamId", params.TeamID)
	values.Set("leagueId", params.LeagueID)
	values.Set("season", params.Season)
	values.Set("date", params.Date)
	values.Set("gameType", params.GameType)
	values.Set("field", params.Fields)

	return httpGet("attendance", values)
}

// Awards : get awards list.
func Awards(params *Params) (body string, err error) {
	values := url.Values{}
	values.Set("sportdId", params.SportsID)
	values.Set("leagueId", params.LeagueID)
	values.Set("season", params.Season)
	values.Set("hydrate", params.Hydrate)
	values.Set("field", params.Fields)

	return httpGet("awards", values)
}

// AwardRecips : get each award recipients.
func AwardRecips(awardID string, params *Params) (body string, err error) {
	var appType = "awards" + "/" + awardID + "/" + "recipients"

	values := url.Values{}
	values.Set("sportdId", params.SportsID)
	values.Set("leagueId", params.LeagueID)
	values.Set("season", params.Season)
	values.Set("hydrate", params.Hydrate)
	values.Set("field", params.Fields)

	return httpGet(appType, values)
}

// Divisions : get division list.
func Divisions(params *Params) (body string, err error) {
	values := url.Values{}
	values.Set("divisionId", params.DivisionID)
	values.Set("leagueId", params.LeagueID)
	values.Set("sportsId", params.SportsID)

	return httpGet("divisions", values)
}

// Roster : get team roster.
func Roster(teamID string, params *Params) (body string, err error) {
	var appType = "teams" + "/" + teamID + "/" + "roster"

	values := url.Values{}
	values.Set("rosterType", params.RosterType)
	values.Set("season", params.Season)
	values.Set("date", params.Date)
	values.Set("hydrate", params.Hydrate)
	values.Set("field", params.Fields)

	return httpGet(appType, values)
}

// BoxScore : get specific game's boxscore.
func BoxScore(gamePK string, params *Params) (body string, err error) {
	var appType = "game" + "/" + gamePK + "/" + "boxscore"

	values := url.Values{}
	values.Set("timecode", params.Timecode)
	values.Set("field", params.Fields)

	return httpGet(appType, values)
}

// LineScore : get specific game's linescore.
func LineScore(gamePK string, params *Params) (body string, err error) {
	var appType = "game" + "/" + gamePK + "/" + "linescore"

	values := url.Values{}
	values.Set("timecode", params.Timecode)
	values.Set("field", params.Fields)

	return httpGet(appType, values)
}

// Standings : get standings.
func Standings(params *Params) (body string, err error) {
	values := url.Values{}
	values.Set("leagueId", params.LeagueID)
	values.Set("season", params.Season)
	values.Set("standingsType", params.StandingsType)
	values.Set("date", params.Date)
	values.Set("hydrate", params.Hydrate)
	values.Set("field", params.Fields)

	return httpGet("standings", values)
}

// TeamLeaders : get team leaders of given stat.
func TeamLeaders(teamID string, params *Params) (body string, err error) {
	var appType = "teams" + "/" + teamID + "/" + "leaders"

	values := url.Values{}
	values.Set("leaderCategories", params.LeaderCategories)
	values.Set("leaderGameTypes", params.LeaderGameTypes)
	values.Set("season", params.Season)
	values.Set("date", params.Date)
	values.Set("hydrate", params.Hydrate)
	values.Set("field", params.Fields)

	return httpGet(appType, values)
}

// LeagueLeaders : get league leaders of given stat.
func LeagueLeaders(params *Params) (body string, err error) {
	values := url.Values{}
	values.Set("leaderCategories", params.LeaderCategories)
	values.Set("season", params.Season)
	values.Set("sportdId", params.SportsID)
	values.Set("leaderGameTypes", params.LeaderGameTypes)
	values.Set("playerPool", params.PlayerPool)
	values.Set("statGroup", params.StatGroup)
	values.Set("hydrate", params.Hydrate)
	values.Set("field", params.Fields)

	return httpGet("stats/leaders", values)
}

// PlayerInfo : get specific player's info.
func PlayerInfo(personID string, params *Params) (body string, err error) {
	var appType = "people" + "/" + personID

	values := url.Values{}
	values.Set("hydrate", params.Hydrate)
	values.Set("field", params.Fields)

	return httpGet(appType, values)
}

// PlayByPlay : get specific game's play-by-play.
func PlayByPlay(gamePK string, params *Params) (body string, err error) {
	var appType = "game" + "/" + gamePK + "/" + "playByPlay"

	values := url.Values{}
	values.Set("timecode", params.Timecode)
	values.Set("field", params.Fields)

	return httpGet(appType, values)
}
