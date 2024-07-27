package diamond

import (
	"errors"
	"fmt"
	"strings"
)

func Gen(char byte) (string, error) {
	if char == 'A' {
		return "A", nil
	}

	if char < 'A' || char > 'Z' {
		return "", errors.New("invalid input")
	}

	diamond := make([]string, 0, 1)
	for i := byte('A'); i <= char; i++ {
		spaceCount := int(char - i)
		spaceLR := strings.Repeat(" ", spaceCount)
		if i == 'A' {
			diamond = append(diamond, fmt.Sprintf("%s%c%s", spaceLR, i, spaceLR))
		} else {
			spaceCenter := strings.Repeat(" ", 2*(int(i)-int('A'))-1)
			diamond = append(diamond, fmt.Sprintf("%s%c%s%c%s", spaceLR, i, spaceCenter, i, spaceLR))
		}
	}

	for i := len(diamond) - 2; i >= 0; i-- {
		diamond = append(diamond, diamond[i])
	}

	return strings.Join(diamond, "\n"), nil
}
