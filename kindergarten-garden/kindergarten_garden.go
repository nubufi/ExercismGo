package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)


type Garden struct {
	diagram   string
	children  []string
}

var PlantMap = map[string]string{
	"V": "violets",
	"R": "radishes",
	"C": "clover",
	"G": "grass",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if !checkDiagramFormat(diagram) {
		return nil, fmt.Errorf("invalid diagram")
	}

	if hasDuplicates(children) {
		return nil, fmt.Errorf("duplicate children")
	}

	return &Garden{diagram, children}, nil
	
}

func (g *Garden) Plants(child string) ([]string, bool) {
	diagram := strings.Trim(g.diagram, "\n")
	rows := strings.Split(diagram, "\n")
	plants := []string{}

	children := make([]string, len(g.children))
	copy(children, g.children)

	sort.Strings(children)
	for i, c := range children {
		if c == child {
			for _, row := range rows {
				row = strings.TrimSpace(row)
				if len(row) != len(children)*2 {
					return nil, false
				}
				for _, p := range row[i*2 : i*2+2] {
					plants = append(plants, PlantMap[string(p)])
				}
			}
			break
		}
	}

	if len(plants) == 0 {
		return nil, false
	}

	return plants, true
}

func checkDiagramFormat(diagram string) bool {
	// check if starts with a newline
	if  !strings.HasPrefix(diagram, "\n") {
		fmt.Println(diagram,string(diagram[0]))
		return false
	}

	rows := strings.Split(diagram, "\n")
	if len(rows) != 3 {
		return false
	}

	// check if each row has the same length
	r1 := len(strings.TrimSpace(rows[1]))
	r2 := len(strings.TrimSpace(rows[2]))

	if r1 != r2 {
		return false
	}

	for _, row := range rows {
		row = strings.TrimSpace(row)
		if len(row)%2 != 0 {
			return false
		}
		for _, p := range row {
			if _, exists := PlantMap[string(p)]; !exists {
				return false
			}
		}
	}
	return true
}

func hasDuplicates(children []string) bool {
	seen := make(map[string]struct{})
	for _, value := range children {
		if _, exists := seen[value]; exists {
			return true
		}
		seen[value] = struct{}{}
	}
	return false
}
