package gogrid

import (
	"errors"
	"fmt"
	"strings"
)

type Alignment int

const (
	Left Alignment = iota
	Center
	Right
)

// Cover is an enumeration of the ways to format the width of columns.
type Cover int

const (
	Cell     Cover = iota // Cell specifies column width should fit longest cell in column.
	Variable              // Variable specifies column width should be between Token and Cell, inclusive.
	Token                 // Token specifies column width should fit longest token.
	Fixed                 // Fixed specifies column should have a fixed width.
)

type Style int

const (
	StyleSpace Style = iota
	StyleBoxNarrow
	StyleBoxWide
)

type Grid struct {
	data             [][]string  // data[row][column] => cell
	colors           [][]string  // colors[row][column] to use for a cell
	ColumnAlignments []Alignment // controls column alignment for each column
	ColumnWidths     []int       // 0 implies use as many text columns as required for fields
	ColumnCovers     []Cover     // tracks how each column width should be determined
	cellWidths       []int       // tracks widest cell for each numbered column
	tokenWidths      []int       // tracks widest token for each numbered column
	HeaderColor      string      // empty implies no special header colorization
	DefaultCellColor string
	DefaultAlignment Alignment
	DefaultCover     Cover
	DefaultWidth     int
	Style            Style // Style specifies what style of table to render
	Narrow           bool  // Narrow specifies whether additional space between columns
}

// AppendColumn appends the specified column to the grid, specifying the new
// column's alignment and maximum width.
func (g *Grid) AppendColumn(alignment Alignment, maxWidth int, column []string) error {
	gridRowCount := len(g.data)
	var cellWidthsMax, tokenWidthsMax int

	if gridRowCount > 0 {
		if thisColumnRowCount := len(column); thisColumnRowCount != gridRowCount {
			return fmt.Errorf("cannot append column with different number of rows than already in grid: %d != %d", thisColumnRowCount, gridRowCount)
		}
		for ri, cell := range column {
			g.colors[ri] = append(g.colors[ri], g.DefaultCellColor)
			g.data[ri] = append(g.data[ri], cell)

			if cw := len(cell); cw > cellWidthsMax {
				cellWidthsMax = cw
			}
			for _, token := range strings.Fields(cell) {
				if tw := len(token); tw > tokenWidthsMax {
					tokenWidthsMax = tw
				}
			}
		}
	} else {
		// grid has no data yet; append one row for each cell provided
		for _, cell := range column {
			g.colors = append(g.colors, []string{g.DefaultCellColor})
			g.data = append(g.data, []string{cell})

			if cw := len(cell); cw > cellWidthsMax {
				cellWidthsMax = cw
			}
			for _, token := range strings.Fields(cell) {
				if tw := len(token); tw > tokenWidthsMax {
					tokenWidthsMax = tw
				}
			}
		}
	}

	g.ColumnAlignments = append(g.ColumnAlignments, alignment)
	g.ColumnWidths = append(g.ColumnWidths, maxWidth)
	g.cellWidths = append(g.cellWidths, cellWidthsMax)
	g.tokenWidths = append(g.tokenWidths, tokenWidthsMax)
	return nil
}

// AppendRow appends the specified columrow
func (g *Grid) AppendRow(column []string) error {
	var tokenWidthsMax int

	if len(g.data) > 0 {
		if lc, le := len(column), len(g.data[0]); lc != le {
			return fmt.Errorf("cannot append row when it has different number of columns than existing: %d != %d", lc, le)
		}
		for ci, cell := range column {
			if cw := len(cell); cw > g.cellWidths[ci] {
				g.cellWidths[ci] = cw
			}
			for _, token := range strings.Fields(cell) {
				if tw := len(token); tw > tokenWidthsMax {
					tokenWidthsMax = tw
				}
			}
		}
	} else {
		// grid has no data yet
		for _, cell := range column {
			for _, token := range strings.Fields(cell) {
				if tw := len(token); tw > tokenWidthsMax {
					tokenWidthsMax = tw
				}
			}

			g.ColumnAlignments = append(g.ColumnAlignments, g.DefaultAlignment)
			g.ColumnWidths = append(g.ColumnWidths, g.DefaultWidth)
			g.cellWidths = append(g.cellWidths, len(cell))
			g.tokenWidths = append(g.tokenWidths, tokenWidthsMax)
		}
	}

	// Default column color? Cell color?
	rowColors := make([]string, len(column))
	for i := range column {
		rowColors[i] = g.DefaultCellColor
	}
	g.colors = append(g.colors, rowColors)

	g.data = append(g.data, column)
	return nil
}

func (g *Grid) ColumnCellWidth(i int) (int, error) {
	if i < len(g.cellWidths) {
		return g.cellWidths[i], nil
	}
	return 0, errors.New("No such column")
}

func (g *Grid) ColumnCount() int {
	return len(g.ColumnAlignments)
}

func (g *Grid) ColumnTokenWidth(i int) (int, error) {
	if i < len(g.tokenWidths) {
		return g.tokenWidths[i], nil
	}
	return 0, errors.New("No such column")
}

func (g *Grid) RowCount() int {
	return len(g.data)
}

// Format returns the formatted grid.
//
// Do we want to be able to return formatted data without lines?
//
// Each outer slice represents one row of logical data. But then each logical
// row needs to include number of number of display rows required for the row,
// which is the longest number of lines for each cell.
//
// # Each inner string slice represents
//
// Requires each row to have same number of columns.
func (g *Grid) Format() []string {
	var columnDelimiter, linePrefix, lineSuffix string
	switch g.Style {
	case StyleBoxNarrow:
		columnDelimiter = "|"
		linePrefix = "|"
		lineSuffix = "|"
	case StyleBoxWide:
		columnDelimiter = " | "
		linePrefix = "| "
		lineSuffix = " |"
	default:
		// The default handles StyleSpace
		columnDelimiter = " "
	}

	lines := make([]string, len(g.data))

	for ri, row := range g.data {
		fields := make([]string, len(row))

		for ci, cell := range row {
			width := g.cellWidths[ci]
			if columnWidth := g.ColumnWidths[ci]; columnWidth > 0 {
				width = columnWidth
			}

			cell = align(g.ColumnAlignments[ci], width, cell)
			cell = withColors(g.colors[ri][ci], cell)

			fields[ci] = cell
		}

		lines[ri] = linePrefix + strings.Join(fields, columnDelimiter) + lineSuffix
	}

	return lines
}
