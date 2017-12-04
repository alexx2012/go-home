package service

import (
	"github.com/alexx2012/go-home/service/device/persistence/model"
	"github.com/alexx2012/go-home/service/device/proto/device"
	"golang.org/x/net/context"
)

func (ref *DeviceService) AddRawDevice(ctx context.Context, in *device.RawDevice) (*device.Unit, error) {
	var (
		parentDeviceId = byte(in.ParentDeviceId)
		subDeviceId    = byte(in.SubDeviceId)
	)

	err := ref.repository.Save(&model.RawDevice{
		Id:             model.RawDeviceId(parentDeviceId, subDeviceId),
		ParentDeviceId: parentDeviceId,
		SubDeviceId:    subDeviceId,
		SubDeviceType:  byte(in.SubDeviceType),
	})

	if err != nil {
		return nil, err
	}

	return &device.Unit{}, nil
}
