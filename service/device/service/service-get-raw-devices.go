package service

import (
	"github.com/alexx2012/go-home/service/device/proto/device"
	"golang.org/x/net/context"
)

func (ref *DeviceService) GetRawDevices(context.Context, *device.Unit) (*device.RawDevices, error) {
	devices, err := ref.repository.FindAll()

	if err != nil {
		return nil, err
	}

	out := &device.RawDevices{
		Devices: make([]*device.RawDevice, len(devices)),
	}

	for i, d := range devices {
		out.Devices[i] = &device.RawDevice{
			ParentDeviceId: int32(d.ParentDeviceId),
			SubDeviceId:    int32(d.SubDeviceId),
			SubDeviceType:  int32(d.SubDeviceType),
		}
	}

	return out, nil
}
