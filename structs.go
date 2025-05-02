package main

type DataFrame struct {
	Header    string `json:"header"`
	DeviceID  uint16 `json:"deviceId"`
	Timestamp uint32 `json:"timestamp"`
	Voltage   uint16 `json:"voltage"`
	Current   uint16 `json:"current"`
	Checksum  string `json:"checksum"`
}
