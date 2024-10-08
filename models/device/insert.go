package models

import (
	db "github.com/LeonardoGrigolettoDev/fly-esp-server-go/database/postgre"
)

func Insert(device Device) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTRO device (mac, actionable, camera, drone) VALUES ($1, $2, $3, $4) RETURN id`

	err = conn.QueryRow(sql, device.MAC, device.Actionable, device.Camera, device.Drone).Scan(id)

	return
}
