package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	Name string
	MP   int
	W    int
	D    int
	L    int
	P    int
}

func Tally(reader io.Reader, writer io.Writer) error {
	text := convertToText(reader)
	lines, err := splitText(text, "\n")
	if err != nil {
		return err
	}

	var teams = make(map[string]Team)

	for _, line := range lines {
		cols, err := splitText(line, ";")
		if err != nil {
			return err
		}
		if len(cols) >= 3 {
			updateTeams(teams, cols)
		}

	}
	sortedTeams := sortTeams(teams)
	result := printTeams(sortedTeams)
	fmt.Println(result)
	writer.Write([]byte(result))

	return nil
}

func convertToText(reader io.Reader) string {
	b := make([]byte, 10240)
	reader.Read(b)

	text := string(b)
	return text
}

func splitText(text, splitter string) ([]string, error) {
	splitted := strings.Split(text, splitter)

	if splitter == "\n" && len(splitted) == 1 {
		return splitted, errors.New("No new line found")
	}

	var result []string
	for _, s := range splitted {
		if len(s) > 1 && len(s) < 100 {
			result = append(result, s)
		}
	}

	return result, nil
}

func updateTeams(teams map[string]Team, cols []string) {
	teamName1 := cols[0]
	teamName2 := cols[1]
	result := cols[2]

	var team1, team2 Team

	if _, ok := teams[teamName1]; ok {
		team1 = teams[teamName1]
	}
	if _, ok := teams[teamName2]; ok {
		team2 = teams[teamName2]
	}
	switch result {
	case "win":
		team1.W++
		team2.L++
		team1.P += 3
	case "loss":
		team1.L++
		team2.W++
		team2.P += 3
	case "draw":
		team1.D++
		team2.D++
		team1.P++
		team2.P++
	}
	team1.MP++
	team2.MP++
	teams[teamName1] = team1
	teams[teamName2] = team2
}

// Sort teams by point
func sortTeams(teams map[string]Team) []Team {
	var sortedTeams []Team
	for name, team := range teams {
		team.Name = name
		sortedTeams = append(sortedTeams, team)
	}

	sort.Slice(sortedTeams, func(i, j int) bool {
		return sortedTeams[i].P > sortedTeams[j].P
	})

	// Sort by name if points are equal
	sort.Slice(sortedTeams, func(i, j int) bool {
		return sortedTeams[i].Name < sortedTeams[j].Name && sortedTeams[i].P == sortedTeams[j].P
	})

	return sortedTeams
}

func printTeams(teams []Team) string {
	var result = "Team                           | MP |  W |  D |  L |  P\n"
	for _, team := range teams {
		line := fmt.Sprintf("%-*s |  %d |  %d |  %d |  %d |  %d\n", 30, team.Name, team.MP, team.W, team.D, team.L, team.P)
		result += line
	}
	return result
}
