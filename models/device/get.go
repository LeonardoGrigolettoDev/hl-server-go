package models

import (
	db "github.com/LeonardoGrigolettoDev/fly-esp-server-go/database/postgre"
)

func Get(id int64) (device Device, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	row := conn.QueryRow(`SELECT * FROM device WHERE id=$1`, id)
	err = row.Scan(&device.ID, &device.MAC, &device.Actionable, &device.Camera, &device.Drone)
	return
}

func GetAll() (devices []Device, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	rows, err := conn.Query(`SELECT * FROM device`)

	if err != nil {
		return
	}

	for rows.Next() {
		var device Device
		err = rows.Scan(&device.ID, &device.MAC, &device.Actionable, &device.Camera, &device.Drone)
		if err != nil {
			continue
		}
		devices = append(devices, device)
	}
	return
}
