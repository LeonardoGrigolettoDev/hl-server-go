package model

type Device struct {
	ID          string `json:"id"`
	Device_type string `json:"device_type"`
	Mac         string `json:"mac_address"`
	Created_at  string `json:"created_at"`
}
