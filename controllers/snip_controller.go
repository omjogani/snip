package controllers

import (
	"fmt"

	"github.com/fatih/color"
)

func GetSnip() {
	// TODO: Get Snip and Print
	fmt.Println("Here is your snips")
}

func AddSnip() {
	// TODO: Add Snip
	color.Blue("Added Sip")
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
