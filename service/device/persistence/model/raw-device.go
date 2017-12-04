package model

import (
	"encoding/binary"
)

type RawDevice struct {
	Id             uint16
	ParentDeviceId byte
	SubDeviceId    byte
	SubDeviceType  byte
}

func RawDeviceId(parentDeviceId byte, subDeviceId byte) uint16 {
	return binary.LittleEndian.Uint16([]byte{parentDeviceId, subDeviceId})
}
