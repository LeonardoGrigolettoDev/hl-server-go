package model

type Device struct {
	ID          int    `json:"id_device"`
	Name        string `json:"name"`
	Device_type string `json:"device_type"`
	Mac         string `json:"mac_address"`
}
