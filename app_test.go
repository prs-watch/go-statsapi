package statsapi

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

// Util function to read answer JSON.
func readAnswer(filePath string) (body string) {
	file, _ := os.Open(filePath)
	bodyBuf, _ := ioutil.ReadAll(file)
	return string(bodyBuf)
}
