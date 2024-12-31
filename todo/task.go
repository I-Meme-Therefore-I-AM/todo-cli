package todo

import (
	"fmt"
	"os"
	"time"
	"todo-cli/util"
)

type TaskTodo interface {
	Add(activity string)
	Delete(id string)
	List() []string
	Complete()
}

type Tasks struct {
	ID          int       `csv:"ID"`
	Description string    `csv:"Descripion"`
	CreatedAt   time.Time `csv:"CreatedAt"`
	IsComplete  bool      `csv:"IsComplete"`
}

const FILENAME string = "todo.csv"

var csvHeader = []string{"ID", "Description", "CreatedAt", "IsComplete"}

func (t *Tasks) Add() {
	var f *os.File
	var err error

	if util.FileExists(FILENAME) {
		f, err = util.LoadFile(FILENAME)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		defer f.Close()
		err = writeToCsv(f, csvHeader)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	} else {
		f, err = util.LoadFile(FILENAME)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	err = appendToCsv(f, *t)

	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
	} else {
		fmt.Fprintln(os.Stdout, "task created successful")
	}
}

func (t *Tasks) List() {
	// list tasks
}
