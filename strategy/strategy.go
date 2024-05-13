package strategy

type Sheets interface {
	Insert(ID int, rows [][]any) error
}

type GoogleSheets struct {
	GoogleSheetAPIClient struct{}
}

func (gs *GoogleSheets) Insert(ID int, rows [][]any) error {
	return nil
}

type Excel struct {
	ExcelAPIClient struct{}
}

func (e *Excel) Insert(ID int, rows [][]any) error {
	return nil
}
