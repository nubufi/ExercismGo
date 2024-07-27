package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	var matches []string
	addFileName := len(files) > 1
	for _, file := range files {
		lines := readFile(file)
		matches = append(matches, grep(file, pattern, flags, lines,addFileName)...)
	}

	return matches
}

func match(pattern, line string,xFlag,vFlag bool) bool {
	if xFlag {
		if vFlag {
			return pattern != line
		}
		return pattern == line
	} else {
		if vFlag {
			return !strings.Contains(line, pattern)
		}
		return strings.Contains(line, pattern)
	}
}

func grep(file,pattern string, flags, lines []string,addFileName bool) []string {
	flags_ := strings.Join(flags, "")
	nFlag := strings.Contains(flags_, "-n")
	lFlag := strings.Contains(flags_, "-l")
	xFlag := strings.Contains(flags_, "-x")
	vFlag := strings.Contains(flags_, "-v")
	iFlag := strings.Contains(flags_, "-i")
	
	var fileName string
	if addFileName {
		fileName = fmt.Sprintf("%s:", file)
	}
	var matches []string
	for i, line := range lines {
		l := line
		if iFlag {
			l = strings.ToLower(line)
			pattern = strings.ToLower(pattern)
		}
		if match(pattern, l, xFlag, vFlag){
			if lFlag {
				matches = append(matches, file)
				return matches	
			}
			if  nFlag {
				line = fmt.Sprintf("%s%d:%s",fileName, i+1, line)
			} else {
				line = fmt.Sprintf("%s%s",fileName, line)
			}
			matches = append(matches, line)
		}
	}

	return matches
}

func readFile(file string) []string {
	var lines []string
	// Open file
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	// Read file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
