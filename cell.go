package spreadsheet

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

// Cell describes a cell data
type Cell struct {
	Row               uint
	Column            uint
	Value             string
	Note              string
	UserEnteredFormat *sheets.CellFormat

	modifiedFields string
}

// Pos returns the cell's position like "A1"
func (cell *Cell) Pos() string {
	return numberToLetter(int(cell.Column)+1) + fmt.Sprintf("%d", cell.Row+1)
}
