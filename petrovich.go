package petrovich

import (
	"fmt"
	"log"
	"strings"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

const (
	//Male gender constant
	Male = "male"
	//Female gender constant
	Female = "female"
	//Androgynous gender constant
	Androgynous = "androgynous"
)

const (
	//Nominative case constant
	Nominative = "nominative"
	//Genitive case constant
	Genitive = "genitive"
	//Dative case constant
	Dative = "dative"
	//Accusative case constant
	Accusative = "accusative"
	//Instrumental case constant
	Instrumental = "instrumental"
	//Prepositional case constant
	Prepositional = "prepositional"
)

//Rules TODO
type Rules struct {
	Lastname   ruleGroup
	Firstname  ruleGroup
	Middlename ruleGroup
}

type ruleGroup struct {
	Exceptions []rule
	Suffixes   []rule
}

type rule struct {
	Gender string
	Test   []string
	Mods   []string
	Tags   []string
}

// ErrPetrovich TODO
type ErrPetrovich string

func (e ErrPetrovich) Error() string {
	return string(e)
}

// FirstName TODO
func (rules Rules) FirstName(firstName string, gender string, caseName string) (string, error) {
	return convertTo(firstName, gender, caseName, &rules.Firstname)
}

// LastName TODO
func (rules Rules) LastName(lastName string, gender string, caseName string) (string, error) {
	return convertTo(lastName, gender, caseName, &rules.Lastname)
}

// MiddleName TODO
func (rules Rules) MiddleName(middleName string, gender string, caseName string) (string, error) {
	return convertTo(middleName, gender, caseName, &rules.Middlename)
}

//LoadRules TODO
func LoadRules(data []byte) Rules {
	t := Rules{}
	err := yaml.Unmarshal(data, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return t
}

func convertTo(name string, gender string, caseName string, ruleGroup *ruleGroup) (string, error) {
	return findAndApply(name, gender, caseName, ruleGroup)
}

func findAndApply(name string, gender string, caseName string, ruleGroup *ruleGroup) (string, error) {
	rule, err := findRule(name, gender, ruleGroup.Exceptions)
	if err != nil {
		rule, err = findRule(name, gender, ruleGroup.Suffixes)
	}
	if err != nil {
		return "", err
	}
	return apply(name, caseName, &rule)
}

func findRule(name string, gender string, rules []rule) (rule, error) {
	for _, rule := range rules {
		if matchRule(name, gender, &rule) {
			return rule, nil
		}
	}
	return rule{}, ErrPetrovich(fmt.Sprintf("Rule not found for name: %s; gender: %s;", name, gender))
}

func matchRule(name string, gender string, rule *rule) bool {
	if (rule.Gender == Male && gender == Female) ||
		(rule.Gender == Female && gender == Male) {
		return false
	}
	name = strings.ToLower(name)
	for _, test := range rule.Test {
		if strings.HasSuffix(name, test) {
			return true
		}
	}
	return false
}

func apply(name string, caseName string, rule *rule) (string, error) {
	caseModifier, err := findCaseModifier(caseName, rule)
	if err != nil {
		return "", err
	}
	for _, c := range caseModifier {
		switch c {
		case '.':
		case '-':
			_, size := utf8.DecodeLastRuneInString(name)
			name = name[:len(name)-size]
		default:
			name += string(c)
		}
	}
	return name, nil
}

func findCaseModifier(caseName string, rule *rule) (string, error) {
	switch caseName {
	case Nominative:
		return "", nil
	case Genitive:
		return rule.Mods[0], nil
	case Dative:
		return rule.Mods[1], nil
	case Accusative:
		return rule.Mods[2], nil
	case Instrumental:
		return rule.Mods[3], nil
	case Prepositional:
		return rule.Mods[4], nil
	default:
		return "", ErrPetrovich(fmt.Sprintf("Unknown grammatical case: %s;", caseName))
	}
}
