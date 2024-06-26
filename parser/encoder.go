package parser

import (
	"fmt"
	"os"
)

// WritePingviewTemplate Create the bulk add template for Pingview manager
//
// tgu: map of tgu and line data
func WritePingviewTemplate(tgu []Data) error {
	var (
		err  error    // error
		file *os.File // file to write to
	)

	// create output file
	file, err = os.Create(fmt.Sprintf("out/template/pingview_template.txt"))
	if err != nil {
		panic(err)
	}

	// defer file close
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// Sort he TGU array by zone and callsign before write it to file
	sortedTgu := sortAscendingWith2Fields(tgu, "Zone", "CallSign")

	// cycle through map and write to file according to template
	for _, linedata := range sortedTgu {
		var baseIp string // base IP address
		// Check if network ip field is a valid ipv4 and strip last octet
		if isValidIpv4(linedata.Network) {
			baseIp = linedata.Network[:len(linedata.Network)-2]
		} else {
			baseIp = "0.0.0"
			fmt.Printf("Error parsing\t-\t%s - %s\n", linedata.CallSign, linedata.City)
		}
		header := fmt.Sprintf("Group: %s %s %s - TGU: %s - %s - %s", linedata.CallSign, linedata.City, linedata.VType, linedata.Tgu, linedata.Address, linedata.Zone)
		gateway := fmt.Sprintf("%s.1 GATEWAY", baseIp)
		voip := fmt.Sprintf("%s.141 VOIP", baseIp)
		pc := fmt.Sprintf("%s.10 PC", baseIp)

		// Write data to file
		// write header
		_, err = file.WriteString(header + "\n")
		if err != nil {
			return err
		}

		// write gateway
		_, err = file.WriteString(gateway + "\n")
		if err != nil {
			return err
		}

		// write voip
		_, err = file.WriteString(voip + "\n")
		if err != nil {
			return err
		}

		// write pc
		_, err = file.WriteString(pc + "\n")
		if err != nil {
			return err
		}

		// write empty line
		_, err = file.WriteString("\n")
		if err != nil {
			return err
		}

		// flush buffer
		err = file.Sync()
		if err != nil {
			return err
		}
	}

	return nil
}
