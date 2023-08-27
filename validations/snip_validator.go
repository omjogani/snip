package validations

import (
	"regexp"

	model "github.com/omjogani/snip/models"
)

func isSnipValid(snip model.Snip) string {
	isSnipEmpty := regexp.MustCompile(`/^(?!\s*$).+/`)

	// Title Validations
	if !isSnipEmpty.MatchString(snip.Title) {
		return "title-required"
	}

	return "OK"
}
