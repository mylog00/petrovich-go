package petrovich

import (
	"io/ioutil"
	"testing"
)

var rulesBytes = readFile("petrovich-rules/rules.yml")

func TestLastname(t *testing.T) {
	petrovich := LoadFromFile(rulesBytes)
	actual, err := petrovich.LastName("Дубовицкая", Female, Nominative)
	examineAnswer("Дубовицкая", actual, err, t)

	actual, err = petrovich.LastName("Дубовицкая", Female, Genitive)
	examineAnswer("Дубовицкой", actual, err, t)

	actual, err = petrovich.LastName("Дубовицкая", Female, Dative)
	examineAnswer("Дубовицкой", actual, err, t)

	actual, err = petrovich.LastName("Дубовицкая", Female, Accusative)
	examineAnswer("Дубовицкую", actual, err, t)

	actual, err = petrovich.LastName("Дубовицкая", Female, Instrumental)
	examineAnswer("Дубовицкой", actual, err, t)

	actual, err = petrovich.LastName("Дубовицкая", Female, Prepositional)
	examineAnswer("Дубовицкой", actual, err, t)
}

func TestFirstnameException(t *testing.T) {
	petrovich := LoadFromFile(rulesBytes)
	firstname := "Пётр"

	actual, err := petrovich.FirstName(firstname, Male, Nominative)
	examineAnswer("Пётр", actual, err, t)

	actual, err = petrovich.FirstName(firstname, Male, Genitive)
	examineAnswer("Петра", actual, err, t)

	actual, err = petrovich.FirstName(firstname, Male, Dative)
	examineAnswer("Петру", actual, err, t)

	actual, err = petrovich.FirstName(firstname, Male, Accusative)
	examineAnswer("Петра", actual, err, t)

	actual, err = petrovich.FirstName(firstname, Male, Instrumental)
	examineAnswer("Петром", actual, err, t)

	actual, err = petrovich.FirstName(firstname, Male, Prepositional)
	examineAnswer("Петре", actual, err, t)
}

func TestDoubleLastname(t *testing.T) {
	petrovich := LoadFromFile(rulesBytes)
	lastname := "Салтыков-Щедрин"

	actual, err := petrovich.LastName(lastname, Male, Nominative)
	examineAnswer("Салтыков-Щедрин", actual, err, t)

	actual, err = petrovich.LastName(lastname, Male, Genitive)
	examineAnswer("Салтыкова-Щедрина", actual, err, t)

	actual, err = petrovich.LastName(lastname, Male, Dative)
	examineAnswer("Салтыкову-Щедрину", actual, err, t)

	actual, err = petrovich.LastName(lastname, Male, Accusative)
	examineAnswer("Салтыкова-Щедрина", actual, err, t)

	actual, err = petrovich.LastName(lastname, Male, Instrumental)
	examineAnswer("Салтыковым-Щедриным", actual, err, t)

	actual, err = petrovich.LastName(lastname, Male, Prepositional)
	examineAnswer("Салтыкове-Щедрине", actual, err, t)
}
func examineAnswer(expected string, actual string, err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Wrong answer expected='%s', actual='%s'", expected, actual)
	}
}
func readFile(fileName string) []byte {
	data, _ := ioutil.ReadFile(fileName)
	return data
}
