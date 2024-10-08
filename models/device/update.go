package models

import db "github.com/LeonardoGrigolettoDev/fly-esp-server-go/database/postgre"

func Update(id int64, device Device) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	res, err := conn.Exec(`UPDATE device SET mac=$2, actionable=$3, camera=$4, drone=$4 WHERE id=$1`)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
