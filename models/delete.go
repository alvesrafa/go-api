package models

import "alvesrafa.dev/study/database"

func Delete(id int64) (int64, error) {

	conn, err := database.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	sqlStatment := `DELETE FROM todos WHERE id=$1`

	res, err := conn.Exec(sqlStatment, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
