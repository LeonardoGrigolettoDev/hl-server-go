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
