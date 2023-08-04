package sheets

type GoogleSheets interface {
	WithSheetID(sheetID string) GoogleSheets
	WithReadRange(readRange string) GoogleSheets
	Get()
}
