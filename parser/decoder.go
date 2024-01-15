package parser

import (
	"encoding/csv"
	"os"
	"strings"
)

// ParseCsv parses a CSV file and returns a slice of strings containing the subset excluding blank lines
func ParseCsv(filename string) ([][]string, error) {
	var (
		err        error       // error
		file       *os.File    // file to parse
		reader     *csv.Reader // csv reader
		rawdata    [][]string  // parsed raw data
		parseddata [][]string  // parsed data
	)

	// open file
	file, err = os.Open(filename)
	if err != nil {
		return nil, err
	}

	// close file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// read CSV raw data
	reader = csv.NewReader(file)
	reader.FieldsPerRecord = -0
	rawdata, err = reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// cycle through raw data and append to parsed data
	for i, row := range rawdata {
		// exclude 1st row
		if i == 0 {
			continue
		}

		city := toUpperCase(row[4])          // trim spaces from city and uppercase
		address := toTitleCase(row[3])       // Trim spaces from address and titlecase
		vtype := strings.TrimSpace(row[1])   // trim spaces from vehicle type
		callsign := toUpperCase(row[0])      // trim spaces for callsign
		tgu := strings.TrimSpace(row[5])     // trim spaces from TGU
		network := strings.TrimSpace(row[6]) // trim spaces from network

		// check if TGU is not blank in the CSV
		if tgu != "" {
			// append to parsed data
			parseddata = append(parseddata, []string{city, vtype, network, tgu, address, callsign})
		}
	}

	return parseddata, nil
}
