package model

import "fmt"

func CreateToDo(name, todo string) error {
	fmt.Println(name, todo)
	insertQ, err := con.Query("INSERT INTO ToDo VALUES(?, ?)", name, todo)
	defer insertQ.Close()

	if err != nil {
		return err
	}

	return nil
}
