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