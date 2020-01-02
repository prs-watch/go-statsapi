# go-statsapi

Golang client for MLB Stats API.

## Getting Started

```golang
package main

import (
    "fmt"
    "github.com/prs-watch/go-statsapi"
)

func main() {
    params := new(statsapi.Params)
    params.Season = "2019"
    params.TeamID = "143"

    atts := statsapi.Attendance(params)
    fmt.Println(atts)
}
```

## Installation

```bash
go get github.com/prs-watch/go-statsapi
```

## Functions

- `statsapi.Attendance()`
    - get attendance data.
- `statsapi.Awards()`
    - get list of awardIds.
- `statsapi.AwardRecips()`
    - get recipients by award you pass as method param.
- `statsapi.Divisions()`
    - get list of MLB and MiLB divisions.
- `statsapi.Roster()`
    - get team roster.
- `statsapi.BoxScore()`
    - get specific game's boxscore.
- `statsapi.LineScore()`
    - get specific game's linescore.
- `statsapi.Standings()`
    - get standings.
- `statsapi.TeamLeaders()`
    - get team leaders of given stat.
- `statsapi.LeagueLeaders()`
    - get league leaders of given stat.
- `statsapi.PlayerInfo()`
    - get specific player's info.
- `statsapi.PlayByPlay()`
    - get specific game's play-by-play.