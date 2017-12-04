package service

import (
	"github.com/alexx2012/go-home/service/device/proto/device"
	"golang.org/x/net/context"
)

func (ref *DeviceService) RemoveRawDevice(ctx context.Context, in *device.RawDeviceId) (*device.Unit, error) {
	if err := ref.repository.RemoveById(uint16(in.Value)); err != nil {
		return nil, err
	}

	return &device.Unit{}, nil
}
