package day19

import (
	"advent-of-code-2023/utils"
	"strings"
)

type Rule struct {
	property    rune
	operator    rune
	threshold   int
	destination string
}

func NewRule(ruleSegment string) Rule {
	if !strings.ContainsRune(ruleSegment, ':') {
		return Rule{
			destination: ruleSegment,
		}
	}
	branchSegments := strings.Split(ruleSegment, ":")
	return Rule{
		property:    rune(ruleSegment[0]),
		operator:    rune(ruleSegment[1]),
		threshold:   utils.ToInt(branchSegments[0][2:len(branchSegments[0])]),
		destination: branchSegments[1],
	}
}

type Workflow struct {
	name  string
	rules []Rule
}

func NewWorkflow(line string) Workflow {
	workflow := Workflow{}
	topLevel := strings.Split(line[0:len(line)-1], "{")
	workflow.name = topLevel[0]
	for _, ruleSegment := range strings.Split(topLevel[1], ",") {
		workflow.rules = append(workflow.rules, NewRule(ruleSegment))
	}
	return workflow
}

func (w Workflow) Process(r *Rating) string {
	for _, rule := range w.rules {
		if rule.property == 0 {
			return rule.destination
		}
		switch rule.property {
		case 'x':
			if rule.operator == '<' && r.x < rule.threshold {
				return rule.destination
			} else if rule.operator == '>' && r.x > rule.threshold {
				return rule.destination
			}
		case 'm':
			if rule.operator == '<' && r.m < rule.threshold {
				return rule.destination
			} else if rule.operator == '>' && r.m > rule.threshold {
				return rule.destination
			}
		case 'a':
			if rule.operator == '<' && r.a < rule.threshold {
				return rule.destination
			} else if rule.operator == '>' && r.a > rule.threshold {
				return rule.destination
			}
		default:
			if rule.operator == '<' && r.s < rule.threshold {
				return rule.destination
			} else if rule.operator == '>' && r.s > rule.threshold {
				return rule.destination
			}
		}
	}
	return "NO"
}

type Rating struct {
	x, m, a, s int
	accepted   bool
}

func (r *Rating) Screen(workflows map[string]Workflow) bool {
	destination := "in"
	for destination != "A" && destination != "R" {
		destination = workflows[destination].Process(r)
	}
	if destination == "A" {
		return true
	}
	return false
}

func (r *Rating) GetSum() int {
	return r.x + r.m + r.a + r.s
}

func NewRating(line string) Rating {
	segments := strings.Split(line[1:len(line)-1], ",")
	return Rating{
		x:        utils.ToInt(strings.Split(segments[0], "=")[1]),
		m:        utils.ToInt(strings.Split(segments[1], "=")[1]),
		a:        utils.ToInt(strings.Split(segments[2], "=")[1]),
		s:        utils.ToInt(strings.Split(segments[3], "=")[1]),
		accepted: false,
	}
}

func parseFile(filename string) (map[string]Workflow, []Rating) {
	workflows := make(map[string]Workflow)
	var ratings []Rating
	parseRatings := false
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			parseRatings = true
			continue
		}
		if parseRatings {
			ratings = append(ratings, NewRating(line))
		} else {
			workflow := NewWorkflow(line)
			workflows[workflow.name] = workflow
		}
	}
	return workflows, ratings
}

func Part1(filename string) int {
	answer := 0
	workflows, ratings := parseFile(filename)
	for _, rating := range ratings {
		if rating.Screen(workflows) {
			answer += rating.GetSum()
		}
	}
	return answer
}
