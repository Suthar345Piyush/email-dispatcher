// producer will read all the emails

package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadRecipient(filePath string, ch chan Recipient) error {
	defer close(ch)

	f, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer f.Close()

	// creating a reader to read the csv data

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		fmt.Println(record)

		// sending data into channel
		ch <- Recipient{
			Name:  record[0],
			Email: record[1],
		}

		// send -> consumer -> unbuffered channel
	}

	return nil

}
