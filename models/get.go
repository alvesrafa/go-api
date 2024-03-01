package models

import "alvesrafa.dev/study/database"

func Get(id int64) (todo Todo, err error) {

	conn, err := database.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sqlStatment := `SELECT * FROM todos WHERE id=$1`

	row := conn.QueryRow(sqlStatment, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
