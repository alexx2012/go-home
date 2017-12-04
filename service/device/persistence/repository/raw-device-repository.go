package repository

import (
	"errors"

	"github.com/alexx2012/go-home/service/device/persistence/model"
)

var (
	ErrRawDeviceNotFound = errors.New("RAW_DEVICE_NOT_FOUND")
)

type RawDeviceRepositoryInterface interface {
	Save(*model.RawDevice) error
	FindAll() ([]*model.RawDevice, error)
	FindById(id uint16) (*model.RawDevice, error)
	RemoveById(id uint16) error
}
