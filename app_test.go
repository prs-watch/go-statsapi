package statsapi

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Util function to read answer JSON.
func readAnswer(filePath string) (body string) {
	file, _ := os.Open(filePath)
	bodyBuf, _ := ioutil.ReadAll(file)
	return string(bodyBuf)
}

func TestAttendance(t *testing.T) {
	params := new(Params)
	params.TeamID = "143"
	params.Season = "2019"

	result, _ := Attendance(params)

	ans := readAnswer("testdata/attendance_ans.json")

	assert.JSONEq(t, result, ans)
}

func TestAwards(t *testing.T) {
	params := new(Params)
	params.Season = "2019"
	result, _ := Awards(params)

	ans := readAnswer("testdata/awards_ans.json")

	assert.JSONEq(t, result, ans)
}

func TestAwardRecips(t *testing.T) {
	params := new(Params)
	params.Season = "2019"
	result, _ := AwardRecips("MLBHOF", params)

	ans := readAnswer("testdata/award_recips_ans.json")

	assert.JSONEq(t, result, ans)
}

func TestDivisions(t *testing.T) {
	params := new(Params)
	params.SportsID = "1"
	result, _ := Divisions(params)

	ans := readAnswer("testdata/divisions_ans.json")

	assert.JSONEq(t, result, ans)
}

func TestRoster(t *testing.T) {
	params := new(Params)
	params.Season = "2019"
	result, _ := Roster("143", params)

	ans := readAnswer("testdata/roster_ans.json")

	assert.JSONEq(t, result, ans)
}

func TestBoxScore(t *testing.T) {
	params := new(Params)
	result, _ := BoxScore("529572", params)

	ans := readAnswer("testdata/boxscore_ans.json")

	assert.JSONEq(t, result, ans)
}

func TestLineScore(t *testing.T) {
	params := new(Params)
	result, _ := LineScore("529572", params)

	ans := readAnswer("testdata/linescore_ans.json")

	assert.JSONEq(t, result, ans)
}

// Timestamp assertion issue exists!
// func TestStandings(t *testing.T) {
// 	params := new(Params)
// 	params.Season = "2019"
// 	params.LeagueID = "103"
// 	result, _ := Standings(params)

// 	ans := readAnswer("testdata/standings_ans.json")

// 	assert.JSONEq(t, result, ans)
// }

func TestTeamLeaders(t *testing.T) {
	params := new(Params)
	params.Season = "2019"
	params.LeaderCategories = "earnedRunAverage"
	result, _ := TeamLeaders("143", params)

	ans := readAnswer("testdata/team_leaders_ans.json")

	assert.JSONEq(t, result, ans)
}

func TestLeagueLeaders(t *testing.T) {
	params := new(Params)
	params.Season = "2019"
	params.LeaderCategories = "earnedRunAverage"
	params.PlayerPool = "qualified"
	result, _ := LeagueLeaders(params)

	ans := readAnswer("testdata/league_leaders_ans.json")

	assert.JSONEq(t, result, ans)
}

func TestPlayerInfo(t *testing.T) {
	params := new(Params)
	result, _ := PlayerInfo("547943", params)

	ans := readAnswer("testdata/player_info_ans.json")

	assert.JSONEq(t, result, ans)
}

func TestPlayByPlay(t *testing.T) {
	params := new(Params)
	result, _ := PlayByPlay("529572", params)

	ans := readAnswer("testdata/play_by_play_ans.json")

	assert.JSONEq(t, result, ans)
}
