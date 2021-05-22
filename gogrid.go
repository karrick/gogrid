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

type Grid struct {
	data             [][]string  // data[row][column] => cell
	colors           []string    // color to use for a row
	ColumnAlignments []Alignment // controls column alignment for each column
	ColumnWidths     []int       // 0 implies use as many text columns as required for fields
	dataWidths       []int       // tracks widest cell for each numbered column
	Delimiter        string      // Delimiter allows specifying a string to put between each column other than a single space.
	HeaderColor      string      // empty implies no special header colorization
	DefaultRowColor  string
	DefaultAlignment Alignment
	DefaultWidth     int
}

// AppendColumn appends the specified column to the grid, specifying the new
// column's alignment and maximum width.
func (g *Grid) AppendColumn(alignment Alignment, maxWidth int, column []string) error {
	gridRowCount := len(g.data)
	var dataWidthsMax int

	if gridRowCount > 0 {
		if thisColumnRowCount := len(column); thisColumnRowCount != gridRowCount {
			return fmt.Errorf("cannot append column with different number of rows than already in grid: %d != %d", thisColumnRowCount, gridRowCount)
		}
		for ri, cell := range column {
			g.data[ri] = append(g.data[ri], cell)
			if dw := len(cell); dw > dataWidthsMax {
				dataWidthsMax = dw
			}
		}
	} else {
		// grid has no data yet; append one row for each cell provided
		for _, cell := range column {
			g.data = append(g.data, []string{cell})
			g.colors = append(g.colors, g.DefaultRowColor)
			if dw := len(cell); dw > dataWidthsMax {
				dataWidthsMax = dw
			}
		}
	}

	g.ColumnAlignments = append(g.ColumnAlignments, alignment)
	g.ColumnWidths = append(g.ColumnWidths, maxWidth)
	g.dataWidths = append(g.dataWidths, dataWidthsMax)
	return nil
}

// AppendRow appends the specified columrow
func (g *Grid) AppendRow(column []string) error {
	if len(g.data) > 0 {
		if lc, le := len(column), len(g.data[0]); lc != le {
			return fmt.Errorf("cannot append row when it has different number of columns than existing: %d != %d", lc, le)
		}
		for ci, cell := range column {
			if l := len(cell); l > g.dataWidths[ci] {
				g.dataWidths[ci] = l
			}
		}
	} else {
		// grid has no data yet
		for _, cell := range column {
			g.ColumnAlignments = append(g.ColumnAlignments, g.DefaultAlignment)
			g.ColumnWidths = append(g.ColumnWidths, g.DefaultWidth)
			g.dataWidths = append(g.dataWidths, len(cell))
		}
	}

	g.colors = append(g.colors, g.DefaultRowColor)
	g.data = append(g.data, column)
	return nil
}

func (g *Grid) ColumnCount() int {
	return len(g.ColumnAlignments)
}

func (g *Grid) ColumnDataWidth(i int) (int, error) {
	if i < len(g.dataWidths) {
		return g.dataWidths[i], nil
	}
	return 0, errors.New("No such column")
}

func (g *Grid) RowCount() int {
	return len(g.data)
}

// Requires each row to have same number of columns.
func (g *Grid) Format() []string {
	lines := make([]string, len(g.data))

	var delim string
	if g.Delimiter == "" {
		delim = " "
	} else {
		delim = g.Delimiter
	}

	for ri, row := range g.data {
		var pre, post string
		if color := g.colors[ri]; color != "" {
			switch strings.ToLower(color) {
			case "bold":
				pre = "\033[1m"
			case "underscore":
				pre = "\033[1;4m"
			case "reverse":
				pre = "\033[1;7m"
			case "red":
				pre = "\033[1;31m"
			case "green":
				pre = "\033[1;32m"
			case "yellow":
				pre = "\033[1;33m"
			case "purple":
				pre = "\033[1;34m"
			case "magenta":
				pre = "\033[1;35m"
			case "teal":
				pre = "\033[1;36m"
			case "white":
				pre = "\033[1;37m"
			default:
				pre = g.colors[ri]
			}
			post = "\033[0m"
		}

		fields := make([]string, len(row))
		for ci, cell := range row {
			dataWidths := g.dataWidths[ci]
			if columnWidth := g.ColumnWidths[ci]; columnWidth > 0 {
				dataWidths = columnWidth
			}
			fields[ci] = align(g.ColumnAlignments[ci], dataWidths, pre, cell, post)
		}
		lines[ri] = strings.Join(fields, delim)
	}

	return lines
}

func align(alignment Alignment, width int, pre, field, post string) string {
	text := pre + field + post
	if width == 0 {
		return text
	}

	needed := width - len(field)
	if needed < 0 {
		return pre + field[:width] + post // trim upstream field
	}

	switch alignment {
	case Left:
		return text + strings.Repeat(" ", needed)
	case Center:
		half := needed >> 1
		double := half << 1
		ws := strings.Repeat(" ", half)
		if double == needed {
			return ws + text + ws
		}
		return ws + text + ws + " " // need extra whitespace on one of the sides
	case Right:
		return strings.Repeat(" ", needed) + text
	}
	panic("NOTREACHED")
}
