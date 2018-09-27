package tournament

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Result is multipropose struct we use to work with game results.
type Result struct {
	Team      string
	Games     int
	Win       int
	Draw      int
	Loss      int
	Points    int
	SortScore int // little cheat
}

// Points is Quick look for game results points.
var Points = map[string]int{
	"draw": 1,
	"win":  3,
	"loss": 0,
}

// Define a function Tally(io.Reader, io.Writer) error.

// Tally write to io.Writer and return error
func Tally(r io.Reader, w io.Writer) error {
	cvsr := csv.NewReader(r)
	records, err := cvsr.ReadAll()
	if err != nil {
		return err
	}

	tableMap := make(map[string]Result)

	for _, row := range records {
		var row = strings.Split(row[0], ";")

		// Rows validation
		if len(row) != 3 {
			// is it comment?
			if row[0][0] == '#' {
				continue
			}

			return errors.New("Wrong number of fields")
		}

		if row[2] != "draw" && row[2] != "win" && row[2] != "loss" {
			return errors.New("Unknown field")
		}

		// Making Map or teams
		var state1, state2 string
		switch row[2] {
		case "win":
			state1, state2 = "win", "loss"
		case "loss":
			state1, state2 = "loss", "win"
		case "draw":
			state1, state2 = "draw", "draw"
		}

		// Creating/Updating values in map
		for i, state := range []string{state1, state2} {

			team := row[0]
			if i == 1 {
				team = row[1]
			}

			var player Result
			var ok bool
			if player, ok = tableMap[team]; !ok {
				player = Result{Team: team}
			}
			// checking stage
			player.Games++
			player.Points += Points[state]

			switch state {
			case "win":
				player.Win++
			case "loss":
				player.Loss++
			case "draw":
				player.Draw++
			}

			tableMap[team] = player
		}
	}

	// Map to slice and sort slice
	var sortedSliceTeams []Result

	for _, results := range tableMap {
		// Little cheat to make sorting simplier.
		results.SortScore = results.Points*10000 + results.Games*1000 + results.Draw*100 - results.Loss*10
		sortedSliceTeams = append(sortedSliceTeams, results)
	}

	// Sorting final Results.
	sort.Slice(sortedSliceTeams, func(i, j int) bool {

		if sortedSliceTeams[i].SortScore > sortedSliceTeams[j].SortScore {
			return true
		}

		if sortedSliceTeams[i].SortScore == sortedSliceTeams[j].SortScore {
			var names = []string{sortedSliceTeams[i].Team, sortedSliceTeams[j].Team}
			sort.Strings(names)
			if names[0] == sortedSliceTeams[i].Team {
				return true
			}
		}

		return false
	})

	// Header and Row holders.
	var header = "%-31s| %-2s | %2s | %2s | %2s | %2s\n"
	var row = "%-31s| %2d | %2d | %2d | %2d | %2d\n"

	w.Write([]byte(fmt.Sprintf(header, "Team", "MP", "W", "D", "L", "P")))

	// Body
	for _, r := range sortedSliceTeams {
		w.Write([]byte(fmt.Sprintf(row, r.Team, r.Games, r.Win, r.Draw, r.Loss, r.Points)))
	}
	return nil
}
