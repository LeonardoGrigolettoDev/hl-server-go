package repository

import (
	"database/sql"
	"fmt"

	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/model"
)

type DeviceRepository struct {
	connection *sql.DB
}

func NewDeviceRepository(connection *sql.DB) DeviceRepository {
	return DeviceRepository{
		connection: connection,
	}
}

func (dr *DeviceRepository) GetDevices() ([]model.Device, error) {
	query := "SELECT * FROM device"
	rows, err := dr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Device{}, err
	}

	var deviceList []model.Device
	var deviceObj model.Device

	for rows.Next() {
		err = rows.Scan(
			&deviceObj.ID,
			&deviceObj.Name,
			&deviceObj.Device_type,
			&deviceObj.Mac,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Device{}, err
		}

		deviceList = append(deviceList, deviceObj)
	}
	rows.Close()
	return deviceList, nil
}
