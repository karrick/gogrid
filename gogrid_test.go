package gogrid

import "testing"

func TestGrid(t *testing.T) {
	t.Run("pre-define column widths", func(t *testing.T) {
		// Should be able to format each line as it's added, but this would
		// require less flexibility. Namely, the user would need to specify
		// column header and width before ingesting any rows.
	})

	// Grid columns are in row major, column minor order. Test cases will always
	// add data cells to fill up a grid using text data that matches typical
	// spreadsheet addresses.
	//
	// A1 B1 C1 D1
	// A2 B2 C2 D2
	// A3 B3 C3 D3
	// A4 B4 C4 D4
	// A5 B5 C5 D5
	// A6 B6 C6 D6

	t.Run("AppendColumn", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			t.Run("first widest", func(t *testing.T) {
				var g Grid

				_, err := g.ColumnCellWidth(0)
				ensureError(t, err, "No such column")

				ensureError(t, g.AppendColumn(Center, 13, []string{"aaa111", "aa22", "a3"}))

				if got, want := g.RowCount(), 3; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnCount(), 1; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}

				// column 0
				if got, want := g.ColumnAlignments[0], Center; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[0], 13; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err := g.ColumnCellWidth(0)
				ensureError(t, err)
				if got, want := cdw, 6; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
			})
			t.Run("middle widest", func(t *testing.T) {
				var g Grid

				_, err := g.ColumnCellWidth(0)
				ensureError(t, err, "No such column")

				ensureError(t, g.AppendColumn(Center, 13, []string{"aa11", "aaa222", "a3"}))

				if got, want := g.RowCount(), 3; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnCount(), 1; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}

				// column 0
				if got, want := g.ColumnAlignments[0], Center; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[0], 13; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err := g.ColumnCellWidth(0)
				ensureError(t, err)
				if got, want := cdw, 6; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
			})
			t.Run("last widest", func(t *testing.T) {
				var g Grid

				_, err := g.ColumnCellWidth(0)
				ensureError(t, err, "No such column")

				ensureError(t, g.AppendColumn(Center, 13, []string{"a1", "aa22", "aaa333"}))

				if got, want := g.RowCount(), 3; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnCount(), 1; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}

				// column 0
				if got, want := g.ColumnAlignments[0], Center; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[0], 13; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err := g.ColumnCellWidth(0)
				ensureError(t, err)
				if got, want := cdw, 6; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
			})
		})

		t.Run("not-empty", func(t *testing.T) {
			t.Run("mismatch rows", func(t *testing.T) {
				t.Run("too few", func(t *testing.T) {
					var g Grid
					ensureError(t, g.AppendColumn(Center, 11, []string{"a1", "a2"}))

					ensureError(t, g.AppendColumn(Right, 13, []string{"b11111"}), "different number of rows")

					if got, want := g.RowCount(), 2; got != want {
						t.Fatalf("GOT: %v; WANT: %v", got, want)
					}
					if got, want := g.ColumnCount(), 1; got != want {
						t.Fatalf("GOT: %v; WANT: %v", got, want)
					}

					// column 0
					if got, want := g.ColumnAlignments[0], Center; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
					if got, want := g.ColumnWidths[0], 11; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
					cdw, err := g.ColumnCellWidth(0)
					ensureError(t, err)
					if got, want := cdw, 2; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
				})

				t.Run("too many", func(t *testing.T) {
					var g Grid
					ensureError(t, g.AppendColumn(Center, 11, []string{"a1", "a2"}))

					ensureError(t, g.AppendColumn(Right, 13, []string{"c11111", "c2", "c3"}), "different number of rows")

					if got, want := g.RowCount(), 2; got != want {
						t.Fatalf("GOT: %v; WANT: %v", got, want)
					}
					if got, want := g.ColumnCount(), 1; got != want {
						t.Fatalf("GOT: %v; WANT: %v", got, want)
					}

					// column 0
					if got, want := g.ColumnAlignments[0], Center; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
					if got, want := g.ColumnWidths[0], 11; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
					cdw, err := g.ColumnCellWidth(0)
					ensureError(t, err)
					if got, want := cdw, 2; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
				})
			})

			t.Run("just right", func(t *testing.T) {
				var g Grid
				ensureError(t, g.AppendColumn(Center, 11, []string{"a1", "aa22", "aaa333"}))

				ensureError(t, g.AppendColumn(Right, 13, []string{"bbbb11111", "bbbbb22222", "bbbbbb333333"}))

				if got, want := g.RowCount(), 3; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnCount(), 2; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}

				// column 0
				if got, want := g.ColumnAlignments[0], Center; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[0], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err := g.ColumnCellWidth(0)
				ensureError(t, err)
				if got, want := cdw, 6; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}

				// column 1
				if got, want := g.ColumnAlignments[1], Right; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[1], 13; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err = g.ColumnCellWidth(1)
				ensureError(t, err)
				if got, want := cdw, 12; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
			})
		})
	})

	t.Run("AppendRow", func(t *testing.T) {
		t.Run("empty", func(t *testing.T) {
			t.Run("first widest", func(t *testing.T) {
				var g Grid
				g.DefaultAlignment = Left
				g.DefaultWidth = 11

				_, err := g.ColumnCellWidth(0)
				ensureError(t, err, "No such column")

				ensureError(t, g.AppendRow([]string{"aaa111", "bb11", "c1"}))

				if got, want := g.RowCount(), 1; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnCount(), 3; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}

				// column 0
				if got, want := g.ColumnAlignments[0], Left; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[0], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err := g.ColumnCellWidth(0)
				ensureError(t, err)
				if got, want := cdw, 6; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}

				// column 1
				if got, want := g.ColumnAlignments[1], Left; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[1], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err = g.ColumnCellWidth(1)
				ensureError(t, err)
				if got, want := cdw, 4; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}

				// column 2
				if got, want := g.ColumnAlignments[2], Left; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[2], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err = g.ColumnCellWidth(2)
				ensureError(t, err)
				if got, want := cdw, 2; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
			})
			t.Run("middle widest", func(t *testing.T) {
				var g Grid
				g.DefaultAlignment = Center
				g.DefaultWidth = 11

				_, err := g.ColumnCellWidth(0)
				ensureError(t, err, "No such column")

				ensureError(t, g.AppendRow([]string{"aa11", "bbb111", "c1"}))

				if got, want := g.RowCount(), 1; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnCount(), 3; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}

				// column 0
				if got, want := g.ColumnAlignments[0], Center; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[0], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err := g.ColumnCellWidth(0)
				ensureError(t, err)
				if got, want := cdw, 4; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}

				// column 1
				if got, want := g.ColumnAlignments[1], Center; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[1], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err = g.ColumnCellWidth(1)
				ensureError(t, err)
				if got, want := cdw, 6; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}

				// column 2
				if got, want := g.ColumnAlignments[2], Center; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[2], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err = g.ColumnCellWidth(2)
				ensureError(t, err)
				if got, want := cdw, 2; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
			})
			t.Run("last widest", func(t *testing.T) {
				var g Grid
				g.DefaultAlignment = Right
				g.DefaultWidth = 11

				_, err := g.ColumnCellWidth(0)
				ensureError(t, err, "No such column")

				ensureError(t, g.AppendRow([]string{"a1", "bb11", "ccc111"}))

				if got, want := g.RowCount(), 1; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnCount(), 3; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}

				// column 0
				if got, want := g.ColumnAlignments[0], Right; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[0], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err := g.ColumnCellWidth(0)
				ensureError(t, err)
				if got, want := cdw, 2; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}

				// column 1
				if got, want := g.ColumnAlignments[1], Right; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[1], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err = g.ColumnCellWidth(1)
				ensureError(t, err)
				if got, want := cdw, 4; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}

				// column 2
				if got, want := g.ColumnAlignments[2], Right; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[2], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err = g.ColumnCellWidth(2)
				ensureError(t, err)
				if got, want := cdw, 6; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
			})
		})

		t.Run("non-empty", func(t *testing.T) {
			t.Run("mismatch rows", func(t *testing.T) {
				t.Run("too few", func(t *testing.T) {
					var g Grid
					g.DefaultAlignment = Right
					g.DefaultWidth = 11
					ensureError(t, g.AppendRow([]string{"a1", "bb11"}))

					ensureError(t, g.AppendRow([]string{"a2"}), "different number of columns")

					t.Logf("%#v\n", g)

					if got, want := g.RowCount(), 1; got != want {
						t.Fatalf("GOT: %v; WANT: %v", got, want)
					}
					if got, want := g.ColumnCount(), 2; got != want {
						t.Fatalf("GOT: %v; WANT: %v", got, want)
					}

					// column 0
					if got, want := g.ColumnAlignments[0], Right; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
					if got, want := g.ColumnWidths[0], 11; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
					cdw, err := g.ColumnCellWidth(0)
					ensureError(t, err)
					if got, want := cdw, 2; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}

					// column 1
					if got, want := g.ColumnAlignments[1], Right; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
					if got, want := g.ColumnWidths[1], 11; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
					cdw, err = g.ColumnCellWidth(1)
					ensureError(t, err)
					if got, want := cdw, 4; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
				})

				t.Run("too many", func(t *testing.T) {
					var g Grid
					g.DefaultAlignment = Right
					g.DefaultWidth = 11
					ensureError(t, g.AppendRow([]string{"a1", "aa22"}))

					ensureError(t, g.AppendRow([]string{"c11111", "c2", "c3"}), "different number of columns")

					if got, want := g.RowCount(), 1; got != want {
						t.Fatalf("GOT: %v; WANT: %v", got, want)
					}
					if got, want := g.ColumnCount(), 2; got != want {
						t.Fatalf("GOT: %v; WANT: %v", got, want)
					}

					// column 0
					if got, want := g.ColumnAlignments[0], Right; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
					if got, want := g.ColumnWidths[0], 11; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
					cdw, err := g.ColumnCellWidth(0)
					ensureError(t, err)
					if got, want := cdw, 2; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}

					// column 1
					if got, want := g.ColumnAlignments[1], Right; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
					if got, want := g.ColumnWidths[1], 11; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
					cdw, err = g.ColumnCellWidth(1)
					ensureError(t, err)
					if got, want := cdw, 4; got != want {
						t.Errorf("GOT: %v; WANT: %v", got, want)
					}
				})
			})

			t.Run("just right", func(t *testing.T) {
				var g Grid
				g.DefaultAlignment = Right
				g.DefaultWidth = 11
				ensureError(t, g.AppendRow([]string{"a1", "bb11", "ccc111"}))

				ensureError(t, g.AppendRow([]string{"a2", "b2", "cccc2222"}))

				if got, want := g.RowCount(), 2; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnCount(), 3; got != want {
					t.Fatalf("GOT: %v; WANT: %v", got, want)
				}

				// column 0
				if got, want := g.ColumnAlignments[0], Right; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[0], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err := g.ColumnCellWidth(0)
				ensureError(t, err)
				if got, want := cdw, 2; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}

				// column 1
				if got, want := g.ColumnAlignments[1], Right; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[1], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err = g.ColumnCellWidth(1)
				ensureError(t, err)
				if got, want := cdw, 4; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}

				// column 2
				if got, want := g.ColumnAlignments[2], Right; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				if got, want := g.ColumnWidths[2], 11; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
				cdw, err = g.ColumnCellWidth(2)
				ensureError(t, err)
				if got, want := cdw, 8; got != want {
					t.Errorf("GOT: %v; WANT: %v", got, want)
				}
			})
		})
	})
}
