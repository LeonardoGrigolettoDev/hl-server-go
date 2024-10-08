package models

import db "github.com/LeonardoGrigolettoDev/fly-esp-server-go/database/postgre"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	res, err := conn.Exec(`DELETE device FROM device WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
