package helpers

import (
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

func PrintTable(header []string, rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(rows) // The data in the table
	table.Render()         // Render the table
}

func LocalTime() string {
	return time.Now().Local().String()
}

func GetTimeBasedOnLoc(c string) (string, error) {
	loc, err := time.LoadLocation(c)
	if err != nil {
		return "", err
	}
	return time.Now().In(loc).String(), nil

}
