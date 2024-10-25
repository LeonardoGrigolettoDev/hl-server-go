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
			&deviceObj.Device_type,
			&deviceObj.Mac,
			&deviceObj.Created_at,
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

func (dr *DeviceRepository) GetDeviceById(id_device int) (*model.Device, error) {
	query, err := dr.connection.Prepare("SELECT * FROM device WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var device model.Device

	err = query.QueryRow(id_device).Scan(
		&device.ID,
		&device.Device_type,
		&device.Mac,
		&device.Created_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()

	return &device, nil
}

func (dr *DeviceRepository) CreateDevice(device model.Device) (string, error) {
	var id string
	query, err := dr.connection.Prepare("INSERT INTO device" +
		"(id, device_type, mac)" +
		" VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = query.QueryRow(device.ID, device.Device_type, device.Mac).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	query.Close()
	return id, nil
}

func (dr *DeviceRepository) UpdateDeviceById(device model.Device) (*model.Device, error) {
	query, err := dr.connection.Prepare("UPDATE device" +
		" SET id = $1, device_type = $2, mac = $3 " +
		"WHERE id = $4" +
		" RETURNING id, device_type, mac, created_at")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = query.QueryRow(device.ID, device.Device_type, device.Mac, device.ID).Scan(
		&device.ID, &device.Device_type, &device.Mac, &device.Created_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()

	return &device, nil
}
