package service

import "github.com/alexx2012/go-home/service/device/persistence/repository"

type DeviceService struct {
	repository repository.RawDeviceRepositoryInterface
}

func NewDeviceService(repository repository.RawDeviceRepositoryInterface) *DeviceService {
	return &DeviceService{
		repository: repository,
	}
}
