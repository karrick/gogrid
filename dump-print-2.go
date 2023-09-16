package gogrid

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Loop through columns once, tracking the total width for two scenarios: when
// the column widths are set to the max token length, and when the column
// widths are set to the max cell length.
//
// After loop through columns, if terminal width is less than sum of token
// widths, then use token widths.
//
// Otherwise, if terminal width is greater than sum of cell widths, use cell
// widths.
//
// Otherwise, when terminal width is closer to total cell widths, start from
// cell widths and scale down.
//
// Otherwise, terminal width must be closer to total token widths. Start from
// token widths and scale up.
//
// When terminal width is closer to sum of token widths, then use the token
// widths slice as the base, and expand each column by granting it a portion of
// the additional width for the terminal. But each not more than its cell width.
//
// When terminal width is closer to sum of cell widths, then use the cell widths
// slice as the base, and ...
func (g *Grid) DumpPrint2(terminalWidth int) {
	increment := 2
	if g.Narrow {
		increment = 1
	}

	totalWithCellWidths := increment  // include columns for "| ".
	totalWithTokenWidths := increment // include columns for "| ".

	columnsLen := len(g.cellWidths) // number of table columns
	resizableColumns := make([]int, 0, columnsLen)

	for ci := 0; ci < columnsLen; ci++ {
		if g.ColumnWidths[ci] > 0 {
			// Set column width to predefined value
			totalWithCellWidths += g.ColumnWidths[ci] + increment
			totalWithTokenWidths += g.ColumnWidths[ci] + increment

			g.cellWidths[ci] = g.ColumnWidths[ci]
			g.tokenWidths[ci] = g.ColumnWidths[ci]
		} else {
			totalWithCellWidths += g.cellWidths[ci] + increment
			totalWithTokenWidths += g.tokenWidths[ci] + increment

			// Track which columns are resizable.
			resizableColumns = append(resizableColumns, ci)
		}
	}

	diffToken := terminalWidth - totalWithTokenWidths // 80 - 24 = 56; 80 - 88 = -8
	diffCell := terminalWidth - totalWithCellWidths   // 80 - 46 = 34; 80 - 128 = -48

	var widths []int
	var scaleUp bool

	if diffToken < 0 {
		fmt.Fprintf(os.Stderr, "TODO use token widths (results will extend beyond max)\n")
		// Alternatively, could reduce by difference, causing only a few
		// places where tokens were split.
		widths = g.tokenWidths
	} else if diffCell > 0 {
		// Everything fits in the terminal width.
		fmt.Fprintf(os.Stderr, "TODO use cell widths\n")
		widths = g.cellWidths
	} else if diffCell < diffToken {
		fmt.Fprintf(os.Stderr, "TODO start from cell widths; scale down\n")
		widths = g.cellWidths
	} else {
		fmt.Fprintf(os.Stderr, "TODO start from token widths; scale up\n")
		widths = g.tokenWidths
		scaleUp = true
	}

	resizableColumnsLen := len(resizableColumns)

	if scaleUp {
		fmt.Fprintf(os.Stderr, "TODO: scale down\n")
		perColumnAddition := diffToken / resizableColumnsLen
		for _, ci := range resizableColumns {
			widths[ci] += perColumnAddition
		}
	} else {
		fmt.Fprintf(os.Stderr, "TODO: scale down\n")
	}

	fmt.Printf("| ")
	for _, width := range widths {
		fmt.Printf("%s |", align(Center, width, strconv.Itoa(width)))
	}
	fmt.Printf("\n")

	// row 0 [
	//     row 0 line 0 [
	//         row 0 line 0 column 0 "one alpha first"
	//         row 0 line 0 column 1 "one bravo first"
	//         row 0 line 0 column 1 "one charlie first"
	//     ]
	//     row 0 line 1 [
	//         row 0 line 0 column 0 "one alpha second"
	//         row 0 line 0 column 1 "one charlie second"
	//     ]
	// ]
	// row 1 [
	//     row 1 line 0 [
	//         row 1 line 0 column 0 "two alpha first"
	//         row 1 line 0 column 1 "two bravo first"
	//         row 1 line 0 column 1 "two charlie first"
	//     ]
	//     row 1 line 1 [
	//         row 1 line 0 column 0 "two alpha second"
	//         row 1 line 0 column 1 "two charlie second"
	//     ]
	// ]

	var rowTerminalLines [][][]string

	for _, row := range g.data {
		rowLines := [][]string{}

		for ci, cell := range row {
			// line wrap this cell to a slice of strings
			cellLines := []string{}
			words := strings.Fields(cell)
			line := words[0]
			words = words[1:]

			for _, word := range words {
				if len(line)+1+len(word) > widths[ci] {
					cellLines = append(cellLines, line)
				} else {
					line += " "
				}
				line += word
			}

			if len(line) > 0 {
				cellLines = append(cellLines, line)
			}

			// POST: cellLines is a slice of strings, one string for each
			// terminal line that this cell requires.

			// ["one alpha first", "one alpha second"]
			rowLines = append(rowLines, cellLines)
		}

		rowTerminalLines = append(rowTerminalLines, invertMatrix(rowLines))
	}

	// return strings.Repeat(" ", needed) + text
	interval := "|"
	for _, width := range widths {
		interval += strings.Repeat("-", width) + "|"
	}
	interval += "\n"

	fmt.Print(interval)

	for ri, rowTerminalLine := range rowTerminalLines {
		fmt.Printf("|")
		for ci := range rowTerminalLine {
			cell := "foo" // text
			cell = align(g.ColumnAlignments[ci], widths[ci], cell)
			cell = withColors(g.colors[ri][ci], cell)
			fmt.Printf("%s|", cell)
		}
		fmt.Printf("\n%s", interval)
	}
}

func invertMatrix(rowLines [][]string) [][]string {
	// POST: rowLines:
	//     [
	//         ["one alpha first", "one alpha second"],
	//         ["one bravo first"],
	//         ["one charlie first", "one charlie second"],
	//     ]

	// TODO: terminalLines:
	//     [
	//         ["one alpha first", "one bravo first", "one charlie first"],
	//         ["one alpha second", "",               "one charlie second"],
	//     ]
	return rowLines
	// var terminalLines [][]string
	// for _, rowLine := range rowLines {
	// 	for _, cellLine := range rowLine {

	// 	}
	// }
}
