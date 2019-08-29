package petrovich

import (
	"io/ioutil"
	"testing"
)

var rulesBytes = readFile("petrovich-rules/rules.yml")

func TestLastname(t *testing.T) {
	petrovich := LoadFromFile(rulesBytes)
	test := testCase{"Дубовицкая", "Дубовицкая", Female, Nominative}
	examineLastname(petrovich, test, t)

	test.expected = "Дубовицкой"
	test.caseName = Genitive
	examineLastname(petrovich, test, t)

	test.expected = "Дубовицкой"
	test.caseName = Dative
	examineLastname(petrovich, test, t)

	test.expected = "Дубовицкую"
	test.caseName = Accusative
	examineLastname(petrovich, test, t)

	test.expected = "Дубовицкой"
	test.caseName = Instrumental
	examineLastname(petrovich, test, t)

	test.expected = "Дубовицкой"
	test.caseName = Prepositional
	examineLastname(petrovich, test, t)
}

func TestFirstnameException(t *testing.T) {
	petrovich := LoadFromFile(rulesBytes)
	test := testCase{"Пётр", "Пётр", Male, Nominative}
	examineFirstname(petrovich, test, t)

	test.expected = "Петра"
	test.caseName = Genitive
	examineFirstname(petrovich, test, t)

	test.expected = "Петру"
	test.caseName = Dative
	examineFirstname(petrovich, test, t)

	test.expected = "Петра"
	test.caseName = Accusative
	examineFirstname(petrovich, test, t)

	test.expected = "Петром"
	test.caseName = Instrumental
	examineFirstname(petrovich, test, t)

	test.expected = "Петре"
	test.caseName = Prepositional
	examineFirstname(petrovich, test, t)
}

func TestDoubleLastname(t *testing.T) {
	petrovich := LoadFromFile(rulesBytes)
	test := testCase{"Салтыков-Щедрин", "Салтыков-Щедрин", Male, Nominative}
	examineLastname(petrovich, test, t)

	test.expected = "Салтыкова-Щедрина"
	test.caseName = Genitive
	examineLastname(petrovich, test, t)

	test.expected = "Салтыкову-Щедрину"
	test.caseName = Dative
	examineLastname(petrovich, test, t)

	test.expected = "Салтыкова-Щедрина"
	test.caseName = Accusative
	examineLastname(petrovich, test, t)

	test.expected = "Салтыковым-Щедриным"
	test.caseName = Instrumental
	examineLastname(petrovich, test, t)

	test.expected = "Салтыкове-Щедрине"
	test.caseName = Prepositional
	examineLastname(petrovich, test, t)
}

func examineFirstname(petrovich *Petrovich, test testCase, t *testing.T) {
	actual, err := petrovich.FirstName(
		test.original,
		test.gender,
		test.caseName)
	examineAnswer(test, actual, err, t)
}

func examineLastname(petrovich *Petrovich, test testCase, t *testing.T) {
	actual, err := petrovich.LastName(
		test.original,
		test.gender,
		test.caseName)
	examineAnswer(test, actual, err, t)
}

func examineAnswer(test testCase, actual string, err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
	if actual != test.expected {
		t.Errorf("Wrong answer Actual: %s, Test Case: %s", actual, test)
	}
}

func readFile(fileName string) []byte {
	data, _ := ioutil.ReadFile(fileName)
	return data
}
