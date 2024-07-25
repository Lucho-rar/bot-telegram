package data

import (
	"fmt"
	"strings"

	"google.golang.org/api/sheets/v4"
)

func ReadData(srv *sheets.Service, spreadsheetId string, readRange string) (*sheets.ValueRange, error) {
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func SearchVerb(srv *sheets.Service, spreadsheetId string, readRange string) (*sheets.ValueRange, error) {
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

func DataToString(data *sheets.ValueRange) string {
	var sb strings.Builder

	if len(data.Values) == 0 {
		sb.WriteString("No data found.")
	} else {
		// sb.WriteString("Data:\n")
		for _, row := range data.Values {
			for _, col := range row {
				//sb.WriteString(fmt.Sprintf("%d: %v", i+1, col))
				sb.WriteString(fmt.Sprintf("%v", col))
			}
		}
	}

	return sb.String()
}

func ReturnData(data *sheets.ValueRange) string {
	return DataToString(data)
}
