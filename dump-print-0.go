package gogrid

func (g *Grid) DumpPrint0(maxTextColumns int) {
	increment := 2
	if g.Narrow {
		increment = 1
	}

	columnsLen := len(g.cellWidths)            // number of table columns
	textColumnsAvailable := maxTextColumns - 2 // account for width of "| " on left of each row
	var ci int                                 // column index

	////////////////
	// Determine width of each column, to later use when wrapping the contents
	// of each cell into multiple lines. Prefer cells just wide enough for
	// their longest token.
	////////////////

	for ci = 0; ci < columnsLen; ci++ {
		if g.ColumnWidths[ci] < g.tokenWidths[ci] {
			g.ColumnWidths[ci] = g.tokenWidths[ci]
		}
		textColumnsAvailable -= g.ColumnWidths[ci] + increment // NOTE: includes " |" on right of each column.
	}

	// If possible, expand text width of as many of the columns as possible to
	// prevent each table column from being long and narrow.
	columnsNotExpanded := make([]int, 0, columnsLen)

	for ci = 0; ci < columnsLen && textColumnsAvailable > 0; ci++ {
		// Number of additional text columns that a cell could consume.
		additionalWidth := g.cellWidths[ci] - g.tokenWidths[ci]

		if additionalWidth <= textColumnsAvailable {
			g.ColumnWidths[ci] = g.cellWidths[ci]
			textColumnsAvailable -= additionalWidth
		} else {
			columnsNotExpanded = append(columnsNotExpanded, ci)
		}
	}
	// POST: ci points to the next table column that has yet to be expanded.

	if textColumnsAvailable > 0 {

		for ci = 0; ci < columnsLen; ci++ {
			if g.ColumnWidths[ci] < g.cellWidths[ci] {
				//
			}
		}

		if columnsLen > ci {
			// Additional width is the number of additional text columns that
			// each remaining column should be given in order to use as much
			// of maxTextColumns that are provided.
			additionalWidth := textColumnsAvailable / (columnsLen - ci)
			for ; ci < columnsLen; ci++ {
				g.ColumnWidths[ci] += additionalWidth
			}
		}
	}

	////////////////
	// Convert from rows to lines, knowing that because cells might need to
	// use more than their column width to hold their entire contents, some
	// cells require multiple lines.
	//
	// Each line is a string slice, where each element represents a single
	// column in that line.
	////////////////

	for ri, row := range g.data {
		_, _ = ri, row // TODO
	}

	// rows: [
	//   // row[0]:
	//   [
	//
	//   ],
	//   // row[1]:
	//   [
	//   ],
	// ]
}
