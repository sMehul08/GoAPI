package model

import (
	"GoAPI/view"
	"fmt"
)

func CreateToDo(name, todo string) error {
	fmt.Println(name, todo)
	insertQ, err := con.Query("INSERT INTO ToDo VALUES(?, ?)", name, todo)
	defer insertQ.Close()

	if err != nil {
		return err
	}

	return nil
}

func ReadAll() ([]view.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM ToDo")
	if err != nil {
		return nil, err
	}

	todos := []view.PostRequest{}

	for rows.Next() {
		data := view.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}

	return todos, err
}

func ReadByName(name string) ([]view.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM ToDo WHERE name = ?", name)
	if err != nil {
		return nil, err
	}

	todos := []view.PostRequest{}

	for rows.Next() {
		data := view.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}

	return todos, err
}
