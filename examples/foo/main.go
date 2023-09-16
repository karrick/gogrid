package main

import (
	"fmt"
	"os"

	"github.com/karrick/gogrid"
)

func main() {
	var grid gogrid.Grid
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultCellColor = "red"
	if err := grid.AppendRow([]string{"red 1a"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultCellColor = "green"
	if err := grid.AppendRow([]string{"green 2a"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultCellColor = "yellow"
	if err := grid.AppendRow([]string{"yellow 3a"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultCellColor = "magenta"
	if err := grid.AppendColumn(gogrid.Center, 0, []string{"magenta 1b", "magenta 2b", "magenta 3b"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultCellColor = "magenta"
	if err := grid.AppendRow([]string{"magenta 4a", "magenta 4b"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultCellColor = "cyan"
	if err := grid.AppendColumn(gogrid.Right, 0, []string{"cyan 1c", "cyan 2c", "cyan 3c", "cyan 4c"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultCellColor = "white"
	if err := grid.AppendColumn(gogrid.Center, 0, []string{"white 1d", "white 2d", "white 3d", "white 4d"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultCellColor = "red"
	if err := grid.AppendColumn(gogrid.Right, 0, []string{"red 1e", "red 2e", "red 3e", "red 4e"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultCellColor = "green"
	if err := grid.AppendRow([]string{"green 5a", "green 5b", "green 5c", "green 5d", "green 5e"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	// grid.HeaderColor = "green"
	grid.Style = gogrid.StyleBoxNarrow
	// grid.Style = gogrid.StyleBoxWide
	lines := grid.Format()
	fmt.Printf("%d LINES\n", len(lines))
	for i, line := range lines {
		fmt.Printf("% 3d: %s\n", i+1, line)
	}
	fmt.Println("EOF")

	grid.DumpPrint2(80)
	fmt.Println("EOF")

	os.Exit(0)
}
