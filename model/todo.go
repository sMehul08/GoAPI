package model

func CreateToDo() error {
	insertQ, err := con.Query("INSERT INTO ToDo VALUES(?, ?)", "Test", "123")
	defer insertQ.Close()

	if err != nil {
		return err
	}

	return nil
}
