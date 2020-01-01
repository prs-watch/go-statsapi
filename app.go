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

// Attendance : Get attendance.
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

// Awards : Get awards list.
func Awards(params *Params) (body string, err error) {
	values := url.Values{}
	values.Set("sportdId", params.SportsID)
	values.Set("leagueId", params.LeagueID)
	values.Set("season", params.Season)
	values.Set("hydrade", params.Hydrade)
	values.Set("field", params.Fields)

	return httpGet("awards", values)
}

// AwardRecips : Get each award recipients.
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

// Divisions : Get division list.
func Divisions(params *Params) (body string, err error) {
	values := url.Values{}
	values.Set("divisionId", params.DivisionID)
	values.Set("leagueId", params.LeagueID)
	values.Set("sportsId", params.SportsID)

	return httpGet("divisions", values)
}
