package todo

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

func getLastIndex(f *os.File) (int, error) {
	rCsv := csv.NewReader(f)

	data, err := rCsv.ReadAll()
	if err != nil {
		return 0, err
	}

	lenData := len(data)
	if lenData == 0 {
		lenData = 1
	}
	return lenData, nil
}
func writeToCsv(f *os.File, data []string) error {
	wCsv := csv.NewWriter(f)

	err := wCsv.Write(data)

	if err != nil {
		return err
	}
	wCsv.Flush()
	return nil
}

func appendToCsv(f *os.File, t Tasks) error {
	wCsv := csv.NewWriter(f)

	index, err := getLastIndex(f)

	if err != nil {
		return err
	}

	t.ID = index
	t.CreatedAt = time.Now()
	t.IsComplete = false

	data := []string{strconv.Itoa(t.ID), t.Description, t.CreatedAt.String(), strconv.FormatBool(t.IsComplete)}
	err = wCsv.Write(data)

	if err != nil {
		return err
	}
	wCsv.Flush()
	return nil
}
