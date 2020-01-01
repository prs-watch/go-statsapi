package statsapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// consts
const (
	BaseURL = "https://statsapi.mlb.com/api/"
	Ver     = "v1"
)

// Params object
type Params struct {
	TeamID       string
	LeagueID     string
	LeagueListid string
	Season       string
	Date         string
	GameType     string
	Fields       string
	SportsID     string
	Hydrade      string
	DivisionID   string
	RosterType   string
	Timecode     string
}

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
	values.Set("hydrade", params.Hydrade)
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
	values.Set("hydrade", params.Hydrade)
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
	values.Set("hydrade", params.Hydrade)
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
