package controllers

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/omjogani/snip/helpers"
	"github.com/omjogani/snip/models"
	"github.com/omjogani/snip/validations"
)

func GetSnip() {
	// TODO: Get Snip and Print
	helpers.GetSnipHelper()
}

func AddSnip() {
	// get data from user

	var title string
	var description string
	var snip models.Snip

	snip.SnipId = 3
	fmt.Printf("Title: ")
	in := bufio.NewReader(os.Stdin)
	title, errTitle := in.ReadString('\n')
	checkNilError(errTitle)
	// fmt.Scanf(&title)
	snip.Title = title
	fmt.Printf("Description: ")
	description, errDescription := in.ReadString('\n')
	checkNilError(errDescription)

	snip.Description = description
	snip.DateTime = time.Now()

	// validate that data
	validationResponse := validations.IsSnipDataValid(snip)
	if validationResponse == "title-required" {
		color.Red("Invalid 'title' data! - Try Again")
		return
	} else if validationResponse == "description-required" {
		color.Red("Invalid 'description' data! - Try Again")
		return
	} else if validationResponse == "OK" {
		if helpers.AddSnipHelper(snip) {
			color.Green("✔ Snip Added!")
		}
	}
	// acknowledge user

}

func RemoveSnip() {
	// TODO: Remove Snip
}

func UpdateSnip() {
	// TODO: Update Snip
}

func AppendSnip() {
	// TODO: Append Snip
}

func checkNilError(err error) {
	if err != nil {
		color.Red("SNIP-ERROR: %s\n", err)
	}
}
