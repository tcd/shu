package shu

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

// Table prints a pretty table of all issues.
// https://github.com/alexeyco/simpletable
func Table(shoes Shoes) {
	data := tableFormat(shoes)

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Repo"},
			{Align: simpletable.AlignCenter, Text: "Number"},
			{Align: simpletable.AlignCenter, Text: "Title"},
		},
	}

	for _, row := range data {
		r := []*simpletable.Cell{
			{Text: row[0].(string)},
			{Text: row[1].(string)},
			{Text: row[2].(string)},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	// table.SetStyle(simpletable.StyleMarkdown)
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

// wrangle Shoes into a type suitable for printing with SimpleTable.
func tableFormat(shoes Shoes) [][]interface{} {
	data := make([][]interface{}, len(shoes.I))

	for i, shu := range shoes.I {
		var shuData = []interface{}{
			shu.Repo(),
			fmt.Sprintf("%d", shu.Number),
			shu.Title,
		}
		data[i] = shuData
	}
	return data
}
