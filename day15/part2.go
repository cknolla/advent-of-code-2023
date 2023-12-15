package day15

import (
	"slices"
	"strconv"
	"strings"
)

var boxes = make([]box, 256)

type lens struct {
	label  string
	box    int
	length int
}

func (l *lens) store() {
	for i, existingLens := range boxes[l.box].lenses {
		if existingLens.label == l.label {
			boxes[l.box].lenses[i] = *l
			return
		}
	}
	boxes[l.box].lenses = append(boxes[l.box].lenses, *l)
}

func retrieve(label string) {
	boxNum := hasher(label)
	for i, existingLens := range boxes[boxNum].lenses {
		if existingLens.label == label {
			boxes[boxNum].lenses = slices.Delete(boxes[boxNum].lenses, i, i+1)
			return
		}
	}
}

func newLens(label string, length int) lens {
	return lens{
		label:  label,
		box:    int(hasher(label)),
		length: length,
	}
}

type box struct {
	lenses []lens
}

func processStep(step string) {
	if strings.Contains(step, "=") {
		parts := strings.Split(step, "=")
		length, _ := strconv.Atoi(parts[1])
		l := newLens(parts[0], length)
		l.store()
	} else {
		label := strings.TrimRight(step, "-")
		retrieve(label)
	}
}

func getPower() (power int) {
	for _, b := range boxes {
		for slot, l := range b.lenses {
			power += (l.box + 1) * (slot + 1) * l.length
		}
	}
	return power
}

func Part2(filename string) int {
	answer := 0
	steps := parseFile(filename)
	for _, step := range steps {
		processStep(step)
	}
	answer = getPower()
	return answer
}
