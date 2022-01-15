package readcsv

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ReadCsv(filePath string, colIndex int) ([]int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var datas []int
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, row := range records {
		value, err := strconv.Atoi(row[colIndex])
		if err != nil {
			return nil, err
		}
		datas = append(datas, value)
	}

	return datas, nil
}
