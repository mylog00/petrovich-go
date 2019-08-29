package petrovich

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

const pathToFile = "./petrovich-eval/"

type testCase struct {
	original string
	expected string
	gender   string
	caseName string
}

func (t testCase) String() string {
	return fmt.Sprintf("Original: %s, Expected: %s, Genger: %s, Case: %s;",
		t.original,
		t.expected,
		t.gender,
		t.caseName)
}

func TestFirstnamesFile(t *testing.T) {
	petrovich := LoadFromFile(rulesBytes)

	file, _ := os.Open(pathToFile + "firstnames.tsv")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		test := prepareTestCase(scanner.Text())
		actual, err := petrovich.FirstName(test.original, test.gender, test.caseName)
		examineAnswer(test.expected, actual, err, t)
	}
}

func prepareTestCase(line string) testCase {
	splitStr := strings.SplitN(line, "\t", 3)
	original := strings.ToLower(splitStr[0])
	expected := strings.ToLower(splitStr[1])
	splitStr = strings.SplitN(splitStr[2], ",", 3)
	gender := parseGender(splitStr[0])
	caseName := parseCaseName(splitStr[2])
	return testCase{original, expected, gender, caseName}
}

func parseGender(gender string) string {
	switch gender {
	case "мр":
		return Male
	case "жр":
		return Female
	default:
		return Androgynous
	}
}

func parseCaseName(caseName string) string {
	switch caseName {
	case "рд":
		return Genitive
	case "дт":
		return Dative
	case "вн":
		return Accusative
	case "тв":
		return Instrumental
	case "пр":
		return Prepositional
	default:
		//case "им"
		return Nominative
	}
}
