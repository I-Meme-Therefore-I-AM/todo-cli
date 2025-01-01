package todo

import (
	"strconv"
	"time"
)

var layOut = "2006-01-02T15:04:05"

func unmarshTasks(t *Tasks) []string {
	data := []string{strconv.Itoa(t.ID),
		t.Description,
		t.CreatedAt.Format(layOut),
		strconv.FormatBool(t.IsComplete)}
	return data
}

func marshTasks(data []string) (*Tasks, error) {
	id, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, err
	}

	b, err := strconv.ParseBool(data[3])
	if err != nil {
		return nil, err
	}

	t, err := time.Parse(layOut, data[2])

	if err != nil {
		return nil, err
	}
	task := &Tasks{ID: id, Description: data[1], CreatedAt: t, IsComplete: b}
	return task, nil
}
