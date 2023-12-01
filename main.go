package main

import (
	"bufio"
)

// challenge is unused at the moment, but considering mapping individual
// parts selectable and executable, so keeping entrypoint signature standard
type challenge struct {
	entrypoint func(scanner *bufio.Scanner) (string, error)
	doExecute  bool
}
