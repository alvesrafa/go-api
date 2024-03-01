package models

import "alvesrafa.dev/study/database"

func Update(id int64, todo Todo) (int64, error) {

	conn, err := database.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	sqlStatment := `UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$4`

	res, err := conn.Exec(sqlStatment, id, todo.Title, todo.Description, todo.Done)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
