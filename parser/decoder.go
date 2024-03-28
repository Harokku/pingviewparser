package parser

import (
	"encoding/csv"
	"os"
	"strings"
)

// ParseCsv parses a CSV file and returns a slice of strings containing the subset excluding blank lines
func ParseCsv(config ParserConfig) ([]Data, error) {
	var (
		err        error       // error
		file       *os.File    // file to parse
		reader     *csv.Reader // csv reader
		rawdata    [][]string  // parsed raw data
		parseddata []Data      // parsed data
	)

	// open file
	file, err = os.Open(config.Filename)
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

		d := Data{
			Zone:     toUpperCase(row[0]),       // trim spaces for zone and uppercase
			CallSign: toUpperCase(row[1]),       // trim spaces for callsign
			VType:    strings.TrimSpace(row[2]), // trim spaces from vehicle type
			City:     toUpperCase(row[5]),       // trim spaces from city and uppercase
			Address:  toTitleCase(row[4]),       // Trim spaces from address and titlecase
			Tgu:      strings.TrimSpace(row[6]), // trim spaces from TGU
			Network:  strings.TrimSpace(row[7]), // trim spaces from network
		}

		// filter out not requested rows
		if toUpperCase(config.Zone) != d.Zone && config.Zone != "all" {
			continue
		}

		// check if TGU is not blank in the CSV
		if d.Tgu != "" {
			// append to parsed data
			parseddata = append(parseddata, d)
		}
	}

	return parseddata, nil
}
