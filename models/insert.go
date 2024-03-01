package models

import "alvesrafa.dev/study/database"

func Insert(todo Todo) (id int64, err error) {

	conn, err := database.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sqlStatment := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sqlStatment, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
