package helpers

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

var supabase *supa.Client

func init() {
	// Grab Credential from .env
	envError := godotenv.Load(".env")
	checkNilError(envError)
	color.Green("🚀 Success: Connected to DB! 🚀")

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase = supa.CreateClient(supabaseUrl, supabaseKey)
}

func GetSnipHelper() {
	// TODO: Get Snip and Print
	fmt.Println("Here is your snips")
}

func AddSnipHelper() {
	// TODO: Add Snip
	color.Blue("Added Sip")
}

func RemoveSnipHelper() {
	// TODO: Remove Snip
}

func UpdateSnipHelper() {
	// TODO: Update Snip
}

func AppendSnipHelper() {
	// TODO: Append Snip
}

func checkNilError(err error) {
	if err != nil {
		color.Red("SNIP-ERROR: %s\n", err)
	}
}
