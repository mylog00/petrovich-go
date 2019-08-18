package petrovich

import (
	"io/ioutil"
	"testing"
)

func TestRule(t *testing.T) {
	rulesFile := readFile("./rules.yml")
	rules := LoadRules(rulesFile)
	actual, err := rules.LastName("Дубовицкая", Female, Nominative)
	examineAnswer("Дубовицкая", actual, err, t)

	actual, err = rules.LastName("Дубовицкая", Female, Genitive)
	examineAnswer("Дубовицкой", actual, err, t)

	actual, err = rules.LastName("Дубовицкая", Female, Dative)
	examineAnswer("Дубовицкой", actual, err, t)

	actual, err = rules.LastName("Дубовицкая", Female, Accusative)
	examineAnswer("Дубовицкую", actual, err, t)

	actual, err = rules.LastName("Дубовицкая", Female, Instrumental)
	examineAnswer("Дубовицкой", actual, err, t)

	actual, err = rules.LastName("Дубовицкая", Female, Prepositional)
	examineAnswer("Дубовицкой", actual, err, t)
}

func examineAnswer(expected string, actual string, err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Wrong answer expected='%s', actual='%s'", actual, expected)
	}
}
func readFile(fileName string) []byte {
	data, _ := ioutil.ReadFile(fileName)
	return data
}
