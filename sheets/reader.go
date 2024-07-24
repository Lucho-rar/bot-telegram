package sheets

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

func ReadData(srv *sheets.Service, spreadsheetId, readRange string) (*sheets.ValueRange, error) {
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PrintData(data *sheets.ValueRange) {
	if len(data.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Data:")
		for _, row := range data.Values {
			for i, col := range row {
				fmt.Printf("Column %d: %v\n", i+1, col)
			}
		}
	}
}
