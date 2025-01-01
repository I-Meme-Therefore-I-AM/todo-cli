package todo

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
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

	data := unmarshTasks(&t)
	err = wCsv.Write(data)

	if err != nil {
		return err
	}
	wCsv.Flush()
	return nil
}

// read csv task
func readTask(f *os.File, all bool) {
	rCsv := csv.NewReader(f)

	r, err := rCsv.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	wr := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(wr, "ID\tDescription\tCreatedAt")
	for index, fr := range r {
		if index == 0 {
			continue
		}
		d, err := marshTasks(fr)

		if err != nil {
			fmt.Fprintln(os.Stdout, err)
		}
		if d.IsComplete == true && all != true {
			fmt.Fprintf(wr, "%d\t%s\t%s\n", d.ID, d.Description, d.CreatedAt)
		} else if d.IsComplete != true && all != true {
			continue
		} else {
			fmt.Fprintf(wr, "%d\t%s\t%s\t%v\n", d.ID, d.Description,
				timediff.TimeDiff(d.CreatedAt),
				d.IsComplete)
		}
	}
	wr.Flush()
}
