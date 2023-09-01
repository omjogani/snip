package helpers

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
	"github.com/omjogani/snip/models"
)

var supabase *supa.Client

func init() {
	// Grab Credential from .env & Connect to Supabase DB
	envError := godotenv.Load(".env")
	checkNilError(envError)

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase = supa.CreateClient(supabaseUrl, supabaseKey)
	color.Green("🚀 Success: Connected to DB! 🚀")
}

func GetSnipHelper() {
	// TODO: Get Snip and Print
	var results map[string]interface{}
	// var snipResult []models.Snip

	err := supabase.DB.From("snips").Select("*").Single().Execute(&results)
	checkNilError(err)
	fmt.Println(results)
	// var list []interface{}
	// for _, v := range results {
	// 	snipResult = append(snipResult, v)
	// }
	// printValuePrettier(snipResult)
}

func AddSnipHelper(snip models.Snip) bool {
	// TODO: Add Snip
	var snipResult []models.Snip
	err := supabase.DB.From("snips").Insert(snip).Execute(&snipResult)
	printValuePrettier(snipResult)
	checkNilError(err)
	return true
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

func printValuePrettier(snip []models.Snip) {
	color.Blue("------------------------------------------")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "TITLE", "DESCRIPTION", "DATE TIME"})
	for i := 0; i < len(snip); i++ {
		t.AppendRow(table.Row{
			snip[i].SnipId,
			snip[i].Title,
			snip[i].Description,
			snip[i].DateTime,
		})

	}
	t.AppendSeparator()
	t.AppendFooter(table.Row{"", "", "TOTAL", len(snip)})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	t.Render()
	color.Blue("------------------------------------------")
}

func checkNilError(err error) {
	if err != nil {

		fmt.Print("Error: ")
		fmt.Println(err)
	}
}
