package main

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
)

func DecodeDataFrame(hexStr string) (*DataFrame, error) {
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}

	if len(data) < 16 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "data frame too short")
	}

	return &DataFrame{
		Header:    string(data[0:4]),
		DeviceID:  binary.BigEndian.Uint16(data[4:6]),
		Timestamp: binary.BigEndian.Uint32(data[6:10]),
		Voltage:   binary.BigEndian.Uint16(data[10:12]),
		Current:   binary.BigEndian.Uint16(data[12:14]),
		Checksum:  hex.EncodeToString(data[14:16]),
	}, nil
}

func ValidateChecksum(frame *DataFrame) bool {
	// Example dummy checksum logic (use your actual logic here)
	calculated := (frame.DeviceID + frame.Voltage + frame.Current) % 256
	received, err := hex.DecodeString(frame.Checksum)
	if err != nil || len(received) != 1 {
		return false
	}
	return calculated == uint16(received[0])
}
