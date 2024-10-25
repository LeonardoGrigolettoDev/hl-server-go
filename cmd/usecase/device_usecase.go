package usecase

import (
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/model"
	"github.com/LeonardoGrigolettoDev/hl-server-go/cmd/repository"
)

type DeviceUsecase struct {
	repository repository.DeviceRepository
}

func NewDeviceUsecase(repo repository.DeviceRepository) DeviceUsecase {
	return DeviceUsecase{repository: repo}
}

func (du *DeviceUsecase) GetDevices() ([]model.Device, error) {
	return du.repository.GetDevices()
}

func (du *DeviceUsecase) CreateDevice(device model.Device) (model.Device, error) {
	id_device, err := du.repository.CreateDevice(device)
	if err != nil {
		return model.Device{}, err
	}

	device.ID = id_device

	return device, nil
}

func (du *DeviceUsecase) GetDeviceById(id_device int) (*model.Device, error) {
	device, err := du.repository.GetDeviceById(id_device)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (du *DeviceUsecase) UpdateDeviceById(device *model.Device) (*model.Device, error) {
	device, err := du.repository.UpdateDeviceById(*device)
	if err != nil {
		return nil, err
	}
	return device, nil
}
