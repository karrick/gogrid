package gogrid

import (
	"fmt"
	"os"
)

func main() {
	var grid Grid
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	if true {
		grid.DefaultRowColor = "red"
		if err := grid.AppendRow([]string{"a1"}); err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

		grid.DefaultRowColor = "teal"
		if err := grid.AppendRow([]string{"a2"}); err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "grid: %v\n", grid)
	} else {
		grid.DefaultRowColor = "green"
		if err := grid.AppendColumn(Left, 13, []string{"a1", "a2"}); err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "grid: %v\n", grid)
	}

	grid.DefaultRowColor = "yellow"
	if err := grid.AppendRow([]string{"a3"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultRowColor = "red"
	if err := grid.AppendColumn(Center, 27, []string{"b1", "b2", "b3"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultRowColor = "teal"
	if err := grid.AppendRow([]string{"a4", "b4"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultRowColor = "green"
	if err := grid.AppendColumn(Right, 7, []string{"c1", "c2", "c3", "c4"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultRowColor = "yellow"
	if err := grid.AppendColumn(Center, 0, []string{"d1", "dd22", "ddd333", "dddd4444"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultRowColor = "red"
	if err := grid.AppendColumn(Center, 6, []string{"e1", "ee22", "eee333", "eeee4444"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.DefaultRowColor = "teal"
	if err := grid.AppendRow([]string{"a5", "b5", "c5", "d5", "e5"}); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "grid: %v\n", grid)

	grid.Delimiter = "|"
	// grid.HeaderColor = "green"
	lines := grid.Format()
	fmt.Printf("%d LINES\n", len(lines))
	for i, line := range lines {
		fmt.Printf("% 3d: %s\n", i+1, line)
	}
	fmt.Println("EOF")
	os.Exit(0)
}
