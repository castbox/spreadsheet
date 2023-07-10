package spreadsheet

import "google.golang.org/api/sheets/v4"

// CellData is data about a specific cell.
type CellData struct {
	UserEnteredValue  ExtendedValue      `json:"userEnteredValue"`
	EffectiveValue    ExtendedValue      `json:"effectiveValue"`
	FormattedValue    string             `json:"formattedValue"`
	UserEnteredFormat *sheets.CellFormat `json:"userEnteredFormat"`
	// EffectiveFormat *CellFormat `json:"effectiveFormat"`
	Hyperlink string `json:"hyperlink"`
	Note      string `json:"note"`
	// TextFormatRuns []*TextFormatRun `json:"textFormatRuns"`
	// DataValidation *DataValidationRule `json:"dataValidation"`
	// PivotTable *PivotTable `json:"pivotTable"`
}
