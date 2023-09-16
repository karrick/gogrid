package gogrid

// | one alpha first | two bravo first |
// | one alpha second | two bravo second |

// For each row, compute column width by using cell width, and for each cell
// subtract its text width from the available text columns.
//
// If available text columns is less than zero, then divide it by number of
// columns to determine how many text columns each table column must yield.
//
// Then loop through table columns, and for each column, if reducing by above
// quotient reduces number of text columns to be lower than that column's max
// token length, then reduce that column width to that max token legth, and
// restart, after eliminating this column from future reductions.
//
// If however, reducing by quotient does not reduce below any remaining
// column's respective max token length, then reduce each column length
// accordingly.

func (g *Grid) DumpPrint1(maxTextColumns int) {
	increment := 2
	if g.Narrow {
		increment = 1
	}

	columnsLen := len(g.cellWidths) // number of table columns
	columns := make([]int, columnsLen)
	columnWidths := make([]int, columnsLen)

	textColumnsAvailable := maxTextColumns - 2 // account for width of "| " on left of each row

	// Determine width of each column without wrapping cell contents to
	// multiple lines.
	for ci := 0; ci < columnsLen; ci++ {
		if g.ColumnWidths[ci] > 0 {
			// Set width to predefined value.
			columnWidths[ci] = g.ColumnWidths[ci]
		} else {
			// Set its width to the width of the cell contents.
			g.ColumnWidths[ci] = g.cellWidths[ci]
		}
		textColumnsAvailable -= columnWidths[ci] + increment // NOTE: includes " |" on right of each column.

		// Create a list of columns indexes available for reduction in next
		// step.
		columns[ci] = ci
	}

outer:
	// NOTE: When textColumnsAvailable is negative, then this needs to reclaim
	// some of the text columns from each table column.
	for textColumnsAvailable < 0 && columnsLen > 0 {
		// quotient is number of text columns that each table column must give
		// up in order to fit each cell. NOTE: integer division will round
		// down.
		quotient := -textColumnsAvailable / columnsLen

		for _, ci := range columns {
			if textColumnsAvailable >= 0 || columnsLen == 0 {
				break outer
			}

			// available is the number of text columns that this table column
			// can yield.
			available := columnWidths[ci] - g.tokenWidths[ci] // 20 - 15 = 5

			if available < quotient {
				// This column may not be reduced beyond its max token width.
				columnWidths[ci] = g.tokenWidths[ci]
				textColumnsAvailable += available

				// Remove the index of this column from the slice of columns
				// available to shrink.
				copy(columns[ci:], columns[ci+1:]) // shift items to the right of ci to the left by one index
				columnsLen--
				columns = columns[:columnsLen] // shrink slice to drop off final duplicated element

				// ??? I am not certain I want to continue outer here, because
				// it potentially shrinks earlier columns by more than later
				// columns.
				//
				// BUT this cannot simply continue on the inner loop because
				// the loop elements have changed.
				continue outer
			}
		}

		// All column indices in columns slice can be reduced by quotient.
		for _, ci := range columns {
			columnWidths[ci] -= quotient
			textColumnsAvailable += quotient
		}
	}

	// POST: Each column has been reduced by as much as needed up to as much as possible.

}
