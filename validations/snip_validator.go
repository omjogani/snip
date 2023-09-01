package validations

import (
	model "github.com/omjogani/snip/models"
)

func IsSnipDataValid(snip model.Snip) string {
	// Title Validations
	if len(snip.Title) < 1 {
		return "title-required"
	}

	// Title Validations
	if len(snip.Description) < 1 {
		return "description-required"
	}

	return "OK"
}
